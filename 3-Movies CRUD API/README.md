# Movies CRUD API

A RESTful API built with Go for managing a movie database. This API provides full CRUD (Create, Read, Update, Delete) operations for movies with director information using the Gorilla Mux router.

## Features

- **Complete CRUD Operations**: Create, Read, Update, and Delete movies
- **RESTful Design**: Follows REST API conventions with appropriate HTTP methods
- **JSON Responses**: All data exchanged in JSON format
- **In-Memory Storage**: Uses Go slices for fast, lightweight data storage
- **Structured Data**: Movies include director information with nested data structures
- **Random ID Generation**: Automatically generates unique IDs for new movies
- **Gorilla Mux Router**: Professional-grade HTTP router for clean URL patterns

## Project Structure

```
3-Movies CRUD API/
├── main.go              # Main application with API endpoints
├── go.mod              # Go module with dependencies
├── go.sum              # Dependency checksums
├── movies-crud-api     # Compiled binary (if built)
└── README.md           # This file
```

## Prerequisites

- Go 1.24.5 or later
- Internet connection (for downloading dependencies)

## Dependencies

- **Gorilla Mux** (`github.com/gorilla/mux v1.8.1`): HTTP router and URL matcher

## Installation & Setup

1. Clone or download this repository
2. Navigate to the project directory:
   ```bash
   cd "3-Movies CRUD API"
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Run the application:
   ```bash
   go run main.go
   ```

Alternatively, build and run the binary:
```bash
go build -o movies-crud-api
./movies-crud-api
```

## API Endpoints

The server runs on port 8000. Once started, you'll see:
```
Starting server at port 8000...
```

### Available Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/movies` | Retrieve all movies |
| GET | `/movies/{id}` | Retrieve a specific movie by ID |
| POST | `/movies` | Create a new movie |
| PUT | `/movies/{id}` | Update an existing movie by ID |
| DELETE | `/movies/{id}` | Delete a movie by ID |

## Data Structure

### Movie Object
```json
{
  "id": "1",
  "isbn": "438734",
  "title": "Movie One",
  "director": {
    "firstname": "John",
    "lastname": "Doe"
  }
}
```

### Fields Description
- **id** (string): Unique identifier for the movie
- **isbn** (string): International Standard Book Number
- **title** (string): Movie title
- **director** (object): Director information
  - **firstname** (string): Director's first name
  - **lastname** (string): Director's last name

## Usage Examples

### 1. Get All Movies
```bash
curl -X GET http://localhost:8000/movies
```

**Response:**
```json
[
  {
    "id": "1",
    "isbn": "438734",
    "title": "Movie One",
    "director": {
      "firstname": "John",
      "lastname": "Doe"
    }
  },
  {
    "id": "2",
    "isbn": "438735",
    "title": "Movie Two",
    "director": {
      "firstname": "Jane",
      "lastname": "Smith"
    }
  }
]
```

### 2. Get Single Movie
```bash
curl -X GET http://localhost:8000/movies/1
```

### 3. Create New Movie
```bash
curl -X POST http://localhost:8000/movies \
  -H "Content-Type: application/json" \
  -d '{
    "isbn": "123456",
    "title": "New Movie",
    "director": {
      "firstname": "Steven",
      "lastname": "Spielberg"
    }
  }'
```

### 4. Update Movie
```bash
curl -X PUT http://localhost:8000/movies/1 \
  -H "Content-Type: application/json" \
  -d '{
    "isbn": "654321",
    "title": "Updated Movie",
    "director": {
      "firstname": "Christopher",
      "lastname": "Nolan"
    }
  }'
```

### 5. Delete Movie
```bash
curl -X DELETE http://localhost:8000/movies/1
```

## Sample Data

The API comes pre-loaded with sample movies:

1. **Movie One**
   - ID: 1
   - ISBN: 438734
   - Director: John Doe

2. **Movie Two**
   - ID: 2
   - ISBN: 438735
   - Director: Jane Smith

## Technical Implementation

### Key Components

- **`Movie` Struct**: Defines the movie data structure with JSON tags
- **`Director` Struct**: Nested structure for director information
- **`getMovies()`**: Returns all movies in JSON format
- **`getMovie()`**: Returns a specific movie by ID
- **`createMovie()`**: Adds a new movie with auto-generated ID
- **`updateMovie()`**: Updates an existing movie by ID
- **`deleteMovie()`**: Removes a movie from the collection

### Router Configuration

Uses Gorilla Mux for clean URL routing:
```go
r.HandleFunc("/movies", getMovies).Methods("GET")
r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
r.HandleFunc("/movies", createMovie).Methods("POST")
r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
```

### Data Storage

- **In-Memory Storage**: Uses Go slices for fast access
- **No Persistence**: Data is lost when the server restarts
- **Thread-Safe Operations**: Basic CRUD operations on slice data structure

## Testing the API

### Using curl (Command Line)
All the examples above use curl commands for testing.

### Using Postman
1. Import the endpoints into Postman
2. Set the base URL to `http://localhost:8000`
3. Test each endpoint with appropriate HTTP methods and JSON payloads

### Using a REST Client Extension
Many code editors have REST client extensions that can test these endpoints directly.

## Development

### Adding New Features

To extend this API:

1. **Add new fields** to the `Movie` or `Director` structs
2. **Implement new endpoints** by adding handler functions
3. **Add validation** for input data
4. **Implement database storage** to replace in-memory storage
5. **Add authentication/authorization** for secure access

### Error Handling

Currently implements basic error handling:
- Returns appropriate HTTP status codes
- Handles JSON encoding/decoding errors
- Manages resource not found scenarios

## Potential Enhancements

- **Database Integration**: Replace in-memory storage with PostgreSQL/MongoDB
- **Input Validation**: Add comprehensive request validation
- **Error Responses**: Standardized error response format
- **Logging**: Add structured logging with different levels
- **Authentication**: JWT-based authentication system
- **Pagination**: Add pagination for large movie collections
- **Search/Filter**: Add search and filtering capabilities
- **Unit Tests**: Comprehensive test suite
- **Docker Support**: Containerization for easy deployment

## Contributing

Feel free to submit issues and enhancement requests!

## License

This project is open source and available under the [MIT License](LICENSE).
