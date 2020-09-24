package models

//人结构体的定义，包含三个字段：name,age,sex,与前端一致
type Person struct {
	Name string
	Age int
	Sex string
}

type Information struct {
	Name string
	Birthday string
	Address string
	Nick string
}
