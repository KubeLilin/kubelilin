package tests

import (
	"fmt"
	"testing"
)

func TestSlices(t *testing.T) {
	slice := make([]int, 100)
	p := &slice

	slice[0] = 111
	slice[1] = 222

	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println((*p)[0])
	fmt.Println((*p)[1])

}

func TestPointer(t *testing.T) {
	name := "deployment"
	var ss *string
	ss = &name
	println(*ss)
}

func TestQuerySC(t *testing.T) {
	//dsn := "root:P@ssW0rd@tcp(47.100.213.41)/sgr_pass?charset=utf8&parseTime=True"
	//db1, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//query := app.queryServiceConnectionList(db1)
	//query1 := query.Where(func(e dto.ServiceConnectionInfo) bool {
	//	return e.ServiceType > 1
	//})
	//assert.Equal(t, query1.Count() > 0, true)
}
