# Crime Stats Analysis Application

This repository contains a proof of concept application that demonstrates the integration of Go, Python, gRPC, MongoDB, Postgres, Redis, and Docker. The application uses the Crime Stats API to fetch and analyze crime data.

## Overview

The application is structured to use gRPC for client-server communication, with the server implemented in Python and clients in both Go and Python. It utilizes MongoDB and Postgres for data storage and Redis for caching. Docker is employed to containerize the various components of the application.

### Components

- **Python gRPC Server:** Handles API requests and communicates with MongoDB, Postgres, and Redis.
- **Go and Python gRPC Clients:** Interact with the gRPC server.
- **MongoDB:** Stores raw and semi-structured data from the Crime Stats API.
- **Postgres:** Stores structured and processed crime data for complex queries and analytics.
- **Redis:** Used for caching frequent queries or data.
- **Docker:** Manages and containerizes the application and database services.

## Getting Started

### Prerequisites

- Docker
- Docker Compose
- Go (for Go client)
- Python 3.x

### Installation

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/marcelluseasley/crime-stats-analysis-app.git
   cd crime-stats-analysis-app
   ```

2. **Build and Run with Docker Compose:**
   ```bash
   docker-compose up --build
   ```

### Directory Structure

```plaintext
my-grpc-app/
├── buf.yaml
├── proto/
├── go/
│   ├── client/
│   ├── server/
│   └── pb/
├── python/
│   ├── client/
│   ├── server/
│   └── pb/
├── tests/
├── .gitignore
├── README.md
└── Dockerfile
```

## Usage

- **Starting the Server:**
  The Python gRPC server starts automatically when you run `docker-compose up`. It listens for incoming RPC calls.

- **Using the Clients:**
  Clients in Go and Python can interact with the server once it's up and running. See client documentation in `go/client/` and `python/client/` for details.

## Development

### Generating gRPC Stubs

- Use `buf` to generate/update the gRPC stubs for Go and Python as needed.

### Adding New Features

- Follow the existing directory structure and coding standards for any new features.

## Testing

- Unit and integration tests are located in the `tests/` directory. Run these tests to ensure application integrity.

## Contributing

Contributions to this project are welcome! Please read our contributing guidelines for details on how to contribute.

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgments

- Crime Stats API: [Crime Data](https://rapidapi.com/jgentes/api/crime-data/)

