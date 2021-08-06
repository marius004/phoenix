CREATE TABLE IF NOT EXISTS submissions (
    id              bigserial   PRIMARY KEY,
    created_at      timestamp   NOT NULL DEFAULT NOW(),

    score           int         NOT NULL DEFAULT 0,
    lang            lang_type   NOT NULL,
    status          status_type NOT NULL DEFAULT 'waiting',
    message         text        NOT NULL DEFAULT '',
    
    source_code text NOT NULL DEFAULT '',
    has_compile_error bool NOT NULL DEFAULT false,

    problem_id      int         NOT NULL REFERENCES problems(id) ON DELETE CASCADE,
    user_id         int         NOT NULL REFERENCES users(id) ON DELETE CASCADE
);