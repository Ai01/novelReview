## ADDED Requirements

### Requirement: Explore feed API
The system SHALL provide an API endpoint that returns a paginated list of book reviews ordered by most recent.

#### Scenario: Fetch first page of explore feed
- **WHEN** client sends `GET /api/v1/explore?cursor=0&limit=20`
- **THEN** the server returns HTTP 200 with `items` array (each containing `comment_id`, `book`, `user`, `content`, `likes`, `created_at`), `next_cursor`, and `has_more`

#### Scenario: Fetch next page with cursor
- **WHEN** client sends `GET /api/v1/explore?cursor=20&limit=20`
- **THEN** the server returns items with `comment_id` > 20, ordered by `created_at` DESC

#### Scenario: Last page
- **WHEN** there are fewer than `limit` items remaining
- **THEN** the server returns `has_more: false` and `next_cursor` equal to the last item's id

#### Scenario: Invalid cursor
- **WHEN** client sends a non-numeric or negative cursor value
- **THEN** the server returns HTTP 400 with error message

#### Scenario: Invalid limit
- **WHEN** client sends limit > 50 or limit <= 0
- **THEN** the server clamps limit to 20 (default) and returns HTTP 200

### Requirement: Explore search API
The system SHALL provide a search endpoint that filters book reviews by keyword matching book title, author, or review content.

#### Scenario: Search by book title keyword
- **WHEN** client sends `GET /api/v1/explore/search?q=三体&cursor=0&limit=20`
- **THEN** the server returns review items where the book title or author matches "三体", or review content contains "三体"

#### Scenario: Search with pagination
- **WHEN** client sends `GET /api/v1/explore/search?q=三体&cursor=20&limit=20`
- **THEN** the server returns the next page of matching results using cursor-based pagination

#### Scenario: Empty search keyword
- **WHEN** client sends `GET /api/v1/explore/search?q=` (empty query)
- **THEN** the server returns HTTP 400 with error message "搜索关键词不能为空"

#### Scenario: No search results
- **WHEN** no books or reviews match the keyword
- **THEN** the server returns HTTP 200 with empty `items` array, `has_more: false`

### Requirement: API response format
The API SHALL use a consistent JSON response envelope for explore endpoints.

#### Scenario: Successful response structure
- **WHEN** a request succeeds
- **THEN** the response body is `{"items": [...], "next_cursor": <number>, "has_more": <boolean>}` and HTTP status is 200

#### Scenario: Each item in the list
- **WHEN** an item is returned in the `items` array
- **THEN** it MUST contain `comment_id` (number), `book` object (id, title, author, cover), `user` object (id, username, avatar), `content` (string), `likes` (number), `created_at` (ISO 8601 string)

### Requirement: No authentication required
The explore endpoints SHALL be publicly accessible without authentication tokens.

#### Scenario: Unauthenticated access
- **WHEN** client sends a request without an Authorization header
- **THEN** the server returns HTTP 200 with data (no 401 error)

### Requirement: Performance
The explore API SHALL respond within 300ms under normal load.

#### Scenario: API response time
- **WHEN** the explore endpoint is called with default limit=20
- **THEN** the server responds within 300ms

### Requirement: Error handling
The explore API SHALL return appropriate HTTP error codes for invalid requests and server errors.

#### Scenario: Database connection failure
- **WHEN** the database is unreachable
- **THEN** the server returns HTTP 503 with `{"error": "服务暂时不可用"}`

#### Scenario: Internal server error
- **WHEN** an unexpected error occurs during query
- **THEN** the server returns HTTP 500 with `{"error": "服务器内部错误"}` and logs the error
