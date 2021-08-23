package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SgrSysMenuMgr struct {
	*_BaseMgr
}

// SgrSysMenuMgr open func
func SgrSysMenuMgr(db *gorm.DB) *_SgrSysMenuMgr {
	if db == nil {
		panic(fmt.Errorf("SgrSysMenuMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SgrSysMenuMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sgr_sys_menu"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SgrSysMenuMgr) GetTableName() string {
	return "sgr_sys_menu"
}

// Get 获取
func (obj *_SgrSysMenuMgr) Get() (result SgrSysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrSysMenu{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SgrSysMenuMgr) Gets() (results []*SgrSysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrSysMenu{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SgrSysMenuMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SgrSysMenu{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_SgrSysMenuMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTenantID tenant_id获取 租户
func (obj *_SgrSysMenuMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithMenuCode menu_code获取 编码
func (obj *_SgrSysMenuMgr) WithMenuCode(menuCode string) Option {
	return optionFunc(func(o *options) { o.query["menu_code"] = menuCode })
}

// WithMenuName menu_name获取 目录名称
func (obj *_SgrSysMenuMgr) WithMenuName(menuName string) Option {
	return optionFunc(func(o *options) { o.query["menu_name"] = menuName })
}

// WithIsRoot is_root获取 是否是根目录
func (obj *_SgrSysMenuMgr) WithIsRoot(isRoot int8) Option {
	return optionFunc(func(o *options) { o.query["is_root"] = isRoot })
}

// WithParentID parent_id获取 父层级id
func (obj *_SgrSysMenuMgr) WithParentID(parentID int64) Option {
	return optionFunc(func(o *options) { o.query["parent_id"] = parentID })
}

// WithSort sort获取 权重，正序排序
func (obj *_SgrSysMenuMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithStatus status获取 状态
func (obj *_SgrSysMenuMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreatetionTime createtion_time获取
func (obj *_SgrSysMenuMgr) WithCreatetionTime(createtionTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["createtion_time"] = createtionTime })
}

// WithUpdateTime update_time获取
func (obj *_SgrSysMenuMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_SgrSysMenuMgr) GetByOption(opts ...Option) (result SgrSysMenu, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SgrSysMenu{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SgrSysMenuMgr) GetByOptions(opts ...Option) (results []*SgrSysMenu, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SgrSysMenu{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_SgrSysMenuMgr) GetFromID(id uint64) (result SgrSysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrSysMenu{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_SgrSysMenuMgr) GetBatchFromID(ids []uint64) (results []*SgrSysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrSysMenu{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容 租户
func (obj *_SgrSysMenuMgr) GetFromTenantID(tenantID int64) (results []*SgrSysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrSysMenu{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找 租户
func (obj *_SgrSysMenuMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*SgrSysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrSysMenu{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromMenuCode 通过menu_code获取内容 编码
func (obj *_SgrSysMenuMgr) GetFromMenuCode(menuCode string) (result SgrSysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrSysMenu{}).Where("`menu_code` = ?", menuCode).Find(&result).Error

	return
}

// GetBatchFromMenuCode 批量查找 编码
func (obj *_SgrSysMenuMgr) GetBatchFromMenuCode(menuCodes []string) (results []*SgrSysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrSysMenu{}).Where("`menu_code` IN (?)", menuCodes).Find(&results).Error

	return
}

// GetFromMenuName 通过menu_name获取内容 目录名称
func (obj *_SgrSysMenuMgr) GetFromMenuName(menuName string) (results []*SgrSysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrSysMenu{}).Where("`menu_name` = ?", menuName).Find(&results).Error

	return
}

// GetBatchFromMenuName 批量查找 目录名称
func (obj *_SgrSysMenuMgr) GetBatchFromMenuName(menuNames []string) (results []*SgrSysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrSysMenu{}).Where("`menu_name` IN (?)", menuNames).Find(&results).Error

	return
}

// GetFromIsRoot 通过is_root获取内容 是否是根目录
func (obj *_SgrSysMenuMgr) GetFromIsRoot(isRoot int8) (results []*SgrSysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrSysMenu{}).Where("`is_root` = ?", isRoot).Find(&results).Error

	return
}

// GetBatchFromIsRoot 批量查找 是否是根目录
func (obj *_SgrSysMenuMgr) GetBatchFromIsRoot(isRoots []int8) (results []*SgrSysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrSysMenu{}).Where("`is_root` IN (?)", isRoots).Find(&results).Error

	return
}

// GetFromParentID 通过parent_id获取内容 父层级id
func (obj *_SgrSysMenuMgr) GetFromParentID(parentID int64) (results []*SgrSysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrSysMenu{}).Where("`parent_id` = ?", parentID).Find(&results).Error

	return
}

// GetBatchFromParentID 批量查找 父层级id
func (obj *_SgrSysMenuMgr) GetBatchFromParentID(parentIDs []int64) (results []*SgrSysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrSysMenu{}).Where("`parent_id` IN (?)", parentIDs).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 权重，正序排序
func (obj *_SgrSysMenuMgr) GetFromSort(sort int) (results []*SgrSysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrSysMenu{}).Where("`sort` = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找 权重，正序排序
func (obj *_SgrSysMenuMgr) GetBatchFromSort(sorts []int) (results []*SgrSysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrSysMenu{}).Where("`sort` IN (?)", sorts).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态
func (obj *_SgrSysMenuMgr) GetFromStatus(status int8) (results []*SgrSysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrSysMenu{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态
func (obj *_SgrSysMenuMgr) GetBatchFromStatus(statuss []int8) (results []*SgrSysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrSysMenu{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreatetionTime 通过createtion_time获取内容
func (obj *_SgrSysMenuMgr) GetFromCreatetionTime(createtionTime time.Time) (results []*SgrSysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrSysMenu{}).Where("`createtion_time` = ?", createtionTime).Find(&results).Error

	return
}

// GetBatchFromCreatetionTime 批量查找
func (obj *_SgrSysMenuMgr) GetBatchFromCreatetionTime(createtionTimes []time.Time) (results []*SgrSysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrSysMenu{}).Where("`createtion_time` IN (?)", createtionTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_SgrSysMenuMgr) GetFromUpdateTime(updateTime time.Time) (results []*SgrSysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrSysMenu{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_SgrSysMenuMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*SgrSysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrSysMenu{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SgrSysMenuMgr) FetchByPrimaryKey(id uint64) (result SgrSysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrSysMenu{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchUniqueByUnCCode primary or index 获取唯一内容
func (obj *_SgrSysMenuMgr) FetchUniqueByUnCCode(menuCode string) (result SgrSysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrSysMenu{}).Where("`menu_code` = ?", menuCode).Find(&result).Error

	return
}
