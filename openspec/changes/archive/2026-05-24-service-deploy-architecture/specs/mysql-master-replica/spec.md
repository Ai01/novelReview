## ADDED Requirements

### Requirement: MySQL master and replica are deployed as Docker services
The system MUST deploy MySQL master and replica via Docker using docker compose configurations stored under the repository-root deploy directory.

#### Scenario: Starting database via deploy
- **WHEN** an operator starts infrastructure from the deploy entry point
- **THEN** the MySQL master and replica containers are started and become reachable

### Requirement: MySQL runs in a master-replica topology
The system SHALL run MySQL in a master-replica topology where the master accepts writes and the replica is used for reads.

#### Scenario: Topology is started
- **WHEN** infrastructure deployment starts the database stack
- **THEN** the master and replica MySQL instances are running and reachable

### Requirement: Replica is configured to replicate from master
The system MUST configure the replica to replicate from the master.

#### Scenario: Replication is active
- **WHEN** a write is performed on the master
- **THEN** the corresponding data becomes visible on the replica after replication catches up

### Requirement: Application receives separate write and read DSNs
The backend MUST accept separate configuration values for write and read database connections.

#### Scenario: Backend configured with two DSNs
- **WHEN** the backend is started with a write DSN and a read DSN
- **THEN** the backend can establish both connections and uses them according to routing rules

### Requirement: Reads default to replica with fallback to master
The backend MUST route reads to the read connection by default and MUST fall back to the write connection for reads when the read connection is unavailable.

#### Scenario: Replica unavailable
- **WHEN** the replica database is unavailable
- **THEN** the backend continues to serve read requests by reading from the master

### Requirement: Backend supports strong-consistency reads via master
The backend SHALL provide a mechanism to force reads to the master for strong consistency when required by a request flow.

#### Scenario: Read-after-write requires strong consistency
- **WHEN** a request performs a write and then requires an immediate read of the written data
- **THEN** the backend reads from the master and returns the latest data
