package cli

import (
	"context"
	"io"
	"os"
	"strings"

	"github.com/4ra1n/y4-lang/chardet"
	"github.com/4ra1n/y4-lang/conf"
	"github.com/4ra1n/y4-lang/core"
	"github.com/4ra1n/y4-lang/lexer"
	"github.com/4ra1n/y4-lang/log"
	"github.com/4ra1n/y4-lang/pre"
)

func start(cancel context.CancelFunc) {
	log.Debug("start y4-lang")
	// 首先是找主函数入口
	var (
		err      error
		mainFile string
	)
	// 如果输入多个文件
	if len(filePath) > 1 {
		// 寻找多个文件里的主函数
		// 没有主函数的多个文件不允许执行
		mainFile, err = pre.SearchMain(filePath)
		if err != nil {
			log.Error(err)
			return
		}
	} else if len(filePath) == 1 {
		// 只输入一个文件允许直接执行
		mainFile = filePath[0]
	} else {
		// 输入文件长度是0实际应该不会到达这里
		log.Error("错误的文件输入")
		return
	}

	// 主文件名二次验证
	mainFile = strings.TrimSpace(mainFile)
	if mainFile == "" {
		log.Error("文件名是空")
		return
	}

	// 检查后缀必须是 Y4
	if !strings.HasSuffix(strings.ToLower(mainFile), ".y4") {
		log.Errorf("文件后缀名必须是.y4")
		return
	}

	// 检查文件存在
	file, err := os.Open(mainFile)
	if err != nil {
		if os.IsNotExist(err) {
			log.Errorf("文件不存在: %s", filePath)
		} else {
			log.Errorf("打开文件失败: %s", err.Error())
		}
		return
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Errorf("关闭文件错误: %s", err.Error())
		}
	}()

	// 检查 UTF-8
	// 只支持 UTF-8 编码
	encoding, err := chardet.DetectFileEncoding(mainFile)
	if encoding != "UTF-8" {
		log.Errorf("输入文件必须是 UTF-8 编码 (%s)", encoding)
		return
	}

	// 预处理器
	// 主要是处理开头的引入部分
	var newReader io.Reader
	if !conf.DisablePreProcess {
		ip := pre.NewIncludeProcessor(mainFile)
		// 引入其他本地文件则替换内容
		// 引入标准库不进行替换
		newReader = ip.Process()
	} else {
		newReader, err = os.Open(mainFile)
		if err != nil {
			log.Errorf("创建文件输入流失败")
			return
		}
	}

	// 词法分析
	l := lexer.NewLexer(newReader)
	// 创建解释器
	i := core.NewInterpreter(l, cancel)

	// 环境设置
	if envSize != 0 {
		log.Debugf("设置环境大小 %d", envSize)
		i.SetEnvSize(envSize)
	}

	if poolSize != 0 {
		log.Debugf("设置协程池大小 %d", poolSize)
		i.SetPoolSize(poolSize)
	}

	// 检测参数
	ok := i.Check()
	if !ok {
		log.Error("启动参数检查失败")
		return
	}

	// 解释执行
	i.Start()
}
