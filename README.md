# API
## Introduction
The RESTful API to CRUD a Security Scan Result (“Result”) written by Go. The structure of Result
- Id: any type of unique id
- Status: "Queued" | "In Progress" | "Success" | "Failure"
- RepositoryName: string
- Findings: JSONB, see [example](https://github.com/guardrailsio/backend-engineer-challenge/blob/master/example-findings.json)
- QueuedAt: timestamp
- ScanningAt: timestamp
- FinishedAt: timestamp
## Diagrams
### General Architect
### ERD
## Project Structures
```
/adapters
/configurations
/controllers
/databases
/deployments
/docs
/documentations
/dtos
/entities
/repositories
/routes
/services
```
## Technology Stack
1. Go Gin (Web server framework)
2. Gorm (ORM package)
3. PosgresQL (RDBMS)
4. Viper (Config loader package)
## How to run
### Prerequisites
1. We need install Go 1.16 or later. For installation instructions, see[Installing Go](https://golang.org/doc/install).
2. [Installing postgresQL](https://www.postgresql.org/download/)
3. Clone the project from https://github.com/xxxle0/guardrails-test
4. Run the command `cd /api && go mod vendor` to install dependencies of the project
### Command
```javascript
make swagger-migrate // Update api documentation
make build // Build docker image
make run // Run Go Project
```