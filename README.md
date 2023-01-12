# 前置条件

* WSL2 + Ubuntu 20.04LTS
* Golang
* Docker Desktop

# 测试环境

* Windows 10 21H2 (x86-64)
* go 1.19.5
* docker 20.10.22

# 目前问题

原论文中的代码前端和后端是分开的，前端没有做完，也没有和后端联系起来；后端看起来比较完整，但是缺少了本地ethereum服务端和ipfs服务端建立的代码。因此如果想要完整地实现出来，需要在本地完成ethereum与ipfs服务端的建立，把前端部分给完成，再恢复与后端的接口。

# 项目结构

整个项目基于go语言，前端使用了gin框架，后端使用了go-ethereum，go-dockerclient，go-ipfs-api等若干组件，来实现以下功能：
1. 搭建区块链平台，并使用区块链来分享docker镜像
2. 使用docker内建的组件实现对docker镜像的解析、漏洞查询
3. 使用同态加密技术把docker镜像与查询出来的CVE列表封包，添加至IPFS
4. 使用SIMD相关的库加速AES运算

# 运行代码

## 获取代码

在原作者的基础上继续完成了缺少的部分，fork地址:
https://github.com/w43322/SDIS

```
git clone https://github.com/w43322/SDIS
cd SDIS
```

## 准备工作

在本地的Docker Desktop上准备好要共享的docker镜像（以图中的`ubuntu`镜像为例）：
![](.\readme\images\docker_image_preperation.PNG)

### 潜在的问题

在接下来的过程，要用到`docker scan`命令来扫描镜像中的CVE，这要调用docker的snyk组件，第一次运行的时候要按y确认，因此可以先使用`docker scan --json ubuntu`命令预先激活snyk组件

此外，在命令行中调用`docker scan`命令时，默认没有登陆状态，只能调用十次。在上传了很多镜像而超过这个限制后，会出现错误提示。此时再手动使用`docker scan --json ubuntu`命令，会弹出浏览器登录窗口，登录后就可以继续执行了。

## 分析、加密镜像并上传至区块链

运行脚本，上传docker镜像

```
go run ".\main.go" "ubuntu"
```

此时，`SDISUtils.SecureDockerImageUpload`会被调用，完成镜像的上传

## 在区块链上分享镜像

## 从区块链上下载镜像
