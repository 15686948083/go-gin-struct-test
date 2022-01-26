package Dao

import (
	"fmt"
	. "gin-web-struct/Model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 接口的入口函数 相当于django url_pattern中的函数名称

func SelectUsersHandler(c *gin.Context) {
	// 实例化参数结构体
	query := UserQueryParam{}
	// 绑定参数
	err := c.BindQuery(&query)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 4001,
			// "msg":  "参数验证错误",
			"msg": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查询用户列表页信息",
			"data": query,
		})
	}
	// // 获取参数
	// username := c.Query("username")
	// if username == "" {
	// 	// 查询所有
	// 	c.JSON(200, gin.H{
	// 		"code": 200,
	// 		"msg":  "查询用户列表页信息",
	// 	})
	// } else {
	// 	// 查询一个用户的信息
	// 	c.JSON(200, gin.H{
	// 		"code": 200,
	// 		"msg":  "查询用户" + username + "信息",
	// 	})
	// }
}

func CreateUserHandler(c *gin.Context) {

	c.JSON(200, CreateUser(1, "alex"))
}
func UpdateUserHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "修改用户信息",
	})
}
func DeleteUserHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "删除用户信息",
	})
}
