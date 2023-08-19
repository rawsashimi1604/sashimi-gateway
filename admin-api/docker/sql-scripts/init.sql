CREATE DATABASE admin_api;

\c admin_api;

CREATE TABLE test (
    id SERIAL PRIMARY KEY,
    test_text TEXT
);

INSERT INTO test(test_text)
VALUES ('hello world from 1');

CREATE TABLE service (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    target_url TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);