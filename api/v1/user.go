package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/l6l6ng/go-web2/model"
	"github.com/l6l6ng/go-web2/utils/errmsg"
	"net/http"
	"strconv"
)

// UserExist 查询用户是否存在
func UserExist(c *gin.Context) {

}

// AddUser 添加用户
func AddUser(c *gin.Context) {
	user := model.User{}
	//err := c.ShouldBindWith(&user, binding.Form)
	err := c.ShouldBind(&user)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//验证用户是否存在
	errCode := model.UserExit(user.UserName)
	if errCode == errmsg.SUCCESS {
		//创建用户
		errCode = model.AddUser(&user)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": errCode,
		"msg":    errmsg.GetErrMsg(errCode),
	})
}

// GetUsers 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pagenum", "1"))

	type responseUser struct {
		Id       uint64 `json:"id"`
		UserName string `json:"user_name"`
	}

	users := model.GetUsers(pageSize, pageNum)
	data := make([]responseUser, len(users))
	for k, v := range users {
		data[k] = responseUser{
			v.ID,
			v.UserName,
		}
	}

	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
		"data":   data,
	})
}

//查询单个用户

// EditUser 编辑用户
func EditUser(c *gin.Context) {

}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {

}
