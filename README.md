# Y4-Lang

![](https://img.shields.io/github/license/4ra1n/y4-lang)
![](https://img.shields.io/github/languages/top/4ra1n/y4-lang)
![](https://img.shields.io/github/v/release/4ra1n/y4-lang)
![](https://img.shields.io/github/downloads/4ra1n/y4-lang/total)
![](https://img.shields.io/github/actions/workflow/status/4ra1n/y4-lang/y4-lang.yml?branch=master)
![](https://img.shields.io/badge/Code%20Lines-8865-blue)

[更新日志 - Change Log](CHANGELOG.md)

注意：v0.0.2 中文版正在开发中，已发布的 v0.0.1 版本不支持中文

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

使用 `Y4-Lang` 实现快排

```text
函数 交换(数组, 甲, 乙) {
    临时 = 数组[甲];
    数组[甲] = 数组[乙];
    数组[乙] = 临时;
}

函数 分区(数组, 低位, 高位) {
    目标 = 数组[高位];
    甲 = 低位 - 1;
    循环 乙=低位; 乙<高位; 乙=乙+1 {
        如果 数组[乙] < 目标 {
            甲 = 甲 + 1;
            交换(数组, 甲, 乙);
        }
    }
    交换(数组, 甲 + 1, 高位);
    返回 甲 + 1;
}

函数 快排(数组, 低位, 高位) {
    如果 低位 < 高位 {
        目标 = 分区(数组, 低位, 高位);
        快排(数组, 低位, 目标-1);
        快排(数组, 目标+1, 高位);
    }
}

数组 = [10, 7, 8, 9, 1, 5, 666, 888, 10000, -50];
数组长 = 长度(数组);
快排(数组, 0, 数组长-1);
打印(数组)
```

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
