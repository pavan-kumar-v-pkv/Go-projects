# Booking App

A command-line conference ticket booking application built with Go. This application allows users to book tickets for the "Go Conference" with real-time ticket availability tracking and email confirmation simulation.

## Features

- **Interactive Booking System**: Command-line interface for ticket booking
- **Input Validation**: Validates user name, email format, and ticket quantity
- **Real-time Ticket Tracking**: Monitors remaining tickets and prevents overbooking
- **Concurrent Processing**: Uses goroutines for simulated email ticket delivery
- **User Data Management**: Stores booking information with structured data types
- **Booking History**: Displays all current bookings and attendee first names

## Project Structure

```
BOOKING_APP/
├── main.go              # Main application with booking logic
├── go.mod              # Go module definition
├── helper/             # Helper package
│   └── helper.go       # Input validation functions
└── README.md           # This file
```

## Prerequisites

- Go 1.24.5 or later
- No external dependencies required (uses Go standard library only)

## Installation & Setup

1. Clone or download this repository
2. Navigate to the project directory:
   ```bash
   cd BOOKING_APP
   ```
3. Run the application:
   ```bash
   go run main.go
   ```

## Usage

When you run the application, you'll be greeted with:

```
Welcome to Go Conference booking application!
We have a total of 50 tickets and 50 are available.
Get your tickets here to attend.
```

### Booking Process

The application will prompt you for the following information:

1. **First Name**: Must be at least 2 characters long
2. **Last Name**: Must be at least 2 characters long  
3. **Email Address**: Must contain both '@' and '.' symbols
4. **Number of Tickets**: Must be between 1 and the number of remaining tickets

### Example Interaction

```
Enter your name: John
Enter your last name: Doe
Enter your email: john.doe@example.com
Enter number of tickets: 3
```

### Successful Booking

Upon successful booking, you'll see:
- Confirmation message with booking details
- Updated remaining ticket count
- List of all current bookings
- Simulated email ticket delivery (after 10 seconds)

### Input Validation

The application validates all inputs and provides specific error messages:

- **Invalid Name**: "Your first name or last name entered is too short. They each must be at least 2 characters long."
- **Invalid Email**: "Your email address must contain an '@' and a '.'"
- **Invalid Ticket Count**: "You can book a maximum of X tickets. You tried to book Y tickets."

## Technical Features

### Concurrency

The application demonstrates Go's concurrency features:
- Uses **goroutines** for ticket delivery simulation
- Implements **sync.WaitGroup** to ensure proper program termination

### Data Structures

- **Struct-based User Data**: Stores booking information in a structured format
- **Dynamic Slices**: Manages bookings with flexible data storage
- **Type Safety**: Uses specific data types (`uint` for tickets, structured data for users)

### Validation Package

The `helper` package provides:
- **Name Validation**: Ensures names are at least 2 characters long
- **Email Validation**: Checks for basic email format requirements
- **Ticket Validation**: Verifies ticket quantity is valid and available

## Code Overview

### Main Components

- **`main()`**: Main application loop and orchestration
- **`greetUsers()`**: Displays welcome message and ticket availability
- **`getUserInput()`**: Handles user input collection
- **`bookTicket()`**: Processes booking and updates ticket count
- **`sendTicket()`**: Simulates email delivery with goroutine
- **`getFirstNames()`**: Extracts first names from all bookings
- **`helper.ValidateUserInput()`**: Validates all user inputs

### Configuration

```go
const conferenceTickets = 50      // Total available tickets
var conferenceName = "Go Conference"  // Conference name
var remainingTickets uint = 50    // Current available tickets
```

## Development

### Adding New Features

To extend this application:

1. **Add new validation rules** in `helper/helper.go`
2. **Modify user data structure** in the `UserData` struct
3. **Add new booking fields** by updating the input collection functions
4. **Enhance ticket delivery** by modifying the `sendTicket()` function

### Testing

Run the application and test various scenarios:
- Valid bookings
- Invalid names (less than 2 characters)
- Invalid email formats
- Booking more tickets than available
- Booking all remaining tickets

## Learning Concepts

This project demonstrates several Go programming concepts:

- **Package Management**: Custom packages and imports
- **Data Types**: Structs, slices, maps, and type declarations
- **Input/Output**: Console input/output operations
- **Concurrency**: Goroutines and WaitGroups
- **Validation**: Input validation and error handling
- **String Manipulation**: Working with user input and formatting

## Contributing

Feel free to submit issues and enhancement requests!

## License

This project is open source and available under the [MIT License](LICENSE).
