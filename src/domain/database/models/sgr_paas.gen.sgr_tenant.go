package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SgrTenantMgr struct {
	*_BaseMgr
}

// SgrTenantMgr open func
func SgrTenantMgr(db *gorm.DB) *_SgrTenantMgr {
	if db == nil {
		panic(fmt.Errorf("SgrTenantMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SgrTenantMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sgr_tenant"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SgrTenantMgr) GetTableName() string {
	return "sgr_tenant"
}

// Get 获取
func (obj *_SgrTenantMgr) Get() (result SgrTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenant{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SgrTenantMgr) Gets() (results []*SgrTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenant{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SgrTenantMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SgrTenant{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_SgrTenantMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTName t_name获取 租户名称
func (obj *_SgrTenantMgr) WithTName(tName string) Option {
	return optionFunc(func(o *options) { o.query["t_name"] = tName })
}

// WithTCode t_code获取 租户编码
func (obj *_SgrTenantMgr) WithTCode(tCode string) Option {
	return optionFunc(func(o *options) { o.query["t_code"] = tCode })
}

// WithStatus status获取 状态
func (obj *_SgrTenantMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreationTime creation_time获取 创建时间
func (obj *_SgrTenantMgr) WithCreationTime(creationTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["creation_time"] = creationTime })
}

// WithUpdateTime update_time获取 修改时间
func (obj *_SgrTenantMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_SgrTenantMgr) GetByOption(opts ...Option) (result SgrTenant, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SgrTenant{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SgrTenantMgr) GetByOptions(opts ...Option) (results []*SgrTenant, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SgrTenant{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_SgrTenantMgr) GetFromID(id uint64) (result SgrTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenant{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_SgrTenantMgr) GetBatchFromID(ids []uint64) (results []*SgrTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenant{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromTName 通过t_name获取内容 租户名称
func (obj *_SgrTenantMgr) GetFromTName(tName string) (results []*SgrTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenant{}).Where("`t_name` = ?", tName).Find(&results).Error

	return
}

// GetBatchFromTName 批量查找 租户名称
func (obj *_SgrTenantMgr) GetBatchFromTName(tNames []string) (results []*SgrTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenant{}).Where("`t_name` IN (?)", tNames).Find(&results).Error

	return
}

// GetFromTCode 通过t_code获取内容 租户编码
func (obj *_SgrTenantMgr) GetFromTCode(tCode string) (result SgrTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenant{}).Where("`t_code` = ?", tCode).Find(&result).Error

	return
}

// GetBatchFromTCode 批量查找 租户编码
func (obj *_SgrTenantMgr) GetBatchFromTCode(tCodes []string) (results []*SgrTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenant{}).Where("`t_code` IN (?)", tCodes).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态
func (obj *_SgrTenantMgr) GetFromStatus(status int8) (results []*SgrTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenant{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态
func (obj *_SgrTenantMgr) GetBatchFromStatus(statuss []int8) (results []*SgrTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenant{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreationTime 通过creation_time获取内容 创建时间
func (obj *_SgrTenantMgr) GetFromCreationTime(creationTime time.Time) (results []*SgrTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenant{}).Where("`creation_time` = ?", creationTime).Find(&results).Error

	return
}

// GetBatchFromCreationTime 批量查找 创建时间
func (obj *_SgrTenantMgr) GetBatchFromCreationTime(creationTimes []time.Time) (results []*SgrTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenant{}).Where("`creation_time` IN (?)", creationTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 修改时间
func (obj *_SgrTenantMgr) GetFromUpdateTime(updateTime time.Time) (results []*SgrTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenant{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 修改时间
func (obj *_SgrTenantMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*SgrTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenant{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SgrTenantMgr) FetchByPrimaryKey(id uint64) (result SgrTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenant{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchUniqueByUnCode primary or index 获取唯一内容
func (obj *_SgrTenantMgr) FetchUniqueByUnCode(tCode string) (result SgrTenant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenant{}).Where("`t_code` = ?", tCode).Find(&result).Error

	return
}
