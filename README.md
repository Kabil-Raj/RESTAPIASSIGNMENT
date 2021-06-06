# AmazonProductScrapperAssignment
Kindly follow below steps

clone the git repository using below command

git clone https://github.com/Kabil-Raj/RESTAPIASSIGNMENT.git

move to project root directory and build application using below command

docker-compose build

launh the application using below command

docker-compose up

Wait for few mins for application to run and connect to the database

Launch Postman or preferred tool, give post request to below url and provide the amazon product url in the body as json

http://localhost:8084/scrape/product

Body: { "url" : "https://www.amazon.in/Yuwell-Portable-Concentrator-Machine-concentration/dp/B093YBRNGY/ref=lp_6504395031_1_1?s=specialty-aps" }

I have comitted postman collection, kindly check for reference

Once the url has been hit, internally it will call second POST service and data will be stored in the database

For this project I have used MySQL DB, use below port to connect to MYsql DB

HOST = localhost

POST = 8307

DB DatabaseName AmazonProductDatabase

DB TableName AmazonProductDetails

DB_USER = admin

DB_PASSWORD = Password1@