package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SgrTenantUserMgr struct {
	*_BaseMgr
}

// SgrTenantUserMgr open func
func SgrTenantUserMgr(db *gorm.DB) *_SgrTenantUserMgr {
	if db == nil {
		panic(fmt.Errorf("SgrTenantUserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SgrTenantUserMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sgr_tenant_user"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SgrTenantUserMgr) GetTableName() string {
	return "sgr_tenant_user"
}

// Get 获取
func (obj *_SgrTenantUserMgr) Get() (result SgrTenantUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUser{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SgrTenantUserMgr) Gets() (results []*SgrTenantUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUser{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SgrTenantUserMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SgrTenantUser{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_SgrTenantUserMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTenantID tenant_id获取 租户
func (obj *_SgrTenantUserMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithUserName user_name获取 用户名
func (obj *_SgrTenantUserMgr) WithUserName(userName string) Option {
	return optionFunc(func(o *options) { o.query["user_name"] = userName })
}

// WithAccount account获取 账号
func (obj *_SgrTenantUserMgr) WithAccount(account string) Option {
	return optionFunc(func(o *options) { o.query["account"] = account })
}

// WithPassword password获取 密码
func (obj *_SgrTenantUserMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithMobile mobile获取 手机
func (obj *_SgrTenantUserMgr) WithMobile(mobile string) Option {
	return optionFunc(func(o *options) { o.query["mobile"] = mobile })
}

// WithEmail email获取 邮箱
func (obj *_SgrTenantUserMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithStatus status获取 状态
func (obj *_SgrTenantUserMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreatetionTime createtion_time获取
func (obj *_SgrTenantUserMgr) WithCreatetionTime(createtionTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["createtion_time"] = createtionTime })
}

// WithUpdateTime update_time获取
func (obj *_SgrTenantUserMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_SgrTenantUserMgr) GetByOption(opts ...Option) (result SgrTenantUser, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUser{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SgrTenantUserMgr) GetByOptions(opts ...Option) (results []*SgrTenantUser, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUser{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_SgrTenantUserMgr) GetFromID(id uint64) (result SgrTenantUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUser{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_SgrTenantUserMgr) GetBatchFromID(ids []uint64) (results []*SgrTenantUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUser{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容 租户
func (obj *_SgrTenantUserMgr) GetFromTenantID(tenantID int64) (results []*SgrTenantUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUser{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找 租户
func (obj *_SgrTenantUserMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*SgrTenantUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUser{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromUserName 通过user_name获取内容 用户名
func (obj *_SgrTenantUserMgr) GetFromUserName(userName string) (results []*SgrTenantUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUser{}).Where("`user_name` = ?", userName).Find(&results).Error

	return
}

// GetBatchFromUserName 批量查找 用户名
func (obj *_SgrTenantUserMgr) GetBatchFromUserName(userNames []string) (results []*SgrTenantUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUser{}).Where("`user_name` IN (?)", userNames).Find(&results).Error

	return
}

// GetFromAccount 通过account获取内容 账号
func (obj *_SgrTenantUserMgr) GetFromAccount(account string) (results []*SgrTenantUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUser{}).Where("`account` = ?", account).Find(&results).Error

	return
}

// GetBatchFromAccount 批量查找 账号
func (obj *_SgrTenantUserMgr) GetBatchFromAccount(accounts []string) (results []*SgrTenantUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUser{}).Where("`account` IN (?)", accounts).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容 密码
func (obj *_SgrTenantUserMgr) GetFromPassword(password string) (results []*SgrTenantUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUser{}).Where("`password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找 密码
func (obj *_SgrTenantUserMgr) GetBatchFromPassword(passwords []string) (results []*SgrTenantUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUser{}).Where("`password` IN (?)", passwords).Find(&results).Error

	return
}

// GetFromMobile 通过mobile获取内容 手机
func (obj *_SgrTenantUserMgr) GetFromMobile(mobile string) (results []*SgrTenantUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUser{}).Where("`mobile` = ?", mobile).Find(&results).Error

	return
}

// GetBatchFromMobile 批量查找 手机
func (obj *_SgrTenantUserMgr) GetBatchFromMobile(mobiles []string) (results []*SgrTenantUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUser{}).Where("`mobile` IN (?)", mobiles).Find(&results).Error

	return
}

// GetFromEmail 通过email获取内容 邮箱
func (obj *_SgrTenantUserMgr) GetFromEmail(email string) (results []*SgrTenantUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUser{}).Where("`email` = ?", email).Find(&results).Error

	return
}

// GetBatchFromEmail 批量查找 邮箱
func (obj *_SgrTenantUserMgr) GetBatchFromEmail(emails []string) (results []*SgrTenantUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUser{}).Where("`email` IN (?)", emails).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态
func (obj *_SgrTenantUserMgr) GetFromStatus(status int8) (results []*SgrTenantUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUser{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态
func (obj *_SgrTenantUserMgr) GetBatchFromStatus(statuss []int8) (results []*SgrTenantUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUser{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreatetionTime 通过createtion_time获取内容
func (obj *_SgrTenantUserMgr) GetFromCreatetionTime(createtionTime time.Time) (results []*SgrTenantUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUser{}).Where("`createtion_time` = ?", createtionTime).Find(&results).Error

	return
}

// GetBatchFromCreatetionTime 批量查找
func (obj *_SgrTenantUserMgr) GetBatchFromCreatetionTime(createtionTimes []time.Time) (results []*SgrTenantUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUser{}).Where("`createtion_time` IN (?)", createtionTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_SgrTenantUserMgr) GetFromUpdateTime(updateTime time.Time) (results []*SgrTenantUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUser{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_SgrTenantUserMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*SgrTenantUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUser{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SgrTenantUserMgr) FetchByPrimaryKey(id uint64) (result SgrTenantUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SgrTenantUser{}).Where("`id` = ?", id).Find(&result).Error

	return
}
