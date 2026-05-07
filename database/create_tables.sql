
CREATE TABLE IF NOT EXISTS music (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    mid TEXT NOT NULL UNIQUE,
    author TEXT,
    name TEXT NOT NULL,
    href TEXT,
    cover TEXT,
    lyc TEXT,
    music_url TEXT,
    local_path TEXT,
    duration_sec INTEGER DEFAULT 0,
    desc TEXT,

    status INTEGER DEFAULT 0,

    create_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    update_time DATETIME DEFAULT CURRENT_TIMESTAMP
);