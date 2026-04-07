# Task Manager API - Stability Team Technical Test

This repository contains an improved version of the Task Manager API built with Go and Fiber. The system has been refactored for better stability, correctness, and performance.

## 1. Issues Found
During the initial audit of the original codebase, several critical issues were identified:
- **Race Conditions**: The original slice-based storage was not thread-safe, leading to potential data corruption or application crashes under concurrent requests.
- **Pointer-to-Loop-Variable Bug**: A common Go mistake in retrieval logic where task references could point to incorrect memory addresses.
- **Unstable Deletion**: Modifying a slice during iteration in `DeleteTask` could cause runtime panics or skipped elements.
- **Poor Input Validation**: The API accepted empty titles and invalid ID formats, which could lead to inconsistent data.
- **Non-Standard HTTP Responses**: Incorrect status codes (e.g., always returning 200 OK) and inconsistent JSON structures.

## 2. Fixes Implemented
The following fixes were applied to resolve the identified issues:
- **Thread-Safe Concurrent Storage**: Replaced the slice with a `map[int]models.Task` protected by a `sync.RWMutex` to ensure safe access from multiple goroutines.
- **Safe ID Retrieval**: Fixed the retrieval logic to return copies of data, eliminating pointer-related bugs.
- **Robust Deletion**: Implemented map-based deletion by key, which is inherently safer and more efficient than slice manipulation.
- **Input Sanitization**: Added checks to ensure IDs are valid integers and task titles are not empty or whitespace-only.
- **RESTful Error Handling**: Standardized HTTP status codes (400 for bad input, 404 for missing resources, 201 for creation).

## 3. Improvements Made
In addition to fixes, several architectural improvements were integrated:
- **Automated ID Management**: The store now manages task IDs internally, preventing conflicts and ensuring uniqueness.
- **Fiber Middleware**: Integrated `Logger` for observability and `Recover` to prevent the server from crashing on unexpected panics.
- **Unified Response Structure**: All API responses follow a consistent JSON format (wrapping data in a `data` field), making it easier for frontend integration.
- **Clean Code Architecture**: Separated concerns between `models`, `store`, and `handlers` for better maintainability.

---

## Setup & Running

### 1. Install dependencies
```bash
go mod tidy
```

### 2. Run the server
```bash
go run main.go
```

### 3. Server access
The server runs at [http://localhost:3000](http://localhost:3000).

## Available Endpoints

| Method | Endpoint   | Description         |
|--------|------------|---------------------|
| GET    | /tasks     | List all tasks      |
| GET    | /tasks/:id | Get a specific task |
| POST   | /tasks     | Create a new task   |
| DELETE | /tasks/:id | Delete a task       |

---

### Candidate Information
- **Name**: Havin Neo Dimas Nugraha
- **Position**: Full Stack Developer Intern
