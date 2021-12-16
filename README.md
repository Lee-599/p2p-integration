# p2p-integration

一个基于p2p的异地联调工具



| 技术栈         | 版本                   |
| ------------- | ----------------------|
| go            | go1.17.2              |
| sqlite        | v1.2.3                |
| gorm          | v1.22.2               |
| wails         | v1.16.8               |
| vue.js        | 3.0.0                 |
| ant-design-vue| 2.2.8                 |

### wails 安装

#### 前置条件
- Ubuntu
```
sudo apt install libgtk-3-dev libwebkit2gtk-4.0-dev
```
- MacOS
```
xcode-select --install
```

#### 安装命令
```
go install github.com/wailsapp/wails/cmd/wails@latest
```

### 项目编译
```
wails build
cd build
./link-server
```
### 开发调试
```
wails serve
cd frontend
npm run serve
```