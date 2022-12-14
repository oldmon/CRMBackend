# CRMBackend

The project represents the backend of a customer relationship management (CRM) web application. As users interact with the app via some user interface, your server will support all of the functionalities:
    Getting a list of all customers
    Getting data for a single customer
    Adding a customer
    Updating a customer's information
    Removing a customer

##installation

Download the main.go and main_test.go from the associated workspace
Download the external 3rd Party Modules
    github.com/google/uuid - Used to generate a unique id
    github.com/gorilla/mux - Used as the http router
go mod init
go mod tidy

##launch

go run main.go

##usage

The application handles the following 5 operations for customers in the "database":
    Getting a single customer through a /customers/{id} path
    Getting all customers through a the /customers path
    Creating a customer through a /customers path
    Updating a customer through a /customers/{id} path
    Deleting a customer through a /customers/{id} path
Each RESTful route is associated with the correct HTTP verb.