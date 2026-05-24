## ADDED Requirements

### Requirement: Deploy directory exists at repository root
The system MUST provide a deploy directory at the repository root (top-level) as the single canonical location for infrastructure deployment assets.

#### Scenario: Operator locates deploy assets
- **WHEN** an operator opens the repository root directory
- **THEN** the deploy directory exists as a first-level folder and contains the deployment entry point

### Requirement: Deploy directory provides a single entry point for infrastructure startup
The system SHALL provide a deploy directory that contains the canonical entry script(s) to start and initialize infrastructure services required by the application.

#### Scenario: Starting infrastructure from deploy entry point
- **WHEN** an operator runs the deploy entry script
- **THEN** required infrastructure services are started in the correct order and the script exits with code 0

### Requirement: Deploy assets are self-contained and environment-agnostic
The deploy directory MUST include all configuration and scripts needed to bootstrap infrastructure, and MUST rely on environment variables for environment-specific values.

#### Scenario: Running deploy on different environments
- **WHEN** the same deploy entry script is executed with different environment variable values
- **THEN** the infrastructure stack starts targeting the environment-specific endpoints and ports

### Requirement: Infrastructure services are deployed via Docker Compose from deploy
The system MUST deploy infrastructure services via Docker using docker compose configurations stored under the deploy directory.

#### Scenario: Starting infrastructure with docker compose
- **WHEN** an operator starts infrastructure using the deploy entry script
- **THEN** the entry script uses docker compose to start the infrastructure services

### Requirement: Deploy entry script is safe to run multiple times
The deploy entry script MUST be idempotent and MUST not fail when run while the infrastructure is already running.

#### Scenario: Re-running deploy while services are up
- **WHEN** an operator runs the deploy entry script a second time
- **THEN** the script completes successfully without breaking running services

### Requirement: Deploy process defines the deployment order contract
The deployment process MUST define the order contract as: infrastructure first, then backend, then frontend.

#### Scenario: Deploy order execution
- **WHEN** an operator follows the deployment process
- **THEN** infrastructure services are started before the backend and frontend are deployed
