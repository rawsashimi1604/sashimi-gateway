CREATE DATABASE admin_api;

\c admin_api;

CREATE TABLE service (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    target_url TEXT UNIQUE NOT NULL,
    path TEXT UNIQUE NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE route (
    id SERIAL PRIMARY KEY,
    service_id INT REFERENCES service(id) NOT NULL,
    method TEXT NOT NULL,
    path TEXT NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

INSERT INTO service(name, target_url, path, description, created_at, updated_at)
VALUES 
    ('Salmon', 'http://localhost:8081', 'salmon', 'The salmon microservice used to learn how to create a golang api gateway infrastructure.', current_timestamp, current_timestamp),
    ('Tuna', 'http://localhost:8082', 'tuna', 'The tuna microservice used to learn how to create a golang api gateway infrastructure.', current_timestamp, current_timestamp);

INSERT INTO route(service_id, method, path, description, created_at, updated_at)
VALUES 
    (1, 'GET', '/', 'Get all salmon dishes', current_timestamp, current_timestamp),
    (1, 'POST', '/', 'Add salmon dish', current_timestamp, current_timestamp),
    (2, 'GET', '/', 'Get all tuna dishes', current_timestamp, current_timestamp),
    (1, 'GET', '/test', 'Test salmon service', current_timestamp, current_timestamp),
    (1, 'GET', '/:id', 'Get salmon dish by id', current_timestamp, current_timestamp);
