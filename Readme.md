# gRPC Go Study Project

A comprehensive collection of gRPC implementations in Go, demonstrating various communication patterns and real-world use cases.

## 📋 Overview

This repository contains three complete gRPC services showcasing different aspects of gRPC development:

- **Greet Service** - Demonstrates all four gRPC communication patterns
- **Calculator Service** - Implements mathematical operations with streaming
- **Blog Service** - Full CRUD application with MongoDB integration

## 🚀 Features

### Communication Patterns

- **Unary RPC** - Single request, single response
- **Server Streaming** - Single request, multiple responses
- **Client Streaming** - Multiple requests, single response
- **Bidirectional Streaming** - Multiple requests and responses

### Services Included

#### 1. Greet Service

- Simple greeting with various streaming patterns
- Deadline/timeout handling
- All four communication patterns implemented

#### 2. Calculator Service

- Prime number decomposition (server streaming)
- Average calculation (client streaming)
- Maximum number finder (bidirectional streaming)
- Square root with error handling

#### 3. Blog Service

- Complete CRUD operations
- MongoDB integration
- Create, Read, Update, Delete blogs
- List all blogs with streaming

## 🛠️ Tech Stack

- **Go** 1.22.1+
- **gRPC** - High-performance RPC framework
- **Protocol Buffers** - Data serialization
- **MongoDB** - Database for blog service
- **Docker** - Containerization support

## 📦 Installation

### Prerequisites

- Go 1.22.1 or higher
- Protocol Buffers compiler (`protoc`)
- MongoDB (for blog service)

### Setup

1. Clone the repository:

```bash
git clone https://github.com/Vova4o/grpc.git
cd grpc
```

2. Install dependencies:

```bash
go mod download
```

3. Generate protocol buffer code:

```bash
make all
```

## 🎯 Usage

### Build All Services

```bash
make all
```

### Build Individual Services

```bash
make greet        # Build greet service
make calculator   # Build calculator service
make blog         # Build blog service
```

### Run Servers

```bash
# Greet server
./bin/greet/server

# Calculator server
./bin/calculator/server

# Blog server (requires MongoDB)
./bin/blog/server
```

### Run Clients

```bash
# Greet client
./bin/greet/client

# Calculator client
./bin/calculator/client

# Blog client
./bin/blog/client
```

## 📁 Project Structure

```
.
├── greet/              # Greet service implementation
│   ├── client/         # Client implementations
│   ├── server/         # Server implementations
│   └── proto/          # Protocol buffer definitions
├── calculator/         # Calculator service implementation
│   ├── client/         # Client implementations
│   ├── server/         # Server implementations
│   └── proto/          # Protocol buffer definitions
├── blog/               # Blog service with MongoDB
│   ├── client/         # Client implementations
│   ├── server/         # Server implementations
│   ├── proto/          # Protocol buffer definitions
│   └── docker-compose.yml
├── ssl/                # SSL/TLS certificates
└── Makefile           # Build automation
```

## 🔧 Available Make Commands

```bash
make all            # Generate protobuf code and build all services
make greet          # Build greet service
make calculator     # Build calculator service
make blog           # Build blog service
make test           # Run tests
make clean          # Clean generated files
make rebuild        # Clean and rebuild everything
make bump           # Update dependencies
make about          # Display build information
make help           # Show available commands
```

## 🌟 Key Learnings

This project demonstrates:

- Protocol buffer schema design
- Code generation from `.proto` files
- Implementation of all gRPC communication patterns
- Error handling and deadlines
- Database integration with gRPC
- Reflection API for debugging
- Best practices for service architecture

## 📝 API Examples

### Greet Service

- `Greet()` - Simple unary call
- `GreetManyTimes()` - Server streaming
- `LongGreet()` - Client streaming
- `GreetEveryone()` - Bidirectional streaming
- `GreetWithDeadline()` - Timeout handling

### Calculator Service

- `Numbers()` - Prime number decomposition
- `AverageOfSumm()` - Calculate average
- `Sqrt()` - Square root with error handling

### Blog Service

- `CreateBlog()` - Create new blog post
- `ReadBlog()` - Read blog by ID
- `UpdateBlog()` - Update existing blog
- `DeleteBlog()` - Delete blog post
- `ListBlogs()` - Stream all blogs

## 🔒 SSL/TLS Support

The project includes SSL certificate generation scripts in the `ssl/` directory for secure communication.

## 🤝 Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## 📄 License

This project is for educational purposes.

## 👤 Author

**Vladimir Gavrilenko** ([@Vova4o](https://github.com/Vova4o))

---

⭐ If you found this project helpful, please give it a star!
