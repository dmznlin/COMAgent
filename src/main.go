// Package main
/******************************************************************************
  作者: dmzn@163.com 2024-11-13 11:08:09
  描述: 沁恒串口服务SDK
******************************************************************************/
package main

import (
	"errors"
	"github.com/dmznlin/znlib-go/znlib"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
)

var (
	// 初始化lib
	_ = znlib.InitLib(nil, nil)

	//udp server
	server net.Conn
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	//release

	router := gin.Default()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			znlib.ErrorCaller(err, "main.ListenAndServe")
		}
	}()

	znlib.WaitSystemExit(func() error {
		if err := srv.Shutdown(znlib.Application.Ctx); err != nil {
			znlib.ErrorCaller(err, "Server Shutdown")
			return err
		}

		znlib.Info("COMAgent.gin Shutdown.")
		return nil
	}, func() error {
		if !znlib.IsNil(server) {
			server.Close()
		}

		znlib.Info("COMAgent.udp Shutdown.")
		return nil
	})
}
