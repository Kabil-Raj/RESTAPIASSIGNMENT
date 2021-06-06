package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/storedata", storeProductData).Methods("POST")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("STOREDATA_PORT"), myRouter))
}

func storeProductData(w http.ResponseWriter, req *http.Request) {
	var productDetailMap map[string]string
	json.NewDecoder(req.Body).Decode(&productDetailMap)
	saveDataInDatabase(productDetailMap["ProductName"], productDetailMap["ProductImageUrl"], productDetailMap["ProductDescription"], productDetailMap["ProductPrice"], productDetailMap["ProductReviews"], time.Now())
}

func main() {
	handleRequests()
}

func connectMySql() *sql.DB {
	mysqlConnectionURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_ROOT_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
	db, err := sql.Open("mysql", mysqlConnectionURL)
	if err != nil {
		logError(err)
	}
	return db
}

func logError(err error) {
	log.Fatal(err)
}

func saveDataInDatabase(productName string, productImageUrl string, productDescription string, productPrice string, productReviews string, createdTime time.Time) {

	// Insert values into database
	sqlInsertStatement, err := connectMySql().Prepare("INSERT INTO AmazonProductDetails (ProductName,ProductImageUrl,ProductDescription,ProductPrice,ProductReviews,CreatedTime) VALUES (?,?,?,?,?,?);")
	if err != nil {
		logError(err)
	}

	_, err = sqlInsertStatement.Exec(productName, productImageUrl, productDescription, productPrice, productReviews, createdTime)
	if err != nil {
		logError(err)
	}

	defer connectMySql().Close()
}
