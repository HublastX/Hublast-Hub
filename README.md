# Project Management System

This project is a web-based project management system built using Go, Gin Gonic for the API, and GORM for the ORM with PostgreSQL as the database. It provides functionalities for managing projects, tasks, and users.

## Features

- Create, read, update, and delete projects
- Create, read, update, and delete tasks
- User management with authentication
- Middleware for logging and authentication
- RESTful API design

## Technologies Used

- Go (Golang)
- Gin Gonic
- GORM
- PostgreSQL
- Docker

## Getting Started

### Prerequisites

- Go installed on your machine
- PostgreSQL installed and running
- Docker (optional, for containerization)

### Installation

1. Clone the repository:

   ```
   git clone <repository-url>
   cd github.com/HublastX/HubLast-Hub
   ```

2. Create a `.env` file based on the `.env.example` provided and configure your database connection settings.

3. Install the necessary Go dependencies:

   ```
   go mod tidy
   ```

### Running the Application

To run the application, you can use the following command:

```
go run cmd/main.go
```

Alternatively, if you are using Docker, you can build and run the Docker container:

```
docker-compose up --build
```

### API Endpoints

- **Projects**

  - `POST /projects` - Create a new project
  - `GET /projects/:id` - Get a project by ID
  - `PUT /projects/:id` - Update a project by ID
  - `DELETE /projects/:id` - Delete a project by ID

- **Tasks**

  - `POST /tasks` - Create a new task
  - `GET /tasks/:id` - Get a task by ID
  - `PUT /tasks/:id` - Update a task by ID
  - `DELETE /tasks/:id` - Delete a task by ID

- **Users**
  - `POST /users` - Create a new user
  - `GET /users/:id` - Get a user by ID
  - `PUT /users/:id` - Update a user by ID
  - `DELETE /users/:id` - Delete a user by ID

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.
