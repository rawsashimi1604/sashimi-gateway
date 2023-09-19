# Sashimi API Gateway

## WORK IN PROGRESS

<p>A Lightweight API Gateway built in Golang to learn more about API infrastructure and microservices</p>

## Architecture
<img width="773" alt="image" src="https://github.com/rawsashimi1604/sashimi-gateway/assets/75880261/d8f3da11-2636-4b74-b59a-110c2e648642">

## Proposed cloud deployment architecture
<img width="506" alt="image" src="https://github.com/rawsashimi1604/sashimi-gateway/assets/75880261/55b60953-9782-4b8b-b0a4-5145a19b61cd">

## Features 
- Analytics
- Authentication
- Rate Limiting
- Admin API (API Onboarding and offboarding)
- Reverse Proxy
- GUI Application 

## Multiple Golang applications
```
go work init
go work use ./admin-api ./salmon-api ./tuna-api
```
