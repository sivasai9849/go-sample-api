# 🚀 Go API

🛠️ A RESTful API built with Go, featuring user management capabilities using Gin framework and PostgreSQL database.

## ✨ Features

- 👥 User CRUD operations
- 🗄️ PostgreSQL database integration
- 📝 Structured logging
- ⚙️ Environment-based configuration
- 🛡️ Error handling middleware
- 💓 Health check endpoint
- ✅ Input validation
- 🏗️ Clean architecture pattern

## 📋 Prerequisites

- 🔧 Go 1.23 or higher
- 🐘 PostgreSQL
- 🔄 Air (for hot reload during development)

## 🚦 Getting Started

1. 📥 Clone the repository
   ```bash
   git clone https://github.com/sivasai9849/go-sample-api.git
   ```

2. 🔑 Set up environment variables
   ```bash
   cp .env.example .env
   ```
3. 📦 Install dependencies
   ```bash
   go mod download
   ```
4. 🏃‍♂️ Run the application

    `# Using go run`
    ```bash
    go run cmd/api/main.go
    ```
    `# Using Air for hot reload`
    ```bash
    air
    ```

## 🛣️ API Endpoints

### 👥 Users
- `POST /api/v1/users` - Create a new user
- `GET /api/v1/users` - List all users
- `GET /api/v1/users/:id` - Get user by ID
- `PUT /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user

### 💓 Health Check
- `GET /health` - Check API health status

## 📁 Project Structure
.
├── 📂 cmd/
│   └── 📂 api/
│       └── 📄 main.go
├── 📂 internal/
│   ├── 📂 config/
│   ├── 📂 domain/
│   ├── 📂 dto/
│   ├── 📂 handler/
│   ├── 📂 middleware/
│   ├── 📂 repository/
│   ├── 📂 service/
│   └── 📂 validation/
└── 📂 pkg/
    ├── 📂 errors/
    └── 📂 logger/

## 🛠️ Development

The project uses Air for hot reloading during development. Configuration can be found in `.air.toml`.

## 🤝 Contributing

1. 🔱 Fork the repository
2. 🌿 Create your feature branch (`git checkout -b feature/amazing-feature`)
3. 💾 Commit your changes (`git commit -m 'Add some amazing feature'`)
4. 📤 Push to the branch (`git push origin feature/amazing-feature`)
5. 🎯 Open a Pull Request