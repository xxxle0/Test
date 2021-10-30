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
![General Architect](https://github.com/xxxle0/guardrails-test/blob/master/General%20Architect.png?raw=true)
*Notes: This is my proposal architect if I build a repo scan service not what I have implemented on this repo*
### ERD
![ERD](https://github.com/xxxle0/guardrails-test/blob/master/ERD.png?raw=true)
*Notes: This is my proposal ERD if I build a repo scan service not what I have implemented on this repo*
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
5. Testify (Testing suite)
## How to run
### Prerequisites
1. We need install Go 1.16 or later. For installation instructions, see[Installing Go](https://golang.org/doc/install).
2. [Installing postgresQL](https://www.postgresql.org/download/)
3. Clone the project from https://github.com/xxxle0/Test
4. Run the command `cd /api && go mod vendor` to install dependencies of the project
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