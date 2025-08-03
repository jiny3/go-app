# go-app

个人 Go 应用程序模板

## 🚀 开始使用

### 安装

1.  点击 "Use this template" 按钮创建你自己的仓库。
2.  将仓库克隆到本地：
    ```bash
    git clone https://github.com/your-username/your-repository.git
    cd your-repository
    ```
3.  开始开发

### 运行应用

```bash
go run main.go
```

## 📁 项目结构

```
.
├── bootstrap
│   ├── config.go # 读取配置逻辑
│   ├── init.go   # 初始化逻辑
│   └── start.go  # 启动逻辑
├── config        # 配置
│   └── app.toml  
├── go.mod
├── go.sum
├── library       # lib 库
├── main.go       # 入口文件
├── README.md
├── resource      # 长连接 client
├── services      # 业务逻辑
│   └── hello.go
├── tools         # 工具库
└── utils         # 公用组件
```
