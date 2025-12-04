# Go URL Shortener

A simple, fast, and efficient URL shortener built with Go (Golang), Fiber, and Redis. This project allows you to shorten long URLs, set custom short codes, and define expiration times. It is fully containerized using Docker.

## Features

- **Shorten URLs**: Convert long URLs into short, manageable links.
- **Custom Short Codes**: Specify your own custom short URL (optional).
- **Expiry**: Set an expiration time for the shortened URL.
- **Rate Limiting**: Built-in rate limiting to prevent abuse (default: quota per 30 minutes).
- **Fast Redirection**: Uses Redis for high-performance URL resolution.
- **Dockerized**: Easy setup and deployment with Docker Compose.

## Tech Stack

- **Language**: Go (Golang)
- **Framework**: [Fiber](https://gofiber.io/)
- **Database**: Redis
- **Containerization**: Docker & Docker Compose

## Prerequisites

- [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/) installed on your machine.
- (Optional) [Go](https://golang.org/) installed if you want to run it locally without Docker.

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/alfredzimmer/url-shortener
cd url-shortener
```

### 2. Environment Setup

Ensure you have a `.env` file in the `api` directory or root (depending on how you run it, the code looks for it). Based on the code, it loads `.env`.

Example `.env` variables (you might need to create this):
```env
APP_PORT=:3000
DB_ADDR=db:6379
DB_PASSWORD=
API_QUOTA=10
DOMAIN=localhost:3000
```

### 3. Run with Docker Compose


```bash
docker-compose up -d
```

This will start two containers:
- `api`: The Go application running on port `3000`.
- `db`: The Redis database running on port `6379`.

## API Endpoints

### 1. Shorten a URL

**Endpoint**: `POST /api/v1`

**Request Body**:
```json
{
  "url": "https://www.google.com",
  "short": "goog", 
  "expiry": 24
}
```
- `url`: The original long URL (required).
- `short`: Custom short code (optional). If not provided, a random UUID-based short code is generated.
- `expiry`: Expiration time in hours (optional).

**Response**:
```json
{
  "url": "https://www.google.com",
  "short": "http://localhost:3000/goog",
  "expiry": 24,
  "rate_limit": 9,
  "rate_limit_reset": 30
}
```

### 2. Resolve a URL

**Endpoint**: `GET /:url`

Redirects the user to the original URL associated with the short code.

**Example**:
Open `http://localhost:3000/goog` in your browser.

### 3. Reset Rate Limit (Dev/Admin)

**Endpoint**: `POST /api/resolve`

Resets the rate limit for the requesting IP address.

**Response**:
```json
{
  "message": "Rate limit reset successfully",
  "ip": "..."
}
```
