<h1 align="center">
A Simple GraphQL Server
</h1>
<p align="center">
    <a href="#" target="_blank">
        <img src="https://img.shields.io/badge/Go-1.17+-00ADD8?style=for-the-badge&logo=go" alt="go version" /> 
    </a>
    &nbsp;
    <a>
        <img src="https://img.shields.io/badge/GraphQl-E10098?style=for-the-badge&logo=graphql&logoColor=white" alt="graphql" />
    </a>
    &nbsp;
    <a href="#" target="_blank">
        <img src="https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white" alt="Postgres" />
    </a>
    &nbsp;
    <a href="#" target="_blank">
        <img src="https://img.shields.io/badge/Docker-2CA5E0?style=for-the-badge&logo=docker&logoColor=white" alt="Postgres" />
    </a>
</p>
<hr/>

### Problem

Build a simple graphql server with postgres database that can perform this simple task. Finding books of all the authors and vice versa. `Using this library : https://github.com/graphql-go/graphql`

### Starting Guide:
#### With Docker:
1. Clone this repository
```shell
git clone https://github.com/milon19/graphql-server-golang.git 
```

2. Create a `.env` file with following data in project root directory
```dotenv
COMPOSE_FILE=build/docker-compose.yml
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_HOST=database
POSTGRES_PORT=5437
DATABASE_NAME=graphql-server
SERVER_PORT=8080
```

3. Run Docker Command
```shell
docker-compose up --build -d
```
Server will start on `127.0.0.0:8080`

#### Run locally without docker:
1. Clone this repository
```shell
git clone https://github.com/milon19/graphql-server-golang.git
```
2. Create a .env file by following `.env.example` file in project root directory
```dotenv
COMPOSE_FILE=build/docker-compose.yml
POSTGRES_USER=<DB_USER>
POSTGRES_PASSWORD=<DB_PASSWORD>
POSTGRES_HOST=<DB_HOST>
POSTGRES_PORT=<DB_PORT>
DATABASE_NAME=<DB_NAME>
SERVER_PORT=8080
```
3. Create a postgres database with the name you have given in `.env`
4. Run **`build/init.sql`** to create tables and populate data
5. Run project
```shell
go run cmd/app/main.go 
```

> **NOTE**: It is recommended to run using docker to avoid unnecessary setup procedure.

### Test the server:
#### Queries:
1. Query for `GetAllBooksWithAuthors`: 

> [http://127.0.0.1:8080/graphql?query={books{id,title,price,isbn_no,authors{id,name,biography}}}](http://127.0.0.1:8080/graphql?query={books{id,title,price,isbn_no,authors{id,name,biography}}})

2. Query for `GetAllTheBooksOf<name>`
> [http://127.0.0.1:8080/graphql?query={authors(name:"Robi"){id,name,biography,books{id,title,price,isbn_no}}}](http://127.0.0.1:8080/graphql?query={authors(name:"Robi"){id,name,biography,books{id,title,price,isbn_no}}})

### File Structure
```shell
├── build
│   ├── docker-compose.yml
│   ├── Dockerfile
│   └── init.sql
├── cmd
│   └── app
│       └── main.go                  // Program entry point
├── go.mod
├── go.sum
├── internal                         // Main source  code
│   └── app
│       ├── adapter                  // Outer layer. All Framework and Database related code
│       │   ├── controller           // Controller funtion
│       │   │   ├── service.go       // DB service and repository provider
│       │   │   └── graphql.go
│       │   ├── db
│       │   │   └── connections.go   // postgre db connection
│       │   ├── graphql              // All graphql related code
│       │   │   ├── dataloader.go
│       │   │   └── types.go
│       │   ├── repository          // Implementation of repository
│       │   │   ├── author.go
│       │   │   └── book.go
│       │   ├── route.go
│       │   └── schema              // Schemas
│       │       ├── author.go
│       │       ├── book.go
│       │       └── graphql.go
│       ├── application             // Middle Layer. mainly deals with business logic
│       │   └── usecase
│       │       ├── author_loader.go
│       │       └── book_loader.go
│       └── domain                  // Inner Layer. All Schemas and repository defination
│           ├── author.go
│           ├── book.go
│           └── repository
│               ├── author.go
│               └── book.go
└── README.md

```