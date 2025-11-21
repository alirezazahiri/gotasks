package taskhandler

type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	DueDateUnix *int64  `json:"due_date_unix"`
}

type UpdateTaskRequest struct {
	ID          string `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Priority    string `json:"priority"`
	DueDateUnix *int64  `json:"due_date_unix"`
}

type TaskResponse struct {
	ID              string `json:"id"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	Status          string `json:"status"`
	Priority        string `json:"priority"`
	DueDateUnix     int64  `json:"due_date_unix"`
	CompletedAtUnix *int64 `json:"completed_at_unix,omitempty"`
	CreatedAtUnix   int64  `json:"created_at_unix"`
	UpdatedAtUnix   int64  `json:"updated_at_unix"`
}
