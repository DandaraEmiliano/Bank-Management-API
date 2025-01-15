package api

import (
	"bank-management/middleware"
	"bank-management/models"
	"bank-management/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/users", func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing password"})
			return
		}
		user.Password = string(hashedPassword)

		if err := db.Create(&user).Error; err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
			return
		}

		c.JSON(http.StatusCreated, user)
	})

	r.POST("/login", func(c *gin.Context) {
		var credentials struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.ShouldBindJSON(&credentials); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}

		var user models.User
		if err := db.Where("email = ?", credentials.Email).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}

		token, err := utils.GenerateToken(user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	})

	auth := r.Group("/auth", middleware.AuthMiddleware())
	{
		auth.GET("/users", func(c *gin.Context) {
			var users []models.User
			if err := db.Find(&users).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, users)
		})

		auth.POST("/accounts", func(c *gin.Context) {
			var account models.Account
			if err := c.ShouldBindJSON(&account); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
				return
			}
			if err := db.Create(&account).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusCreated, account)
		})

		auth.GET("/accounts", func(c *gin.Context) {
			var accounts []models.Account
			if err := db.Find(&accounts).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, accounts)
		})

		auth.POST("/accounts/transfer", func(c *gin.Context) {
			type TransferRequest struct {
				FromAccountID uint    `json:"from_account_id"`
				ToAccountID   uint    `json:"to_account_id"`
				Amount        float64 `json:"amount"`
			}

			var transfer TransferRequest
			if err := c.ShouldBindJSON(&transfer); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
				return
			}

			var fromAccount, toAccount models.Account
			if err := db.First(&fromAccount, transfer.FromAccountID).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Source account not found"})
				return
			}

			if err := db.First(&toAccount, transfer.ToAccountID).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Target account not found"})
				return
			}

			if fromAccount.Balance < transfer.Amount {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient balance"})
				return
			}

			fromAccount.Balance -= transfer.Amount
			toAccount.Balance += transfer.Amount

			db.Save(&fromAccount)
			db.Save(&toAccount)

			c.JSON(http.StatusOK, gin.H{"message": "Transfer completed successfully"})
		})
	}
}
