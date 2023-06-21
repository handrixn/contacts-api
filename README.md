# Contact Management RESTful API

This project is a Contact Management RESTful API implemented in Golang. It provides CRUD (Create, Read, Update, Delete) operations for managing contacts. The API supports sorting, filtering, and pagination.

## Features

- Create a new contact
- Retrieve a contact by ID
- Update an existing contact
- Delete a contact
- Get all contacts with sorting, filtering, and pagination

## Technologies Used

- Golang
- RESTful API principles
- JSON file storage

## Getting Started

### Prerequisites

- Go (version 1.20 or higher) installed
- Git installed (optional)

### Installation

1. Clone the repository:

   `git clone https://github.com/handrixn/contacts-api.git`

   Alternatively, you can download the project files manually.

2. Change to the project directory:

   `cd contacts-api`

3. Build and run the project:

   `make build`

4. Run the project:

   `./make run`

   The API will start running on `http://localhost:8080`.

### API Documentation

The API documentation (Swagger) is available at `http://localhost:8080/docs`.

### Usage

- Create a new contact:
  - Method: POST
  - Endpoint: `/contacts`
  - Request Body:
    ```json
    {
      "name": "John Doe",
      "gender": "male",
      "phone": "123456789",
      "email": "john.doe@example.com"
    }
    ```
- Retrieve a contact by ID:
  - Method: GET
  - Endpoint: `/contacts/{id}`
- Update an existing contact:
  - Method: PUT
  - Endpoint: `/contacts/{id}`
  - Request Body:
    ```json
    {
      "name": "Updated Name",
      "phone": "987654321"
    }
    ```
- Delete a contact:
  - Method: DELETE
  - Endpoint: `/contacts/{id}`
- Get all contacts with sorting, filtering, and pagination:
  - Method: GET
  - Endpoint: `/contacts`

### Build and Run (Alternative Method)

If you prefer to manually build and run the project without using the Makefile, you can follow these steps:

1. Build the project:

   `go build -o contact-api`

2. Run the project:

   `./contact-api`

Remember to have the `go` command available in your system's PATH.

For cleaning the generated binary file, you can use the command:

`make clean`
