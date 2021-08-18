CREATE TABLE IF NOT EXISTS blog_posts (
    id            bigserial   PRIMARY KEY,
    created_at    timestamp   NOT NULL DEFAULT NOW(),

    title          text  NOT NULL DEFAULT '',
    message        text  NOT NULL DEFAULT '',

    author_id bigint NOT NULL DEFAULT 1,

    FOREIGN KEY(author_id) REFERENCES users(id) ON DELETE SET DEFAULT
);