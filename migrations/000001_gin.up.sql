CREATE TABLE IF NOT EXISTS book{
    id serial PRIMARY KEY,
    title VARCHAR(100),
    author VARCHAR(50),
    publishedyear int
};