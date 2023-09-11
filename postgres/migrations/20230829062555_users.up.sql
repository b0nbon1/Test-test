CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    age INT NOT NULL,
    email TEXT NOT NULL,
    phone TEXT NOT NULL,
    address TEXT NOT NULL,
    city TEXT NOT NULL
);

