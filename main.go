// 程序的入口函数
package main

import (
	. "gin-web-struct/Dao"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routerGroup := r.Group("/api/v1")
	{
		// 查询用户列表页信息
		routerGroup.GET("/users", SelectUsersHandler)
		// 添加用户信息
		routerGroup.POST("/users", CreateUserHandler)
		// 修改用户信息
		routerGroup.PUT("/user/:user_id", UpdateUserHandler)
		// 删除用户信息
		routerGroup.DELETE("/user/:user_id", DeleteUserHandler)
	}
	r.Run()
}
