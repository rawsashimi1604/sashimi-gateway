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
