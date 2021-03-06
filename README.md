# API

## Introduction

The Web service RESTful API to CRUD a Security Scan Result (“Result”) written by Go.
The structure of Result

- Id: any type of unique id
- Status: "Queued" | "In Progress" | "Success" | "Failure"
- RepositoryName: string
- Findings: JSONB, see [example](https://github.com/guardrailsio/backend-engineer-challenge/blob/master/example-findings.json)
- QueuedAt: timestamp
- ScanningAt: timestamp
- FinishedAt: timestamp

## Diagrams

### General Architect

![General Architect](https://github.com/xxxle0/guardrails-test/blob/master/General%20Architect.png?raw=true)
_Notes: This is my architect proposal if I build a scalable scan service not what I have implemented on this repo_

### ERD

![ERD](https://github.com/xxxle0/guardrails-test/blob/master/ERD.png?raw=true)
_Notes: This is my ERD proposal if I build a scalable scan service not what I have implemented on this repo_

### API Documentation

![API documentation](https://github.com/xxxle0/Test/blob/master/API%20Documentation.png?raw=true)
*After run the project success, you can access API documentation via http://localhost:8080/swagger/index.html#/default/*
## Project Structures

```
/adapters        // Map structure between entity and dto
/configurations  // Config
/controllers     // Route Handler
/databases       // Database Config
/deployments     // Deployment stuff like helm template, Dockerfile
/docs            // Swagger Generate
/documentations  // Swagger Config
/dtos            // Dto
/entities        // Define Entity Model
/repositories    // Data Access Layer
/routes          // Router Definition
/services        // Contains Business Logic
```

## Technology Stack

1. Go Gin (HTTP Web framework)
2. Gorm (ORM)
3. PosgresQL (RDBMS)
4. Viper (Config loader)
5. Testify (Testing suite)
6. Swaggo (API documentation)

## How to run

### Prerequisites

1. We need install Go 1.16 or later. For installation instructions, see [Installing Go](https://golang.org/doc/install).
2. [Installing postgresQL](https://www.postgresql.org/download/)
3. Clone the project from https://github.com/xxxle0/Test
4. `cd /api` then change `app.env.example` -> `app.env`
5. Run the command `go mod vendor` to install dependencies of the project

### Command

```javascript
cd ./api
make swagger-migrate // Update api documentation
make build // Build docker image
make run // Run Project
make test // Run unit test
```

You can also use docker-compose to run because I have pushed the docker image of API to public docker repository `duybkit13/guard-test`

```javascript
docker compose up
```
