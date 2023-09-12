# Sashimi Gateway Admin API

## Building

- start docker container
  - load psql image
  - load sql files

## Environment file

put `.env` into root directory of `admin-api`

```
POSTGRES_URL=postgresql://postgres:password123@localhost:5432/admin_api
MANAGER_URL=http://localhost:5173

# Gateway metadata
SASHIMI_GATEWAY_NAME="Sushi Gateway"
SASHIMI_HOSTNAME="http://localhost:8080"
SASHIMI_TAGLINE="hello from sashimi gateway"

# Gateway config
# interval in seconds between each caching request
SASHIMI_LOCAL_PORT=8080
SASHIMI_REQUEST_INTERVAL=5

# JWT Auth for admin Api
SASHIMI_ADMIN_JWT_KEY=<generate your own 256 bit key, there is a function in utils file jwt_key.go to help with this>
```
