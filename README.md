# E-Commerce Microservices Platform

A production-ready microservices e-commerce platform with complete CI/CD pipeline.

## 🏗️ Architecture

This project demonstrates a modern microservices architecture with:
- **4 Microservices**: User, Product, Order, Payment
- **Multiple Tech Stacks**: Node.js, Python (FastAPI), Go
- **Databases**: PostgreSQL, MongoDB, Redis
- **Container Orchestration**: Docker, Kubernetes
- **CI/CD**: GitHub Actions, Jenkins
- **Monitoring**: Prometheus, Grafana

## 📋 Services

| Service | Tech Stack | Port | Description |
|---------|-----------|------|-------------|
| User Service | Node.js + Express | 3001 | User authentication and management |
| Product Service | Python + FastAPI | 3002 | Product catalog management |
| Order Service | Go + Gorilla Mux | 3003 | Order processing |
| Payment Service | Node.js + Express | 3004 | Payment processing |

## 🚀 Quick Start

### Prerequisites
- Docker & Docker Compose
- Node.js 18+
- Python 3.11+
- Go 1.21+
- kubectl (for Kubernetes deployment)
- Helm 3+ (for Kubernetes deployment)

### Run Locally with Docker Compose
```bash
# Clone the repository
git clone https://github.com/YOUR_USERNAME/ecommerce-microservices.git
cd ecommerce-microservices

# Start all services
docker-compose up -d

# Check service health
curl http://localhost:3001/health  # User Service
curl http://localhost:3002/health  # Product Service
curl http://localhost:3003/health  # Order Service

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

### Run Individual Services

#### User Service (Node.js)
```bash
cd services/user-service
npm install
npm run dev
```

#### Product Service (Python)
```bash
cd services/product-service
pip install -r requirements.txt
python -m app.main
```

#### Order Service (Go)
```bash
cd services/order-service
go mod download
go run main.go
```

## 🔧 Development

### Project Structure
