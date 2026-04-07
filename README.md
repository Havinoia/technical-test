# Task Manager API - Stability Team Technical Test

This repository contains an improved version of the Task Manager API built with Go and Fiber. The system has been refactored for better stability, correctness, and performance.

## Key Improvements & Fixes

### 1. Stability & Correctness Fixes
- **Thread-Safe Storage**: Replaced the original slice-based storage with a `map[int]models.Task` protected by a `sync.RWMutex`. This ensures the API remains stable and data-consistent under concurrent load.
- **Fixed Logical Bugs**: 
    - Resolved a common Go "pointer-to-loop-variable" bug in `GetTaskByID` that could lead to returning incorrect task references.
    - Fixed a slice manipulation bug in `DeleteTask` where modifying a slice during iteration could cause runtime errors or skipped elements.
- **Robust Error Handling**: Added validation for task IDs (ensuring they are valid integers) and implemented explicit error checks for all storage operations.
- **RESTful HTTP Standards**: Updated handlers to return correct HTTP status codes:
    - `404 Not Found` for missing resources.
    - `400 Bad Request` for invalid input (e.g., non-numeric IDs, empty titles).
    - `201 Created` for successfully added tasks.

### 2. General Enhancements
- **Input Validation**: Implemented mandatory title checks for new tasks (rejecting empty or whitespace-only titles).
- **Middleware Integration**: Added Fiber's `Logger` and `Recover` middlewares for enhanced request observability and application resilience against unexpected panics.
- **Standardized API Responses**: All responses now follow a consistent structure, wrapping data in a "data" field or providing clear "error" messages.
- **Automatic ID Management**: Task IDs are now managed internally by the store to prevent conflicts and ensure uniqueness.

## Setup & Running

1. **Install dependencies**:
   ```bash
   go mod tidy
   ```

2. **Run the server**:
   ```bash
   go run main.go
   ```

3. **Server access**:
   The server runs at [http://localhost:3000](http://localhost:3000).

## Available Endpoints

| Method | Endpoint      | Description           |
|--------|---------------|-----------------------|
| GET    | /tasks        | List all tasks        |
| GET    | /tasks/:id    | Get a specific task   |
| POST   | /tasks        | Create a new task     |
| DELETE | /tasks/:id    | Delete a task         |

## Technologies Used
- **Go** (v1.24+)
- **Fiber** (v2)
- **Standard sync package** for concurrency control
