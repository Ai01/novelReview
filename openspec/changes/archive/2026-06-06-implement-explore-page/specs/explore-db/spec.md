## ADDED Requirements

### Requirement: Books table
The system SHALL have a `books` table to store novel/book metadata.

#### Scenario: Table creation
- **WHEN** schema migration runs
- **THEN** a `books` table is created with columns: `id` (BIGINT AUTO_INCREMENT PK), `title` (VARCHAR(255) NOT NULL), `author` (VARCHAR(255) NOT NULL), `cover` (VARCHAR(512)), `description` (TEXT), `category` (VARCHAR(100)), `tags` (VARCHAR(255)), `status` (VARCHAR(50) DEFAULT '连载中'), `last_updated` (DATETIME), `created_at` (DATETIME DEFAULT CURRENT_TIMESTAMP), `updated_at` (DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP)

#### Scenario: Title index
- **WHEN** searching books by title
- **THEN** a FULLTEXT index on `(title, author)` enables efficient text search

### Requirement: Comments table
The system SHALL have a `comments` table to store user reviews on books.

#### Scenario: Table creation
- **WHEN** schema migration runs
- **THEN** a `comments` table is created with columns: `id` (BIGINT AUTO_INCREMENT PK), `book_id` (BIGINT NOT NULL), `user_id` (BIGINT NOT NULL), `content` (TEXT NOT NULL), `likes` (INT DEFAULT 0), `created_at` (DATETIME DEFAULT CURRENT_TIMESTAMP)

#### Scenario: Foreign key constraints
- **WHEN** a comment references a book or user
- **THEN** `book_id` references `books(id)` and `user_id` references `users(id)` with foreign key constraints

#### Scenario: Query performance indexes
- **WHEN** querying comments by book or by time
- **THEN** indexes exist on `(book_id, created_at)` for per-book comment queries and `(created_at)` for global timeline queries

### Requirement: Seed data
The system SHALL include seed data with at least 20 books and 40 comments to populate the explore page for development and testing.

#### Scenario: Books seed data
- **WHEN** the seed script runs
- **THEN** at least 20 book records are inserted with realistic Chinese novel data including titles, authors, categories, and statuses

#### Scenario: Comments seed data
- **WHEN** the seed script runs
- **THEN** at least 40 comment records are inserted, each associated with a valid book and user, with varied content and timestamps

#### Scenario: Idempotent seeding
- **WHEN** the seed script runs more than once
- **THEN** it does not duplicate existing data (uses INSERT IGNORE or checks before insert)
