package app

import (
	"github.com/wailsapp/wails"
	"link-server/module/account"
	"link-server/module/p2p"
)

type Config struct {
	PublicKey string
	Port      int
	PeerId    string
}

type Setting struct {
	log            *wails.CustomLogger
	p2pService     p2p.Service
	accountService account.Service
}

func NewSettingApp(service p2p.Service, accountService account.Service) Setting {
	return Setting{
		p2pService:     service,
		accountService: accountService,
	}
}

func (s *Setting) WailsInit(runtime *wails.Runtime) error {
	s.log = runtime.Log.New("Setting")
	return nil
}

// GetSetting 查看配置信息
func (s *Setting) GetSetting() (*Config, error) {
	config := &Config{}
	//查询用户信息中的公钥
	info, err := s.accountService.GetAccount()
	if err != nil {
		return config, err
	}
	config.PublicKey = info.PublicKey
	//查询P2P设置中的设置信息
	p2pConfig, err := s.p2pService.GetSetting()
	if err != nil {
		return config, nil
	}
	config.Port = p2pConfig.Port
	config.PeerId = p2pConfig.PeerId
	return config, nil
}

// Setting 设置公钥信息
func (s *Setting) Setting(publicKey string) (bool, error) {
	accountInfo, err := s.accountService.GetAccount()
	if err != nil {
		return false, err
	}
	//设置用户公钥
	accountInfo.PublicKey = publicKey
	s.accountService.SaveAccount(&accountInfo)
	return true, nil
}

//InitP2pSetting 初始化p2p设置
func (s *Setting) InitP2pSetting() (bool, error) {
	//初始化配置
	err := s.p2pService.InitSetting()
	if err != nil {
		return false, err
	}
	return true, nil
}
