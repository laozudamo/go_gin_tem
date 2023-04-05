package controller

import (
	"encoding/json"
	"fmt"
	response "goGinTem/Response"
	"goGinTem/dao"
	"goGinTem/forms"
	"goGinTem/models"
	"goGinTem/utils"

	"github.com/gin-gonic/gin"
)

// @Summary      账号登录接口
// @Description  get accounts
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        q    query     string  false  "name search by q"  Format(email)
// @Success      200  {array}   model.Account
// @Failure      400  {object}  httputil.HTTPError
// @Router       /login [post]
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

func GetUserInfo(c *gin.Context) {
	value, _ := c.Get("userId")

	user, ok := dao.GetUserInfo(value)
	info := models.UserInfo{}

	if err := json.Unmarshal([]byte(user.UserInfo), &info); err != nil {
		return
	}

	if ok {
		response.Success(c, 200, "获取成功", info)
	} else {
		response.Err(c, 200, 500, "查询数据错误", "")
	}

}

func UpdateUserInfo(c *gin.Context) {
	id, _ := c.Get("userId")

	userInfo := models.UserInfo{}

	if err := c.ShouldBindJSON(&userInfo); err != nil {
		fmt.Printf("err1: %v\n", err)
		panic(err)
	}

	jsonStringInfo, err := json.Marshal(userInfo)

	if err != nil {
		panic(err)
	}
	_, ok := dao.UpdateUserInfo(string(jsonStringInfo), id)

	if ok {
		response.Success(c, 200, "用户信息更新成功", "")
	}

}
