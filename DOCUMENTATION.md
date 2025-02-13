**DOCUMENTATION.md**  
# eLibrary System Documentation

## Table of Contents
1. [Project Overview](#project-overview)
2. [System Architecture](#system-architecture)
3. [Core Components](#core-components)
4. [Database Schema](#database-schema)
5. [API Reference](#api-reference)
6. [External Services](#external-services)
7. [Frontend Structure](#frontend-structure)
8. [Environment Setup](#environment-setup)
9. [Development Guide](#development-guide)

---

## Project Overview
A web-based library management system with:
- Book catalog management
- User subscription/notification system
- Book borrowing/returning workflows
- Integration with **Google Books API** (ISBN validation)
- Email notifications via **SendGrid**
- PostgreSQL database

## System Architecture 

├── cmd/
│ └── server/
│ └── main.go # Application entry point
├── internal/
│ ├── config/ # Configuration loader
│ ├── handlers/ # HTTP handlers
│ ├── models/ # Data models
│ ├── repository/ # Database operations
│ └── services/ # External services
├── migrations/ # Database schema
├── web/ # Frontend assets
│ ├── static/ # CSS/JS files
│ └── templates/ # HTML templates
└── pkg/utils/ # Utility functions

## Core Components

### 1. Handlers
| Handler               | Endpoints                                 | Description                     |
|-----------------------|-------------------------------------------|---------------------------------|
| `BookHandler`         | `GET /api/books/{id}`, `POST /api/borrow` | Manage book operations          |
| `UserHandler`         | `POST /api/subscribe` (WIP: Unsubscribe)  | Handle user subscriptions       |
| `NotificationHandler` | (Manual trigger endpoint planned)         | Email notifications             |

### 2. Models
```go
// Book Model
type Book struct {
    ID        int       // Unique identifier
    Title     string    // Book title
    ISBN      string    // Validated ISBN format
    Available bool      // Borrowing status
}

// User Model
type User struct {
    Email      string    // Unique email
    Subscribed bool      // Subscription status
}

3. Repositories
Repository	      Key Methods                       	Description
BookRepository	  GetAllBooks(), GetBookByID()	      Database operations for books
UserRepository	  FindUserByEmail(), CreateUser()	    User management

4. Services
Service	Functionality
GoogleBooksService	Fetch book metadata using ISBN
NotificationService	Send emails via SendGrid API

## Database Schema

-- Books Table
CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    isbn VARCHAR(20) UNIQUE NOT NULL,
    available BOOLEAN DEFAULT TRUE
);

-- Users Table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    subscribed BOOLEAN DEFAULT FALSE
);

## API Reference

1. Get Book by ID
GET /api/books/{id}

Response:
{
  "id": 42,
  "title": "The Go Programming Language",
  "isbn": "978-0134190440",
  "available": true
}

2. Borrow Book
POST /api/borrow

Request Body:
{
  "user_id": 1,
  "book_id": 42
}

## External Services
Google Books Integration
  - Validates ISBNs and fetches metadata
  - Configure via GOOGLE_BOOKS_API_KEY in .env

SendGrid Email Service
  - Sends notifications when books become available
  - Requires SENDGRID_API_KEY in .env

## Frontend Structure
Key Files:
web/
├── static/
│   ├── app.js      # Dynamic UI interactions
│   └── style.css   # Responsive styling
└── templates/
    └── index.html  # Book listing template

Features:
  - Real-time book availability display
  - Subscription form with AJAX
  - Responsive grid layout for books

## Environment Setup
Database: Run migrations:
psql -U postgres -d elibrary -f migrations/001_initial_schema.sql

Environment Variables (.env):
DATABASE_URL="postgres://user:pass@localhost:5432/elibrary"
SENDGRID_API_KEY="your_sendgrid_key"
GOOGLE_BOOKS_API_KEY="your_google_key"

## Development Guide
Running the Server
go run cmd/server/main.go

Testing
# Run all tests
go test ./...

Contribution
  - Create feature branches (git checkout -b feature/xyz)
  - Follow Go style guidelines
  - Update documentation for new features
