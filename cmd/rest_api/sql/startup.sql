drop table if exists users;
CREATE TABLE IF NOT EXISTS users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    user_age INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- Insert statements for the users table
-- Will create a flag to load dummy data only when flag activated
INSERT INTO users (username, user_age, created_at, updated_at)
VALUES ('alice', 23, '2024-03-11 07:43:21', '2024-03-11 07:43:21');

INSERT INTO users (username, user_age, created_at, updated_at)
VALUES ('bob', 42, '2024-03-25 07:43:21', '2024-03-25 07:43:21');

INSERT INTO users (username, user_age, created_at, updated_at)
VALUES ('charlie', 30, '2024-03-20 07:43:21', '2024-03-20 07:43:21');

INSERT INTO users (username, user_age, created_at, updated_at)
VALUES ('dave', 28, '2024-03-15 07:43:21', '2024-03-15 07:43:21');

INSERT INTO users (username, user_age, created_at, updated_at)
VALUES ('eve', 37, '2024-03-29 07:43:21', '2024-03-29 07:43:21');

INSERT INTO users (username, user_age, created_at, updated_at)
VALUES ('frank', 55, '2024-03-18 07:43:21', '2024-03-18 07:43:21');

INSERT INTO users (username, user_age, created_at, updated_at)
VALUES ('grace', 33, '2024-03-29 07:43:21', '2024-03-29 07:43:21');

INSERT INTO users (username, user_age, created_at, updated_at)
VALUES ('helen', 47, '2024-03-14 07:43:21', '2024-03-14 07:43:21');

INSERT INTO users (username, user_age, created_at, updated_at)
VALUES ('ian', 25, '2024-03-22 07:43:21', '2024-03-22 07:43:21');

INSERT INTO users (username, user_age, created_at, updated_at)
VALUES ('jane', 31, '2024-03-27 07:43:21', '2024-03-27 07:43:21');


