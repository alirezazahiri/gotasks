package taskhandler

type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	DueDateUnix *int64  `json:"due_date_unix"`
}

type ListTasksRequest struct {
	Page     int64 `json:"page" binding:"required"`
	PageSize int64 `json:"page_size" binding:"required"`
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

type ListTasksResponse struct {
	Tasks      []*TaskResponse `json:"tasks"`
	Total      int64           `json:"total"`
	Page       int64           `json:"page"`
	PageSize   int64           `json:"page_size"`
	TotalPages int64           `json:"total_pages"`
}