# API REST with Golang

This project is a RESTful API built with Golang. It provides a set of endpoints to perform CRUD operations.

## Features

- Create, read, update, and delete resources
- JSON-based API
- Simple and easy to extend

## Requirements

- Go 1.16 or higher
- A running instance of a database (e.g., PostgreSQL, MySQL)

## Installation

1. Clone the repository:
  ```sh
  git clone https://github.com/yourusername/api-rest.git
  ```
2. Navigate to the project directory:
  ```sh
  cd api-rest
  ```
3. Install dependencies:
  ```sh
  go mod tidy
  ```

## Configuration

Add datastore-service_account.json file in service_accounts folder

Create a `.env` file in the root directory and add your database configuration:
```env
DATASTORE_SERVICE_ACCOUNT_PATH=your/path/to/file/
```

## Running the Application

To start the server, run:
```sh
go run main.go
```

The server will start on `http://localhost:8080`.

## API Endpoints

- `GET /players` - List all players
- `GET /players/{id}` - Get a specific resource
- `POST /players` - Create a new resource
- `PUT /players/{id}` - Update a specific resource
- `DELETE /players/{id}` - Delete a specific resource

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.