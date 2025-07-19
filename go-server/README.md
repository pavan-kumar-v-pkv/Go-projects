# Go Web Server

A simple HTTP web server built with Go that serves static files and handles form submissions.

## Features

- **Static File Serving**: Serves HTML files and other static content from the `./static` directory
- **Form Handler**: Processes POST requests from HTML forms and displays submitted data
- **Hello Endpoint**: Simple GET endpoint that returns a greeting message
- **Error Handling**: Includes proper HTTP error responses for invalid paths and methods

## Project Structure

```
go-server/
├── main.go           # Main server application
├── go.mod           # Go module definition
├── static/          # Static files directory
│   ├── index.html   # Welcome page
│   └── form.html    # Form submission page
└── README.md        # This file
```

## Prerequisites

- Go 1.24.5 or later
- No external dependencies required (uses Go standard library only)

## Installation & Setup

1. Clone or download this repository
2. Navigate to the project directory:
   ```bash
   cd go-server
   ```
3. Run the server:
   ```bash
   go run main.go
   ```

## Usage

Once the server is running, it will start on port 8080. You'll see the message:
```
Starting server at port 8080
```

### Available Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/` | GET | Serves static files (index.html by default) |
| `/form` | POST | Handles form submissions and displays submitted data |
| `/hello` | GET | Returns a simple "Hello!" message |

### Accessing the Application

1. **Home Page**: Visit [http://localhost:8080](http://localhost:8080) to see the welcome page
2. **Form Page**: Visit [http://localhost:8080/form.html](http://localhost:8080/form.html) to access the form
3. **Hello Endpoint**: Visit [http://localhost:8080/hello](http://localhost:8080/hello) for a simple greeting

### Form Submission

The form page (`form.html`) includes:
- Name field (text input)
- Email field (email input) 
- Address field (textarea)

When submitted, the form data is processed by the `/form` endpoint and the submitted values are displayed back to the user.

## Code Overview

### Main Components

- **`formHandler`**: Processes form submissions, parses form data, and returns the submitted values
- **`helloHandler`**: Simple handler that returns "Hello!" for GET requests to `/hello`
- **`main`**: Sets up the server with route handlers and starts listening on port 8080

### Error Handling

- Returns 404 for invalid paths in the hello handler
- Returns 405 (Method Not Allowed) for non-GET requests to `/hello`
- Handles form parsing errors gracefully

## Development

To modify this server:

1. Edit `main.go` to add new endpoints or modify existing handlers
2. Add new static files to the `./static` directory
3. Restart the server to see changes

## Contributing

Feel free to submit issues and enhancement requests!

## License

This project is open source and available under the [MIT License](LICENSE).
