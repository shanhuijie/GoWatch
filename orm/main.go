package main

import(
	"GoWatch/orm/qa"
	_"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)



func main(){
	tru := qa.Astruct{}
	tru.GetContent(2)
	// tru.GetContent(15)
	fmt.Println(tru)
}