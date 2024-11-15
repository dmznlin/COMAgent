// Package main
/******************************************************************************
  作者: dmzn@163.com 2024-11-13 11:08:09
  描述: 沁恒串口服务SDK
******************************************************************************/
package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"github.com/dmznlin/znlib-go/znlib"
	"github.com/dmznlin/znlib-go/znlib/restruct"
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
	//gin.SetMode(gin.ReleaseMode)
	//release

	router := gin.Default()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
			"name":    "dmzn",
		})
	})

	router.GET("/test", func(c *gin.Context) {
		cfg := DEVICEHW_CONFIG{
			Modulename: "11111",
			Username:   "22222",
			PassWord:   "33333",
			DevIP:      [4]uchar{192, 168, 1, 2},
		}

		str, err := json.Marshal(cfg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		znlib.Info(string(str))
		//log

		buf, err := restruct.Pack(binary.BigEndian, &cfg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		c.String(http.StatusOK, "data: %v", buf)
	})

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			znlib.ErrorCaller(err, "main.ListenAndServe")
		}
	}()

	znlib.WaitSystemExit(func() error {
		if err := srv.Shutdown(znlib.Application.Ctx); err != nil {
			znlib.ErrorCaller(err, "main.ServerShutdown")
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
