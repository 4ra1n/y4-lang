# Y4-Lang

![](https://img.shields.io/github/license/4ra1n/y4-lang)
![](https://img.shields.io/github/languages/top/4ra1n/y4-lang)
![](https://img.shields.io/github/v/release/4ra1n/y4-lang)
![](https://img.shields.io/github/actions/workflow/status/4ra1n/y4-lang/y4-lang.yml?branch=master)
![](https://img.shields.io/badge/Code%20Lines-7929-blue)

[英文版本 - English Version](README.md)

`Y4-Lang` 是一个基于 `Golang` 的编程语言（解释型脚本语言）

不使用任何第三方库，仅依赖 `Golang` 标准库实现，语法类似 `Python` 简单易用

```python
def hello(a, b) {
    if a < b {
        print("hello world");
    } else {
        print("error");    
    }
}

array = [1, 2, 3];
b = 2;
hello(array[0], b);
```

主要特性：
- 从零实现词法分析，语法分析，遍历 `AST` 解释执行的过程
- 类型包含 `int/float/bool/string` 和 `object`
- 支持 `if/else/while/continue/break` 等基本语法
- 支持数组类型以及 `list` 和 `map` 等高级结构
- 支持通过 `#include` 语法导入多个脚本文件执行
- 支持通过 `def` 语法定义函数和执行
- 支持 `http` 和 `base64` 等常见的库（可扩展）

使用 `Github Actions` 编译多版本可执行文件：
- windows (arm/arm64/386/amd64)
- darwin (arm64/amd64)
- linux (arm/arm64/386/amd64)

即将支持 `VSCode` 插件一键安装运行

## 快速开始

（1）从 `Release` 页面下载你需要的二进制文件

（2）使用 `-f` 参数指定 `y4` 文件（要求后缀必须是`.y4`）

将 `README` 开头部分代码保存到 `test.y4` 并输入

```shell
./y4lang -f test.y4
```

输出

```text
__  ____ __        __
\ \/ / // /       / /  ____ _____  ____ _
 \  / // /_______/ /  / __ `/ __ \/ __ `/
 / /__  __/_____/ /__/ /_/ / / / / /_/ /
/_/  /_/       /_____|__,_/_/ /_/\__, /
                                /____/
y4-lang v0.0.1
y4-lang is a script language based on golang
project: https://github.com/4ra1n/y4-lang
hello world
y4-lang run finish
```

（3）可选参数

- 使用 `--log-level` 指定日志级别（默认 `error` 级别）
- 使用 `--pool-size` 指定线程池大小（不建议改动）
- 使用 `--env-size` 指定默认环境大小（不建议改动）

使用 `--log-level debug` 查看 `AST` 结构，用于反馈 `BUG` 或调试开发
