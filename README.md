# 🍼 Baby Starter: Your First Go Backend

Welcome to Baby Starter! This project is designed to be a simple and educational starting point for new Go developers venturing into backend development. It provides a clean, organized, and easy-to-understand structure that you can build upon.

## ✨ Features

*   **Live Reloading:** Uses [Air](https://github.com/cosmtrek/air) for automatic rebuilding and restarting of the application when you make changes.
*   **Structured Logging:** Implements `zerolog` for structured, leveled logging.
*   **Easy Configuration:** Uses `godotenv` and `env` to manage environment variables.
*   **Database Ready:** Comes with `gorm` and `sqlite3` for a simple database setup.
*   **Request Validation:** Includes an example of request validation using the `zog` library.
*   **Clear Project Structure:** The project is organized into logical directories for easy navigation and extension.

## 🚀 Getting Started

### Prerequisites

*   [Go](https://golang.org/dl/)
*   [Air](https://github.com/cosmtrek/air) (for live reloading)

### Installation & Running

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/your-username/baby_starter.git
    ```
2.  **Navigate to the project directory:**
    ```bash
    cd baby_starter
    ```
3.  **Install dependencies:**
    ```bash
    go mod tidy
    ```
4.  **Run the application with live reloading:**
    ```bash
    air
    ```

The application will be running at `http://localhost:8080`.

## 📂 Project Structure

The project is structured to separate concerns and make it easy to find what you're looking for.

```
.
├── app/
│   └── app.go           # Core application setup (logging, env vars)
├── database/
│   ├── database.go      # Database initialization and migrations
│   └── model/
│       └── user.go      # GORM model for the User
├── handler/
│   ├── get_index.go     # Handler for the GET / route
│   ├── post_user.go     # Handler for the POST /api/user route
│   └── handler.go       # Helper for request parsing and validation
├── server/
│   ├── middlewares.go   # Middlewares for the Echo server
│   ├── routes.go        # Route definitions
│   └── server.go        # Echo server setup and startup
├── util/
│   └── password.go      # Password hashing utilities
├── .air.toml            # Configuration for Air (live reloading)
├── .gitignore           # Files and directories to be ignored by Git
├── go.mod & go.sum      # Go module dependencies
└── main.go              # Main entry point of the application
```

## 💡 Core Concepts

This section explains the key parts of the application and how they work together.

### ⚙️ Configuration (`app/app.go`)

The application is configured through environment variables. The `app.Init()` function, called from `main.go`, does the following:

1.  **Initializes `zerolog`:** Sets up a console writer for structured logging.
2.  **Loads `.env` file:** Uses `godotenv` to load environment variables from a `.env` file (if it exists).
3.  **Parses Environment Variables:** Uses the `caarlos0/env` library to parse environment variables into the `app.ENV` struct. This provides default values and a clear structure for your configuration.

### 🌐 Routing (`server/routes.go` & `server/server.go`)

The application uses the [Echo](https://echo.labstack.com/) web framework.

*   **`server.Start()`:** This function, called from `main.go`, creates a new Echo instance, initializes middlewares, sets up routes, and starts the server.
*   **`initRoutes(e *echo.Echo)`:** This is where you define your application's routes. It maps HTTP methods and paths to specific handler functions.

### 📦 Handlers (`handler/`)

Handlers are the functions that process incoming HTTP requests and generate responses.

#### `handler.GetIndex`

This is a simple handler that returns a "Hello World" string. It demonstrates the basic structure of a handler.

```go
func GetIndex(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}
```

#### `handler.PostUser`

This handler is a more complete example that demonstrates:

1.  **Request Validation:** It uses a helper function `ParseReq` to validate the incoming JSON request body against a schema defined with the `zog` library.

    *   The schema in `PostUser` defines the expected shape of the request body, including data types, required fields, and constraints (e.g., min/max length, email format).
    *   If validation fails, it returns a `400 Bad Request` with a JSON object containing the validation errors.

2.  **Password Hashing:** It uses the `util.HashPassword` function to securely hash the user's password before storing it.

3.  **Database Interaction:**
    *   It checks if a user with the given email already exists.
    *   It creates a new `model.User` and saves it to the database.

4.  **JSON Response:** It returns a `200 OK` with the newly created user object as JSON (excluding the password).

### ✅ Validating Requests (`handler/handler.go`)

The `ParseReq` function is a generic helper that simplifies request validation.

```go
func ParseReq(c echo.Context, schema *z.StructSchema, destPtr any) (sanitizedErrs map[string][]string, ok bool) {
    // ...
}
```

You can use it in your handlers to validate any request body by providing a `zog` schema. This promotes code reuse and keeps your handlers clean.

### 🗄️ Database (`database/`)

*   **`database.Init()`:** This function, called from `main.go`, initializes the database connection.
    *   It creates the data directory if it doesn't exist.
    *   It opens a connection to the SQLite database using `gorm`.
    *   It runs `AutoMigrate` to automatically create or update the database schema based on your GORM models.
*   **`database/model/user.go`:** This file defines the `User` model using GORM struct tags. These tags are used to define the table schema, constraints, and relationships.

### 쿼리 (Querying the Database)

This project uses GORM to interact with the database. Here are some examples of how you can query the database:

*   **Get the first record:**

    ```go
    user := model.User{}
    if err := gorm.G[model.User](app.DB).Where("email = ?", "test@example.com").First(c.Request().Context(), &user); err != nil {
        // Handle error
    }
    ```

*   **Get all records:**

    ```go
    var users []model.User
    if err := gorm.G[model.User](app.DB).Find(c.Request().Context(), &users); err != nil {
        // Handle error
    }
    ```

*   **Get a record by ID:**

    ```go
    var user model.User
    if err := gorm.G[model.User](app.DB).First(c.Request().Context(), &user, c.Param("id")); err != nil {
        // Handle error
    }
    ```

## 🔧 Extending the Application

Here are some examples of how you can extend this starter project.

### Add a New Route and Handler

1.  **Create a new handler function** in a new file in the `handler/` directory (e.g., `handler/get_users.go`). This handler could fetch all users from the database.

    ```go
    // handler/get_users.go
    package handler

    import (
        "baby_starter/app"
        "baby_starter/database/model"
        "net/http"

        "github.com/labstack/echo/v4"
        "gorm.io/gorm"
    )

    func GetUsers(c echo.Context) error {
        var users []model.User
        if err := gorm.G[model.User](app.DB).Find(c.Request().Context(), &users); err != nil {
            return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch users"})
        }
        return c.JSON(http.StatusOK, users)
    }
    ```

2.  **Register the new route** in `server/routes.go`:

    ```go
    // server/routes.go
    func initRoutes(e *echo.Echo) {
        e.GET("/", handler.GetIndex)
        e.POST("/api/user", handler.PostUser)
        e.GET("/api/users", handler.GetUsers) // Add this line
    }
    ```

### Add a New Model

1.  **Define the new model** in a new file in `database/model/` (e.g., `database/model/product.go`).

    ```go
    // database/model/product.go
    package model

    import "gorm.io/gorm"

    type Product struct {
        gorm.Model
        Name  string
        Price uint
    }
    ```

2.  **Add the new model to the `AutoMigrate`** call in `database/database.go`:

    ```go
    // database/database.go
    func Init() {
        // ...
        if err := db.AutoMigrate(&model.User{}, &model.Product{}); err != nil { // Add &model.Product{}
            app.LOG.Fatal().Err(err).Msg("failed to run migrations")
        }
    }
    ```

## 🤔 What's Next?

Now that you have a basic understanding of the project, here are some ideas for what you can do next:

*   **Add Authentication:** Implement user authentication using JWTs (JSON Web Tokens).
*   **Write Tests:** Write unit and integration tests for your handlers and database logic.
*   **Add More Models:** Extend the application with more complex data models and relationships.
*   **Error Handling:** Implement a more robust error handling strategy.
*   **Deployment:** Deploy your application to a cloud provider like Heroku or AWS.

## 🤝 Contributing

Contributions are welcome! If you have any ideas, suggestions, or improvements, feel free to open an issue or submit a pull request.
