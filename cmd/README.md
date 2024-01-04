#### How to add a new command
```sh
cobra-cli.exe add COMMAND_NAME
```
Before that, you should install `cobra-cli.exe`.
```sh
go install github.com/spf13/cobra-cli@latest
```

### Commands
- `init` 初始化, 不存在配置文件则生成测试配置文件, 存在则无提示无需初始化.
- `report` 直接生成测试报告
- `start` 启动WebUI
- `clean` 清理项目中生成的测试中间文件`.FILE_NAME_test.go`

---
[HOME](../README.md)