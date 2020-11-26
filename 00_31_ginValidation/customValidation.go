package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
)

//Booking aaa
type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,checkdate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

//カスタマValidation作成
// var checkdate validator.Func = func(f1 validator.FieldLevel) bool {
func checkdate(f1 validator.FieldLevel) bool {
	date, ok := f1.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}

func main() {
	// router作成
	r := gin.Default()
	//カスタマValidationのレジスト
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("checkdate", checkdate)
	}

	//Get受け取り
	r.GET("/bookable", getBookable)
	r.Run()
}

//get EndPoint
func getBookable(c *gin.Context) {
	var b Booking
	//Queryの値だけチェックの場合
	// if err := c.ShouldBindWith(&b, binding.Query); err == nil {
	//post queryの値全部チェック
	if err := c.ShouldBind(&b); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
