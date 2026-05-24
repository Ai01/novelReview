## ADDED Requirements

### Requirement: Observability services are deployed as Docker services
The system MUST deploy Prometheus and Grafana via Docker using docker compose configurations stored under the repository-root deploy directory.

#### Scenario: Starting observability via deploy
- **WHEN** an operator starts infrastructure from the deploy entry point
- **THEN** Prometheus and Grafana containers are started and become reachable on the internal network

### Requirement: Backend exposes Prometheus-compatible metrics
The backend service SHALL expose a Prometheus-compatible metrics endpoint.

#### Scenario: Scraping backend metrics
- **WHEN** Prometheus scrapes the backend metrics endpoint
- **THEN** the backend responds with a valid Prometheus text exposition format payload

### Requirement: Prometheus scrapes the backend metrics endpoint
The system MUST configure Prometheus to scrape the backend metrics endpoint on a defined interval.

#### Scenario: Prometheus target is up
- **WHEN** Prometheus starts with the provided configuration
- **THEN** the backend target appears as up in Prometheus target status

### Requirement: Grafana connects to Prometheus as a datasource
The system MUST configure Grafana to connect to Prometheus as a datasource.

#### Scenario: Grafana datasource is available
- **WHEN** Grafana starts with provisioning enabled
- **THEN** the Prometheus datasource exists and queries can be executed successfully

### Requirement: Minimal dashboard exists for service health
The system SHALL provide a minimal dashboard that visualizes service health signals (availability and request latency/error rate) derived from available metrics.

#### Scenario: Opening the health dashboard
- **WHEN** an operator opens Grafana and selects the application health dashboard
- **THEN** the dashboard loads panels populated with Prometheus data

### Requirement: Observability stack starts independently of application deploy
Prometheus and Grafana MUST be startable as part of infrastructure deployment without requiring backend or frontend deployment to succeed.

#### Scenario: Starting observability first
- **WHEN** an operator starts the observability stack without starting backend/frontend
- **THEN** Prometheus and Grafana start successfully and remain running
