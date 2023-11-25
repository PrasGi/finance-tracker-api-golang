package financecontroller

import (
	helpers "finance-tracker-api/Helpers"
	models "finance-tracker-api/Models"
	financerequest "finance-tracker-api/Requests/FinanceRequest"
	financeresponse "finance-tracker-api/Responses/FinanceResponse"
	financeservice "finance-tracker-api/Services/FinanceService"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Index(context *gin.Context) {
	var finance []models.Finance
	err := financeservice.Index(&finance)

	if helpers.JsonIfErr(err, context, 500) {
		return
	}

	context.JSON(200, gin.H{
		"status_code": 200,
		"message":     "Success get all finance",
		"data":        financeresponse.FinanceResponseMany(finance),
	})
}

func Store(context *gin.Context) {
	var body financerequest.BodyFinanceRequest

	user := context.MustGet("user").(models.User)

	body.UserId = user.ID

	if helpers.JsonIfErr(context.Bind(&body), context, 403) {
		return
	}

	if helpers.JsonIfErr(body.ValidateCreateRequest(), context, 400) {
		return
	}

	finance := models.Finance{
		Title:      body.Title,
		Amount:     body.Amount,
		UserId:     user.ID,
		CategoryId: body.CategoryId,
	}

	if helpers.JsonIfErr(financeservice.Store(&finance), context, 500) {
		return
	}

	context.JSON(200, gin.H{
		"status_code": 200,
		"message":     "Success creating finance",
		"data":        financeresponse.FinanceResponseOne(finance),
	})
}

func Show(context *gin.Context) {
	var finance models.Finance
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if helpers.JsonIfErr(err, context, 400) {
		return
	}

	if helpers.JsonIfErr(financeservice.Show(&finance, id), context, 500) {
		return
	}

	context.JSON(200, gin.H{
		"status_code": 200,
		"message":     "Success get finance",
		"data":        financeresponse.FinanceResponseOne(finance),
	})
}

func Update(context *gin.Context) {
	var body financerequest.BodyFinanceRequest

	user := context.MustGet("user").(models.User)

	body.UserId = user.ID

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

	finance := models.Finance{
		Title:      body.Title,
		Amount:     body.Amount,
		CategoryId: body.CategoryId,
		UserId:     user.ID,
	}

	if helpers.JsonIfErr(financeservice.Update(&finance, id), context, 500) {
		return
	}

	context.JSON(200, gin.H{
		"status_code": 200,
		"message":     "Success updating finance",
		"data":        financeresponse.FinanceResponseOne(finance),
	})
}

func Destroy(context *gin.Context) {
	var finance models.Finance
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if helpers.JsonIfErr(err, context, 400) {
		return
	}

	if helpers.JsonIfErr(financeservice.Destroy(&finance, id), context, 500) {
		return
	}

	context.JSON(200, gin.H{
		"status_code": 200,
		"message":     "Success deleting finance",
	})
}
