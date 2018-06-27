package posts

import (
	"net/http"
	"strconv"

	"Kerncheh/medium-api/db/connection"

	"Kerncheh/medium-api/models"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/orm"
)

func Index(c *gin.Context) {
	pageString := c.DefaultQuery("page", "1")
	showNotPublishedParam := c.DefaultQuery("shownotpublished", "false")
	page, _ := strconv.Atoi(pageString)

	posts := make([]*models.Post, 0)

	var pager = &orm.Pager{}
	pager.Limit = 100
	pager.SetPage(page)

	query := connection.DBCon.Model(&posts)

	if showNotPublished, err := strconv.ParseBool(showNotPublishedParam); err != nil || !showNotPublished {
		query = query.Where("published = ?", true)
	}

	if err := query.Order("id asc").Apply(pager.Paginate).Select(); err == nil {
		c.JSON(http.StatusOK, posts)
	} else {
		RespondWithError(c, err)
	}
}

func Show(c *gin.Context) {
	//TODO: Secure endpoint for unpublished posts
	post := &models.Post{}

	if err := connection.DBCon.Model(post).Where("id = ?", c.Param("id")).Select(); err == nil {
		c.JSON(http.StatusOK, post)
	} else {
		RespondWithError(c, err)
	}
}

func Create(c *gin.Context) {
	var post models.Post

	if err := c.ShouldBindJSON(&post); err == nil {
		if err := connection.DBCon.Insert(&post); err == nil {
			c.JSON(http.StatusCreated, post)
		} else {
			RespondWithError(c, err)
		}
	} else {
		RespondWithError(c, err)
	}
}

func Update(c *gin.Context) {
	post := &models.Post{}

	if err := connection.DBCon.Model(post).Where("id = ?", c.Param("id")).Select(); err == nil {
		if err := c.ShouldBindJSON(&post); err == nil {
			if err := connection.DBCon.Update(post); err == nil {
				c.JSON(http.StatusOK, post)
			} else {
				RespondWithError(c, err)
			}
		} else {
			RespondWithError(c, err)
		}
	} else {
		RespondWithError(c, err)
	}
}

func RespondWithError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
}
