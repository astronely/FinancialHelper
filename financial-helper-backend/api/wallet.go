package api

import (
	db "FinancialHelper/db/sqlc"
	util "FinancialHelper/utils"
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

type NullInt64 struct {
	Val     int64
	IsValid bool
}

type AddWalletRequest struct {
	Name     string `json:"wallet_name" binding:"required"`
	Balance  string `json:"balance" binding:"required"`
	Currency string `json:"currency" binding:"required"`
}

type AddWalletResponse struct {
	Id        int64     `json:"id"`
	Name      string    `json:"wallet_name" binding:"required"`
	Balance   float64   `json:"balance" binding:"required"`
	Currency  string    `json:"currency" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"required"`
}

type DeleteWalletRequest struct {
	WalletName string `json:"name" binding:"required"`
}

type UpdateWalletRequest struct {
	Name     string `json:"name" binding:"required"`
	Value    string `json:"value" binding:"required"`
	Currency string `json:"currency" binding:"required"`
}

type UpdateWalletResponse struct {
	Id       int64   `json:"id"`
	Name     string  `json:"wallet_name"`
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
}

func (server *Server) addWallet(ctx *gin.Context) {
	var req AddWalletRequest
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

	balance, _ := strconv.ParseFloat(req.Balance, 10)
	arg := db.CreateWalletParams{
		Owner: int64(claims.ID),
		Name:  req.Name,
		Balance: sql.NullFloat64{
			Float64: balance,
			Valid:   true,
		},
		Currency: req.Currency,
	}

	wallet, err := server.store.CreateWallet(ctx, arg)
	if err != nil {
		printSqlError(ctx, err)
		return
	}

	rsp := AddWalletResponse{
		Id:        wallet.ID,
		Name:      wallet.Name,
		Balance:   wallet.Balance.Float64,
		Currency:  wallet.Currency,
		CreatedAt: wallet.CreatedAt,
	}

	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) getWallets(ctx *gin.Context) {
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

	wallets, err := server.store.GetWalletsByOwner(ctx, int64(claims.ID))
	if err != nil {
		printSqlError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, wallets)
}

func (server *Server) deleteWallet(ctx *gin.Context) {
	var req DeleteWalletRequest
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
	data := db.DeleteWalletParams{
		Owner: int64(claims.ID),
		Name:  req.WalletName,
	}
	log.Println(req.WalletName)
	err = server.store.DeleteWallet(ctx, data)
	if err != nil {
		printSqlError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Result": "Wallet deleted"})
}

func (server *Server) updateWallet(ctx *gin.Context) {
	var req UpdateWalletRequest

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

	balance, _ := strconv.ParseFloat(req.Value, 10)
	arg := db.UpdateWalletParams{
		Owner: int64(claims.ID),
		Name:  req.Name,
		Balance: sql.NullFloat64{
			Float64: balance,
			Valid:   true,
		},
	}

	updatedWallet, err := server.store.UpdateWallet(ctx, arg)
	if err != nil {
		printSqlError(ctx, err)
		return
	}
	rsp := UpdateWalletResponse{
		Id:       updatedWallet.ID,
		Name:     updatedWallet.Name,
		Balance:  updatedWallet.Balance.Float64,
		Currency: updatedWallet.Currency,
	}

	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) decreaseWalletBalance(ctx *gin.Context) {
	var req UpdateWalletRequest

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

	balance, _ := strconv.ParseFloat(req.Value, 10)
	log.Println(balance)
	arg := db.DecreaseWalletBalanceParams{
		Owner: int64(claims.ID),
		Name:  req.Name,
		Balance: sql.NullFloat64{
			Float64: balance,
			Valid:   true,
		},
	}

	log.Println(arg.Balance)
	updatedWallet, err := server.store.DecreaseWalletBalance(ctx, arg)
	if err != nil {
		printSqlError(ctx, err)
		return
	}
	rsp := UpdateWalletResponse{
		Id:       updatedWallet.ID,
		Name:     updatedWallet.Name,
		Balance:  updatedWallet.Balance.Float64,
		Currency: updatedWallet.Currency,
	}

	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) increaseWalletBalance(ctx *gin.Context) {
	var req UpdateWalletRequest

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

	balance, _ := strconv.ParseFloat(req.Value, 10)
	arg := db.AddWalletBalanceParams{
		Owner: int64(claims.ID),
		Name:  req.Name,
		Balance: sql.NullFloat64{
			Float64: balance,
			Valid:   true,
		},
	}

	updatedWallet, err := server.store.AddWalletBalance(ctx, arg)
	if err != nil {
		printSqlError(ctx, err)
		return
	}
	rsp := UpdateWalletResponse{
		Id:       updatedWallet.ID,
		Name:     updatedWallet.Name,
		Balance:  updatedWallet.Balance.Float64,
		Currency: updatedWallet.Currency,
	}

	ctx.JSON(http.StatusOK, rsp)
}
