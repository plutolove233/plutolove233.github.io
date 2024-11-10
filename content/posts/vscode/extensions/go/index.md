+++
title = 'vscode Go环境设置'
date = 2024-11-10T20:11:24+08:00
draft = false
tags = ["golang"]
categories = ["vscode插件配置设置"]
+++

首先你需要按照[这个教程](https://plutolove233.github.io/go-install)安装好Go环境。

之后在vscode插件商城中，搜寻`Go`，选择第一个进行安装。

![](./install.png)

安装好插件后，按下`Ctrl+Shift+P`快捷键，输入go tool，并安装所有工具。

![](./tool.png)
![](./tools2.png)

观察在`${GOPATH}/bin`目录下是否存在这些工具的可执行文件。

![](./ok.png)

重启vscode，即可。