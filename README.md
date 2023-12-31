# gograce
## 项目架构图
![架构图](<./pictures/屏幕截图 2024-01-02 103324.png>)

## 效果参考 C-tessy
![Alt text](./pictures/2a55a16f93e71756400e3db17d9ac58.png)
![Alt text](./pictures/51daac1d3264f639e06ff7dad34be81.png)

## Development Plan
### V1.0
####  Backend
- [x] 实现源文件AST分析
- [ ] 聚合YAML测试配置文件和AST树
- [ ] 聚合测试配置后生成前端显示数据
- [ ] CMD初始化可视化界面, 参数化启动
- [ ] 测试配置修改
- [ ] 根据聚合后的测试配置生成`.SOURCE_FILE_test.go`文件
- [ ] 执行单元测试(All / Current)
- [ ] ~~监控源文件变化~~(暂时采用手动测试)
- [ ] ~~测试用例不匹配项设置暂存区~~(配置与AST不匹配时, 在下一次同步配置到YAML中时, 自动忽略)
- [ ] 测试结果生成与导出

#### Frontend
- [ ] 根据现有测试信息进行显示
- [ ] 选择是否忽略文件或者函数, 并添加忽略备注
- [ ] ~~测试用例分组~~(暂时不支持, 需要求函数功能单一, 且颗粒度足够细)
- [ ] 新增测试用例
- [ ] 编辑测试用例
- [ ] 变量/函数打桩(暂时不支持Mock)
- [ ] 测试结果展示与导出

### APIs
- Refer to - [APIs](./doc/API.md)
  

### Provided Commands
- Refer to - [Commands](./cmd/README.md)

### How To Use
- Clone this project to your local machine and go into the project directory `gograce/web/client/go-grace`, execute below command to build the vue app.
    ```sh
    npm install
    npm run build 
    vue build
    ```
- Build the backend server by executing below command.
  ```sh
  go build -o gograce
  ```
- Start the backend server by executing below command.
  ```sh
  gograce start
  ```
- And then open the browser and visit `localhost:8080` to see the web app.
