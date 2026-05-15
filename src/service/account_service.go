package service

import (
	"geji/dao"
	"geji/data"
	"geji/model"
	"geji/util"
	"time"

	"github.com/gin-gonic/gin"
)

type AccountService struct {
}

var AccSvr AccountService

func init() {
	AccSvr = AccountService{}
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
		Avater:     req.Avatar,
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

}
