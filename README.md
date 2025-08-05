# GoChat

**GoChat** is a modular chat application written in Go, designed with scalability and deployment in mind. It leverages containerization, CI/CD pipelines, and Helm-based Kubernetes deployment to support modern development workflows.

---

## 📁 Project Structure

    GoChat/
    │
    ├── charts/ # Helm charts for Kubernetes deployment
    │ └── chatapp/
    │ └── templates/
    │
    ├── cicd/ # CI/CD configurations
    │ └── github/ # GitHub Actions workflows
    │
    ├── config/ # App-level configuration files
    ├── pkg/ # Reusable Go packages
    ├── platforms/
    │ └── db/ # Database setup and schema
    ├── scripts/ # Utility scripts (e.g., migrations)
    ├── services/ # Microservices (chat, auth, etc.)
    │
    ├── Makefile # Commands to build, test, and deploy
    ├── docker-compose.yml # Local multi-service orchestration
    ├── go.mod / go.sum # Go modules and dependencies
    └── README.md # Project documentation


---

## Features

- Modular service structure using Go
- Containerized via Docker
- CI/CD with GitHub Actions
- Kubernetes deployment using Helm
- Scripted database setup and service scaffolding

---

## Getting Started

### Prerequisites

- Go 1.21+
- Docker & Docker Compose
- Helm (for Kubernetes deployments)
- Git

### Local Development (via Docker Compose)

```bash

# Clone the repository
git clone -b dev https://github.com/Maxcillius/GoChat.git
cd GoChat

# Start all services
docker-compose up --build

# Manual Build (Go)
make build          # Compile all services
make run            # Run services locally

# Kubernetes Deployment
cd charts/chatapp

# Install Helm dependencies (if any)
helm dependency update

# Deploy to cluster
helm install gochat ./

```

### CI/CD

GitHub Actions workflows are located in cicd/github/. These include:

- Build validation
- Linting
- Testing
- Docker image publishing (optional)

Workflows trigger on pushes and PRs to the dev branch.

### Contributing

Contributions are welcome!

- Fork the repo and create a new branch from dev.
- Commit and push your changes.
- Open a Pull Request targeting dev.

Please ensure code is formatted and tests are added where necessary.

### License

This project is currently unlicensed.

### Contact

For support or collaboration, open an issue or reach out via GitHub Discussions.
