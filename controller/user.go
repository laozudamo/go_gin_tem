package controller

import (
	response "goGinTem/Response"
	"goGinTem/dao"
	"goGinTem/forms"
	"goGinTem/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PasswordLogin(c *gin.Context) {
	PasswordLoginForm := forms.PasswordLoginForm{}

	if err := c.ShouldBind(&PasswordLoginForm); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}
	//查询数据库是否有该用户
}

// @Tags 首页
// @Success 200 {string} Welcome
// @Router /index [get]
func GetIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "这是数据",
		"data": "",
	})

}

func GetUserList(c *gin.Context) {
	// 获取参数
	UserListForm := forms.UserListForm{}
	if err := c.ShouldBind(&UserListForm); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}
	// 获取数据
	total, userlist := dao.GetUserListDao(UserListForm.Page, UserListForm.PageSize)
	// 判断
	if (total + len(userlist)) == 0 {
		response.Err(c, 400, 400, "未获取到到数据", map[string]interface{}{
			"total":    total,
			"userlist": userlist,
		})
		return
	}
	response.Success(c, 200, "获取用户列表成功", map[string]interface{}{
		"total":    total,
		"userlist": userlist,
	})
}

func ResgeterUser(c *gin.Context) {
	registerUser := &forms.RegisterUser{}

	if err := c.ShouldBindJSON(&registerUser); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}

	_, hasUser := dao.FindUser(registerUser.Tel)

	if hasUser {
		response.Success(c, 4001, "用户已经存在", nil)
		return
	}

	ok := dao.CreateUser(registerUser.Tel)

	if ok {
		// 生成token

		response.Success(c, 200, "注册用户成功", nil)
	}

}
