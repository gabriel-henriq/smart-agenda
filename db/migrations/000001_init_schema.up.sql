CREATE TABLE IF NOT EXISTS tablets
(
    id           SERIAL PRIMARY KEY NOT NULL,
    name         TEXT,
    created_at   TIMESTAMP DEFAULT NOW(),
    updated_at   TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS professors
(
    id           SERIAL PRIMARY KEY NOT NULL,
    name         TEXT,
    label_color  TEXT,
    created_at   TIMESTAMP DEFAULT NOW(),
    updated_at   TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS rooms
(
    id           SERIAL PRIMARY KEY NOT NULL,
    name         TEXT,
    label_color  TEXT,
    created_at   TIMESTAMP DEFAULT NOW(),
    updated_at   TIMESTAMP DEFAULT NOW()
);


CREATE TABLE IF NOT EXISTS aulas
(
    id           SERIAL PRIMARY KEY NOT NULL,
    tablet_id    INT,
    professor_id INT,
    room_id      INT,
    student_name TEXT,
    meet_start   TIMESTAMP,
    meet_end     TIMESTAMP,
    observation  TEXT,
    created_at   TIMESTAMP DEFAULT NOW(),
    updated_at   TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (tablet_id) REFERENCES tablets (id) ON DELETE CASCADE,
    FOREIGN KEY (professor_id) REFERENCES professors (id) ON DELETE CASCADE,
    FOREIGN KEY (room_id) REFERENCES rooms (id) ON DELETE CASCADE
);
