A minimalistic, modern API for tracking habits, built with **Go**, **Gin**, **Ent ORM**, and **PostgreSQL**. Features a simple HTML interface (with htmx support) and auto-generated Swagger docs. Perfect for learning, hacking, or as a foundation for your own productivity tools!

---

## 🚀 Features

- **CRUD** operations for habits
- Clean HTML interface (with htmx for snappy UX)
- RESTful API endpoints
- Auto-generated Swagger UI (`/swagger/index.html`)
- Built with Go, Gin, Ent, PostgreSQL

---

## 🛠️ Quickstart

### 1. Clone & Setup

```bash
git clone https://github.com/yourusername/habit-tracker-api.git
cd habit-tracker-api
```

### 2. Configure PostgreSQL

Make sure you have a local PostgreSQL instance running. Update the connection string in `main.go` if needed:

```bash
docker compose up -d db
```

### 3. Install Dependencies

```bash
go mod tidy
```

### 4. Run the Server

```bash
go run main.go
```

Visit [http://localhost:8080](http://localhost:8080) in your browser.

---

## 📚 API Docs

Swagger UI is auto-generated and available at:

```
http://localhost:8080/swagger/index.html
```

Explore, test, and play with the API right from your browser!

---

## 🧩 Endpoints

| Method | Endpoint           | Description            |
|--------|--------------------|------------------------|
| GET    | `/`                | List all habits        |
| POST   | `/habits`          | Create a new habit     |
| GET    | `/habits/:id`      | Show edit form         |
| PUT    | `/habits/:id`      | Update a habit         |
| DELETE | `/habits/:id`      | Delete a habit         |

---

## 🏗️ Tech Stack

- [Go](https://golang.org/)
- [Gin Web Framework](https://gin-gonic.com/)
- [Ent ORM](https://entgo.io/)
- [PostgreSQL](https://www.postgresql.org/)
- [Swagger (swaggo)](https://github.com/swaggo/swag)
- [htmx](https://htmx.org/) (for partial HTML updates)

---

## 💡 Development Notes

- HTML templates live in `/templates`
- Swagger docs generated with [swag](https://github.com/swaggo/swag) (`swag init`)
- Edit connection string or ORM schema as needed for your setup
- Ran into the following issue with [swagger](https://github.com/swaggo/swag/issues/1622), solved by refactoring to use named handlers


---

## 📜 License

MIT License. See [LICENSE](LICENSE) for details.

---
