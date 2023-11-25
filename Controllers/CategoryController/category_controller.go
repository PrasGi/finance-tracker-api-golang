package categorycontroller

import (
	helpers "finance-tracker-api/Helpers"
	models "finance-tracker-api/Models"
	categoryrequest "finance-tracker-api/Requests/CategoryRequest"
	categoryresponse "finance-tracker-api/Responses/CategoryResponse"
	categoryservice "finance-tracker-api/Services/CategoryService"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Index(context *gin.Context) {
	var categories []models.Category
	err := categoryservice.Index(&categories)

	if helpers.JsonIfErr(err, context, 500) {
		return
	}

	context.JSON(200, gin.H{
		"status_code": 200,
		"message":     "Success get all category",
		"data":        categoryresponse.CategoryResponseMany(categories),
	})
}

func Store(context *gin.Context) {
	var body categoryrequest.BodyCategoryRequest

	if helpers.JsonIfErr(context.Bind(&body), context, 403) {
		return
	}

	if helpers.JsonIfErr(body.ValidateCreateRequest(), context, 400) {
		return
	}

	category := models.Category{
		Name:        body.Name,
		Description: body.Description,
	}

	if helpers.JsonIfErr(categoryservice.Store(&category), context, 500) {
		return
	}

	context.JSON(200, gin.H{
		"status_code": 200,
		"message":     "Success creating category",
		"data":        categoryresponse.CategoryResponseOne(category),
	})
}

func Show(context *gin.Context) {
	var category models.Category
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if helpers.JsonIfErr(err, context, 400) {
		return
	}

	if helpers.JsonIfErr(categoryservice.Show(&category, id), context, 500) {
		return
	}

	context.JSON(200, gin.H{
		"status_code": 200,
		"message":     "Success get one category",
		"data":        categoryresponse.CategoryResponseOne(category),
	})
}

func Update(context *gin.Context) {
	var body categoryrequest.BodyCategoryRequest

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if helpers.JsonIfErr(err, context, 400) {
		return
	}

	if helpers.JsonIfErr(context.Bind(&body), context, 403) {
		return
	}

	if helpers.JsonIfErr(body.ValidateCreateRequest(), context, 400) {
		return
	}

	category := models.Category{
		Name:        body.Name,
		Description: body.Description,
	}
	if helpers.JsonIfErr(categoryservice.Update(&category, id), context, 500) {
		return
	}

	context.JSON(200, gin.H{
		"status_code": 200,
		"message":     "Success updating category",
		"data":        categoryresponse.CategoryResponseOne(category),
	})
}

func Destroy(conext *gin.Context) {
	var category models.Category
	id, err := strconv.ParseInt(conext.Param("id"), 10, 64)

	if helpers.JsonIfErr(err, conext, 400) {
		return
	}

	if helpers.JsonIfErr(categoryservice.Destroy(&category, id), conext, 500) {
		return
	}

	conext.JSON(200, gin.H{
		"status_code": 200,
		"message":     "Success deleting category",
	})
}
