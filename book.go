package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBookList(ctx *gin.Context) {
	// var bookList []Book
	// if err := maria.DB.Preload("Category").Find(&bookList).Error; err != nil {
	// 	ctx.JSON(http.StatusBadRequest, err)
	// }
	// fmt.Println(&bookList, bookList[0].Category.Name)
	// ctx.JSON(http.StatusOK, bookList)
	var bookList []Book
	query := maria.DB.Table("book b").Select(
		"b.id AS book_id, b.name AS book_name, c.name as category").Joins("JOIN category c on b.category_id = c.id")
	query.Find(&bookList)
	fmt.Println(bookList)
}

func GetOneBook(ctx *gin.Context) {
	book_id := ctx.Param("id")
	book := Book{}

	category := Category{}
	resData := GetOneBookResFormat{}
	if err := maria.DB.Where("id = ?", book_id).Related(&category, "Category").Find(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, "no such book")
		return
	}
	resData.Book = book
	ctx.JSON(http.StatusOK, resData)

}

func CreateBook(ctx *gin.Context) {
	book := Book{}
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	reqData := CreateBookFormat{}
	println(string(body))
	if err := json.Unmarshal(body, &reqData); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	if err := maria.DB.Create(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusCreated, "success")

}
