CREATE DATABASE admin_api;

\c admin_api;

CREATE TABLE api_method (
    id SERIAL PRIMARY KEY,
    method TEXT NOT NULL
);

INSERT INTO api_method (method)
VALUES
    ('GET'), ('POST'), ('PUT'), ('DELETE');

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
    path TEXT NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE route_method (
    route_id INT NOT NULL,
    method_id INT NOT NULL,
    PRIMARY KEY (route_id, method_id),
    FOREIGN KEY (route_id) REFERENCES route (id) ON DELETE CASCADE,
    FOREIGN KEY (method_id) REFERENCES api_method (id) ON DELETE CASCADE
);

INSERT INTO service(name, target_url, path, description, created_at, updated_at)
VALUES 
    ('Salmon', 'http://localhost:8081', 'salmon', 'The salmon microservice used to learn how to create a golang api gateway infrastructure.', current_timestamp, current_timestamp),
    ('Tuna', 'http://localhost:8082', 'tuna', 'The tuna microservice used to learn how to create a golang api gateway infrastructure.', current_timestamp, current_timestamp);

INSERT INTO route(service_id, path, description, created_at, updated_at)
VALUES 
    (1, '/', 'Get all salmon dishes', current_timestamp, current_timestamp),
    (1, '/', 'Add salmon dish', current_timestamp, current_timestamp),
    (2, '/', 'Get all tuna dishes', current_timestamp, current_timestamp);

INSERT INTO route_method(route_id, method_id)
VALUES 
    (1, 1),
    (2, 2),
    (3, 1);

