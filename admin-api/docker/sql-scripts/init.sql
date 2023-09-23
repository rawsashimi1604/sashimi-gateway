CREATE DATABASE admin_api;

\c admin_api;

CREATE TABLE service (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    target_url TEXT UNIQUE NOT NULL,
    path TEXT UNIQUE NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    health_check_enabled BOOLEAN NOT NULL,
    health TEXT NOT NULL
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

CREATE TABLE api_request (
    id TEXT PRIMARY KEY,
    service_id INT REFERENCES service(id) NOT NULL,
    route_id INT REFERENCES route(id) NOT NULL,
    path TEXT NOT NULL,
    method TEXT NOT NULL,
    time TIMESTAMP NOT NULL,
    code INT NOT NULL,
    duration INT NOT NULL
);

CREATE TABLE consumer (
    id TEXT PRIMARY KEY,
    username TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE jwt_credentials (
    id TEXT PRIMARY KEY,
    key TEXT NOT NULL,
    secret TEXT NOT NULL,
    name TEXT NOT NULL,
    consumer_id TEXT REFERENCES consumer(id) NOT NULL,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE consumers_has_services (
    consumer_id TEXT REFERENCES consumer(id) NOT NULL,
    service_id INT REFERENCES service(id) NOT NULL
);


INSERT INTO service(name, target_url, path, description, created_at, updated_at, health_check_enabled, health)
VALUES 
    ('Salmon', 'http://localhost:8081', 'salmon', 'The salmon microservice used to learn how to create a golang api gateway infrastructure.', current_timestamp, current_timestamp, true, 'startup'),
    ('Tuna', 'http://localhost:8082', 'tuna', 'The tuna microservice used to learn how to create a golang api gateway infrastructure.', current_timestamp, current_timestamp, true, 'startup');

INSERT INTO route(service_id, method, path, description, created_at, updated_at)
VALUES 
    (1, 'GET', '/', 'Get all salmon dishes', current_timestamp, current_timestamp),
    (1, 'POST', '/', 'Add salmon dish', current_timestamp, current_timestamp),
    (2, 'GET', '/', 'Get all tuna dishes', current_timestamp, current_timestamp),
    (1, 'GET', '/test', 'Test salmon service', current_timestamp, current_timestamp),
    (1, 'GET', '/:id', 'Get salmon dish by id', current_timestamp, current_timestamp),
    (1, 'GET', '/healthz', 'Salmon Healthz', current_timestamp, current_timestamp),
    (2, 'GET', '/healthz', 'Tuna Healthz', current_timestamp, current_timestamp);

INSERT INTO api_request(id, service_id, route_id, path, method, time, code, duration)
VALUES 
    (gen_random_uuid(), 1, 1, '/salmon', 'GET', current_timestamp, 200, 5),
    (gen_random_uuid(), 2, 3, '/tuna', 'GET', current_timestamp, 200, 10);

INSERT INTO consumer(id, username, created_at, updated_at)
VALUES
    ('e757b713-62e2-457e-8320-0e2fc4ac3a12', 'user1', current_timestamp, current_timestamp),

INSERT INTO jwt_credentials(id, key, secret, name, consumer_id, created_at)
VALUES
    (gen_random_uuid(), 'someKey', 'someSecret', 'user1Auth', 'e757b713-62e2-457e-8320-0e2fc4ac3a12', current_timestamp);

INSERT INTO consumers_has_services(consumer_id, service_id)
VALUES 
    ('e757b713-62e2-457e-8320-0e2fc4ac3a12', 1);