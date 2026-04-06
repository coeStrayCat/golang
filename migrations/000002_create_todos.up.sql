CREATE TYPE todo_status AS ENUM ('todo', 'progress', 'done');

CREATE TABLE todos (
    id          BIGSERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    description TEXT,
    status      todo_status NOT NULL DEFAULT 'todo',
    created_at  TIMESTAMPTZ DEFAULT NOW(),
    updated_at  TIMESTAMPTZ DEFAULT NOW()
);