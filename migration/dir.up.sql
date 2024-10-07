\c hw6;

CREATE TABLE IF NOT EXISTS users (
                                     Id SERIAL PRIMARY KEY,
                                     Username VARCHAR NOT NULL UNIQUE,
                                     Email VARCHAR NOT NULL UNIQUE,
                                     HashPassword VARCHAR NOT NULL UNIQUE
);

GRANT ALL PRIVILEGES ON SCHEMA public TO test;

GRANT ALL PRIVILEGES ON TABLE users TO test;