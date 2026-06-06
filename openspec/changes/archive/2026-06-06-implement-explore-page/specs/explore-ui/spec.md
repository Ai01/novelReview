## ADDED Requirements

### Requirement: Explore page layout
The system SHALL render the explore page with a search box at the top, an infinite-scroll review feed in the middle, and a fixed bottom navigation bar.

#### Scenario: User lands on explore page
- **WHEN** user navigates to `/explore`
- **THEN** the page displays title "探索", a search input field, and the first page of review cards

#### Scenario: Desktop layout
- **WHEN** viewport width is 768px or wider
- **THEN** the content area is max-width 640px and centered horizontally, bottom nav is a fixed bar

#### Scenario: Mobile layout
- **WHEN** viewport width is less than 768px
- **THEN** the content area is full width and the bottom navigation bar is fixed to the viewport bottom

### Requirement: Infinite scroll review feed
The system SHALL load review cards incrementally as the user scrolls down using cursor-based pagination.

#### Scenario: Initial load
- **WHEN** explore page mounts
- **THEN** it fetches `GET /api/v1/explore?cursor=0&limit=20` and renders the returned items as review cards

#### Scenario: Scroll to bottom triggers next page
- **WHEN** user scrolls near the bottom of the list (sentinel element enters viewport)
- **THEN** it fetches the next page using the `next_cursor` from the previous response and appends new cards

#### Scenario: No more data
- **WHEN** the API returns `has_more: false`
- **THEN** no further requests are made and a "没有更多了" indicator is shown at the bottom

### Requirement: Review card display
Each review card SHALL display the book name, book author, book cover, reviewer username, reviewer avatar, review content, like count, and review time.

#### Scenario: Render standard review card
- **WHEN** a review item is rendered
- **THEN** it shows book cover image, book title, book author, user avatar, username, review text, like count, and relative time (e.g., "3小时前")

#### Scenario: Card click navigation
- **WHEN** user taps/clicks a review card
- **THEN** the app navigates to `/book/:bookId` (placeholder for future detail page)

### Requirement: Search functionality
The system SHALL allow users to search books and reviews by keyword via the search box at the top of the explore page.

#### Scenario: User performs a search
- **WHEN** user types a keyword in the search box and submits
- **THEN** the review feed is replaced with results from `GET /api/v1/explore/search?q=<keyword>&cursor=0&limit=20`

#### Scenario: Clear search
- **WHEN** user clears the search input
- **THEN** the explore page reverts to the default feed

#### Scenario: Empty search results
- **WHEN** a search returns zero results
- **THEN** an empty state message "暂无搜索结果" is displayed

### Requirement: Bottom navigation bar
The system SHALL render a fixed bottom navigation bar with five tabs: 探索, 书库, +, 书单, 个人.

#### Scenario: Navigation tabs display
- **WHEN** explore page is rendered
- **THEN** bottom bar shows five icon tabs: 探索 (active), 书库, +, 书单, 个人

#### Scenario: Tab switching
- **WHEN** user clicks a tab other than 探索
- **THEN** the app navigates to the corresponding page route (`/library`, `/booklists`, `/profile`); clicking + opens a bottom sheet modal

### Requirement: Loading and error states
The system SHALL handle loading, error, and empty states gracefully.

#### Scenario: Loading state
- **WHEN** data is being fetched
- **THEN** a loading spinner or skeleton placeholder is displayed

#### Scenario: Network error
- **WHEN** the API request fails
- **THEN** an error message "加载失败，请重试" is displayed with a retry button

#### Scenario: Empty feed (default)
- **WHEN** the explore API returns zero items on initial load
- **THEN** an empty state message "暂无内容" is displayed
