## ADDED Requirements

### Requirement: Schema migration deployment
The system SHALL apply the new `books` and `comments` table schema as part of the MySQL bootstrap process without disrupting existing data.

#### Scenario: Fresh deployment
- **WHEN** a fresh deployment runs `infra-up.sh` and `app-up.sh`
- **THEN** the MySQL bootstrap script creates new `books` and `comments` tables alongside existing tables, and inserts seed data

#### Scenario: Rolling update on existing deployment
- **WHEN** deploying to an environment with existing data
- **THEN** the schema migration uses `CREATE TABLE IF NOT EXISTS` so existing data is preserved

### Requirement: Backend service update
The system SHALL deploy the updated Go backend with new explore API endpoints as part of the standard `app-up.sh` process.

#### Scenario: Backend rebuild and restart
- **WHEN** `app-up.sh` is executed after code changes
- **THEN** the backend Docker image is rebuilt and the container is restarted with the new explore endpoints available at `/api/v1/explore` and `/api/v1/explore/search`

#### Scenario: Health check after deploy
- **WHEN** the backend container restarts
- **THEN** the `/readyz` endpoint returns 200 once the database connection is established

### Requirement: Frontend deployment
The system SHALL build and serve the updated React frontend with the new explore page via nginx.

#### Scenario: Frontend rebuild
- **WHEN** `app-up.sh` is executed
- **THEN** the frontend is rebuilt with `npm run build` and the static assets are served by the nginx container

#### Scenario: SPA routing
- **WHEN** user navigates directly to `/explore` in the browser
- **THEN** nginx serves `index.html` (not a 404) so React Router handles the route client-side

### Requirement: No infrastructure changes
The explore page deployment SHALL NOT require new Docker services, containers, or infrastructure components.

#### Scenario: docker-compose unchanged
- **WHEN** the explore page is deployed
- **THEN** no new services are added to `docker-compose.yml`, and existing services (backend, frontend, nginx, mysql, prometheus, grafana, loki, promtail) remain unchanged in structure

### Requirement: Monitoring compatibility
The explore API endpoints SHALL be compatible with existing Prometheus metrics collection and Loki log aggregation.

#### Scenario: Metrics exposure
- **WHEN** the explore endpoints are called
- **THEN** request metrics (duration, status code) are automatically captured by the existing Gin/Prometheus instrumentation

#### Scenario: Log aggregation
- **WHEN** the backend logs explore-related requests
- **THEN** logs are collected by Promtail and available in Grafana via Loki (no configuration changes needed)
