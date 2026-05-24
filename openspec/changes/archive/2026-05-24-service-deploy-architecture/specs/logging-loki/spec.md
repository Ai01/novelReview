## ADDED Requirements

### Requirement: Logging services are deployed as Docker services
The system MUST deploy Loki and the log collection component via Docker using docker compose configurations stored under the repository-root deploy directory.

#### Scenario: Starting logging via deploy
- **WHEN** an operator starts infrastructure from the deploy entry point
- **THEN** Loki and the log collector containers are started and become reachable on the internal network

### Requirement: Loki provides centralized log storage and query
The system SHALL provide Loki as the centralized store for application and infrastructure logs with query capability.

#### Scenario: Loki starts as part of infrastructure deployment
- **WHEN** an operator starts the infrastructure stack
- **THEN** the Loki service is running and reachable on the internal network

### Requirement: Logs are collected and shipped to Loki
The system MUST provide a log collection component that ships logs to Loki.

#### Scenario: Collecting backend logs
- **WHEN** the backend emits logs to standard output or configured log files
- **THEN** the log collection component ships log entries to Loki and they become queryable

### Requirement: Grafana connects to Loki as a datasource
The system MUST configure Grafana to connect to Loki as a datasource.

#### Scenario: Loki datasource is available
- **WHEN** Grafana starts with provisioning enabled
- **THEN** the Loki datasource exists and log queries can be executed successfully

### Requirement: Logs are queryable by service and environment labels
The log pipeline MUST attach labels that allow querying logs by at least service name and environment.

#### Scenario: Query logs by service label
- **WHEN** an operator queries logs filtered by a service label
- **THEN** Grafana returns only log entries for that service

### Requirement: Logging stack starts independently of application deploy
Loki and the log collection component MUST be startable as part of infrastructure deployment without requiring backend or frontend deployment to succeed.

#### Scenario: Starting logging first
- **WHEN** an operator starts the logging stack without starting backend/frontend
- **THEN** Loki and the log collector start successfully and remain running
