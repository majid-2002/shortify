
# Shortify

Shortify is a simple URL shortening service built with Go and Gin, leveraging GORM for database interactions. This project allows users to shorten URLs, redirect to the original URL using the shortened version, update or delete shortened URLs, and view the statistics of how many times the shortened URL has been accessed.

## Features

- **Shorten URLs**: Users can send a POST request to shorten a URL.
- **Redirect**: A GET request to a shortened URL redirects to the original URL.
- **Update**: Users can update the original URL for a shortened URL.
- **Delete**: Users can delete a shortened URL.
- **Stats**: Get access statistics for a shortened URL.

## Requirements

- Go (1.21 or later)
- PostgreSQL
- GORM (used for database interactions)
- Gin (used for building the web framework)

## Setup

### 1. Clone the repository:

```bash
git clone https://github.com/majid-2002/shortify.git
cd shortify
```

### 2. Install dependencies:

```bash
go mod tidy
```

### 3. Set up PostgreSQL:

Make sure you have a PostgreSQL database running. You can use the following environment variables or modify the connection string in `database/db.go`:

```bash
host=localhost user=<username> password=<password> dbname=shortify_db port=5432 sslmode=disable
```

Run the following SQL commands to create the necessary table:

```sql
CREATE DATABASE shortify_db;

-- Adjust the following SQL according to your database setup
-- This command is auto-migrated by the application as well
```

### 4. Run the application:

```bash
go run main.go
```

By default, the server will run on `http://localhost:8080`.

## API Endpoints

### 1. **POST /shorten**

Shorten a URL by sending a `POST` request with a JSON body:

```json
{
  "original": "http://example.com"
}
```

Response:

```json
{
  "original": "http://example.com",
  "shortened": "181cff0647b574a3"
}
```

### 2. **GET /{shortened}**

Redirects to the original URL. For example, to visit the URL shortened as `181cff0647b574a3`, you can make a `GET` request to:

```
http://localhost:8080/181cff0647b574a3
```

### 3. **PUT /{shortened}**

Update the original URL of a shortened URL by sending a `PUT` request with a JSON body:

```json
{
  "original": "http://newexample.com"
}
```

### 4. **DELETE /{shortened}**

Delete a shortened URL:

```bash
DELETE /181cff0647b574a3
```

### 5. **GET /{shortened}/stats**

Get the access statistics for a shortened URL:

```bash
GET /181cff0647b574a3/stats
```

Response:

```json
{
  "shortened": "181cff0647b574a3",
  "access_count": 5
}
```
