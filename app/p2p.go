package app

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails"
	"link-server/module/p2p"
	"link-server/module/resource"
	"os"
	"time"
)

type P2p struct {
	log             *wails.CustomLogger
	p2pServer       p2p.Service
	resourceService resource.Service
	ctx             context.Context
	cancel          func()
}

type ShareResult struct {
	IsShare bool
	Port    string
}

func NewP2pApp(service p2p.Service, resourceService resource.Service) P2p {
	return P2p{
		p2pServer:       service,
		resourceService: resourceService,
	}
}

func (s *P2p) WailsInit(runtime *wails.Runtime) error {
	s.ctx, s.cancel = context.WithCancel(context.Background())
	s.log = runtime.Log.New("P2P")
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("child process interrupt...")
				return
			default:
				runtime.Events.Emit("Links", s.p2pServer.GetLinks())
				time.Sleep(5 * time.Second)
			}
		}
	}(s.ctx)
	return nil
}

// IsP2PSetting 判断p2p信息是否配置
func (s *P2p) IsP2PSetting() bool {
	config, err := s.p2pServer.GetSetting()
	if err != nil {
		return false
	}
	return config.PrivateKey != ""
}

// Link P2P链接
func (s *P2p) Link(port int, peerId string) (bool, error) {
	//进行P2P链接
	err := s.p2pServer.Link(port, peerId)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *P2p) IsSharing() (*ShareResult, error) {
	p2pConfig, err := s.p2pServer.GetSetting()
	if err != nil {
		return nil, err
	}
	peerId := p2pConfig.PeerId
	isShare, port := s.p2pServer.IsShare(peerId)
	return &ShareResult{
		IsShare: isShare,
		Port:    port,
	}, nil
}

// Share 分享p2p连接
func (s *P2p) Share(target, hostname string) (bool, error) {

	s.log.Infof("target is : %s", target)

	p2pConfig, err := s.p2pServer.GetSetting()
	if err != nil {
		return false, err
	}
	if hostname == "" {
		hostname, err = os.Hostname()
		if err != nil {
			return false, err
		}
	}

	r := resource.Resource{
		PeerId:  p2pConfig.PeerId,
		Creator: hostname,
		VmType:  hostname,
	}
	err = s.p2pServer.Share(target)
	if err != nil {
		return false, err
	}
	err = s.resourceService.RegisterResource(r)
	if err != nil {
		return false, err
	}
	return true, nil
}

// UnShare 取消分享p2p连接
func (s *P2p) UnShare() (bool, error) {
	p2pConfig, err := s.p2pServer.GetSetting()
	if err != nil {
		return false, err
	}

	_, err = s.p2pServer.Close(fmt.Sprintf("/p2p/%s", p2pConfig.PeerId))
	if err != nil {
		s.log.Error("关闭p2p共享失败")
		return false, err
	}
	err = s.resourceService.UnRegisterResource(p2pConfig.PeerId)
	if err != nil {
		s.log.Error("取消注册p2p共享失败")
	}
	return err != nil, err
}

// CloseLink 关闭链接
func (s *P2p) CloseLink(target string) (int, error) {
	//断开P2P
	return s.p2pServer.Close(target)
}

// GetLinkStatus 查询P2P链接状态
func (s *P2p) GetLinkStatus() *[]p2p.LinkInfo {
	return s.p2pServer.GetLinks()
}

// WailsShutdown 关闭链接
func (s *P2p) WailsShutdown() {
	s.log.Info("正在关闭，销毁所有p2p连接")
	_ = s.p2pServer.Destroy()
	_, _ = s.UnShare()
	s.cancel()
}
