# Leave Management System – Backend

This repository contains the **backend service** for the ABC Company Leave Management System.

The backend is built using **Golang** and provides REST APIs to manage employees, leave requests, approvals, and dashboard data.

---

# Technologies Used

- Golang
- PostgreSQL
- REST API
- godotenv (for environment variables)

---

# Project Structure

```
leave-management-system-backend
│
├ main.go
├ go.mod
├ README.md
├ .env
│
├ handlers
│   ├ auth_handler.go
│   ├ dashboard_handler.go
│   ├ employee_handler.go
│   └ leave_handler.go
│
├ models
│   ├ auth.go
│   ├ dashboard.go
│   ├ employee.go
│   └ leave.go
│
├ routes
│   └ routes.go
│
└ database
    ├ database.go
    └ schema.sql
```

---

# Prerequisites

Make sure the following software is installed:

- Go (version 1.20 or higher)
- PostgreSQL
- Git

Check installations:

```bash
go version
psql --version
```

---

# Step 1 – Clone the Repository

```bash
git clone https://github.com/MirominaS/leave-management-system-backend
cd leave-management-system-backend
```

---

# Step 2 – Install Dependencies

Install Go modules required for the project.

```bash
go mod tidy
```

---

# Step 3 – Create PostgreSQL Database

Open PostgreSQL terminal:

```bash
psql -U postgres
```

Create the database:

```sql
CREATE DATABASE leave_management;
```

Exit PostgreSQL:

```
\q
```

---

# Database Configuration

This project connects to PostgreSQL using environment variables stored in a `.env` file.

The backend loads these values using the **godotenv** package.

---

## Step 4 – Create `.env` File

Inside the root backend folder create a file named:

```
.env
```

Example project structure:

```
leave-management-system-backend
│
├ main.go
├ go.mod
├ .env
│
├ handlers
├ models
└ database
```

---

## Step 5 – Add Database Credentials

Add the following values to the `.env` file:

```
PORT = portnumber - replace your port here

DB_HOST=localhost
DB_PORT=5432
DB_USERNAME=postgres
DB_PASSWORD=yourpassword
DB_NAME=leave_management
```

Replace `yourpassword` with your PostgreSQL password.

---

# Step 6 – Create Database Tables

Run the SQL schema file included in the project:

```bash
psql -U postgres -d leave_management -f database/schema.sql
```

This command will create the required tables:

- roles
- employees
- leave_types
- leave_status
- leave_request

---

# Step 7 – Run the Backend Server

Start the backend server using:

```bash
go run main.go
```

The server will start at:

```
http://localhost:3300
```

If the connection is successful, you will see:

```
Successfully connected to PostgreSQL!
```

---

# Example API Endpoints

The backend exposes the following REST API endpoints.

---

## Authentication

Login to the system.

```
POST /login
```

---

## Employee Management

Get all employees

```
GET /employees
```

Create a new employee

```
POST /employees/create
```

Get employee profile

```
GET /employees/profile
```

---

## Leave Management

Get all leave requests

```
GET /leave
```

Create a leave request

```
POST /leave/create
```

Cancel a leave request

```
POST /leave/cancel
```

Approve a leave request (Admin/Manager only)

```
POST /leave/approve
```

Reject a leave request (Admin/Manager only)

```
POST /leave/reject
```

Get available leave types

```
GET /leave-types
```

---

## Dashboard

Get dashboard statistics

```
GET /dashboard
```

Get recent leave activity

```
GET /activities
```
| Method | Endpoint           | Description          |
|--------|--------------------|----------------------|
| POST   | /login             | User login           |
| GET    | /employees         | Get all employees    |
| POST   | /employees/create  | Create employee      |
| GET    | /leave             | Get leave requests   |
| POST   | /leave/create      | Create leave request |
| POST   | /leave/approve     | Approve leave        |
| POST   | /leave/reject      | Reject leave         |
| GET    | /dashboard         | Get dashboard data   |
---

# Notes

- Make sure PostgreSQL service is running before starting the backend.
- Start the backend server before running the frontend application.
- The backend runs on port given by you in the env file or it takes **8080**  by default.

---

# Author

Software Engineer Intern Technical Assessment  
Miromina Sritharan
