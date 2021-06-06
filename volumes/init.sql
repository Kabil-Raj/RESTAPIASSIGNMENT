CREATE DATABASE IF NOT EXISTS AmazonProductDatabase;
USE AmazonProductDatabase;
CREATE TABLE IF NOT EXISTS AmazonProductDetails(id int NOT NULL AUTO_INCREMENT, ProductName varchar(255), ProductImageUrl varchar(255), ProductDescription varchar(10000), ProductPrice varchar(255), ProductReviews varchar(255),CreatedTime DATETIME,PRIMARY KEY (id));

