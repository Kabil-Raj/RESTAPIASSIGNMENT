package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/gorilla/mux"
)

type ProductDetail struct {
	ProductName        string `json:"ProductName"`
	ProductImageUrl    string `json:"ProductImageUrl"`
	ProductDescription string `json:"ProductDescription"`
	ProductPrice       string `json:"ProductPrice"`
	ProductReviews     string `json:"ProductReviews"`
}

var ProductDetails []ProductDetail

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/scrap/product", scrapAmazonProduct).Methods("POST")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("SCRAPPER_PORT"), myRouter))
}

func scrapAmazonProduct(w http.ResponseWriter, req *http.Request) {
	productUrl := req.URL.Query().Get("url")
	if productUrl == "" {
		fmt.Fprintf(w, "Please provide proper url")
	} else {
		getProductDetails(productUrl)
		json.NewEncoder(w).Encode(ProductDetails)
		fmt.Println("Endpoint Hit : return product details")
	}
}

func main() {

	handleRequests()
}

func getProductDetails(productUrl string) {

	var productName string

	var productImageUrl string

	var productPrice string

	var productReviews string

	var productDescription string

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML("#productDescription", func(e *colly.HTMLElement) {
		productDescription = strings.TrimSpace(e.DOM.Children().Text())
	})

	c.OnHTML("#acrCustomerReviewText", func(e *colly.HTMLElement) {
		productReviews = e.Text
	})

	c.OnHTML("#desktop_unifiedPrice", func(e *colly.HTMLElement) {
		// normal amazon price
		e.DOM.Find("#priceblock_ourprice").Each(func(i int, s *goquery.Selection) {
			productPrice = strings.TrimPrefix(s.Text(), "₹")
		})
		// deal price
		e.DOM.Find("#priceblock_dealprice").Each(func(i int, s *goquery.Selection) {
			productPrice = s.Text()
		})

	})

	c.OnHTML("h1", func(e *colly.HTMLElement) {
		var getProductName string
		getProductName = getProductName + e.DOM.Find("span").Text()
		for _, character := range getProductName {
			if character != 10 {
				productName = productName + string(character)
			}
		}

	})

	c.OnHTML("#imgTagWrapperId", func(e *colly.HTMLElement) {
		imgSource, isSourceImage := e.DOM.Children().Attr("src")
		if isSourceImage {
			productImageUrl = imgSource
		}
	})

	c.Visit(productUrl)

	ProductDetails = []ProductDetail{
		{ProductName: productName, ProductImageUrl: productImageUrl, ProductDescription: productDescription, ProductPrice: productPrice, ProductReviews: productReviews},
	}

	productData := map[string]string{
		"ProductName":        productName,
		"ProductImageUrl":    productImageUrl,
		"ProductDescription": productDescription,
		"ProductPrice":       productPrice,
		"ProductReviews":     productReviews,
	}

	json_data, err := json.Marshal(productData)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		http.Post("http://scrapperdatahandler:"+os.Getenv("STOREDATA_PORT")+"/storedata", "application/json", bytes.NewBuffer(json_data))
	}

}
