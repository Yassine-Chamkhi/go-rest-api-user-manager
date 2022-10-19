# RESTful user management CRUD API with Go

This is a **Golang** backend app that communicates with a **PostgreSQL** database using **pgx** driver, and that provides a web server and RESTful API endpoints using **gin-gonic**. Included in this repository are **unit tests** and also a Dockerfile and a github actions workflow for **Continuous Integration** and **Containerization**. This app was built to run inside a **Kubernetes** cluster, that is why you will find a **readiness probe** that consists in creating a file /var/ready to signal the readiness to K8s, and a /metrics endpoint that is exposed for monitoring purposes with **Prometheus**.

## Architecture

This app was built with SOLID priciples in mind. It is divided into smaller modules -each in its own package- that interact with each other using dependencies abstracted away by interfaces.
- In **models** we find the User model we used.
- The **repository** is the element responsible for communicating with the database.
- In **services** we find the user management service, namely UserService, that uses the repository and deals with some input validation.
- In **http/handlers** we find the HTTP Handlers that convert user requests and input and use the appropriate service method to deal with it.
- The **server** is the element that contains the gin engine (web server) and runs it, consequently accepting http requests. Checkout the server's InitRoutes method to find out what the available routes are.

## Testing, Containerization and Continuous Integration
The github actions workflow specified for this app is composed of two steps; 
1. Running the unit tests.
2. Building the docker image (from Dockerfile) and pushing it to Docker Hub.

The tests were written using **stretchr/testify** to create assertions and mocks and **bxcodec/faker** to create fake data.

## Running the app
To run the app locally, you need to have a PostgreSQL database set up and running (an sql migrate up operation runs automatically to create the table and populate it with 100 fake entries). You also need to create a .env file where you specify the following variables:
- DATABASE_USERNAME
- DATABASE_PASSWORD
- DATABASE_HOST
- DATABASE_PORT
- DATABASE_NAME

And then simply run ```go run .```
You can verify the app is running by accessing the url **http://localhost:8080/** ; this will be handled by UserHandler's Greet method (The '/' endpoint is also the liveness probe used by K8s to check the app's health).
