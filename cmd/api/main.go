package main

import (
	"test/internal/api/api"
	"test/internal/api/store"
	"test/internal/config"
)

func main() {
	port := "9090"
	//初始化代码配置
	_, err := config.InitConfig()
	if err != nil {
		panic("配置初始化失败" + err.Error())
	}
	//初始化数据层
	factory, err := store.GetFactory()
	if err != nil {
		panic("数据层初始化失败" + err.Error())
	}
	//初始化路由
	router, err := api.RouterInit(factory)
	err = router.Run(":" + port)
	if err != nil {
		panic(err)
	}
}
