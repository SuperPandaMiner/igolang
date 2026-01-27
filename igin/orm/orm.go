package orm

import "github.com/igolang/iorm"

func Init() {
	iorm.Init()
}

func Close() {
	iorm.Close()
}
