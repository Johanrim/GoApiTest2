package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetOneCategory(ctx *gin.Context) {
	category_id := ctx.Param("id")

	mariaDB := MariaDB{}
	mariaDB.connect()
	defer mariaDB.close()

	categoryid, err := strconv.ParseUint(category_id, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	categoryId := uint(categoryid)

	var res GetCategoryResFormat

	category := Category{}

	err = mariaDB.DB.Where("id = ?", category_id).Find(&category).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	res.ID = categoryId
	res.Name = category.Name
	ctx.JSON(http.StatusOK, res)
}
