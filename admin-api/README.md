# Sashimi Gateway Admin API

## Building

- start docker container
  - load psql image
  - load sql files

## Environment file

put `.env` into root directory of `admin-api`

```
POSTGRES_URL=postgresql://postgres:password123@localhost:5432/admin-api
MANAGER_URL=http://localhost:5173
```
