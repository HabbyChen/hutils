package interfaces

import "gorm.io/gorm"

//DbInterface 数据库需要实现的忌口
type DbInterface interface {
	//返回一个gorm的db实例
	DB() *gorm.DB
	//如果有多库的情况，可以根据name返回响应的实例
	GetDB(name string) *gorm.DB
}
