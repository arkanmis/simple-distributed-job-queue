# Simple Distributed Job Queue Simulation
a simple job queue system built by golang

---

## Part 1 — Job Queue Backend (GraphQL)

### Requirement:
-> Go >= 1.20

### How to Run:
-> on VsCode -> Run(Pick Go-Debug)
-> open localhost:58579/graphiql
-> should be look like this
![Graphiql](./docs/graphiql.jpg)

### Requirements
make sure you complete all the function stated at first you load graphiql page
1. SimultaneousCreateJob -> build a simultaneous job
2. SimulateUnstableJob -> special case for task "unstable-job" fails twice before passing
3. GetAllJobs -> get all jobs that already registered
4. GetJobById -> get job by the id that been created by enqueue job
5. GetAllJobStatus -> get the stats of all jobs that been processed

### Evaluation Criteria:
* **Correctness**: Job creation, execution, status updates are accurate.
* **Concurrency Safety**: Multiple jobs created/processed at once → no race, no corruption.
* **Idempotency Handling**: Same job/task with same ID or token → doesn't process twice.
* **Retry Logic**: Failing job retries up to N times with delay.
* **In-memory Safety**: Maps/lists used safely under concurrent access.
* **Code Quality**: Idiomatic Go, good naming, clean package layout.
* **Clean Architecture**: Separation of domain, repository, resolver, GraphQL models.
* **Performance Awareness**: Handles 50–100 concurrent jobs without crash or slowdown.
* **Logging and Debugging**: Logs meaningful events.
* **Graceful Failure Handling**: No panics; job failure doesn't crash system.

---

## Part 2 — HTMX Dashboard

### Overview
Build a server-rendered dashboard at `/jobqueue/dashboard` using [HTMX](https://htmx.org/). The backend handler lives in `delivery/htmx/`; the HTML templates live in `web/htmx/`. No JavaScript framework is needed — HTMX attributes drive all partial page updates.

### Routes

| Method | Path | Purpose |
|--------|------|---------|
| `GET` | `/jobqueue/dashboard` | Full HTML page shell |
| `GET` | `/jobqueue/dashboard/message` | Initial HTMX fragment entry point |

> Add more fragment endpoints under `/jobqueue/dashboard/...` as each section is built out.

---

### Backend (Go — `delivery/htmx/`)

Architecture rules:
- Handlers live in `delivery/htmx/` — the delivery layer only.
- Handlers that need data must call the service layer through the interfaces in `interface/`.
- No business logic inside handlers; handlers only bridge HTTP ↔ service.

#### Fragment Endpoints to Implement

| Endpoint | Description |
|----------|-------------|
| `POST /jobqueue/dashboard/jobs/create` | Runs `SimultaneousCreateJob` with Job1/Job2/Job3 from the form, returns updated job list fragment |
| `POST /jobqueue/dashboard/jobs/unstable` | Runs `SimulateUnstableJob`, returns updated job list fragment |
| `GET /jobqueue/dashboard/status` | Returns the status summary fragment (polled every 2 s) |
| `GET /jobqueue/dashboard/jobs` | Returns the full jobs table fragment (polled every 2 s) |
| `GET /jobqueue/dashboard/jobs/:id` | Returns the job detail fragment for a single job |

---

### Frontend (HTMX — `web/htmx/`)

#### 1. Action Bar
Three trigger buttons at the top of the page:

| Button | Fires | Notes |
|--------|-------|-------|
| **Create 3 Jobs** | `POST /jobqueue/dashboard/jobs/create` | Reads Job1/Job2/Job3 from the Variables Form |
| **Create Unstable Job** | `POST /jobqueue/dashboard/jobs/unstable` | Hard-coded `task: "unstable-job"` |
| **Refresh** *(optional)* | Re-polls `/jobqueue/dashboard/status` and `/jobqueue/dashboard/jobs` | Manual trigger |

#### 2. Variables Form
Inline form whose values feed the **Create 3 Jobs** action. Defaults must match `web/variables.json`:

| Field | Default |
|-------|---------|
| Job1 | `JobTest1` |
| Job2 | `JobTest2` |
| Job3 | `JobTest3` |

#### 3. Status Summary
Live badge/card row — auto-polled every 2 seconds via `hx-trigger="every 2s"` targeting `#status-summary`:

- **Pending** / **Running** / **Failed** / **Completed**

#### 4. Jobs Table
Table auto-polled every 2 seconds targeting `#jobs-table`:

| Column | Field |
|--------|-------|
| ID | `job.id` |
| Task | `job.task` |
| Status | `job.status` |
| Attempts | `job.attempts` |
| View | Button → loads Job Detail Panel for that row's ID |

#### 5. Job Detail Panel
Panel or side section targeting `#job-detail`:
- Triggered by clicking **View** in any table row, or by typing a job ID into a search input.
- Displays: `id`, `task`, `status`, `attempts`.
- Swapped in via `hx-get="/jobqueue/dashboard/jobs/:id"` without reloading the rest of the page.

### How to Access
1. Start the server (see **How to Run** above).
2. Open `http://localhost:58579/jobqueue/dashboard`.
3. Status summary and job list auto-refresh every 2 seconds via HTMX polling.

# Good Luck Guys
