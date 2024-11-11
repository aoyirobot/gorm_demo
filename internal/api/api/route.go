package api

import (
	"github.com/gin-gonic/gin"
	"test/internal/api/store"
	"test/internal/pkg/middle"
)

func RouterInit(factory store.Factory) (*gin.Engine, error) {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//全局异常监控
	router.Use(middle.ErrHandler())
	router.Use(middle.CORS())

	store.SetClient(factory)
	return router, nil
}
