# Y4-Lang

![](https://img.shields.io/github/license/4ra1n/y4-lang)
![](https://img.shields.io/github/languages/top/4ra1n/y4-lang)
![](https://img.shields.io/github/v/release/4ra1n/y4-lang)
![](https://img.shields.io/github/downloads/4ra1n/y4-lang/total)
![](https://img.shields.io/github/actions/workflow/status/4ra1n/y4-lang/y4-lang.yml?branch=master)
![](https://img.shields.io/badge/Code%20Lines-8865-blue)

[更新日志 - Change Log](CHANGELOG.md)

`Y4-Lang` 是一个基于 `Golang` 的中文编程语言（解释型脚本语言）

不使用任何第三方库，仅依赖 `Golang` 标准库实现，语法类似 `Python` 简单易用

主要特性：
- 不使用任何库从零实现词法分析，语法分析，解释执行
- 类型包含 `int/float/bool/string` 和 `object`
- 支持 `if/else/while/continue/break` 等基本语法
- 支持数组类型以及 `list` 和 `map` 等高级结构
- 支持通过 `#include` 语法导入多个脚本文件执行
- 支持通过 `def` 语法定义函数和执行
- 支持 `http` 和 `base64` 等常见的库（可扩展）

一个简单的端口扫描

```text
#引入 "骚客特"

函数 扫端口(主机, 端口) {
    连接 = 骚客特.连接(主机, 端口);
    如果 连接 == 空的 {
        返回 假的;
    }
    骚客特.关闭(连接);
    打印("[+] 端口 " + 端口 + " 开启");
    返回 真的;
}

函数 主函数() {
    目标IP = "127.0.0.1";
    目标开始端口 = 8070;
    目标结束端口 = 8090;
    打印("[*] 开始TCP端口扫描 " + 目标IP);
    循环 端口 = 目标开始端口; 端口 <= 目标结束端口; 端口=端口+1 {
        打印("[*] 开始扫描端口: " + 端口);
        扫端口(目标IP, 端口);
    }
    打印("[*] 目标IP: " + 目标IP + " 扫描完成");
}
```