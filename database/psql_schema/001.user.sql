CREATE TABLE IF NOT EXISTS users(
    id          bigserial PRIMARY KEY,
    created_at  timestamp NOT NULL DEFAULT NOW(),

    username text    NOT NULL UNIQUE,
    email    text    NOT NULL UNIQUE,
    password text    NOT NULL UNIQUE,
    bio      text    NOT NULL DEFAULT '',
    visible  boolean NOT NULL DEFAULT true,

    is_admin     boolean NOT NULL DEFAULT false,
    is_proposer  boolean NOT NULL DEFAULT false,
    is_banned    boolean NOT NULL DEFAULT false,

    verified_email              boolean NOT NULL DEFAULT false,
    email_verification_sent_at  timestamp DEFAULT NULL
);