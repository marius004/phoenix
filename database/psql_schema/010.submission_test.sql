CREATE TABLE IF NOT EXISTS submission_tests (
    id              bigserial   PRIMARY KEY,
    evaluated_at    timestamp   NOT NULL DEFAULT NOW(),

    score          int         NOT NULL DEFAULT 0,
    
    time           float       NOT NULL DEFAULT 0,
    memory         int         NOT NULL DEFAULT 0,
    
    message         text        NOT NULL DEFAULT '',
    exit_code       int         NOT NULL DEFAULT 0,

    submission_id   int         NOT NULL REFERENCES submissions(id) ON DELETE CASCADE,
    test_id         int         NOT NULL REFERENCES tests(id) ON DELETE CASCADE
);