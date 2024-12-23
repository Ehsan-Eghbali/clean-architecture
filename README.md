
# Clean Architecture in Go

This repository provides a sample Go project implementing the **Clean Architecture** approach. By separating responsibilities into distinct layers, the code becomes more readable, maintainable, and extensible.

## Project Structure

```
clean-architecture-go
├── cmd
│   └── main.go
├── internal
│   ├── domain
│   │   └── user.go
│   ├── usecase
│   │   ├── user_usecase.go
│   │   └── user_usecase_impl.go
│   ├── repository
│   │   ├── user_repository.go
│   │   └── memory
│   │       └── user_repository_memory.go
│   └── delivery
│       └── http
│           ├── handler.go
│           └── router.go
└── go.mod
```

- **cmd/**  
  Contains the entry point (`main.go`) for the application.

- **internal/domain/**  
  Holds the core **domain entities** (e.g., `User`) that represent the business objects.

- **internal/usecase/**  
  Contains the **business logic** and **use cases** that operate on domain entities (e.g., user creation, retrieval).

- **internal/repository/**  
  Defines **interfaces** and **implementations** for data access (e.g., database, in-memory storage).

- **internal/delivery/**  
  Handles the **presentation/delivery** layer. In this sample, we use [gorilla/mux](https://github.com/gorilla/mux) to build a simple HTTP API.

## Prerequisites

- **Go** version 1.20 or higher  
- Internet access (for fetching third-party libraries such as Gorilla Mux)

## Getting Started

1. **Clone** or **download** this repository and unzip it if necessary.  
2. Navigate to the project directory:
   ```bash
   cd clean-architecture-go
   ```
3. If needed, install dependencies:
   ```bash
   go mod tidy
   ```
4. Run the application:
   ```bash
   go run ./cmd/main.go
   ```
5. The server will start on port `8080`. Use any REST client (e.g., `curl`, Postman) to send requests.

## API Endpoints

- **Create a New User (POST)**  
  ```
  POST /users
  Content-Type: application/json
  
  {
      "Username": "testuser",
      "Email": "test@example.com"
  }
  ```
  **Sample Successful Response**:
  ```json
  {
    "ID": 1,
    "Username": "testuser",
    "Email": "test@example.com"
  }
  ```

- **Get All Users (GET)**  
  ```
  GET /users
  ```
  **Sample Successful Response**:
  ```json
  [
    {
      "ID": 1,
      "Username": "testuser",
      "Email": "test@example.com"
    }
  ]
  ```

- **Get a User by ID (GET)**  
  ```
  GET /users/{id}
  ```
  **Example**:  
  ```
  GET /users/1
  ```
  **Sample Successful Response**:
  ```json
  {
    "ID": 1,
    "Username": "testuser",
    "Email": "test@example.com"
  }
  ```

## Important Notes

1. This sample uses **in-memory** storage. Data is lost when the application restarts. You can switch to a real database by modifying the repository layer (e.g., using PostgreSQL, MySQL, MongoDB, or an ORM like GORM).
2. Clean Architecture aims to **decouple** domain logic from external details (like frameworks or databases). Hence, the domain layer contains only the business entities and core rules, while implementation details live in the lower layers.
3. Be sure to update the module name in `go.mod` to a valid name if you change the project structure.
4. To avoid naming conflicts between the standard `net/http` package and the `internal/delivery/http` package, we use an **alias** (e.g., `deliveryHttp`) in some of the files.

## Contributing

- Feel free to **fork** this repository and make changes or add new features.
- If you find any issues or have suggestions for improvement, please open an **Issue** in this repository.

---

Thank you for checking out this Clean Architecture sample project in Go! We hope it helps illustrate the key principles of Clean Architecture.
