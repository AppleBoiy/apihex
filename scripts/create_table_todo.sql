CREATE TABLE todos
(
    id         INTEGER PRIMARY KEY,
    title      TEXT,
    done       INTEGER,
    created_at TEXT, -- ISO 8601 format: YYYY-MM-DD HH:MM:SS.SSS
    updated_at TEXT,
    deleted_at TEXT
);