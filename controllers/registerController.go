package controllers

import (
	"beegofile/db_mysql"
	"beegofile/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type RegisterController struct {
	beego.Controller
}

/*
	请求路径：/register
	请求类型：POST
	包含的数据：json数据格式提交
		name, birthday, address, nick,
	后台：接收并打印解析到的请求数据

*/
func (r *RegisterController) Post() {
	//1、接收前端传递的数据
	//body,err := r.Ctx.Request.GetBody()
	//if err != nil {
	//	r.Ctx.WriteString("数据接收错误")
	//	return
	//}
	dataBytes2, err := ioutil.ReadAll(r.Ctx.Request.Body)
	if err != nil {
		r.Ctx.WriteString("您好，数据接受错误")
		return
	}

	var user models.User
	err = json.Unmarshal(dataBytes2,&user)
	if err != nil {
		r.Ctx.WriteString("抱歉，数据解析错误，请重试")
		return
	}
	
	//解析用户数据并保存
	id, err := db_mysql.InserUser(user)
	if err != nil {
		fmt.Println(err.Error())
		r.Ctx.WriteString("存储信息失败")
		return
	}

	fmt.Println(id)

	result := models.ResponseResult{
		Code:    1,
		Message: "保存成功",
		Data:    nil,
	}
	r.Data["json"] = &result
	r.ServeJSON()
	//var user models.User
	//err := r.ParseForm(&user)
	//if err != nil {
	//	r.Ctx.WriteString("数据解析错误")
	//}

	//fmt.Printf("姓名：%s\n生日：%s\n地址：%s\n昵称：%s\n",user.User,user.Birthday,user.Address,user.Nick)
	//
	//r.Ctx.WriteString("数据解析成功，恭喜你啊")
}

