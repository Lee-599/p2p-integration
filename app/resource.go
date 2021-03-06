package app

import (
	"github.com/wailsapp/wails"
	"link-server/module/account"
	"link-server/module/resource"
)

type Resource struct {
	log             *wails.CustomLogger
	resourceService resource.Service
	accountService  account.Service
}

func NewResourceApp(resourceService resource.Service, accountService account.Service) Resource {
	return Resource{
		resourceService: resourceService,
		accountService:  accountService,
	}
}

func (s *Resource) WailsInit(runtime *wails.Runtime) error {
	s.log = runtime.Log.New("Resource")
	return nil
}

func (s *Resource) GetResources() ([]resource.Resource, error) {
	list, err := s.resourceService.GetResourceList()
	return list, err
}
