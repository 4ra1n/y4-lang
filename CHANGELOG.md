## Y4-Lang Change Log

### v0.0.2

更新日志：

- [功能] 支持`main`入口
- [BUG] 解决当最后一行不包含换行符时的报错
- [优化] 取消`-f`指定文件改为直接输入文件名
- [优化] 输入多个文件名时寻找其中的`main`入口
- [其他] 升级了多个`github actions`版本

所有可供下载的文件都由 `Github Actions` 构建，提供以下多种:

- windows (arm/arm64/386/amd64)
- darwin (arm64/amd64)
- linux (arm/arm64/386/amd64)

注意：默认使用了 `upx` 压缩，如果报毒不放心使用可以自行编译

### v0.0.1

第一个版本发布

更新日志：

- 无

所有可供下载的文件都由 `Github Actions` 构建，提供以下多种:

- windows (arm/arm64/386/amd64)
- darwin (arm64/amd64)
- linux (arm/arm64/386/amd64)

注意：默认使用了 `upx` 压缩，如果报毒不放心使用可以自行编译