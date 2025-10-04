# Visa Tracker System

## 📌 Overview

The **Visa Tracker System** is a backend service developed to automate the monitoring of visa expiration dates for foreign citizens entering Kazakhstan. This system replaces manual tracking processes, reduces human error, and enhances data accuracy for migration officers by providing a reliable, automated solution built with modern technologies.

---

## 🛠️ Technology Stack

| Component        | Technology         |
|------------------|--------------------|
| Backend Framework| Go (Fiber)         |
| ORM              | GORM               |
| Database         | PostgreSQL         |
| Containerization | Docker             |
| Orchestration    | Docker Compose     |
| Reverse Proxy    | Nginx              |

---

## 🏗️ System Architecture

The system follows a microservice-inspired containerized architecture:

- **Backend Service**: REST API built with Go Fiber.
- **Database Service**: PostgreSQL managed via Docker with persistent volume.
- **Reverse Proxy**: Nginx routes external requests to the backend.

All services are orchestrated using **Docker Compose** for easy local and production deployment.

---

## 🧩 Core Functionality

### Database Model (`Migrant`)
```go
type Migrant struct {
    ID           uint
    FullName     string
    Passport     string
    Nationality  string
    EntryDate    time.Time
    StayType     string
    VisaExpiry   *time.Time
    AllowedDays  *int
    CreatedAt    time.Time
    UpdatedAt    time.Time
}
```

### API Endpoints

| Endpoint               | Method | Description                     |
|------------------------|--------|---------------------------------|
| `/api/migrants`        | GET    | Retrieve all migrant records    |
| `/api/expired`         | GET    | Get migrants with expired stays |
| `/api/migrants`        | POST   | Add a new migrant               |
| `/api/migrants/:id`    | PUT    | Update migrant record           |
| `/api/migrants/:id`    | DELETE | Delete migrant record           |

---

## 🐳 DevOps & Deployment

### Docker Setup
- Two main services:
  - `db`: PostgreSQL container with persistent volume.
  - `backend`: Go Fiber API container.
- Environment variables are managed via `.env`.

### Deployment Workflow
1. Clone the repository.
2. Create a `.env` file based on `.env.example`.
3. Run:
   ```bash
   docker-compose up --build
   ```
4. The API will be accessible via `http://localhost` (proxied through Nginx).
