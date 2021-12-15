package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sgr/domain/business/tenant"
	"testing"
)

func TestRoleMenuList(t *testing.T) {
	dsn := "root:1234abcd@tcp(cdb-amqub3mo.bj.tencentcdb.com:10042)/sgr_paas?charset=utf8&parseTime=True"
	db1, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	service := tenant.NewSysMenuService(db1)

	dd := service.GetRoleMenuIdList(1)

	assert.Equal(t, len(dd) > 0, true)
}

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
