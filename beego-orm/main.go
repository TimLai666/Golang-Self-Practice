package main

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:password@/chapter4?charset=utf8")
	orm.RegisterModel(new(BeegoUser))
}

type BeegoUser struct {
	Id    int
	Name  string
	Phone string
}

func main() {
	orm.RunSyncdb("default", false, true)

	o := orm.NewOrm()
	user := BeegoUser{
		Name:  "tim",
		Phone: "123456",
	}
	fmt.Println(o.Insert(&user))
}
