CREATE TABLE IF NOT EXISTS tests (
    id bigserial PRIMARY KEY,
    created_at timestamp NOT NULL DEFAULT NOW(),

    problem_id int NOT NULL REFERENCES problems(id) ON DELETE CASCADE,
    score int NOT NULL
);