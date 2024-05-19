
# Email-Rate App

This application provides an API to get the current USD to UAH exchange rate and to subscribe emails for daily updates.

## Setup

1. Clone the repository
2. Go to [currencylayer.com](https://currencylayer.com/) and get the API key
3. Create `.env` file. Use `.env.example` as the example
4. Run `docker-compose up` to start the application

## Endpoints

- `GET /api/rate` - Get the current USD to UAH exchange rate
- `POST /api/subscribe` - Subscribe an email to daily rate updates

