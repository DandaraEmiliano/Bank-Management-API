## Bank Management API

## Overview

The Bank Management API enables the management of users, bank accounts, and secure financial transactions. It provides endpoints to handle user accounts, perform money transfers, and ensure authentication with JWT tokens.

## Features

- CRUD operations for users and accounts.
- Perform secure fund transfers between accounts.
- Authentication and authorization using JWT.
- Password hashing with bcrypt for enhanced security.

## Tech Stack

- Go (Golang)
- Gin (Web Framework)
- PostgreSQL (Database)
- GORM (ORM)
- JWT (Authentication)

## Installation

Prerequisites:
- Go installed (download here)
- PostgreSQL installed and running
- A PostgreSQL database created (bank)

## Clone the Repository

git clone https://github.com/your-repo/bank-management-api.git

cd bank-management-api

## Install Dependencies:

go mod tidy

## Run Migrations

The database schema will be created automatically on server startup using GORM. Ensure your database is running before proceeding.

## Start the Server

- go run main.go
- The API will be available at http://localhost:8080.

## Endpoints

/api/users
- GET: Retrieve all users.
- POST: Create a new user.
- DELETE: Delete a user by ID.
  
/api/accounts
- GET: Retrieve all accounts.
- POST: Create a new account.
- DELETE: Delete an account by ID.

/api/accounts/transfer
- POST: Transfer funds between accounts.

/api/login
- POST: Authenticate a user and return a JWT token.
