# 🔗 EncurtadorDeURL

A simple and efficient **URL shortener** built with **Go (Golang)** using the [Chi Router](https://github.com/go-chi/chi).  
This service receives a long URL and returns a shortened code for redirection.

---

## 🚀 Features

- 🔒 Accepts and validates URL inputs
- ✂️ Generates unique short codes
- 🔁 Redirects users to the original URL
- 📦 Clean JSON API responses

---

## 🛠 Tech Stack

- [Go](https://golang.org/)
- [Chi Router](https://github.com/go-chi/chi)
- `net/http`, `encoding/json`, `math/rand`
- `log/slog` for structured logging

---

## ⚙️ Getting Started

### 1. Clone the repository:
```bash
git clone https://github.com/joaopedroldavid-del/EncurtadorDeURL.git
cd EncurtadorDeURL
```
### 2. Run the application:
```bash
go run main.go
```

## 📤 Example Requests

### Request:
```bash
curl -X POST http://localhost:8080/api/shortn \
-H "Content-Type: application/json" \
-d '{"url":"https://example.com"}'
```
### Response:
```bash
{
  "data": "a8B2xZ"
}
```
### GET /{code}

Redirects to the original long URL stored under the given code.