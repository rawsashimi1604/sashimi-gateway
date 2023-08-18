CREATE DATABASE admin_api;

\c admin_api;

CREATE TABLE test (
    id SERIAL PRIMARY KEY,
    test_text TEXT
);