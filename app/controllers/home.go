package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type articles struct {
	Title string
	Name  string
}

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "main.html", gin.H{"model": "cate", "m": "home"})
}
