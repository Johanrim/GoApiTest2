package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategoryList(ctx *gin.Context) {
	// var res GetCategoryListResFormat
	// category := Category{}
	// err := maria.DB.Find(&category).Error //&category
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, err)
	// 	return
	// }
	// res.Category = append(res.Category, category)
	// ctx.JSON(http.StatusOK, res)
	var categoryList []Category
	if err := maria.DB.Find(&categoryList).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, &categoryList)
}

func GetOneCategory(ctx *gin.Context) {
	category_id := ctx.Param("id")
	// categoryid, err := strconv.ParseUint(category_id, 10, 32)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// categoryId := uint(categoryid)

	var res GetCategoryResFormat

	category := Category{}

	err := maria.DB.Where("id = ?", category_id).Find(&category).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "no such category")
		return
	}

	// res.ID = categoryId
	// res.Name = category.Name
	res.Category = category
	ctx.JSON(http.StatusOK, res)
}

func CreateCategory(ctx *gin.Context) {
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	reqData := CreateCategoryFormat{}

	if err := json.Unmarshal(body, &reqData); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	// requestdata validation

	category := Category{
		Name: reqData.Name,
	}
	resData := CreateCategoryResFormat{}
	if err := maria.DB.Save(&category).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	resData.Message = "Create Success"
	resData.Category = category
	ctx.JSON(http.StatusCreated, resData)
}

func UpdateCategory(ctx *gin.Context) {
	category_id := ctx.Param("id")
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	reqData := UpdateCategoryFormat{}
	resData := UpdateCategoryResFormat{}

	if err := json.Unmarshal(body, &reqData); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	category := Category{}
	if notFound := maria.DB.Where("id = ?", category_id).First(&category).RecordNotFound(); notFound {
		ctx.JSON(http.StatusBadRequest, "no such category")
		return
	}

	category.Name = reqData.Name

	if err := maria.DB.Save(&category).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	resData.Message = "Update Success"
	resData.Category = category
	ctx.JSON(http.StatusOK, resData)
}

func DeleteCategory(ctx *gin.Context) {
	category_id := ctx.Param("id")
	category := Category{}
	resData := RemoveCategoryResFormat{}
	if err := maria.DB.Where("id = ?", category_id).Delete(&category).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	resData.Message = "Delete Success"
	ctx.JSON(http.StatusOK, resData)
}
