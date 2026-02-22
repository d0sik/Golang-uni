CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255),
    age INT
);

INSERT INTO users (name, email, age) VALUES ('John Doe', 'john@gmail.com', 20);