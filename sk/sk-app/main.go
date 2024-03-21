package main

import (
	"fmt"
	//"github.com/JunChen832/my_sk/common/discover"
	//"github.com/longjoy/micro-go-book/ch13-seckill/pkg/bootstrap"
	"github.com/JunChen832/my_sk/sk/sk-app/setup"
)

func main() {
	//mysql.InitMysql(conf.MysqlConfig.Host, conf.MysqlConfig.Port, conf.MysqlConfig.User, conf.MysqlConfig.Pwd, conf.MysqlConfig.Db) // conf.MysqlConfig.Db
	setup.InitZk()
	//setup.InitRedis()
	//setup.InitServer(bootstrap.HttpConfig.Host, bootstrap.HttpConfig.Port)
	fmt.Println("hello world")
}
