# Sashimi API Gateway

<img width="1506" alt="image" src="https://github.com/rawsashimi1604/sashimi-gateway/assets/75880261/ba4e4282-1faa-404d-804a-cdca5ada0fbd">

## WORK IN PROGRESS

<p>A Lightweight API Gateway built in Golang to learn more about API infrastructure and microservices</p>

## Architecture

<img width="773" alt="image" src="https://github.com/rawsashimi1604/sashimi-gateway/assets/75880261/d8f3da11-2636-4b74-b59a-110c2e648642">

## Proposed cloud deployment architecture

<img width="506" alt="image" src="https://github.com/rawsashimi1604/sashimi-gateway/assets/75880261/55b60953-9782-4b8b-b0a4-5145a19b61cd">

## Features

- Request Analytics
- Authentication Provider
- Rate Limiting
- Admin API (API Onboarding and offboarding)
- Reverse Proxy
- API Health Checks
- GUI Application

## Functional Requirements

### Request Routing

- The API gateway should be able to route incoming HTTP(S) requests to specific services based on URL paths, HTTP methods, and other request attributes.
- The routing mechanism should support versioning so that different versions of a service can be addressed separately.

### Authentication & Authorization

- The gateway should be able to validate tokens (e.g., JWT) and decide if a request should proceed.
- Role-based access control (RBAC) might be implemented to restrict endpoints based on user roles.

### Rate Limiting

- To prevent abuse, the gateway should limit the number of requests a client can make in a given time frame. This could be global, per user, per IP, or per service.

### Logging & Monitoring

- All requests and responses should be logged for auditing and debugging purposes.
- Integrations with monitoring tools to track response times, error rates, and other important metrics.

### Service Discovery Integration

- As services can be dynamically added or removed, the API gateway should integrate with service discovery mechanisms to know where to route requests.

### Security

- DDoS protection: Recognizing and mitigating large-scale attempts to overwhelm the system.
- Support for HTTPS with SSL/TLS termination.
- Integration with Web Application Firewalls (WAF) to block malicious requests.

### Health Checks

- The gateway should implement periodic health checks to verify the status and availability of integrated API services.
- The health check mechanism should be configurable, allowing for adjustments to the frequency and conditions of checks.

## To do

- [ ] Requires big infrastructure refactor (code is messy and in a POC state) can consider clean architecture in golang or domain driven design
- [ ] Split the rev. proxy, cron jobs, admin api into different microservices, write docker-compose to start all containers easily
- [ ] Introduce redis as caching (lazy loading or write through)
- [ ] Introduce rate limiting into rev. proxy
- [ ] Incorporate credentials into manager GUI and test functionality
- [ ] Write unit tests, integration tests and e2e tests
- [ ] Write simple script to startup service with configured environment variables OR kubernetes/helm to deploy

## Multiple Golang applications

```
go work init
go work use ./admin-api ./salmon-api ./tuna-api
```
