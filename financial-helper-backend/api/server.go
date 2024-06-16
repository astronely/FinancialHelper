package api

import (
	db "FinancialHelper/db/sqlc"
	"FinancialHelper/middlware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"POST, GET, DELETE", "PUT"},
		AllowHeaders:     []string{"Access-Control-Request-Method, Access-Control-Allow-Methods, Access-Control-Request-Headers, Content-Type, origin, token"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	router.RedirectTrailingSlash = true
	api := router.Group("/api")
	{
		api.POST("/signup", server.createUser)
		api.POST("/signin", server.loginUser)
		api.GET("/signout", server.logoutUser)
		api.POST("/wallets/add", server.addWallet)
		api.GET("/wallets/get_wallets", server.getWallets)
		api.PUT("/wallets/update_wallet", server.updateWallet)
		api.PUT("/wallets/decrease_balance", server.decreaseWalletBalance)
		api.PUT("/wallets/increase_balance", server.increaseWalletBalance)
		api.DELETE("/wallets/delete", server.deleteWallet)
		api.POST("/expenses/add", server.addExpense)
		api.DELETE("/expenses/delete", server.deleteExpense)
		api.GET("/expenses/get_expenses", server.getExpenses)
	}

	user := router.Group("/user").Use(middlware.Auth())
	{
		user.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{"Message": "Success"})
		})
	}

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
