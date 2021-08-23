package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SgrTenantUserRoleMgr struct {
	*_BaseMgr
}

// SgrTenantUserRoleMgr open func
func SgrTenantUserRoleMgr(db *gorm.DB) *_SgrTenantUserRoleMgr {
	if db == nil {
		panic(fmt.Errorf("SgrTenantUserRoleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SgrTenantUserRoleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sgr_tenant_user_role"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SgrTenantUserRoleMgr) GetTableName() string {
	return "sgr_tenant_user_role"
}

// Get 获取
func (obj *_SgrTenantUserRoleMgr) Get() (result SgrTenantUserRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUserRole{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SgrTenantUserRoleMgr) Gets() (results []*SgrTenantUserRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUserRole{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SgrTenantUserRoleMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SgrTenantUserRole{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_SgrTenantUserRoleMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserID user_id获取 用户id
func (obj *_SgrTenantUserRoleMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithRoleID role_id获取 角色id
func (obj *_SgrTenantUserRoleMgr) WithRoleID(roleID int64) Option {
	return optionFunc(func(o *options) { o.query["role_id"] = roleID })
}

// WithCreationTime creation_time获取
func (obj *_SgrTenantUserRoleMgr) WithCreationTime(creationTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["creation_time"] = creationTime })
}

// WithUpdateTime update_time获取
func (obj *_SgrTenantUserRoleMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_SgrTenantUserRoleMgr) GetByOption(opts ...Option) (result SgrTenantUserRole, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUserRole{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SgrTenantUserRoleMgr) GetByOptions(opts ...Option) (results []*SgrTenantUserRole, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUserRole{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_SgrTenantUserRoleMgr) GetFromID(id uint64) (result SgrTenantUserRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUserRole{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_SgrTenantUserRoleMgr) GetBatchFromID(ids []uint64) (results []*SgrTenantUserRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUserRole{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户id
func (obj *_SgrTenantUserRoleMgr) GetFromUserID(userID int64) (results []*SgrTenantUserRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUserRole{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找 用户id
func (obj *_SgrTenantUserRoleMgr) GetBatchFromUserID(userIDs []int64) (results []*SgrTenantUserRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUserRole{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromRoleID 通过role_id获取内容 角色id
func (obj *_SgrTenantUserRoleMgr) GetFromRoleID(roleID int64) (results []*SgrTenantUserRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUserRole{}).Where("`role_id` = ?", roleID).Find(&results).Error

	return
}

// GetBatchFromRoleID 批量查找 角色id
func (obj *_SgrTenantUserRoleMgr) GetBatchFromRoleID(roleIDs []int64) (results []*SgrTenantUserRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUserRole{}).Where("`role_id` IN (?)", roleIDs).Find(&results).Error

	return
}

// GetFromCreationTime 通过creation_time获取内容
func (obj *_SgrTenantUserRoleMgr) GetFromCreationTime(creationTime time.Time) (results []*SgrTenantUserRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUserRole{}).Where("`creation_time` = ?", creationTime).Find(&results).Error

	return
}

// GetBatchFromCreationTime 批量查找
func (obj *_SgrTenantUserRoleMgr) GetBatchFromCreationTime(creationTimes []time.Time) (results []*SgrTenantUserRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUserRole{}).Where("`creation_time` IN (?)", creationTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_SgrTenantUserRoleMgr) GetFromUpdateTime(updateTime time.Time) (results []*SgrTenantUserRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUserRole{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_SgrTenantUserRoleMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*SgrTenantUserRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUserRole{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SgrTenantUserRoleMgr) FetchByPrimaryKey(id uint64) (result SgrTenantUserRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUserRole{}).Where("`id` = ?", id).Find(&result).Error

	return
}
