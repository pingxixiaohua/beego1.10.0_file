package controllers

import (
	"beegofile/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}

type SecondController struct {
	beego.Controller
}

/*
	请求路径：/register
	请求类型：POST
	包含的数据：json数据格式提交
		name, birthday, address, nick,
	后台：接收并打印解析到的请求数据

*/
func (c *SecondController) Post() {
	var information models.Information
	dataBytes2, err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Ctx.WriteString("您好，数据解析错误")
		return
	}
	err = json.Unmarshal(dataBytes2,&information)
	if err != nil {
		c.Ctx.WriteString("抱歉，数据解析错误，请重试")
		return
	}

	fmt.Printf("姓名：%s\n生日：%s\n地址：%s\n昵称：%s\n",information.Name,information.Birthday,information.Address,information.Nick)

	c.Ctx.WriteString("数据解析成功，恭喜你啊")
}

func (c *MainController) Get() {
	userName := c.Ctx.Input.Query("user")
	password := c.Ctx.Input.Query("psd")
	//使用固定数据进行数据校验
	if userName != "admin" || password != "123456"{
		c.Ctx.ResponseWriter.Write([]byte("对不起，数据校验错误"))
		return
	}

	c.Ctx.ResponseWriter.Write([]byte("恭喜，数据校验成功"))

	c.Data["Website"] = "www.baidu.com"
	c.Data["Email"] = "1608386461@qq.com"
	c.TplName = "index.tpl"
}



//func (c * MainController) Post(){
////	//1.接受解析post请求的参数
////	name := c.Ctx.Request.FormValue("name")
////	age := c.Ctx.Request.FormValue("age")
////	sex := c.Ctx.Request.FormValue("sex")
////	fmt.Println(name,age,sex)
////	//数据校验
////	if name != "admin" && age != "15"{
////		//c.Ctx.ResponseWriter.Write([]byte(""))
////		c.Ctx.WriteString("数据校验失败")
////		return
////	}
////	c.Ctx.WriteString("数据校验成功")
////}


//该方法用于处理post类型的请求

func (c * MainController) Post(){
	//解析前端提交的json格式的数据
	var person models.Person
	dataBytes, err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Ctx.WriteString("数据解析错误")
		return
	}
	err = json.Unmarshal(dataBytes,&person)
	if err != nil {
		c.Ctx.WriteString("数据解析错误，请重试")
		return
	}
	fmt.Println("姓名:",person.Name)
	fmt.Println("年龄:",person.Age)
	fmt.Println("性别:",person.Sex)
	c.Ctx.WriteString("数据解析成功")
}


func (c * MainController) Delete(){

}