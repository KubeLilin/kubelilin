package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SgrTenantRoleMgr struct {
	*_BaseMgr
}

// SgrTenantRoleMgr open func
func SgrTenantRoleMgr(db *gorm.DB) *_SgrTenantRoleMgr {
	if db == nil {
		panic(fmt.Errorf("SgrTenantRoleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SgrTenantRoleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sgr_tenant_role"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SgrTenantRoleMgr) GetTableName() string {
	return "sgr_tenant_role"
}

// Get 获取
func (obj *_SgrTenantRoleMgr) Get() (result SgrTenantRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantRole{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SgrTenantRoleMgr) Gets() (results []*SgrTenantRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantRole{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SgrTenantRoleMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SgrTenantRole{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_SgrTenantRoleMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithRoleCode role_code获取 角色编码
func (obj *_SgrTenantRoleMgr) WithRoleCode(roleCode string) Option {
	return optionFunc(func(o *options) { o.query["role_code"] = roleCode })
}

// WithRoleName role_name获取 角色名称
func (obj *_SgrTenantRoleMgr) WithRoleName(roleName string) Option {
	return optionFunc(func(o *options) { o.query["role_name"] = roleName })
}

// WithStatus status获取 状态
func (obj *_SgrTenantRoleMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithTenantID tenant_id获取 租户
func (obj *_SgrTenantRoleMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithCreatetionTime createtion_time获取
func (obj *_SgrTenantRoleMgr) WithCreatetionTime(createtionTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["createtion_time"] = createtionTime })
}

// WithUpdateTime update_time获取
func (obj *_SgrTenantRoleMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_SgrTenantRoleMgr) GetByOption(opts ...Option) (result SgrTenantRole, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantRole{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SgrTenantRoleMgr) GetByOptions(opts ...Option) (results []*SgrTenantRole, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantRole{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_SgrTenantRoleMgr) GetFromID(id uint64) (result SgrTenantRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantRole{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_SgrTenantRoleMgr) GetBatchFromID(ids []uint64) (results []*SgrTenantRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantRole{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromRoleCode 通过role_code获取内容 角色编码
func (obj *_SgrTenantRoleMgr) GetFromRoleCode(roleCode string) (results []*SgrTenantRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantRole{}).Where("`role_code` = ?", roleCode).Find(&results).Error

	return
}

// GetBatchFromRoleCode 批量查找 角色编码
func (obj *_SgrTenantRoleMgr) GetBatchFromRoleCode(roleCodes []string) (results []*SgrTenantRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantRole{}).Where("`role_code` IN (?)", roleCodes).Find(&results).Error

	return
}

// GetFromRoleName 通过role_name获取内容 角色名称
func (obj *_SgrTenantRoleMgr) GetFromRoleName(roleName string) (results []*SgrTenantRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantRole{}).Where("`role_name` = ?", roleName).Find(&results).Error

	return
}

// GetBatchFromRoleName 批量查找 角色名称
func (obj *_SgrTenantRoleMgr) GetBatchFromRoleName(roleNames []string) (results []*SgrTenantRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantRole{}).Where("`role_name` IN (?)", roleNames).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态
func (obj *_SgrTenantRoleMgr) GetFromStatus(status int8) (results []*SgrTenantRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantRole{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态
func (obj *_SgrTenantRoleMgr) GetBatchFromStatus(statuss []int8) (results []*SgrTenantRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantRole{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容 租户
func (obj *_SgrTenantRoleMgr) GetFromTenantID(tenantID int64) (results []*SgrTenantRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantRole{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找 租户
func (obj *_SgrTenantRoleMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*SgrTenantRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantRole{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromCreatetionTime 通过createtion_time获取内容
func (obj *_SgrTenantRoleMgr) GetFromCreatetionTime(createtionTime time.Time) (results []*SgrTenantRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantRole{}).Where("`createtion_time` = ?", createtionTime).Find(&results).Error

	return
}

// GetBatchFromCreatetionTime 批量查找
func (obj *_SgrTenantRoleMgr) GetBatchFromCreatetionTime(createtionTimes []time.Time) (results []*SgrTenantRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantRole{}).Where("`createtion_time` IN (?)", createtionTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_SgrTenantRoleMgr) GetFromUpdateTime(updateTime time.Time) (results []*SgrTenantRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantRole{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_SgrTenantRoleMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*SgrTenantRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantRole{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SgrTenantRoleMgr) FetchByPrimaryKey(id uint64) (result SgrTenantRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantRole{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchUniqueIndexByUnRoleCodeName primary or index 获取唯一内容
func (obj *_SgrTenantRoleMgr) FetchUniqueIndexByUnRoleCodeName(roleCode string, roleName string) (result SgrTenantRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantRole{}).Where("`role_code` = ? AND `role_name` = ?", roleCode, roleName).Find(&result).Error

	return
}
