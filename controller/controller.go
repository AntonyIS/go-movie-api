package controller

import (
	"net/http"

	"github.com/AntonyIS/GoProject/models"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Dresses)
}
