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
		// utils.HandleValidatorError(c, err)
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

func Resgeter(c *gin.Context) {

}
