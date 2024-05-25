package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TbTest struct {
	ID       uint      // Standard field for the primary key
	Name     string    // 一个常规字符串字段
	Age      uint8     // 一个未签名的8位整数
	Birthday time.Time // A pointer to time.Time, can be null
}

func main() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:123456@tcp(172.16.56.130:3306)/tb_test?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         256,                                                                                 // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                                // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                                // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                                // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                               // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})

	if err != nil {
		fmt.Printf("数据库连接失败%s\n", err)
	}

	fmt.Println("数据库连接成功")

	user := TbTest{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

	result := db.Create(&user) // 通过数据的指针来创建

	fmt.Println(result.RowsAffected)
}
