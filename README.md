# Financial Transfers Application

A **Golang microservice** for **financial transactions** between accounts, using **Gin + GORM + PostgreSQL**.

---

## 📜 Features

- Create accounts with initial balances  
- Query account balances  
- Transfer funds between accounts (double-entry: credit & debit)  
- Transactional integrity with **row-level locking**  
- Dockerized for seamless local development

---

## API Specifications

### Create Account

- **POST** `/accounts`
- **Request Body:**
```json
{
  "account_id": 123,
  "balance": "100.23344"
}
```

### Get Account by ID

- **GET** `/accounts/123`
- **Response**
```json
{
  "account_id": 123,
  "balance": "100.23344"
}
```

### Create Transaction

This creates 2 transaction record, Credit & Debit, with same reference-id

- **POST** `/transactions`
- **Request Body:**
```json
{
  "source_account_id": 123,
  "destination_account_id": 432,
  "amount": "100.23344"
}
```

## Running the Application

### Prerequisites
- Docker
- Docker Compose

### Configure Environment
Edit your .env file:

```
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=exchange
DB_HOST=db
DB_PORT=5432
APP_PORT=8080
```

### Run with Docker Compose
Build and start:

Run the following command from the project root folder
```
docker-compose up --build
```

### Running Tests

Run all tests:

Run the following command from the project root folder
```
go test ./...
```

## Assumptions

- Single currency system (no FX)
- No authentication or authorization (internal trusted service)
- Transaction amounts are positive and valid decimals
- Uses GORM transactions with row-level locking to prevent double-spend
- Single-instance architecture (no distributed transactions)

## Potential Improvements

- Add authentication and role-based access control.
- Use UUIDs for account and transaction IDs for scalability.
- Erro codes can be used instead of human readable message (helps in internalization)
- The `amount` in create transaction payload could be float instead of string
- Senstive env variables (DB_PASSWORD) to be stored in key-valut
- Organize DB operations into a separate package `repository`
- Organize business logic in a separate package `services`
- Use DB change management tool (like liquibase) to mange schema changes