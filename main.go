package main

import (
"encoding/hex"
"fmt"
"github.com/PonyWilliam/go-arcsoft/RfidUtils"
"github.com/gin-gonic/gin"
"net/http"
)

func main(){
	router := gin.Default()
	router.GET("",Cors(), func(context *gin.Context) {
		//返回rfid信息
		Rfid := RfidUtils.GetNearRfid()
		if Rfid!=nil{
			fmt.Println(hex.EncodeToString(Rfid[0]))
			context.JSON(200,gin.H{
				"code":200,
				"data":hex.EncodeToString(Rfid[0]),
			})
			return
		}
		context.JSON(200,gin.H{
			"code":500,
			"msg":"附近没有rfid标签",
		})
	})
	router.Use(Cors())
	gin.SetMode(gin.ReleaseMode)
	_ = router.Run("0.0.0.0:5888")
}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}