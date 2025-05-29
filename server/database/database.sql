-- 1. Users 👤​
CREATE TABLE IF NOT EXISTS users (
    id TEXT UNIQUE PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    age INTEGER NOT NULL,
    gender TEXT NOT NULL,
    firstName TEXT NOT NULL,
    lastName TEXT NOT NULL,
    password_hash TEXT NOT NULL
);