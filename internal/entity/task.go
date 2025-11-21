package entity

type Task struct {
	ID          string `gorm:"column:id;primaryKey"`
	Title       string `gorm:"column:title"`
	Description string `gorm:"column:description"`
	Status      string `gorm:"column:status"`
	Priority    string `gorm:"column:priority"`
	DueDate     *int64 `gorm:"column:due_date_unix"`
	CompletedAt *int64 `gorm:"column:completed_at_unix"`
	CreatedAt   int64  `gorm:"column:created_at_unix"`
	UpdatedAt   int64  `gorm:"column:updated_at_unix"`
}

func (Task) TableName() string {
	return "tasks"
}
