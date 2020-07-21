package game

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type playerDatabase struct {
	db *gorm.DB
}

func newPlayerDatabase() *playerDatabase {
	name := conf.Mysql.PlayerDatabase
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", conf.Mysql.Username, conf.Mysql.Password, conf.Mysql.Host, conf.Mysql.Port, name))
	// defer accountDB.Close()	Close 之后对数据库的操作无效且不报错..
	if err != nil {
		panic(err)
	}
	pdb = new(playerDatabase)
	pdb.db = db
	return pdb
}
