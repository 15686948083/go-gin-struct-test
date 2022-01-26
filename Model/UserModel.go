package model

// 用户数据库交互及查询方法封装
type User struct {
	ID       int    `json:"id" form:"id"`
	UserName string `json:"username" form:"username"`
}

// 定义get请求查询的参数的结构体
type UserQueryParam struct {
	// 加入form的tag才能映射到SelectUsersHandler.query
	UserName string `json:"username" form:"username" binding:"required"`
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"pagesize" form:"pagesize"`
}

// 绑定Dao层的json response
func CreateUser(id int, username string) User {
	return User{id, username}
}
