CREATE TABLE "groups" (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE songs (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    release_date DATE NOT NULL,
    lyrics TEXT NOT NULL,
    link TEXT NOT NULL,
    group_id INT REFERENCES groups(id) ON DELETE CASCADE
);
