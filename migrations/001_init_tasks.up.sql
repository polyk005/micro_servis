CREATE TABLE tasks (
    id           UUID PRIMARY KEY,
    type         TEXT NOT NULL,
    status       TEXT NOT NULL CHECK (status IN ('pending', 'completed', 'failed')),
    params       JSONB,
    result       TEXT,
    created_at   TIMESTAMP NOT NULL DEFAULT NOW(),
    completed_at TIMESTAMP
);