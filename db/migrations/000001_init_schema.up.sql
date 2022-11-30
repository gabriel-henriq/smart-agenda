CREATE TABLE IF NOT EXISTS tablets
(
    id           SERIAL PRIMARY KEY NOT NULL,
    name         TEXT NOT NULL,
    label_color  TEXT NOT NULL,
    created_at   TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS professors
(
    id           SERIAL PRIMARY KEY NOT NULL,
    name         TEXT NOT NULL,
    label_color  TEXT NOT NULL,
    created_at   TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS rooms
(
    id           SERIAL PRIMARY KEY NOT NULL,
    name         TEXT NOT NULL,
    label_color  TEXT NOT NULL,
    created_at   TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS aulas
(
    id           SERIAL PRIMARY KEY NOT NULL,
    tablet_id    INT NOT NULL,
    professor_id INT NOT NULL,
    room_id      INT NOT NULL,
    student_name TEXT NOT NULL,
    observation  TEXT NOT NULL,
    meet_start   TIMESTAMP NOT NULL,
    meet_end     TIMESTAMP NOT NULL,
    created_at   TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (tablet_id) REFERENCES tablets (id) ON DELETE CASCADE,
    FOREIGN KEY (professor_id) REFERENCES professors (id) ON DELETE CASCADE,
    FOREIGN KEY (room_id) REFERENCES rooms (id) ON DELETE CASCADE
);
