package app

import (
	"github.com/wailsapp/wails"
	"link-server/module/account"
	"link-server/module/p2p"
)

type Login struct {
	log            *wails.CustomLogger
	accountService account.Service
	p2pService     p2p.Service
}

func NewLoginApp(service account.Service, p2pService p2p.Service) Login {
	return Login{
		accountService: service,
		p2pService:     p2pService,
	}
}

func (s *Login) WailsInit(runtime *wails.Runtime) error {
	s.log = runtime.Log.New("Login")
	return nil
}

// Login 用户登录
func (s *Login) Login(mobile string, smsCode string) (*account.Account, error) {
	return s.accountService.Login(mobile, smsCode)
}

// Logout 用户登出
func (s *Login) Logout() {
	//清除用户信息
	s.accountService.Logout()
	//断掉P2P链接
	_ = s.p2pService.Destroy()
}

// GetCode 发送验证码
func (s *Login) GetCode(mobile string) (bool, error) {
	err := s.accountService.GetCode(mobile)
	if err != nil {
		return false, err
	}
	return true, nil
}
