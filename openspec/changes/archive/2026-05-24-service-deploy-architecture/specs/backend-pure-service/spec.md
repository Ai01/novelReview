## ADDED Requirements

### Requirement: Backend startup has no cross-process side effects
The backend service process MUST start without performing environment bootstrap tasks such as starting/stopping databases, applying migrations, or importing seed data.

#### Scenario: Backend starts in a restricted environment
- **WHEN** the backend runs in an environment where it cannot manage infrastructure services
- **THEN** the backend still starts and serves requests as long as its configured dependencies are reachable

### Requirement: Backend depends on configuration-injected infrastructure endpoints
The backend MUST obtain infrastructure endpoints (database DSNs and other service URLs) via configuration and MUST not assume local infrastructure defaults.

#### Scenario: Backend started with configured DSNs
- **WHEN** the backend is started with explicit database configuration
- **THEN** it connects using the provided values without attempting to discover or start databases

### Requirement: Backend provides readiness signal based on dependency connectivity
The backend SHALL expose a readiness signal that reflects whether critical dependencies (at minimum, the write database connection) are reachable.

#### Scenario: Database unavailable at startup
- **WHEN** the backend starts while the write database is unreachable
- **THEN** the readiness signal indicates not ready

#### Scenario: Database becomes reachable
- **WHEN** the write database becomes reachable after startup
- **THEN** the readiness signal transitions to ready

### Requirement: Backend supports read/write connection separation
The backend MUST support using separate read and write database connections according to routing rules.

#### Scenario: Serving a write operation
- **WHEN** a request performs a write operation
- **THEN** the backend uses the write database connection

#### Scenario: Serving a read operation
- **WHEN** a request performs a non-strong-consistency read operation
- **THEN** the backend uses the read database connection when available
