package resource

import (
	"fmt"
	"link-server/utils"
	"testing"
)

func TestHttpService_GetResourceList(t *testing.T) {
	service := NewServiceImpl(nil, utils.NewHttp())
	list, err := service.GetResourceList()
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(list)
}
