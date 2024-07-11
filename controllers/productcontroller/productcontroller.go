package filmcontroller

import (
	"go-simple-blog/config"
	"go-simple-blog/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var films []models.Film

	config.DB.Find(&films)
	c.JSON(http.StatusOK, gin.H{"films": films})
}

func Show(c *gin.Context) {
	var film models.Film
	id := c.Param("id")

	if err := config.DB.First(&film, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"film": film})
}

func Create(c *gin.Context) {
	var film models.Film

	if err := c.ShouldBindJSON(&film); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	config.DB.Create(&film)
	c.JSON(http.StatusOK, gin.H{"film": film})
}

func Update(c *gin.Context) {
	var film models.Film
	id := c.Param("id")

	if err := c.ShouldBindJSON(&film); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if config.DB.Model(&film).Where("id = ?", id).Updates(&film).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot update table !!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully update data !!"})
}

func Delete(c *gin.Context) {
	var film models.Film
	var input struct {
		ID int64 `json:"id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	id := input.ID
	if config.DB.Delete(&film, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak bisa menghapus data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully delete data !!"})
}
