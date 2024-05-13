<p style="display:flex;justify-content:center;align-items:center;" align="center">
  <a href="https://go.dev/" target="blank"><img src="https://go.dev/images/go-logo-white.svg" width="100" alt="Remix Logo"/></a>
</p>

## Bank Account Api

This API project simulates a simple bank account system using Golang's net/http library.

## Features

- **Golang:** The project leverages Golang for efficient backend development.

## Prerequisites

- **Golang:** Ensure you have Golang version 1.22.3 or higher installed. You can verify the installation by running `go version` in your terminal. 

## Running the application

1. **Clone the repository:**

   ```bash
   git clone https://github.com/julianopedraca/bank-account-api.git
   ```

2. **Navigate to the project directory:**

   ```bash
   cd bank-account-api
   ```

3. **Run the application:**

   ```bash
   go run .
   ```

This will start the API server.

## API Routes

The API provides several functionalities for managing bank accounts:

**1. Get Balance for Non-Existing Account:**

* **Method:** GET
* **Endpoint:** `/balance?account_id={account_id}`
* **Request:** Provides an `account_id` parameter in the query string.
* **Response:** If the account ID doesn't exist, returns a 404 Not Found status code with a message indicating the error.

**2. Create Account with Initial Balance:**

* **Method:** POST
* **Endpoint:** `/event`
* **Request Body:** JSON object with the following properties:
    * `type`: Set to `"deposit"`.
    * `destination`: The account ID (string) for the new account.
    * `amount`: The initial deposit amount (number).
* **Response:** On successful creation, returns a 201 Created status code with a JSON object containing the created account details:
    * `destination`: An object with properties:
        * `id`: The account ID (string).
        * `balance`: The current balance (number), which will be the initial deposit amount.

**3. Deposit into Existing Account:**

* **Method:** POST
* **Endpoint:** `/event`
* **Request Body:** JSON object with the following properties:
    * `type`: Set to `"deposit"`.
    * `destination`: The existing account ID (string).
    * `amount`: The amount to deposit (number).
* **Response:** On successful deposit, returns a 201 Created status code with a JSON object containing the updated account details:
    * `destination`: An object with properties:
        * `id`: The account ID (string).
        * `balance`: The updated balance (number) reflecting the deposit.

**4. Get Balance for Existing Account:**

* **Method:** GET
* **Endpoint:** `/balance?account_id={account_id}`
* **Request:** Provides an `account_id` parameter in the query string.
* **Response:** If the account exists, returns a 200 OK status code with the current balance (number) for the account.

**5. Withdraw from Non-Existing Account:**

* **Method:** POST
* **Endpoint:** `/event`
* **Request Body:** JSON object with the following properties:
    * `type`: Set to `"withdraw"`.
    * `origin`: The non-existing account ID (string).
    * `amount`: The amount to withdraw (number).
* **Response:** Since the account doesn't exist, returns a 404 Not Found status code with a message indicating the error.

**6. Withdraw from Existing Account:**

* **Method:** POST
* **Endpoint:** `/event`
* **Request Body:** JSON object with the following properties:
    * `type`: Set to `"withdraw"`.
    * `origin`: The existing account ID (string).
    * `amount`: The amount to withdraw (number).
* **Response:** On successful withdrawal, returns a 201 Created status code with a JSON object containing the updated account details:
    * `origin`: An object with properties:
        * `id`: The account ID (string).
        * `balance`: The updated balance (number) reflecting the withdrawal.

**7. Transfer from Existing Account:**

* **Method:** POST
* **Endpoint:** `/event`
* **Request Body:** JSON object with the following properties:
    * `type`: Set to `"transfer"`.
    * `origin`: The existing account ID (string) from which to transfer funds.
    * `amount`: The amount to transfer (number).
    * `destination`: The destination account ID (string) to receive the transferred funds.
* **Response:** On successful transfer, returns a 201 Created status code with a JSON object containing details