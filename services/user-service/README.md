# User Service

User management microservice for the e-commerce platform.

## Features
- User registration and authentication
- User profile management
- JWT token generation

## Tech Stack
- Node.js + Express
- PostgreSQL
- JWT authentication

## API Endpoints
- `GET /health` - Health check
- `GET /api/users` - Get all users
- `POST /api/users` - Create new user

## Environment Variables
- `DB_HOST` - PostgreSQL host
- `DB_PORT` - PostgreSQL port
- `DB_NAME` - Database name
- `DB_USER` - Database user
- `DB_PASSWORD` - Database password

## Running Locally
```bash
npm install
npm run dev
```

## Running Tests
```bash
npm test
```
