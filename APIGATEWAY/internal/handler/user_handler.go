package handler

import (
	"apigateway/internal/middleware"
	svc "github.com/AskatNa/apis-gen-user-service/service/frontend/client/v1"
	"strconv"

	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes(r *gin.Engine, client svc.CustomerServiceClient) {
	r.POST("/register", func(c *gin.Context) {
		var req svc.RegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res, err := client.Register(context.Background(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"user_id": res.Id})
	})

	r.POST("/login", func(c *gin.Context) {
		var req svc.LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res, err := client.Login(context.Background(), &req)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	})

	protected := r.Group("/users")
	protected.Use(middleware.JWTAuthMiddleware())

	protected.GET("/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
			return
		}

		res, err := client.Get(c, &svc.GetRequest{Id: id})
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res.Customer)
	})

	// Add other required endpoints to match the CustomerServiceClient interface
	protected.POST("/refresh-token", func(c *gin.Context) {
		var req svc.RefreshTokenRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res, err := client.RefreshToken(c, &req)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	})

	protected.PUT("/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
			return
		}

		var req svc.UpdateRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		req.Id = id

		res, err := client.Update(c, &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	})
}
