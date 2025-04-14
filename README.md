# Go Feedback Service ![Go Version](https://img.shields.io/badge/Go-1.21+-brightgreen)

A containerized feedback collection API built with Go, deployed with Kubernetes, and automated with GitHub Actions CI/CD.

---

## 🚀 Features

- RESTful API with JSON support
- Anonymous feedback submission with optional ratings
- Dockerized for consistent environments
- Ready for Kubernetes with deployment manifests
- Automated CI/CD pipeline via GitHub Actions
- Lightweight and fast (single binary deployment)
- Structured logging for easy debugging

---

## 📁 Project Structure - 
```
.
├── k8s/ # Kubernetes deployment files
│ ├── deployment.yaml # App deployment configuration
│ └── service.yaml # Service exposure configuration
├── .github/workflows/ # CI/CD automation
│ └── deploy.yaml # GitHub Actions workflow
├── main.go # Primary application code
├── Dockerfile # Container build instructions
├── go.mod # Go module dependencies
└── README.md # You are here :)

```

---

## 🛠️ Tech Stack

- **Language**: Go 1.21+
- **Containerization**: Docker ![Docker](https://img.shields.io/badge/Docker-ready-blue)
- **Orchestration**: Kubernetes (Minikube compatible) ![Kubernetes](https://img.shields.io/badge/Kubernetes-compatible-blue)
- **CI/CD**: GitHub Actions ![CI/CD](https://img.shields.io/badge/CI%2FCD-GitHub%20Actions-yellow)
- **Package Management**: Go Modules

---

## 🏁 Quick Start

### Prerequisites

- Go 1.21+
- Docker
- kubectl
- Minikube (for local Kubernetes)

### Run Locally

```bash
# Clone repository
git clone https://github.com/maxwell-alan-josseph/go-feedback-service.git
cd go-feedback-service

# Install dependencies
go mod tidy

# Run application
go run main.go
```

Access the API at: http://localhost:8080

🐳 Docker Deployment
```
# Build image
docker build -t go-feedback-service .

# Run container
docker run -p 8080:8080 go-feedback-service
```
☸️ Kubernetes Deployment
```
# Start Minikube cluster
minikube start

# Deploy application
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml

# Access service
minikube service go-feedback-service
```
📡 API Reference

Submit Feedback
```
curl -X POST http://localhost:8080/feedback \
  -H "Content-Type: application/json" \
  -d '{"message":"Great service!","rating":5}'
```
Response:
```
{
  "id": "abc123",
  "message": "Great service!",
  "rating": 5,
  "created_at": "2023-11-15T10:00:00Z"
}
```
List All Feedback
```
curl http://localhost:8080/feedbacks
```
Response:
```
[
  {
    "id": "abc123",
    "message": "Great service!",
    "rating": 5,
    "created_at": "2023-11-15T10:00:00Z"
  }
]
```

🔄 CI/CD Pipeline
GitHub Actions workflow includes:
1.On push to main branch:
  - Go vet and linting
  - Unit tests (if any)
  - Docker image build
  - Push to GitHub Container Registry
  - Deployment to Kubernetes cluster

🤝 Contributing
We welcome contributions! Please follow these steps:
1. Fork the repository
2. Create your feature branch (git checkout -b feature/amazing-feature)
3. Commit your changes (git commit -m 'Add some amazing feature')
4. Push to the branch (git push origin feature/amazing-feature)
5. Open a Pull Request
