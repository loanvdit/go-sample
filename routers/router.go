package routers

import (
	"fujitech/web/web-go/go-challenge/middleware/jwt"
	"fujitech/web/web-go/go-challenge/routers/api"

	v1 "fujitech/web/web-go/go-challenge/routers/api/v1"

	"github.com/gin-gonic/gin"

	_ "fujitech/web/web-go/go-challenge/docs"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Register
	r.POST("/account/register", api.RegisterAccount)
	r.GET("/account/verify/:token", api.VerifyAccount)

	// Authentication
	r.POST("/auth", api.GetAuth)
	r.POST("/auth/:id", api.GetAuthVerify)

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Upload image
	r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		// Partner
		apiv1.GET("/partners", v1.GetPartners)
		apiv1.POST("/partners", v1.RegisterPartner)
		apiv1.PUT("/partners/:id", v1.StorePartner)
		apiv1.DELETE("/partners/:id", v1.DeletePartner)

		// Transaction
		apiv1.GET("/transactions", v1.GetTransactions)
		apiv1.GET("/transactions/:id", v1.GetTransaction)
		apiv1.GET("/transactions/search/:start_date/:end_date", v1.GetTransactionByDate)
		apiv1.POST("/transactions", v1.RegisterTransaction)
		apiv1.POST("/transactions/status/:id/:status", v1.StoreStatusTransaction)
		apiv1.PUT("/transactions/:id", v1.StoreTransaction)
		apiv1.DELETE("/transactions/:id", v1.DeleteTransaction)

		// Report
		apiv1.GET("/statistic/:id", v1.GetStatisticReport)
	}

	return r
}
