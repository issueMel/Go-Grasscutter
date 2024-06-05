# Go-Grasscutter

[EN](../README.md) | [简中](README_zh-CN.md)
****
this project is imitate [Grasscutter](https://github.com/Grasscutters/Grasscutter).<br>
It only support for version CN 4.0.

## Current features

- [x] 登录
- [ ] 战斗
- [ ] 好友
- [ ] 传送
- [ ] 祈愿
- [ ] 多人游戏 部分 可用
- [ ] 从控制台生成魔物
- [ ] 背包功能（接收或升级物品、角色等）

### Quick Start (automatic)

- 获取 [Go 1.22](https://go.dev/dl/)
- 获取 [MongoDB社区版](https://www.mongodb.com/try/download/community)
- 获取游戏4.0正式版 (如果你没有4.0的客户端，可以在这里找到）：
  [4.0.x Client-github](https://github.com/JRSKelvin/GenshinRepository/blob/main/Version%204.0.0.md)
  [4.0.x Client-cloud drive](https://www.123pan.com/s/HoqUVv-U7SBA.html)
- 下载 [最新的Cultivation版本](https://github.com/Grasscutters/Cultivation/releases/latest)（使用以“.msi”为后缀的安装包）。
- 以管理员身份打开Cultivation，按右上角的下载按钮。
- 点击“下载 Grasscutter 一体化”
- 点击右上角的齿轮
- 将游戏安装路径设置为你游戏所在的位置。
- 保持所有其它设置为默认值
- 打开编译好的Golang程序
- 点击“启动”按钮。
- 随便想一个用户名登录，不需要密码。

### Building

- [Go 1.22](https://go.dev/dl/) 或更高版本
- [Git](https://git-scm.com/downloads)
- [NodeJS](https://nodejs.org/en/download) (Optional, for building the handbook)

##### Clone

```shell
git clone --recurse-submodules https://github.com/issueMel/Go-Grasscutter.git
cd Go-Grasscutter
```

##### Compile

```shell
go build Go-Grasscutter
```

##### Compiling the Handbook (Manually)

With NPM:

```shell
cd src/handbook
npm install
npm run build
```

你可以在项目的根目录找到编译好的程序。

## Contribute

欢迎为该项目做出贡献。

### TIPS

在issue提出你想做的事情，避免与其他贡献者发生冲突。<br>
也可以在寻找项目中的TODO进行完善。

### TODO

- CHECK: 会在后期做改动或移除
- ENHANCE: 性能或占用需要改善
- INCOMPLETE: 功能未完善
