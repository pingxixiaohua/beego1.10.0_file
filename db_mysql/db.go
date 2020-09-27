package db_mysql

import (
	"beegofile/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init(){
	fmt.Println("连接mysql数据库")
	config := beego.AppConfig

	dbDriver := config.String("db_driverName")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")
	//root:wph18407850471@tcp(127.0.0.1:3306)/lol_hero?charset=utf8
	connUrl := dbUser +":" + dbPassword + "@tcp("+dbIp+")/"+dbName+"?charset=utf8"
	db, err := sql.Open(dbDriver,connUrl)
	if err != nil {
		fmt.Println(err.Error())
		panic("数据库连接失败，请检查配置")
	}

	Db = db

}

func InserUser(user models.User) (int64, error){
	//哈希加密
	hashMd5 := md5.New()
	hashMd5.Write([]byte(user.Password))
	bytes := hashMd5.Sum(nil)
	user.Password = hex.EncodeToString(bytes)
	//fmt.Printf("用户名：%s\n，密码：%s\n",user.Nick,user.Password)
	result, err := Db.Exec("insert into user(users,birthday,address,nick,password) values(?,?,?,?,?)",user.Users,user.Birthday,user.Address,user.Nick,user.Password)
	if err != nil {
		return -1,err
	}
	id, err := result.RowsAffected()
	if err != nil {
		return -1,err
	}
	return id,nil

}

func QueryUser()  {
	Db.QueryRow("select * from ")
}
