package main

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var mtx sync.Mutex

// Transaction is for deposit or withdrawal
type Transaction struct {
	Id     int     `json:"id"`     //transaction id
	Amount float64 `json:"amount"` //transaction amount
	Type   string  `json:"Type"`   //deposit or withdrawal transaction
}

var transactions []Transaction //stores transaction histry
var balance float64            //stores balance

// add new transactions
func addNewTransaction(c *gin.Context) {
	var t Transaction

	//parsing json request body
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction format"})
		return
	}

	mtx.Lock()
	t.Id = len(transactions) + 1 //assign transaction Id

	//deposit or withdrawal cases handling
	switch t.Type {
	case "deposit": //deposit case handling
		balance += t.Amount
	case "withdrawal": //withdrawal case handeling
		if t.Amount > balance {
			mtx.Unlock()
			c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient balance"})
			return
		}
		balance -= t.Amount
	default: //Other than deposit and withdrawal case handling which generates error
		mtx.Unlock()
		c.JSON(http.StatusBadRequest, gin.H{"error": "Transaction type must be deposit or withdrawal only"})
		return
	}

	transactions = append(transactions, t) //adding transaction to the list
	mtx.Unlock()

	c.JSON(http.StatusOK, gin.H{"message": "Transaction successfully done", "transaction": t}) //FOR transaction successfully
}

// returns the current account balance
func getCurrentBalance(c *gin.Context) {
	mtx.Lock()
	defer mtx.Unlock()
	c.JSON(http.StatusOK, gin.H{"balance": balance}) //successfully request served
}

// returns the list of all transactions
func getAllTransactions(c *gin.Context) {
	mtx.Lock()
	defer mtx.Unlock()
	c.JSON(http.StatusOK, gin.H{"transactions": transactions}) //successfully request served
}

// setting up web server
func main() {
	r := gin.Default() //creates gin router

	r.POST("/transaction", addNewTransaction)  //Creates new transaction
	r.GET("/balance", getCurrentBalance)       //get current balance
	r.GET("/transactions", getAllTransactions) //lists all transactions

	r.Run(":8080") //starting server on 8080 port
}
