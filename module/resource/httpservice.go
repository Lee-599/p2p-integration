package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/wailsapp/wails"
	"github.com/wailsapp/wails/lib/logger"
	"gorm.io/gorm"
	"io"
	"link-server/config"
	"link-server/utils"
	"log"
	"net/http"
)

type HttpService struct {
	log      *wails.CustomLogger
	db       *gorm.DB
	httpUtil *utils.HttpUtil
}

func NewServiceImpl(db *gorm.DB, httpUtil *utils.HttpUtil) HttpService {
	log := logger.NewCustomLogger("Module_P2P")
	return HttpService{log, db, httpUtil}
}

// GetResourceList 查询我的资源列表
func (r *HttpService) GetResourceList() ([]Resource, error) {
	r.log.Info("start GetResourceList")
	var resources []Resource
	// 通过http请求获得我的资源列表
	res, err := r.httpUtil.NewRequest().
		SetResult(&resources).
		Get(config.HttpGetResource)
	if err != nil {
		r.log.Errorf("GetResourceList http error: %s\n", err)
	}
	if !res.IsSuccess() {
		r.log.Errorf("GetResourceList Response error: %s\n", res.Status())
	}
	return resources, err
}

// RegisterResource 注册资源
func (r *HttpService) RegisterResource(resource Resource) error {

	err := r.UnRegisterResource(resource.PeerId)
	if err != nil {
		r.log.Warn("注册前取消注册失败")
	}

	data, err := json.Marshal(resource)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", config.ResourceApiGateway+"/api/resources", bytes.NewBuffer(data))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	r.log.Info("注册p2p地址成功")
	return nil
}

// UnRegisterResource 取消注册资源
func (r *HttpService) UnRegisterResource(peerid string) error {

	resoure, err := r.LoadRegistryInfoFromChain(peerid)

	if err != nil {
		return err
	}

	if resoure == nil {
		return nil
	}

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/resources/%d", config.ResourceApiGateway, resoure.ID), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// 从链上加载注册信息
func (r *HttpService) LoadRegistryInfoFromChain(peerid string) (*Resource, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", config.ResourceApiGateway+"/api/resource/peer?peerId="+peerid, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	//解析body
	var info Resource
	err = json.Unmarshal(body, &info)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &info, nil
}
