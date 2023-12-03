+++
title = '容器运行方式'
date = 2023-11-30T14:37:10+08:00
draft = false
pre = "/docker-cmd"
categories=['Docker']
+++

### 如何运行容器

在之前的几篇文章中，我们学习了如何使用docker的指令以及如何进行容器打包，接下来我们就需要了解，该如何运行容器。

##### 最简单的方式：

使用`docker run`或`docker start`指令来运行镜像或容器

但是这种方式在面对需要进行容器的端口映射、镜像卷以及容器间的网络通信时，会比较麻烦，不易操作。

同时，如果某个容器额外需要依赖另一个容器的存在，即容器之间存在关联的关系时，容器的启动顺序就比较重要了。举个简单的例子，我们手头有个web服务容器，这个容器需要依赖另一个MySQL容器的存在，这时，我们就需要先在web服务容器运行前，先运行MySQL容器。也许你会认为这种情况下，容器的先后顺序并不会对你造成困扰，但是如果容器数量多了怎么办，容器运行的相关参数特别多怎么办。亦或者经过了较长时间，你忘记了这个依赖关系怎么办？

这时候我们就需要有一个工具来管理容器运行的相关参数以及依赖关系。

### docker-compose工具

##### 如何安装

- Windows环境下，因为docker-desktop自带docker-compose，所以不需要关注如何安装

- Linux环境下，访问docker-compose的下载地址[Releases · docker/compose (github.com)](https://github.com/docker/compose/releases)，选择适合系统的以及相应架构的可运行程序。运行`sudo chmod +x docker-compose`赋予其可运行能力，让后将该可执行文件放入环境变量的目录下即可

##### 编写docker-compose.yml文件

当我们已经将容器打包完成之后，我们可以编写docker-compose.yml文件，从而能够运行容器。

我们举个简单的例子：在这个例子中，我们只需要申明其镜像卷的挂载即可

```yaml
verison: '3'
services:
  cliffwalking:
    image: demo:v1
    container_name: reinforce_learning
    volumes:
      - /home/yizhigopher/cliffwalk/ouput:/cliffwalking/output
      - /home/yizhigopher/cliffwalk/model:/cliffwalking/model
```

我们只需要在当前`docker-compose.yml`的目录下,运行`docker-compose up`就能够运行一个名为`reinforce_learning`的容器,并且将容器内部的output和model文件夹的内容和容器外部的对应的文件夹相关联.

当然,docker-compose能够识别的关键字不止以上这些内容,我们可以通过访问[Docker Compose overview | Docker Docs](https://docs.docker.com/compose/)以及菜鸟教程[Docker Compose | 菜鸟教程 (runoob.com)](https://www.runoob.com/docker/docker-compose.html)获得相关的更多信息.以上只是一个简单的示例,希望能够通过这些能够让读者一窥docker-compose容器编排的高效以及优越性

### 结语

至此,docker的相关知识基本分享完毕,事实上,仅仅通过短短的三个博文是不能说明白docker的,也不可能就仅仅读这些内容就能掌握docker.这三篇博文,仅仅能够让大家对docker有个基本的了解,以及基本操作的熟悉.想要熟练掌握docker,还需要经过自己大量的实践练习和操作才能真正熟悉.
