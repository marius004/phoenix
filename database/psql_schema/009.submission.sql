CREATE TABLE IF NOT EXISTS submissions (
    id              bigserial   PRIMARY KEY,
    created_at      timestamp   NOT NULL DEFAULT NOW(),

    score           int         NOT NULL DEFAULT 0,
    lang            lang_type   NOT NULL,
    status          status_type NOT NULL DEFAULT 'waiting',
    compile_message text        NOT NULL DEFAULT '',
    compile_error   text        NOT NULL DEFAULT '',

    problem_id      int         NOT NULL REFERENCES problems(id) ON DELETE CASCADE,
    user_id         int         NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

ALTER TABLE submissions 
ADD COLUMN has_compile_error boolean;

ALTER TABLE submissions
ADD COLUMN source_code text NOT NULL DEFAULT '';