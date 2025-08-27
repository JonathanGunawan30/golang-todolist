# Go Clean Architecture Todolist API

A simple REST API for a Todolist application built with Go (Golang) and the Fiber web framework. This project serves as a practical example of implementing a clean, layered architecture for educational purposes.

The primary goal of this repository is to demonstrate how to structure a Go application in a way that is modular, scalable, and easy to test by separating concerns into distinct layers.

---
## ## Features

* **Layered Architecture**: Follows a clean architecture pattern (Handler, Usecase/Service, Repository).
* **RESTful API**: Provides full CRUD (Create, Read, Update, Delete) endpoints for managing activities.
* **Structured Logging**: Integrated with Logrus for JSON-formatted logs.
* **Configuration Management**: External configuration using Viper (`config.yaml`).
* **Database Pooling**: Utilizes Go's built-in `database/sql` for efficient connection pooling with PostgreSQL.
* **Testing**: Integration tests.
* **API Documentation**: A complete OpenAPI 3.0 specification is provided in `apispec.yaml`.

---
## ## Tech Stack & Libraries

* **Go**: The core programming language.
* **PostgreSQL**: SQL database for data persistence.
* **Fiber**: A fast and expressive web framework for building the API.
* **Logrus**: A structured logging library.
* **Viper**: A complete configuration solution.
* **Go Playground Validator**: For validating incoming request data.
* **Testify**: A comprehensive toolkit for Go testing, used for assertions, mocking (`testify/mock`), and test suites (`testify/suite`).
* **lib/pq**: The standard Go driver for PostgreSQL.

---
## ## Getting Started

### Prerequisites

* Go (version 1.18 or newer)
* PostgreSQL

### Installation & Setup

1.  **Clone the repository:**
    ```bash
    git clone [https://github.com/YOUR_USERNAME/todolist-v1.git](https://github.com/YOUR_USERNAME/todolist-v1.git)
    cd todolist-v1
    ```

2.  **Set up the Database:**
    Connect to your PostgreSQL instance and run the following SQL to create the required types and table:
    ```sql
    CREATE TYPE status AS ENUM ('NEW', 'ON PROGRESS', 'EXPIRED');
    CREATE TYPE category_type AS ENUM ('TASK', 'EVENT');

    CREATE TABLE activities (
        id SERIAL PRIMARY KEY,
        title VARCHAR(250) NOT NULL,
        category category_type NOT NULL,
        description TEXT NOT NULL,
        activity_date TIMESTAMPTZ NOT NULL,
        status status DEFAULT 'NEW'
    );
    ```

3.  **Configure the Application:**
    Create a `config.yaml` file in the root of the project and add your configuration details:
    ```yaml
    server:
      port: "3000"

    database:
      url: "postgresql://user:password@host:port/database_name"
    ```

4.  **Install Dependencies:**
    ```bash
    go mod tidy
    ```

5.  **Run the Application:**
    ```bash
    go run main.go
    ```
    The server will start on the port specified in your `config.yaml`.

---
## ## API Endpoints

| Method | Endpoint              | Description              |
|--------|-----------------------|--------------------------|
| `GET`  | `/api/activities`     | Get all activities       |
| `POST` | `/api/activities`     | Create a new activity    |
| `PUT`  | `/api/activities/{id}`| Update an existing activity |
| `DELETE`| `/api/activities/{id}`| Delete an activity       |

---
## ## Running Tests

This project includes integration tests.

```bash
go test ./... -v
