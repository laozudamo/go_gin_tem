package forms

type RegisterUser struct {
	Tel       string `form:"tel" json:"tel" binding:"required"`
	Captcha   string `form:"captcha" json:"captcha" binding:"required,min=5,max=5"` // 验证码
	CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required"`       // 验证码id
	PassWord  string `form:"password" json:"password" binding:"required,min=3,max=20"`
	CheckPwd  string `form:"check_pwd" json:"check_pwd" binding:"required,min=3, max=20"`
}

type PasswordLoginForm struct {
	Tel      string `form:"tel" json:"tel" binding:"required"`
	PassWord string `form:"password" json:"password" binding:"required,min=3,max=20"`
}

type UserListForm struct {
	// 页数
	Page int `form:"page" json:"page" binding:"required"`
	// 每页个数
	PageSize int `form:"pageSize" json:"pageSize" binding:"required"`
}
