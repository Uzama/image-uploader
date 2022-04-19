# image-uploader
A simple Golang application to upload image files. The application code is organized using **Clean Architecture**


### How to start 

- Make sure database is up and running and you have update the ```configurations/database.yaml``` file with relevant values.
- Run the ```database.sql``` query to get database and table created.
- Clone the service locally and run the service by typing ```go run main.go```
- make sure service is up and running. 
- Now you can send request to the service.

#### Upload file

```http
    http://localhost:8080/upload
```