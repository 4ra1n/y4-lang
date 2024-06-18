package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/4ra1n/y4-lang/cli"
	"github.com/4ra1n/y4-lang/color"
)

// 全局的版本号和构建时间
func init() {
	// VERSION
	cli.Version = "v0.0.3"
	// BUILD TIME
	cli.BuildTime = "2024/02/01"
}

// Y4-Lang
// y4-lang is a script language based on golang
// Y4 Lang 是一个基于 Golang 的解释型脚本语言
// 不使用任何第三方的库完全手撸，主打一个造轮子
func main() {
	// 这个地方的问题是 Windows 老 CMD 不支持彩色
	// 咱们直接给他设置 ENABLE_VIRTUAL_TERMINAL_PROCESSING
	if !color.IsSupported() {
		// 最终如果还是失败直接禁用颜色
		color.DisableColor()
	}
	// 创建 Context 全局控制
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()
	// 启动
	go cli.Start(cancel)
	// 优雅退出
	sigChan := make(chan os.Signal, 1)
	// 主要是处理 Ctrl+C 的退出
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		// 一个是退出信号
		case <-sigChan:
			if !cli.QuietFlag {
				fmt.Println("ctrl+c stop")
			}
			return
		// 另外是 ctx 结束信号
		case <-ctx.Done():
			if !cli.QuietFlag {
				fmt.Println("y4-lang run finish")
			}
			return
		}
	}
}
