package main

import (
	"fmt"
	"hutils"
	"hutils/boot"
	"hutils_test/aaaaa"
	"os"
	"path/filepath"
	"time"
)

func main() {
	str, _ := os.Getwd()
	hutils.Init(filepath.Join(str, "yml"))

	boot.Log.Infof("fuck %v", "我爱我家啦啦啦 啦啦啦 啦啦啦啦")
	boot.Log.Errorf("fuck %v", time.Now().Format("2016"))

	var uu aaaaa.ExamListSearchLog
	boot.Mysql.DB().Find(&uu)
	fmt.Printf("fuck: %+v", uu.TargetIDCard)
	//for true {

	//	time.Sleep(time.Duration(5) * time.Second)
	//}
	//s := hutils.Config().CusInit("yml/private.yml")
	//type Aaa struct {
	//	Aaa string
	//}
	//type Person struct {
	//	Fuck Aaa
	//	Ccc  string
	//}
	//
	//var person Person
	////将 map 转换为指定的结构体
	//if err := mapstructure.Decode(s, &person); err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Printf("map2struct后得到的 struct 内容为:%+v", person)
}
