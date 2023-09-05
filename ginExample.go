package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CreateInput struct {
	Id int `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Email string `json:"email"`
}

type UpdateInput struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

var infoList []CreateInput

func readInfo(c *gin.Context) {
	if infoList == nil {
		c.JSON(http.StatusNoContent, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data": infoList,
	})
}

func createInfo(c *gin.Context) {
	input := CreateInput{}

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}
	infoList = append(infoList, input)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data": input,
	})
}

func updateInfo(c *gin.Context) {
	id, isExist := c.Params.Get("id")

	if !isExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "No ID",
		})
		return
	}

	n, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "No ID",
		})
		return
	}

	input := UpdateInput{}

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}

	for i, v := range infoList {
		if n == v.Id {
			if input.Name != "" {
				infoList[i].Name = input.Name
			}

			if input.Email != "" {
				infoList[i].Email = input.Email
			}

			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"data": infoList[i],
			})

			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"status": "Not Found",
	})
}

func deleteInfo(c *gin.Context) {
	id, isExist := c.Params.Get("id")

	if !isExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "No ID",
		})
		return
	}

	n, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "No ID",
		})
		return
	}

	for i, v := range infoList {
		if n == v.Id {
			if len(infoList) > 1 {
				infoList = append(infoList[:i], infoList[i+1:]...)
			} else {
				infoList = []CreateInput{}
			}

			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"data": infoList,
			})
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"status": "Not Found",
	})
}

func main() {
	r := gin.Default()
	r.GET("/info", readInfo)
	r.POST("/info", createInfo)
	r.PUT("/info/:id", updateInfo)
	r.DELETE("/info/:id", deleteInfo)

	r.Run("localhost:8080")
}