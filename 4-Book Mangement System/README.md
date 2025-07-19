# Book Management System

A RESTful API built with Go for managing a book library system. This application provides complete CRUD operations for books with MySQL database integration using GORM ORM.

## Features

- **Complete CRUD Operations**: Create, Read, Update, and Delete books
- **RESTful API Design**: Follows REST conventions with appropriate HTTP methods
- **MySQL Database Integration**: Persistent storage using MySQL with GORM ORM
- **Auto-Migration**: Automatic database table creation and schema updates
- **JSON API**: All data exchange in JSON format
- **Modular Architecture**: Clean separation of concerns with MVC pattern
- **Error Handling**: Comprehensive error handling and HTTP status codes

## Project Structure

```
4-Book Management System/
├── cmd/main/
│   └── main.go              # Application entry point
├── pkg/
│   ├── config/
│   │   └── app.go          # Database configuration
│   ├── controllers/
│   │   └── book-controller.go  # HTTP request handlers
│   ├── models/
│   │   └── book.go         # Book model and database operations
│   ├── routes/
│   │   └── bookstore-routes.go  # Route definitions
│   └── utils/
│       └── utils.go        # Utility functions
├── go.mod                  # Go module definition
├── go.sum                  # Dependency checksums
└── README.md              # This file
```

## Prerequisites

- Go 1.24.5 or later
- MySQL 5.7 or later
- Git (for cloning the repository)

## Dependencies

- **GORM** (`github.com/jinzhu/gorm v1.9.16`): ORM library for Go
- **Gorilla Mux** (`github.com/gorilla/mux v1.8.1`): HTTP router and URL matcher
- **MySQL Driver** (`github.com/go-sql-driver/mysql v1.5.0`): MySQL driver for Go

## Installation & Setup

### 1. Clone the Repository
```bash
git clone <repository-url>
cd "4-Book Management System"
```

### 2. Install Dependencies
```bash
go mod tidy
```

### 3. Setup MySQL Database

#### Install MySQL (if not already installed)
```bash
# On macOS using Homebrew
brew install mysql

# Start MySQL service
brew services start mysql
```

#### Create Database and Setup User
```bash
# Connect to MySQL
mysql -u root

# Create database
CREATE DATABASE simplerest;

# Set root password (if needed)
ALTER USER 'root'@'localhost' IDENTIFIED BY 'root123';
FLUSH PRIVILEGES;

# Exit MySQL
exit
```

### 4. Configure Database Connection

The database configuration is in `pkg/config/app.go`:
```go
"root:root123@tcp(127.0.0.1:3306)/simplerest?charset=utf8&parseTime=true&loc=Local"
```

Update the connection string if your MySQL setup differs:
- **Username**: `root`
- **Password**: `root123`
- **Host**: `127.0.0.1`
- **Port**: `3306`
- **Database**: `simplerest`

### 5. Run the Application
```bash
go run cmd/main/main.go
```

The server will start on `localhost:9010`.

## API Endpoints

### Base URL
```
http://localhost:9010
```

### Available Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/book/` | Retrieve all books |
| GET | `/book/{bookId}` | Retrieve a specific book by ID |
| POST | `/book/` | Create a new book |
| PUT | `/book/{bookId}` | Update an existing book by ID |
| DELETE | `/book/{bookId}` | Delete a book by ID |

## Data Structure

### Book Object
```json
{
  "ID": 1,
  "CreatedAt": "2025-01-19T19:30:00Z",
  "UpdatedAt": "2025-01-19T19:30:00Z",
  "DeletedAt": null,
  "name": "The Go Programming Language",
  "author": "Alan Donovan",
  "publication": "Addison-Wesley"
}
```

### Fields Description
- **ID** (uint): Auto-generated unique identifier
- **CreatedAt** (time): Timestamp when the book was created
- **UpdatedAt** (time): Timestamp when the book was last updated
- **DeletedAt** (time): Soft delete timestamp (GORM feature)
- **name** (string): Book title
- **author** (string): Book author
- **publication** (string): Publisher name

## Usage Examples

### 1. Get All Books
```bash
curl -X GET http://localhost:9010/book/
```

**Response:**
```json
[
  {
    "ID": 1,
    "CreatedAt": "2025-01-19T19:30:00Z",
    "UpdatedAt": "2025-01-19T19:30:00Z",
    "DeletedAt": null,
    "name": "The Go Programming Language",
    "author": "Alan Donovan",
    "publication": "Addison-Wesley"
  }
]
```

### 2. Get Single Book
```bash
curl -X GET http://localhost:9010/book/1
```

### 3. Create New Book
```bash
curl -X POST http://localhost:9010/book/ \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Clean Code",
    "author": "Robert C. Martin",
    "publication": "Prentice Hall"
  }'
```

### 4. Update Book
```bash
curl -X PUT http://localhost:9010/book/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Updated Book Title",
    "author": "Updated Author",
    "publication": "Updated Publisher"
  }'
```

### 5. Delete Book
```bash
curl -X DELETE http://localhost:9010/book/1
```

## Architecture Overview

### Model-View-Controller (MVC) Pattern

#### Models (`pkg/models/book.go`)
- Defines the `Book` struct with GORM annotations
- Contains database operations (CRUD methods)
- Handles database connections and auto-migration

#### Controllers (`pkg/controllers/book-controller.go`)
- HTTP request handlers
- Business logic processing
- JSON serialization/deserialization
- HTTP response formatting

#### Routes (`pkg/routes/bookstore-routes.go`)
- URL route definitions
- HTTP method mappings
- Controller function bindings

#### Configuration (`pkg/config/app.go`)
- Database connection management
- Connection pooling
- Environment-specific configurations

#### Utilities (`pkg/utils/utils.go`)
- Common helper functions
- Request body parsing
- Reusable utility methods

## Database Schema

The application automatically creates the following table structure:

```sql
CREATE TABLE books (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  created_at datetime(3) DEFAULT NULL,
  updated_at datetime(3) DEFAULT NULL,
  deleted_at datetime(3) DEFAULT NULL,
  name longtext,
  author longtext,
  publication longtext,
  PRIMARY KEY (id),
  KEY idx_books_deleted_at (deleted_at)
);
```

## Error Handling

The API returns appropriate HTTP status codes:

- **200 OK**: Successful GET, PUT operations
- **201 Created**: Successful POST operations
- **400 Bad Request**: Invalid request format
- **404 Not Found**: Resource not found
- **500 Internal Server Error**: Server-side errors

## Testing the API

### Using curl (Command Line)
All examples above use curl for testing endpoints.

### Using Postman
1. Import endpoints into Postman
2. Set base URL to `http://localhost:9010`
3. Test each endpoint with appropriate methods and JSON payloads

### Sample Test Data
```json
{
  "name": "Design Patterns",
  "author": "Gang of Four",
  "publication": "Addison-Wesley"
}
```

## Development

### Adding New Features

1. **Add new fields** to the Book model in `pkg/models/book.go`
2. **Create new endpoints** by adding handlers in `pkg/controllers/`
3. **Define routes** in `pkg/routes/bookstore-routes.go`
4. **Add validation** in controller methods
5. **Update database schema** (GORM auto-migration handles this)

### Code Style Guidelines

- Follow Go naming conventions
- Use meaningful variable and function names
- Add comments for exported functions
- Handle errors appropriately
- Use GORM best practices for database operations

## Troubleshooting

### Common Issues

1. **Database Connection Failed**
   - Check if MySQL is running: `brew services list | grep mysql`
   - Verify database exists: `mysql -u root -p -e "SHOW DATABASES;"`
   - Check connection credentials in `pkg/config/app.go`

2. **Port Already in Use**
   - Change port in `cmd/main/main.go` if 9010 is occupied
   - Check running processes: `lsof -i :9010`

3. **Module Dependencies**
   - Run `go mod tidy` to resolve dependencies
   - Update Go modules: `go get -u ./...`

4. **GORM Version Compatibility**
   - This project uses GORM v1, which is different from GORM v2
   - Refer to GORM v1 documentation for specific usage

## Performance Considerations

- **Connection Pooling**: GORM handles database connection pooling automatically
- **Indexing**: Add database indexes for frequently queried fields
- **Pagination**: Implement pagination for large datasets
- **Caching**: Consider adding Redis for frequently accessed data

## Security Considerations

- **Input Validation**: Add comprehensive input validation
- **SQL Injection**: GORM provides protection, but validate inputs
- **Authentication**: Implement JWT-based authentication
- **Authorization**: Add role-based access control
- **HTTPS**: Use TLS in production environments

## Deployment

### Production Checklist

1. **Environment Variables**: Use environment variables for sensitive config
2. **Database Migration**: Plan database schema migrations
3. **Logging**: Implement structured logging
4. **Health Checks**: Add health check endpoints
5. **Monitoring**: Set up application monitoring
6. **Docker**: Create Dockerfile for containerization

### Docker Example
```dockerfile
FROM golang:1.24.5-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main cmd/main/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 9010
CMD ["./main"]
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is open source and available under the [MIT License](LICENSE).

## Support

For support and questions:
- Create an issue in the repository
- Check existing documentation
- Review Go and GORM documentation for technical details
