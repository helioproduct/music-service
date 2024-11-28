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

-- indexes
CREATE UNIQUE INDEX idx_groups_name ON "groups" (name);
CREATE INDEX idx_songs_group_id ON songs (group_id);
CREATE INDEX idx_songs_name ON songs (name);
CREATE INDEX idx_songs_release_date_group_id ON songs (release_date, group_id);


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
    E'Hey Jude, don''t make it bad.\nTake a sad song and make it better.\nRemember to let her into your heart,\nThen you can start to make it better.\n\nHey Jude, don''t be afraid.\nYou were made to go out and get her.\nThe minute you let her under your skin,\nThen you begin to make it better.\n\nAnd anytime you feel the pain, hey Jude, refrain,\nDon''t carry the world upon your shoulders.\nFor well you know that it''s a fool who plays it cool\nBy making his world a little colder.\n\nHey Jude, don''t let me down.\nYou have found her, now go and get her.\nRemember to let her into your heart,\nThen you can start to make it better.\n\nSo let it out and let it in, hey Jude, begin,\nYou''re waiting for someone to perform with.\nAnd don''t you know that it''s just you, hey Jude, you''ll do,\nThe movement you need is on your shoulder.\n\nHey Jude, don''t make it bad.\nTake a sad song and make it better.\nRemember to let her under your skin,\nThen you''ll begin to make it\nBetter better better better better better, oh.\n\nNa na na nananana, nannana, hey Jude...\n(repeat X number of times, fade)', 
     'https://example.com/hey-jude', 
     (SELECT id FROM "groups" WHERE name = 'The Beatles')),
    
    ('Breathe (In The Air)', '1973-03-01', 
    E'[Instrumental Intro]\n\n[Verse 1]\nBreathe, breathe in the air\nDon''t be afraid to care\nLeave, but don''t leave me\nLook around and choose your own ground\n\n[Chorus]\nFor long you live and high you fly\nAnd smiles you''ll give and tears you''ll cry\nAnd all you touch and all you see\nIs all your life will ever be\n\n[Verse 2]\nRun, rabbit, run\nDig that hole, forget the sun\nAnd when at last the work is done\nDon''t sit down, it''s time to dig another one\n\n[Chorus]\nFor long you live and high you fly\nBut only if you ride the tide\nAnd balanced on the biggest wave\nYou race towards an early grave', 
     'https://pink.com/breathe-in-the-air', 
     (SELECT id FROM "groups" WHERE name = 'Pink Floyd')),

    ('Stairway to Heaven', '1971-11-08', 
    E'[Instrumental Intro]\n\n[Verse 1]\nThere''s a lady who''s sure all that glitters is gold\nAnd she''s buying a stairway to Heaven\nWhen she gets there she knows if the stores are all closed\nWith a word she can get what she came for\nOoh-ooh, ooh-ooh, and she''s buying a stairway to Heaven\n\n[Verse 2]\nThere''s a sign on the wall, but she wants to be sure\n''Cause you know sometimes words have two meanings\nIn a tree by the brook, there''s a songbird who sings\nSometimes all of our thoughts are misgiven\n\n[Chorus]\nOoh, it makes me wonder\nOoh, makes me wonder\n\n[Verse 3]\nThere''s a feeling I get when I look to the West\nAnd my spirit is crying for leaving\nIn my thoughts I have seen rings of smoke through the trees\nAnd the voices of those who stand looking\n\n[Chorus]\nOoh, it makes me wonder\nOoh, really makes me wonder\n\n[Verse 4]\nAnd it''s whispered that soon if we all call the tune\nThen the piper will lead us to reason\nAnd a new day will dawn for those who stand long\nAnd the forests will echo with laughter\n\n[Interlude]\nOh-oh-oh-oh-woah\n\n[Verse 5]\nIf there''s a bustle in your hedgerow, don''t be alarmed now\nIt''s just a spring clean for the May queen\nYes, there', 
     'https://helio.com/stairway-to-heaven', 
     (SELECT id FROM "groups" WHERE name = 'Led Zeppelin')),

    ('Rape Me', '1993-09-21', 
    E'[Verse 1]\nRape me\nRape me, my friend\nRape me\nRape me again\n\n[Chorus]\nI''m not the only one, ah-ah\nI''m not the only one, ah-ah\nI''m not the only one, ah-ah\nI''m not the only one\n\n[Verse 2]\nHate me\nDo it and do it again\nWaste me\nRape me, my friend\n\n[Chorus]\nI''m not the only one, ah-ah\nI''m not the only one, ah-ah\nI''m not the only one, ah-ah\nI''m not the only one\n\n[Bridge]\nMy favorite inside source\nI''ll kiss your open sores\nI appreciate your concern\nYou''re gonna stink and burn\n\n[Verse 1]\nRape me\nRape me, my friend\nRape me\nRape me again\n\n[Chorus]\nI''m not the only one, ah-ah\nI''m not the only one, ah-ah\nI''m not the only one, ah-ah\nI''m not the only one\n\n[Outro]\nRape me! (Rape me!)\nRape me! (Rape me!)\nRape me! (Rape me!)\nRape me! (Rape me!)\nRape me! (Rape me!)\nRape me! (Rape me!)\nRape me!', 
     'https://genius.com/rape-me', 
     (SELECT id FROM "groups" WHERE name = 'Nirvana')),

    ('Karma Police', '1997-08-25', 
    E'[Verse 1]\nKarma police, arrest this man\nHe talks in maths\nHe buzzes like a fridge\nHe''s like a detuned radio\n\n[Verse 2]\nKarma police, arrest this girl\nHer Hitler hairdo is making me feel ill\nAnd we have crashed her party\n\n[Chorus]\nThis is what you get\nThis is what you get\nThis is what you get\nWhen you mess with us\n\n[Verse 3]\nKarma police, I''ve given all I can, it''s not enough\nI''ve given all I can, but we''re still on the payroll\n\n[Chorus]\nThis is what you get\nThis is what you get\nThis is what you get\nWhen you mess with us\n\n[Outro]\nFor a minute there\nI lost myself, I lost myself\nPhew, for a minute there\nI lost myself, I lost myself\nFor a minute there\nI lost myself, I lost myself\nPhew, for a minute there\nI lost myself, I lost myself', 
     'https://example.com/karma-police', 
     (SELECT id FROM "groups" WHERE name = 'Radiohead')),

    ('Sympathy for the Devil', '1968-12-06', 
    E'[Intro]\nYeow\nYeow\nYeow\n\n[Verse 1]\nPlease allow me to introduce myself\nI''m a man of wealth and taste\nI''ve been around for a long, long year\nStole many a man''s soul and faith\nAnd I was ''round when Jesus Christ\nHad his moment of doubt and pain\nMade damn sure that Pilate\nWashed his hands and sealed his fate\n\n[Chorus]\nPleased to meet you, hope you guess my name\nBut what''s puzzlin'' you is the nature of my game\n\n[Verse 2]\nStuck around St. Petersburg\nWhen I saw it was a time for a change\nKilled the Czar and his ministers\nAnastasia screamed in vain\nI rode a tank, held a general''s rank\nWhen the Blitzkrieg raged and the bodies stank\n\n[Chorus]\nPleased to meet you, hope you guess my name\nOh yeah\nAh, what''s puzzlin'' you is the nature of my game\nAww yeah\n\n[Verse 3]\nI watched with glee while your kings and queens\nFought for ten decades for the gods they made\nI shouted out, "Who killed the Kennedys?"\nWhen, after all, it was you and me\nLet me please introduce myself\nI''m a man of wealth and taste\nAnd I laid traps for troubadours\nWho get killed before they reach Bombay\n\n[Chorus]\nPleased to meet you, hope you guess my name\nOh yeah\nBut what''s puzzlin'' you is the nature of my game\nAww yeah\n(Uh, get down heavy!)\n\n[Guitar Solo]\n\n[Chorus]\nPleased to meet you, hope you''ll guess my name\nAww yeah\nBut what''s confusin'' you is just the nature of my game\nMmm yeah\n\n[Verse 4]\nJust as every cop is a criminal\nAnd all the sinners saints\nAs heads is tails, just call me Lucifer\n''Cause I''m in need of some restraint\nSo if you meet me, have some courtesy\nHave some sympathy and some taste\nUse all your well-learned politesse\nOr I''ll lay your soul to waste\nMmm, yeah\n\n[Chorus]\nPleased to meet you, hope you guessed my name\nMmm, yeah\nBut what''s puzzlin'' you is the nature of my game\nMmm, mean it, get down\n\n[Outro w/Guitar Solo]\nWoo hoo!\nAww yeah\nGet on down\nOh yeah\nBa bum bum, ba ba bum, ba ba bum, ba ba bum, ba ba bum, ba ba bum, ba ba bum, ba ba bum, ba ba bum, yeah\nAh yeah\nTell me, baby, what''s my name?\nTell me, honey, can you guess my name?\nTell me, baby, what''s my name?\nI''ll tell you one time, you''re to blame\nOoo hoo, ooo hoo, oooo hoo\nAll right\nOoo hoo hoo, ooo hoo hoo, ooo hoo hoo\nAh yeah\nOoo hoo hoo, ooo hoo hoo\nAh yes, what''s my name?\nTell me, baby, ah, what''s my name?\nTell me, sweetie, what''s my name?\nOoo hoo hoo, ooo hoo hoo, ooo hoo hoo, ooo hoo hoo\nOoo hoo hoo, ooo hoo hoo, ooo hoo hoo\nAh yeah\nWhat''s my name...', 
     'https://musicly.com/sympathy-for-the-devil', 
     (SELECT id FROM "groups" WHERE name = 'The Rolling Stones')),

    ('Incredible', '2008-08-08', 
    E'[Chorus: Madonna]\nJust one of those things\nWhen everything goes incredible\nAnd all is beautiful\n(Can''t, can''t, can''t get my head around it, I need to think about it)\n(Can''t get my head around, I-I need to think about it)\nAnd all of those things that used to get you down now have no effect at all\n''Cause life is beautiful\n(Can''t, can''t, can''t get my head around it, I need to think about it)\n(Can''t get my head around, I-I need to think about it)\n\n[Verse 1: Madonna]\nRemembering the very first time\nYou caught that someone special''s eye\nAnd all of your care dropped\nAnd all of the world just stopped\n\n[Pre-Chorus: Madonna]\n(I hope)\nI wanna go back to then\nGotta figure out how, gotta remember where\nI felt it (Hmm)\nIt thrilled me (Hmm)\nI want it (Hmm)\nTo fill me\n\n[Chorus: Madonna]\nJust one of those things\nWhen everything goes incredible\nAnd all is beautiful\n(Can''t, can''t, can''t get my head around it, I need to think about it)\n(Can''t get my head around, I-I need to think about it)\nAnd all of those things that used to get you down don''t have no effect at all\n''Cause life is beautiful',
     'https://example.com/incredible', 
     (SELECT id FROM "groups" WHERE name = 'Madonna'));
