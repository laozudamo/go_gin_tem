package forms

type RegisterUser struct {
	Tel       string `form:"tel" json:"tel" binding:"required"`
	Captcha   string `form:"captcha" json:"captcha" binding:"required,min=5,max=5"` // 验证码
	CaptchaId string `form:"captchaId" json:"captchaId" binding:"required"`         // 验证码id
	PassWord  string `form:"password" json:"password" binding:"required,min=3,max=20"`
	CheckPwd  string `form:"checkPwd" json:"checkPwd" inding:"required,min=3,max=20"`
}

type PasswordLoginForm struct {
	Tel      string `form:"tel" json:"tel" binding:"required"`
	PassWord string `form:"password" json:"password" binding:"required,min=3,max=20"`
}

type FixPwdForm struct {
	Tel      string `form:"tel" json:"tel" binding:"required"`
	Captcha  string `form:"captcha" json:"captcha" binding:"required,min=5,max=5"`
	Oldpwd   string `form:"oldpwd" json:"oldpwd" binding:"required,min=3,max=20"`
	PassWord string `form:"password" json:"password" binding:"required,min=3,max=20"`
}

type QueryInfoForm struct {
	ID       uint   `form:"id" json:"id"`
	Username string `form:"username" json:"username"`
	Desc     string `form:"desc" json:"desc"`
	Gender   string `form:"gender" json:"gender"`
	Role     int    `form:"role" json:"role"`
	Email    string `form:"email" json:"email"`
	Address  string `form:"address" json:"address"`
	Tel      string `json:"tel"`
}

type UpdateInfoForm struct {
	Username string `form:"username" json:"username"`
	Desc     string `form:"desc" json:"desc"`
	Gender   string `form:"gender" json:"gender"`
	Role     int    `form:"role" json:"role"`
	Email    string `form:"email" json:"email"`
	Address  string `form:"address" json:"address"`
}

type UserListForm struct {
	// 页数
	Page int `form:"page" json:"page" binding:"required"`
	// 每页个数
	PageSize int `form:"pageSize" json:"pageSize" binding:"required"`
}
