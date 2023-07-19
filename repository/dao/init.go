package dao

import (
	"bytetrack_server/repository/dal"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _db *gorm.DB

func MySQLInit() {
	dsn := "root:root@tcp(127.0.0.1:3306)/bytetrack_server?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}))
	if err != nil {
		panic(err)
	}

	_db = db
	dal.SetDefault(db)
	migration()
}
