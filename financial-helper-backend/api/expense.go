package api

import (
	db "FinancialHelper/db/sqlc"
	util "FinancialHelper/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

type AddExpenseRequest struct {
	WalletId       string `json:"wallet" binding:"required"`
	WalletName     string `json:"wallet_name" binding:"required"`
	WalletCurrency string `json:"wallet_currency" binding:"required"`
	Value          string `json:"price" binding:"required"`
	Name           string `json:"shop_name" binding:"required"`
	Category       string `json:"category" binding:"required"`
	Date           string `json:"date" binding:"required"`
}

type AddExpenseResponse struct {
	Id             int64       `json:"ID" binding:"required"`
	WalletId       int64       `json:"Wallet" binding:"required"`
	WalletName     string      `json:"WalletName" binding:"required"`
	WalletCurrency string      `json:"WalletCurrency" binding:"required"`
	Value          float64     `json:"Value" binding:"required"`
	Name           string      `json:"Name" binding:"required"`
	Category       interface{} `json:"Category" binding:"required"`
	Date           time.Time   `json:"Date" binding:"required"`
	CreatedAt      time.Time   `json:"CreatedAt" binding:"required"`
}

type DeleteExpenseRequest struct {
	Id int64 `json:"id" binding:"required"`
}

func (server *Server) addExpense(ctx *gin.Context) {
	var req AddExpenseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Println(errorResponse(err))
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	clientToken, err := ctx.Cookie("token")
	if err != nil {
		printCookieError(ctx, err)
		return
	}

	claims, errMsg := util.ValidateToken(clientToken)
	if errMsg != "" {
		log.Println(errorResponse(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
		ctx.Abort()
		return
	}
	owner := int64(claims.ID)
	walletId, _ := strconv.ParseInt(req.WalletId, 10, 64)
	value, _ := strconv.ParseFloat(req.Value, 10)
	date, _ := time.Parse("2006-01-02", req.Date)

	arg := db.CreateExpenseParams{
		Owner:      owner,
		Wallet:     walletId,
		WalletName: req.WalletName,
		Currency:   req.WalletCurrency,
		Value:      value,
		Name:       req.Name,
		Category:   db.ExpenseCategory(req.Category),
		Date:       date,
	}

	expense, err := server.store.CreateExpense(ctx, arg)
	if err != nil {
		printSqlError(ctx, err)
		return
	}

	rsp := AddExpenseResponse{
		Id:             expense.ID,
		WalletId:       expense.Wallet,
		WalletName:     expense.WalletName,
		WalletCurrency: expense.Currency,
		Value:          expense.Value,
		Name:           expense.Name,
		Category:       expense.Category,
		Date:           expense.Date,
		CreatedAt:      expense.CreatedAt,
	}

	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) getExpenses(ctx *gin.Context) {
	clientToken, err := ctx.Cookie("token")
	if err != nil {
		printCookieError(ctx, err)
		return
	}

	claims, errMsg := util.ValidateToken(clientToken)
	if errMsg != "" {
		log.Println(errorResponse(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
		ctx.Abort()
		return
	}
	expenses, err := server.store.GetExpensesByOwner(ctx, int64(claims.ID))
	if err != nil {
		printSqlError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, expenses)
}

func (server *Server) deleteExpense(ctx *gin.Context) {
	var req DeleteExpenseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Println(errorResponse(err))
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	clientToken, err := ctx.Cookie("token")
	if err != nil {
		printCookieError(ctx, err)
		return
	}

	_, errMsg := util.ValidateToken(clientToken)
	if errMsg != "" {
		log.Println(errorResponse(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
		ctx.Abort()
		return
	}

	err = server.store.DeleteExpense(ctx, req.Id)
	if err != nil {
		printSqlError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Result": "Expense deleted"})
}
