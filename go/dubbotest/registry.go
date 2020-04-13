package main

import (
	"fmt"
	"strconv"

	dubbocommon "github.com/mosn/registry/dubbo/common"
	dubboconsts "github.com/mosn/registry/dubbo/common/constant"
	"github.com/mosn/registry/dubbo/common/logger"
	registry "github.com/mosn/registry/dubbo/zookeeper"
)

func main() {
	logger.InitLog("./log.log")
	var (
		dubboPathInRegistry = fmt.Sprintf("dubbo://127.0.0.1:20000/%v.%v", "com.test.cch", "GetUser")
		registryPath        = fmt.Sprintf("registry://%v", "127.0.0.1:2181")
	)

	registryURL, err := dubbocommon.NewURL(registryPath,
		dubbocommon.WithParamsValue(dubboconsts.ROLE_KEY, strconv.Itoa(dubbocommon.PROVIDER)),
	)

	reg, err := registry.NewZkRegistry(&registryURL)

	if err != nil {
		fmt.Println("pub fail")
		return
	}

	url, err := dubbocommon.NewURL(dubboPathInRegistry,
		//dubbocommon.WithParamsValue(dubboconsts.CLUSTER_KEY, "cluster"), // need to read from user config
		dubbocommon.WithParamsValue("serviceid", "com.test.aaa"),
		dubbocommon.WithMethods([]string{"GetUser"}))
	if err != nil {
		fmt.Println("pub fail")
		return
	}

	fmt.Println("#####", url)
	err = reg.Register(url)

}
