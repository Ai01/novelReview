## ADDED Requirements

### Requirement: Preparation scripts live under repository-root deploy directory
The system MUST place preparation scripts under the repository-root deploy directory so operators can execute bootstrap from a single canonical location.

#### Scenario: Running bootstrap from deploy
- **WHEN** an operator runs the bootstrap script from the deploy directory
- **THEN** the bootstrap completes successfully without requiring scripts from other locations

### Requirement: Preparation scripts are the only place for environment bootstrap
The system SHALL provide preparation scripts that perform environment bootstrap tasks (database initialization/migration, optional seed data, and service dependency checks) without requiring the backend service process to perform these tasks at runtime.

#### Scenario: Bootstrap executed before starting backend
- **WHEN** an operator runs the preparation entry script
- **THEN** the script completes all required bootstrap steps and exits with code 0
- **AND THEN** the backend can start without performing bootstrap tasks

### Requirement: Preparation scripts are idempotent and retryable
Preparation scripts MUST be safe to execute multiple times and MUST be retryable after partial failure.

#### Scenario: Re-running bootstrap after success
- **WHEN** an operator re-runs the same preparation entry script after it already succeeded
- **THEN** the script completes without duplicating data or breaking existing state

#### Scenario: Re-running bootstrap after partial failure
- **WHEN** the preparation entry script fails part-way through
- **THEN** an operator can re-run the script and it completes successfully without manual cleanup

### Requirement: Preparation scripts provide deterministic exit codes and logs
Preparation scripts MUST write progress and error information to standard output/error and MUST return a non-zero exit code on failure.

#### Scenario: Failing bootstrap step
- **WHEN** a required dependency is unavailable or a migration step fails
- **THEN** the script exits with a non-zero code and prints an actionable error message

### Requirement: Preparation scripts are configurable via environment variables
Preparation scripts MUST accept configuration via environment variables so the same scripts can run in local, staging, and production environments.

#### Scenario: Using environment variables for database endpoints
- **WHEN** an operator provides database connection variables in the environment
- **THEN** the script uses those variables to connect to the correct database instances
