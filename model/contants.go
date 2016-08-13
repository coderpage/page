package model

const (
	// User
	UserGroupNoActived = "nonactived" // 未激活状态
	UserGroupCommon    = "common"     // 一般
	UserGropAdmin      = "admin"      // 管理员
	// Auth
	AuthTypeUserActive   = "u-active"   // 激活类型 Token
	AuthTypeUserSignIn   = "u-signIn"   // 登录类型 Token
	AuthTypeUserFindPwd  = "u-findPwd"  // 找回密码类型 Token
	AuthTypeUserResetPwd = "u-resetPwd" // 重置密码类型 Token
	AuthStatusOK         = "ok"         // auth 状态为 OK
	// Article
	ArticleStatusPublish = "publish" // article 为发布状态
	ArticleStatusDraft   = "draft"   // article 为草稿状态

)
