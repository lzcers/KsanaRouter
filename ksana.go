package ksana

import (
	"strings"
)

type cmdParams struct {
	addr string
	port string
}
type hookFunc func() error

// 框架运行模式
const (
	VERSION = "0.0.1" // 框架版本
	DEV     = "dev"   // 开发模式
	PROD    = "prod"  // 生产模式
)

var (
	hooks  = make([]hookFunc, 0)
	params = new(cmdParams)
)

// AddAPPStartHook 挂载框架启动的模块
func AddAPPStartHook(hf hookFunc) {
	hooks = append(hooks, hf)
}

// 从参数中取出主机地址和端口号
//  host:port
func parseCMDParams(args []string) {
	if len(args) > 0 && args[0] != "" {
		strs := strings.Split(args[0], ":")
		// 主机地址
		if len(strs) > 0 && strs[0] != "" {
			params.addr = strs[0]
		}
		// 端口号
		if len(strs) > 1 && strs[1] != "" {
			params.port = strs[1]
		}
	}
}

// Run 框架启动
func Run(params ...string) {
	// 1. 获取命令行参数
	parseCMDParams(params)
	// 2. 基于获取到的参数，执行一堆模块的初始化工作
	for _, hk := range hooks {
		if err := hk(); err != nil {
			panic(err)
		}
	}
	// return http.ListenAndServe(params.addr, params.port, )
}
