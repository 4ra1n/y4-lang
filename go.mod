module github.com/4ra1n/y4-lang

go 1.21

// Y4-Lang 不依赖外部库
// 部分工具（编码/构建）库内置
// 仅引入 testify 库做测试
require github.com/stretchr/testify v1.9.0

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
