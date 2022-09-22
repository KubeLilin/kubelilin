package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"kubelilin/domain/business/app"
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
	dsn := "root:P@ssW0rd@tcp(47.100.213.41)/sgr_pass?charset=utf8&parseTime=True"
	db1, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sc := app.NewServiceConnectionService(db1)
	te, e0 := sc.GetTest()

	pe, e := sc.GetPipelineEngine()
	img, e1 := sc.GetImageHub()
	sys, e2 := sc.GetSystemCallback()

	fmt.Println(te)
	fmt.Println(pe)
	fmt.Println(img)
	fmt.Println(sys)
	assert.Equal(t, e0 != nil, true)
	assert.Equal(t, e == nil, true)
	assert.Equal(t, e1 == nil, true)
	assert.Equal(t, e2 == nil, true)
}
