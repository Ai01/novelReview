## ADDED Requirements

### Requirement: Domain layer isolation with entity-level splitting
The `internal/domain` package SHALL be split into sub-packages by domain entity (user, book, comment), each containing only the GORM model struct for that entity with zero dependencies on other internal packages.

#### Scenario: Domain entity sub-package location
- **WHEN** a developer needs to add a new database entity (e.g., BookList)
- **THEN** they create `internal/domain/booklist/booklist.go` with the struct definition, and it imports nothing from `repository`, `service`, or `handler`

#### Scenario: Domain model reuse across layers
- **WHEN** repository or service or handler needs a User/Book/Comment struct
- **THEN** they import from `internal/domain/user`, `internal/domain/book`, or `internal/domain/comment` respectively

#### Scenario: Comment model references other domains
- **WHEN** the Comment struct has foreign-key associations to User and Book
- **THEN** it SHALL import `internal/domain/user` and `internal/domain/book` directly (domain-to-domain dependency is allowed)

### Requirement: Repository layer with interface abstraction
The `internal/repository` package SHALL define interfaces for data access and provide GORM-based implementations, depending only on `internal/domain`.

#### Scenario: Repository interface definition
- **WHEN** a new data access pattern is needed
- **THEN** the repository defines an interface (e.g., `CommentRepository`) and a private implementation struct

#### Scenario: Repository implementation uses GORM
- **WHEN** a repository method queries the database
- **THEN** it SHALL use only GORM methods (no raw SQL), with no knowledge of HTTP or business logic

#### Scenario: Service depends on repository interface
- **WHEN** a service needs data access
- **THEN** it accepts the repository interface as a constructor parameter, not the concrete implementation

### Requirement: Service layer with business logic
The `internal/service` package SHALL contain business logic orchestration, depending on `internal/domain` and `internal/repository` interfaces.

#### Scenario: Service constructor
- **WHEN** a service is instantiated
- **THEN** it accepts all dependencies (repository interfaces, database connections) via a `NewXxxService()` constructor function

#### Scenario: Business logic in service
- **WHEN** a use case requires multiple repository calls or data transformation
- **THEN** the orchestration logic lives in the service, not in the handler

### Requirement: Handler layer with HTTP concerns only, split by domain
The `internal/handler` package SHALL handle HTTP parameter parsing, validation, response serialization, and error codes. Handler files SHALL be split by domain entity (auth, explore/book/comment). They SHALL NOT contain business logic or direct database queries.

#### Scenario: Handler delegates to service
- **WHEN** a handler receives a request
- **THEN** it parses query/body parameters, calls the corresponding service method, and serializes the result to JSON

#### Scenario: Handler error handling
- **WHEN** the service returns an error
- **THEN** the handler maps the error to the appropriate HTTP status code and JSON error response

#### Scenario: Handler file naming by domain
- **WHEN** a new API endpoint is added for book operations
- **THEN** the handler code goes into `internal/handler/book_handler.go`, not into a single monolithic file

### Requirement: Router layer centralizes route registration
The `internal/router` package SHALL register all API routes in a single function, accepting handler dependencies as parameters.

#### Scenario: Route registration
- **WHEN** the application starts
- **THEN** `router.Setup(engine, handlerDeps)` registers all `GET`/`POST` routes under `/api/v1/`

#### Scenario: New route addition
- **WHEN** a new API endpoint is added
- **THEN** only the router file and the corresponding handler/service/repository files need to change; `main.go` is not modified

### Requirement: Infrastructure package for external dependencies
The `internal/infra` package SHALL manage database connections, providing `DBForWrite()` and `DBForRead()` functions used by `main.go` for dependency injection.

#### Scenario: Database initialization
- **WHEN** `main.go` calls `infra.InitDB()`
- **THEN** write and read database connections are established with retry logic

#### Scenario: Read/write separation
- **WHEN** a repository needs a database connection
- **THEN** `main.go` injects the connection (or the service uses a helper) without the repository package importing `infra` directly

### Requirement: main.go as composition root
`main.go` SHALL be the sole composition root, constructing all dependencies and wiring them together. It SHALL NOT contain any business logic or route definitions directly.

#### Scenario: Application startup
- **WHEN** the application starts
- **THEN** `main.go` initializes infra, creates repositories, creates services, creates handlers, sets up the router, and starts the HTTP server

#### Scenario: Build verification
- **WHEN** `go build -o main .` is run from the `backend/` directory
- **THEN** the binary compiles successfully with all `internal/` subpackages

### Requirement: README.md reflects new architecture
The project README.md SHALL document the backend layered architecture with a diagram or description of each layer's responsibility.

#### Scenario: README architecture section
- **WHEN** a developer reads README.md
- **THEN** there is a "后端架构" section listing domain/repository/service/handler/router/infra layers with one-line descriptions
