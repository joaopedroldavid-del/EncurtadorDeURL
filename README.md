# 🔗 EncurtadorDeURL

A simple and efficient **URL shortener** built with **Go (Golang)** using the [Chi Router](https://github.com/go-chi/chi).  
Now powered by **Redis** for fast and persistent storage.

---

## 🚀 Features

- 🔒 Accepts and validates URL inputs
- ✂️ Generates unique short codes
- 🔁 Redirects users to the original URL
- 💾 Persists data using Redis
- 📦 Clean JSON API responses
- 🧹 Refactored for better structure and maintainability

---

## 🛠 Tech Stack

- [Go](https://golang.org/)
- [Chi Router](https://github.com/go-chi/chi)
- [Redis](https://redis.io/) (via [go-redis](https://github.com/redis/go-redis))
- `net/http`, `encoding/json`
- `log/slog` for structured logging

---

## ⚙️ Getting Started

### 1. Clone the repository:
```bash
git clone https://github.com/joaopedroldavid-del/EncurtadorDeURL.git
cd EncurtadorDeURL
```
### 2. Start Redis

Make sure you have a Redis server running locally (default port 6379).

### 3. Run the application:
```bash
go run main.go
```

---

## 📤 Example Requests

### POST /api/shortn

**Request:**
```bash
curl -X POST http://localhost:8080/api/shortn \
-H "Content-Type: application/json" \
-d '{"url":"https://example.com"}'
```
**Response:**
```json
{
  "data": "a8B2xZ"
}
```

### GET /{code}

Redirects to the original long URL stored under the given code.

---