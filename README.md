# AmazonProductScrapperAssignment
Kindly follow below steps

clone the git repository using below command
git clone https://github.com/Kabil-Raj/AmazonProductScrapperAssignment

move to project root directory and build application using below command
docker-compose build

launh the application using below command
docker-compose up

Wait for few mins for application to run and connect to the database

Launch Postman or preferred tool, give post request to below url and provide the amazon product url which has to be scrapped
http://localhost:8084/scrapproduct?url=
example: http://localhost:8084/scrapproduct?url=https://www.amazon.in/Yuwell-Portable-Concentrator-Machine-concentration/dp/B093YBRNGY/ref=lp_6504395031_1_1?s=specialty-aps

I have comitted postman collection

Once the url has been hit, internally it will call another POST service and data will be stored in the database
For this project I have used MySQL DB, connect to below port
localhost and port 8307




