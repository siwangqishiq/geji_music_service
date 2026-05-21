package service

import (
	"fmt"
	"geji/component"
	"geji/dao"
	"geji/data"
	"geji/model"
	"geji/util"
	"time"

	"github.com/gin-gonic/gin"
)

type AccountService struct {
	IdTokens map[int64]string
}

var AccSvr AccountService

func init() {
	AccSvr = AccountService{
		IdTokens: make(map[int64]string),
	}
}

func (a *AccountService) AccountCreate(c *gin.Context, req *model.CreateAccountReq) {
	acc, err := dao.QueryAccountByAccount(dao.DB, req.Account)
	if err != nil {
		util.Loge("account error %v", err.Error())
		util.Fail(c, data.ERR_CODE_DATABASE_ERROR, err.Error())
		return
	}

	if acc != nil {
		util.Loge("account error %v", "account 重复")
		util.Fail(c, data.ERR_CODE_DATABASE_ERROR, acc.Account+"已存在")
		return
	}

	var nowTime = time.Now()
	var account dao.AccountModel = dao.AccountModel{
		Account:    req.Account,
		Password:   req.Password,
		Nickname:   req.Nickname,
		Avatar:     req.Avatar,
		Status:     0,
		CreateTime: nowTime,
		UpdateTime: nowTime,
	}

	err = dao.InsertAccount(dao.DB, &account)
	if err != nil {
		util.Fail(c, data.ERR_CODE_DATABASE_ERROR, err.Error())
		return
	}

	util.Logi("account %s register success uid %d", account.Account, account.Uid)
	util.Success(c, "注册成功")
}

func (a *AccountService) Login(c *gin.Context, req *model.LoginReq) {
	account, err := dao.QueryAccountByAccount(dao.DB, req.Account)
	if err != nil {
		util.Fail(c, data.ERR_CODE_DATABASE_ERROR, err.Error())
		return
	}

	if account == nil || account.Uid <= 0 {
		util.Fail(c, data.ERR_CODE_ACCOUNT_NOT_EXIST, "账户不存在")
		return
	}

	util.Logi("account req %s origin %s", req.Password, account.Password)
	if account.Password != req.Password {
		util.Fail(c, data.ERR_CODE_LOGIN_PWDERROR, "密码错误")
		return
	}

	token, err := component.GenerateToken(account.Uid)
	if err != nil {
		util.Fail(c, data.ERR_CODE_LOGIN_GEN_TOKEN_FAILED, "登录token生成错误")
		return
	}

	var resp = model.LoginResp{
		Token:    token,
		Uid:      account.Uid,
		Account:  account.Account,
		Avatar:   account.Avatar,
		Age:      account.Age,
		Nickname: account.Nickname,
		Remark:   account.Remark,
	}

	util.Logi("account %s %d 登录成功 nickname %v avatar %v",
		resp.Account, resp.Uid, resp.Nickname, resp.Avatar)

	// AccSvr.IdTokens[resp.Uid] = resp.Token
	KVSvr.Put(fmt.Sprintf(data.KVCACHE_ONLINE, resp.Uid), resp.Token)
	util.Success(c, resp)
}

func (a *AccountService) Logout(c *gin.Context, userClaims component.UserClaims, reqToken string) {
	util.Logi("logout uid : %v req token: %v", userClaims.Uid, reqToken)
	onlineToken := KVSvr.Get(fmt.Sprintf(data.KVCACHE_ONLINE, userClaims.Uid))
	util.Logi("online Token :%s", onlineToken)

	if onlineToken == reqToken {
		KVSvr.Del(fmt.Sprintf(data.KVCACHE_ONLINE, userClaims.Uid))
		util.Success(c, "注销成功")
	} else {
		util.Fail(c, data.ERR_CODE_TOKEN_MISTAKEN, "token不一致")
	}
}
