CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT,
    email TEXT,
    gender TEXT,
    birth_date DATE
);

CREATE TABLE user_friends (
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    friend_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, friend_id),
    CHECK (user_id != friend_id)
);