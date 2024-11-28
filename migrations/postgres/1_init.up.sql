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


CREATE SEQUENCE IF NOT EXISTS groups_id_seq START 1;
ALTER TABLE groups ALTER COLUMN id SET DEFAULT nextval('groups_id_seq');


---------------------
---Тестовые данные---
---------------------


INSERT INTO "groups" (name)
VALUES
    ('The Beatles'),
    ('Pink Floyd'),
    ('Led Zeppelin'),
    ('Nirvana'),
    ('Radiohead'),
    ('The Rolling Stones'),
    ('Madonna');

INSERT INTO songs (name, release_date, lyrics, link, group_id)
VALUES
    ('Hey Jude', '1968-08-30', 
     E'Hey Jude, don''t make it bad.\nTake a sad song and make it better...', 
     'https://example.com/hey-jude', 
     (SELECT id FROM "groups" WHERE name = 'The Beatles')),
    
    ('Breathe (In The Air)', '1973-03-01', 
     E'[Instrumental Intro]\n\n[Verse 1]\nBreathe, breathe in the air...', 
     'https://pink.com/breathe-in-the-air', 
     (SELECT id FROM "groups" WHERE name = 'Pink Floyd')),

    ('Stairway to Heaven', '1971-11-08', 
     E'[Instrumental Intro]\n\n[Verse 1]\nThere''s a lady who''s sure...', 
     'https://helio.com/stairway-to-heaven', 
     (SELECT id FROM "groups" WHERE name = 'Led Zeppelin')),

    ('Rape Me', '1993-09-21', 
     E'[Verse 1]\nRape me\nRape me, my friend...', 
     'https://genius.com/rape-me', 
     (SELECT id FROM "groups" WHERE name = 'Nirvana')),

    ('Karma Police', '1997-08-25', 
     E'[Verse 1]\nKarma police, arrest this man...', 
     'https://example.com/karma-police', 
     (SELECT id FROM "groups" WHERE name = 'Radiohead')),

    ('Sympathy for the Devil', '1968-12-06', 
     E'[Intro]\nYeow\nYeow\nYeow...', 
     'https://musicly.com/sympathy-for-the-devil', 
     (SELECT id FROM "groups" WHERE name = 'The Rolling Stones')),

    ('Incredible', '2008-08-08', 
     E'[Chorus: Madonna]\nJust one of those things...', 
     'https://example.com/incredible', 
     (SELECT id FROM "groups" WHERE name = 'Madonna'));
