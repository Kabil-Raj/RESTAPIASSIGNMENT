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
	myRouter.HandleFunc("/storedata", productScrappedData).Methods("POST")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("STOREDATA_PORT"), myRouter))
}

func productScrappedData(w http.ResponseWriter, req *http.Request) {
	var result map[string]string
	json.NewDecoder(req.Body).Decode(&result)
	saveDataInDatabase(result["ProductName"], result["ProductImageUrl"], result["ProductDescription"], result["ProductPrice"], result["ProductReviews"], time.Now())
}

func main() {
	handleRequests()
}

func mysqlConnectionURL() string {
	mysqlConnectionURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_ROOT_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
	return mysqlConnectionURL
}

func logError(err error) {
	log.Fatal(err)
}

func saveDataInDatabase(productName string, productImageUrl string, productDescription string, productPrice string, productReviews string, createdTime time.Time) {

	// Connecting to database
	db, err := sql.Open("mysql", mysqlConnectionURL())
	if err != nil {
		logError(err)
	}

	// Insert values into database
	sqlInsertStatement, err := db.Prepare("INSERT INTO AmazonProductDetails (ProductName,ProductImageUrl,ProductDescription,ProductPrice,ProductReviews,CreatedTime) VALUES (?,?,?,?,?,?);")
	if err != nil {
		logError(err)
	}

	_, err = sqlInsertStatement.Exec(productName, productImageUrl, productDescription, productPrice, productReviews, createdTime)
	if err != nil {
		logError(err)
	}

	defer db.Close()
}
