## Requirements
1. Install git
2. Clone the repo
3. Install Go 1.24
4. Install make

## Usage
- Start: `make up`
- Stop: `make down`

## About
This app is being refactored from a monolith to microservices using gRPC for efficient inter-service communication. Each microservice handles a specific domain, improving scalability, maintainability, and deployment flexibility.

**Benefits:**  
- Independent scaling  
- Fault isolation  
- Flexible tech stack  
- Faster development

**Note:** Ensure service discovery, load balancing, and security during migration.
