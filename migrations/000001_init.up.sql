CREATE TABLE tasks (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    status TEXT NOT NULL,
    priority TEXT NOT NULL,
    due_date_unix BIGINT NOT NULL,
    completed_at_unix BIGINT NOT NULL,
    created_at_unix BIGINT NOT NULL,
    updated_at_unix BIGINT NOT NULL
);

CREATE INDEX idx_tasks_id ON tasks (id);
CREATE INDEX idx_tasks_title ON tasks (title);
CREATE INDEX idx_tasks_description ON tasks (description);