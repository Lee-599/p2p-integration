package account

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/wailsapp/wails/lib/logger"
	"gorm.io/gorm"
	"link-server/config"
	"link-server/utils"
)

type ServiceImpl struct {
	log      *logger.CustomLogger
	db       *gorm.DB
	httpUtil *utils.HttpUtil
}

func NewServiceImpl(db *gorm.DB, httpUtil *utils.HttpUtil) ServiceImpl {
	log := logger.NewCustomLogger("Module_Account")
	return ServiceImpl{log, db, httpUtil}
}

// GetAccount 获取账户信息
func (a *ServiceImpl) GetAccount() (Account, error) {
	var account Account
	result := a.db.First(&account)
	if result.Error != nil {
		a.log.Error("GetAccount error")
	}
	return account, result.Error
}

// GetCode 获取验证码
func (a *ServiceImpl) GetCode(mobile string) error {
	res, err := a.httpUtil.NewRequest().
		SetQueryParam("mobile", mobile).
		Get(config.HttpGetCode)
	if err != nil {
		a.log.Errorf("GetResourceList http error: %s\n", err)
		return err
	}
	if !res.IsSuccess() {
		a.log.Errorf("GetResourceList Response error: %s\n", res)
		return fmt.Errorf("GetResourceList Response error: %s", res)
	}
	return nil
}

// SaveAccount 保存账户信息
func (a *ServiceImpl) SaveAccount(account *Account) {
	u, _ := a.GetAccount()
	//保存或更新账户
	u.Username = account.Username
	u.Phone = account.Phone
	u.UserId = account.UserId
	u.PublicKey = account.PublicKey
	a.db.Save(&u)
}

// DeleteAccount 删除账户信息
func (a *ServiceImpl) DeleteAccount() {
	a.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Account{})
}

// Login 用户登录
func (a *ServiceImpl) Login(mobile string, smsCode string) (*Account, error) {
	// 进行登陆
	var result map[string]interface{}
	res, err := a.httpUtil.NewRequest().
		SetHeader("Authorization", "Basic dG50bGlua2luZzp0bnRsaW5raW5nMTIzKio=").
		SetQueryParam("grant_type", "sms_code").
		SetQueryParam("scope", "all").
		SetQueryParam("smsCode", smsCode).
		SetQueryParam("mobile", mobile).
		SetQueryParam("device", "PC").
		SetResult(&result).
		Post(config.HttpLogin)
	if err != nil {
		a.log.Errorf("GetResourceList http error: %s\n", err)
		return nil, err
	}
	if !res.IsSuccess() {
		a.log.Errorf("GetResourceList Response error: %s\n", res)
		return nil, fmt.Errorf("GetResourceList Response error: %s", res)
	}
	token := fmt.Sprint(result["access_token"])
	userId, err := a.getUserIdFromToken(token)
	account, _ := a.GetAccount()
	//登陆成功后保存账户
	account.Username = mobile
	account.Phone = mobile
	account.UserId = userId
	a.db.Save(&account)
	return &account, nil
}

// Logout 用户登出
func (a *ServiceImpl) Logout() {
	a.DeleteAccount()
}

func (a *ServiceImpl) getUserIdFromToken(accessToken string) (string, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(accessToken, jwt.MapClaims{})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return fmt.Sprint(claims["user_id"]), nil
	}
	return "", errors.New("cannot parse accessToken")
}
