# tinyledger
# Assumption
**Ledger Dataset Arrangement:** The API manages entities like accounts and transactions. There may be several accounts associated with each transaction, and each account contains a providing balance.
**Data Preservation:** The ledger can either save information in RAM (for temporary instances) or send it to persistent storage like a databases for the ledger state is preserved throughout the server restarts.
**Concurrency**: The API can manage basic concurrency scenarios like simultaneous requests for account or transaction data.
**Time Stamps:** Transactions can be timestamped and reflect the time when they were created. 
**Authentication**: The API assumes a minimal level of authentication, such as token based, or API key based authentication, for granting access to the ledger system.
**Scalability**: The API coverage for moderate volume of transactions and users efficiently. Depending on your needs, the system can be expanded to support higher loads.

# Installation
**Clone the repository:**
  git clone https://github.com/jain-bhawna/tinyledger/tinyledger.git
**Navigate to the folder**
  cd tinyledger 
**How to Run:**
Install Go.
Install dependencies:
  go mod init tinyledger
  go get github.com/gin-gonic/gin
Run the server:
  go run main.go

# Example requests using postman tool
**Transaction deposit**
Request type: POST
"Content-type" in headers: "application/json"
URL: http://localhost:8080/transaction
Body: [Select raw and Json]
{
    "amount": 100,
    "Type": "deposit"
}
Response:
{
    "message": "Transaction successfully done",
    "transaction": {
        "id": 3,
        "amount": 100,
        "Type": "deposit"
    }
}
**Transaction withdraw**
Request type: POST
"Content-type" in headers: "application/json"
URL: http://localhost:8080/transaction
Body: [Select raw and Json]
{
    "amount": 250,
    "Type": "withdrawal"
}
Response:
{
    "message": "Transaction successfully done",
    "transaction": {
        "id": 4,
        "amount": 50,
        "Type": "withdrawal"
    }
}
**Check balance**
Request type: GET
URL: http://localhost:8080/balance
Response:
{
    "balance": 100
}
**Check all transactions**
Request type: GET
URL: http://localhost:8080/transactions
Response:
{
    "transactions": [
        {
            "id": 1,
            "amount": 100,
            "Type": "deposit"
        },
        {
            "id": 2,
            "amount": 50,
            "Type": "withdrawal"
        },
        {
            "id": 3,
            "amount": 100,
            "Type": "deposit"
        },
        {
            "id": 4,
            "amount": 50,
            "Type": "withdrawal"
        }
    ]
}
