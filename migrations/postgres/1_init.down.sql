CREATE TABLE Group (
    id SERIAL PRIMARY KEY,     
    name TEXT NOT NULL UNIQUE  
);

CREATE TABLE Song (
    id SERIAL PRIMARY KEY,         
    release_date DATE NOT NULL,    
    text TEXT NOT NULL,            
    link TEXT NOT NULL,            
    group_id INT NOT NULL,         
    CONSTRAINT fk_group FOREIGN KEY (group_id) REFERENCES GroupDetail (id)
        ON DELETE CASCADE          
);


