package dao

import (
	"encoding/json"
	"goGinTem/forms"
	"goGinTem/global"
	"goGinTem/models"
)

// var users []models.User
var user models.User
var users []models.User

// GetUserList 获取用户列表(page第几页,page_size每页几条数据)
func GetUserListDao(page int, page_size int) (int, []interface{}) {
	// 分页用户列表数据
	userList := make([]interface{}, 0, len(users))
	// 计算偏移量
	offset := (page - 1) * page_size
	// 查询所有的user
	result := global.DB.Offset(offset).Limit(page_size).Find(&users)
	// 查不到数据时
	if result.RowsAffected == 0 {
		return 0, userList
	}
	// 获取user总数
	total := len(users)
	// 查询数据
	result.Offset(offset).Limit(page_size).Find(&users)
	//
	for _, useSingle := range users {

		userItemMap := handleUserItem(useSingle)

		userList = append(userList, userItemMap)
	}
	return total, userList
}

func handleUserItem(user models.User) *forms.QueryInfoForm {
	userInfo := forms.QueryInfoForm{}
	userInfo.Tel = user.Tel
	userInfo.ID = user.ID
	if err := json.Unmarshal([]byte(user.UserInfo), &userInfo); err != nil {
		return &userInfo
	}
	return &userInfo
}

// UsernameFindUserInfo 通过username找到用户信息

func FindUser(tel string) (*models.User, bool) {
	user := models.User{}
	if err := global.DB.Where("tel = ?", tel).First(&user).Error; err != nil {
		return &user, false
	}
	return &user, true
}

func GetUserInfo(id any) (*models.User, bool) {
	if err := global.DB.Where("id=?", id).First(&user).Error; err != nil {
		return &user, false
	}
	return &user, true
}

func CreateUser(tel string, password string) (*models.User, bool) {
	user := models.User{}
	user.Tel = tel
	user.Password = password
	if err := global.DB.Create(&user).Error; err != nil {
		return &user, false
	}
	return &user, true
}

func UpdateUserInfo(userInfo string, id any) (*models.User, bool) {

	if err := global.DB.Model(&user).Where("id = ?", id).Update("user_info", userInfo).Error; err != nil {
		return &user, false
	}
	return &user, true
}
