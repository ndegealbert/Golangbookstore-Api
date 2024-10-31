# Book Store API

A simple RESTful API built with Go, designed for managing a bookstore's collection.

## Features

- **CRUD Operations**: Create, retrieve, update, and delete books.
- **JSON API**: All responses are returned in JSON format for easy frontend integration.
- **Database Support**: Easily configurable for SQLite, MySQL, or PostgreSQL.
  
## Endpoints

- `GET /books` - List all books
- `GET /books/{id}` - Get a book by ID
- `POST /books` - Add a new book
- `PUT /books/{id}` - Update an existing book
- `DELETE /books/{id}` - Delete a book by ID

## Getting Started

### Prerequisites

- Go installed on your machine
- Database setup (optional for development)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/bookstore-api.git
   cd bookstore-api
