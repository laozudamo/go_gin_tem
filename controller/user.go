package controller

import (
	response "goGinTem/Response"
	"goGinTem/dao"
	"goGinTem/forms"
	"goGinTem/utils"

	"github.com/gin-gonic/gin"
)

func PasswordLogin(c *gin.Context) {
	PasswordLoginForm := forms.PasswordLoginForm{}

	if err := c.ShouldBind(&PasswordLoginForm); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}
	user, ok := dao.FindUser(PasswordLoginForm.Tel)

	if !ok {
		response.Err(c, 200, 4002, "用户名或者密码错错误", "")
		return
	}

	if user.Password != utils.HashString(PasswordLoginForm.PassWord) {
		response.Err(c, 200, 4002, "用户名或者密码错错误", "")
		return
	}

	token := utils.CreateToken(c, int(user.ID), user.Tel)

	response.Success(c, 200, "登录成功", token)

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

	if !store.Verify(registerUser.CaptchaId, registerUser.Captcha, true) {
		response.Err(c, 400, 400, "验证码错误", nil)
		return
	}

	if registerUser.PassWord != registerUser.CheckPwd {
		response.Err(c, 400, 400, "密码不一致,请重新输入", nil)
		return
	}

	_, hasUser := dao.FindUser(registerUser.Tel)

	if hasUser {
		response.Success(c, 4001, "用户已经存在", nil)
		return
	}

	// user := models.User{}
	hashPwd := utils.HashString(registerUser.PassWord)
	_, ok := dao.CreateUser(registerUser.Tel, hashPwd)

	if ok {
		response.Success(c, 200, "注册用户成功", nil)
	}

}
