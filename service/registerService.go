package service

import(
	"douyin/dao"
)

type user struct{
	email string
	password string
	user_id string
	token string
}

func registerService(email string, password string){
	var User user
	//验证email  password格式  正则

	//生成user_id

	//生成token

	dao.InsertUser(User)
}