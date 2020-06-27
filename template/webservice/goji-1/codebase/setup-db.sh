#this file captures db installation using docker , if you dont want this then make sure database details are changed according to your database (environment.go,query.go)


docker run -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=my-secret-password  --name ginRest-database mysql


#create database users;
#create table user(email varchar(255) , name varchar(255), age integer,id varchar(255));
#create table auth (email varchar(255),password varchar(255));