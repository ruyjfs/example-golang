package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
}

func (controller *Controller) Json(c *gin.Context, result *gorm.DB, data interface{}) {
	if result.Error == nil {
		controller.JsonSuccess(c, http.StatusOK, gin.H{"data": data})
		return
	}
	controller.JsonFail(c, http.StatusBadRequest, result.Error.Error())
	return
}

func (controller *Controller) JsonSuccess(c *gin.Context, status int, h gin.H) {
	h["status"] = "success"
	c.JSON(status, h)
	return
}

func (controller *Controller) JsonFail(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"status":  "fail",
		"message": message,
		"data":    nil,
		"code":    status,
	})
}
