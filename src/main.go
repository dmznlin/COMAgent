// Package main
/******************************************************************************
  作者: dmzn@163.com 2024-11-13 11:08:09
  描述: 沁恒串口服务SDK
******************************************************************************/
package main

import (
	"github.com/dmznlin/znlib-go/znlib"
)

// 初始化lib
var _ = znlib.InitLib(nil, nil)

func main() {
	znlib.Info("hello,sdk")
}
