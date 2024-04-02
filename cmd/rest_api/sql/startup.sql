drop table if exists users;
CREATE TABLE IF NOT EXISTS users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    user_age INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
