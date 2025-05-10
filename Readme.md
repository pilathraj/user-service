# User Management Service
This project is a user management service for a library system. It provides APIs for user registration, authentication, profile management, role-based access control, and user preferences management.

---

## Features
- Role-based access control
- User profile management
- User preferences management
- PostgreSQL database integration

---

## Prerequisites
- Docker
- Minikube (for Kubernetes deployment)
- Go (for running tests locally)

---

## Setup Instructions

### 1. Build the Docker Image
```bash
make build
```

### 2. Run the Application Locally
```bash
make run
```

### 3. Stop the Application
```bash
make stop
```

### 4. Deploy to Minikube
```bash
make minikube
```

### Stop Minikube Deployment
```bash 
make minikube-stop
```

### Run Tests
```bash
make test
```

## API Endpoints

### 1. Register a User
- URL: /api/users
```json
{
    "username": "pilathraj",
    "password": "admin123"
}
```
- Response: User registered successfully.

### 2. List All Users 
- Method: GET
- Headers:
  - Authorization: Bearer <JWT_TOKEN>
- Response: List of all users.

### 3. Get User Details
- Method: GET
- URL: /api/users/{id}
- Headers:
  - Authorization: Bearer <JWT_TOKEN>
- Response: User details.

### 4. Update User
- Method: PUT
- URL: /api/users/{id}
- Headers:
  - Authorization: Bearer <JWT_TOKEN>
- Request Body:
```json
{
    "username": "admin",
    "email": "pilathraja@gmail.com",
    "password": "admin123"
}
```
- Response: User updated successfully.

### 5. Delete User
- URL: /api/users/{id}
- Headers:
  - Authorization: Bearer <JWT_TOKEN>
- Response: User deleted successfully.

### 6. Get User Preferences
- Method: GET
- Headers:
  - Authorization: Bearer <JWT_TOKEN>
- Response: User preferences.

### 7. Update User Preferences
- URL: /api/users/{id}/preferences
- Headers:
  - Authorization: Bearer <JWT_TOKEN>
- Request Body:
```json
[
    {
        "value": "on"
    },
    {
        "value": "off"
    },
    {
        "key": "Theme",
        "value": "dark"
    }
]
```
