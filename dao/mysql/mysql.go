package mysql

import (
	"fmt"
	"web_app/settings"

	"gorm.io/gorm/schema"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(cfg *settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName)
	if db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false, // 跳过默认事务
		NamingStrategy: schema.NamingStrategy{ // 命名策略
			TablePrefix:   "t_",  // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: false, // 使用复数表名，此时，`User` 的表名应该是 `t_users`
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 拒绝创建外键约束
	}); err != nil {
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		return
	}
	// 设置最大连接数
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	// 设置连接池中空闲的最大连接数
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	return sqlDB.Ping()
}
