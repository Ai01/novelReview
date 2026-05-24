## ADDED Requirements

### Requirement: Nginx gateway is deployed as a Docker service
The system MUST deploy the Nginx gateway via Docker using docker compose configurations stored under the repository-root deploy directory.

#### Scenario: Starting gateway via deploy
- **WHEN** an operator starts infrastructure from the deploy entry point
- **THEN** the Nginx gateway container is started and becomes reachable

### Requirement: Nginx provides a single external entry point
The system SHALL provide an Nginx gateway that exposes a single external entry point for the application.

#### Scenario: External client accesses application via gateway
- **WHEN** a client sends requests to the gateway base URL
- **THEN** the gateway routes the request to the correct upstream (frontend or backend) based on path rules

### Requirement: Gateway routes API traffic to backend under /api
The gateway MUST route requests under the /api path prefix to the backend service.

#### Scenario: Routing an API request
- **WHEN** a client sends a request to /api/*
- **THEN** Nginx forwards the request to the backend upstream and returns the backend response

### Requirement: Gateway serves frontend content for non-API paths
The gateway MUST serve frontend static content for requests that are not under /api.

#### Scenario: Serving frontend index
- **WHEN** a client requests /
- **THEN** Nginx returns the frontend entry resource

### Requirement: Gateway provides health endpoint for infrastructure checks
The gateway SHALL expose a lightweight health endpoint suitable for liveness checks.

#### Scenario: Health check request
- **WHEN** a client requests the gateway health endpoint
- **THEN** the gateway returns a successful status without requiring backend availability

### Requirement: Gateway preserves and forwards standard proxy headers
The gateway MUST forward standard proxy headers (including X-Forwarded-For and X-Forwarded-Proto) to upstream services.

#### Scenario: Forwarding proxy headers
- **WHEN** a request passes through the gateway to the backend
- **THEN** the backend receives the forwarded proxy headers reflecting the original client request
