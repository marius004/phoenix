CREATE TABLE IF NOT EXISTS problems (
    id          bigserial PRIMARY KEY,
    created_at  timestamp NOT NULL DEFAULT NOW(),

    name              text               NOT NULL UNIQUE,
    description       text               NOT NULL,
    short_description text               NOT NULL DEFAULT '',
    stream            stream_flag        NOT NULL DEFAULT 'console',
    author_id         bigint             NOT NULL DEFAULT 1,
    visible           boolean            NOT NULL DEFAULT false,
    difficulty        problem_difficulty NOT NULL DEFAULT 'easy',
    grade             problem_grade      NOT NULL,
    credits           text               NOT NULL DEFAULT '',

    time_limit   float  NOT NULL DEFAULT 0.1,
    memory_limit int    NOT NULL DEFAULT 65536, -- 64mb
    stack_limit  int    NOT NULL DEFAULT 16384, -- 16mb
    source_size  int    NOT NULL DEFAULT 10000, -- 10kb

    FOREIGN KEY(author_id)  REFERENCES users(id) ON DELETE SET DEFAULT
);