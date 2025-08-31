PRAGMA foreign_keys = ON;

CREATE TABLE users (
    id          TEXT PRIMARY KEY,
    username    TEXT NOT NULL UNIQUE,
    created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE bots (
    id          TEXT PRIMARY KEY CHECK (id <> ''),
    author      TEXT NOT NULL,
    name        TEXT NOT NULL,
    description TEXT NOT NULL,
    status      TEXT NOT NULL DEFAULT 'pending', -- "pending", "approved", "denied"
    created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (author) REFERENCES users(id) ON DELETE CASCADE
);
