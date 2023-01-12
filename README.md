# 前置条件

* WSL2 + Ubuntu 20.04LTS
* Golang
* Docker Desktop
* kubo (ipfs节点的go语言实现，已集成在仓库中)
* ganache
* 全局透明代理（go、docker和相关组件需要代理才能正常工作）

# 测试环境

* Windows 10 21H2 (x86-64)
* go 1.19.5
* docker 20.10.22
* kubo 0.17.0
* node.js 18.13.0
* ganache 2.5.4

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

### 准备docker镜像

启动本地的Docker Desktop上，并准备好要共享的docker镜像（以图中的`ubuntu`镜像为例）：
![](.\readme\images\docker_image_preperation.PNG)

#### 潜在的问题

在接下来的过程，要用到`docker scan`命令来扫描镜像中的CVE，这要调用docker的snyk组件，第一次运行的时候要按y确认，因此可以先使用`docker scan --json ubuntu`命令预先激活snyk组件

此外，在命令行中调用`docker scan`命令时，默认没有登陆状态，只能调用十次。在上传了很多镜像而超过这个限制后，会出现错误提示。此时再手动使用`docker scan --json ubuntu`命令，会弹出浏览器登录窗口，登录后就可以继续执行了。

### 安装并启用kubo daemon（ipfs节点）

把`./kubo/`添加到系统环境变量

在命令行中输入`ipfs init`，初始化ipfs节点，得到以下输出（keypair会不同）：

```
PS C:\Github\SDIS> ipfs init
generating ED25519 keypair...done
peer identity: 12D3KooWKfV5gaX8DP83Q3HZv7Unp6pUQD9PJLLoZZSnoZEYoC8y
initializing IPFS node at C:\Users\raywa\.ipfs
to get started, enter:

        ipfs cat /ipfs/QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc/readme
```

按照提示，输入`ipfs cat /ipfs/...`命令，然后看到如下输出则安装成功

```
PS C:\Github\SDIS> ipfs cat /ipfs/QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc/readme
Hello and Welcome to IPFS!

██╗██████╗ ███████╗███████╗
██║██╔══██╗██╔════╝██╔════╝
██║██████╔╝█████╗  ███████╗
██║██╔═══╝ ██╔══╝  ╚════██║
██║██║     ██║     ███████║
╚═╝╚═╝     ╚═╝     ╚══════╝

If you're seeing this, you have successfully installed
......
```

输入`ipfs daemon --init`命令，以默认配置启动ipfs daemon，成功启动后会输出类似如下的内容：

```
PS C:\Github\SDIS> ipfs daemon --init
Initializing daemon...
Kubo version: 0.17.0
Repo version: 12
System version: amd64/windows
Golang version: go1.19.1
......
API server listening on /ip4/127.0.0.1/tcp/5001
WebUI: http://127.0.0.1:5001/webui
Gateway (readonly) server listening on /ip4/127.0.0.1/tcp/8080
Daemon is ready
```

### 安装并启用ganache（ethereum智能合约daemon）

用下面的链接下载安装ganache
https://trufflesuite.com/ganache/

新建workspace，并在truffle-projects中添加`./truffle-config.js`

在`chain`选项卡中把`gas price`调整为`0`

点击`save workspace`，启动daemon

## 分析、加密镜像并上传至区块链

运行脚本，上传docker镜像

```
go run ".\main.go" "ubuntu"
```

此时，`SDISUtils.SecureDockerImageUpload`会被调用，完成镜像的分析、封包、加密、上传

成功后返回结果：
```
go run ".\main.go" "ubuntu"
Start Secure docker image upload
Scanning the docker image: ubuntu
Error in dockerScanCmdout
2023/01/12 16:58:30 exit status 1
Changing the docker cve list to json file
Successfully resulted in json file.
Title : Out-of-Bounds  Severity : low  CvssScore : 0  CVE : [CVE-2022-3715]
Title : Improper Input Validation  Severity : low  CvssScore : 6.5  CVE : [CVE-2016-2781]
Title : Allocation of Resources Without Limits or Throttling  Severity : low  CvssScore : 7.5  CVE : [CVE-2016-20013]
Title : Allocation of Resources Without Limits or Throttling  Severity : low  CvssScore : 7.5  CVE : [CVE-2016-20013]
Title : CVE-2022-3219  Severity : low  CvssScore : 0  CVE : [CVE-2022-3219]
Title : CVE-2022-3219  Severity : low  CvssScore : 0  CVE : [CVE-2022-3219]
Title : Integer Overflow or Wraparound  Severity : low  CvssScore : 7.5  CVE : [CVE-2018-5709]
Title : Integer Overflow or Wraparound  Severity : medium  CvssScore : 8.8  CVE : [CVE-2022-42898]
Title : Integer Overflow or Wraparound  Severity : low  CvssScore : 7.5  CVE : [CVE-2018-5709]
Title : Integer Overflow or Wraparound  Severity : medium  CvssScore : 8.8  CVE : [CVE-2022-42898]
Title : Integer Overflow or Wraparound  Severity : low  CvssScore : 7.5  CVE : [CVE-2018-5709]
Title : Integer Overflow or Wraparound  Severity : medium  CvssScore : 8.8  CVE : [CVE-2022-42898]
Title : Integer Overflow or Wraparound  Severity : low  CvssScore : 7.5  CVE : [CVE-2018-5709]
Title : Integer Overflow or Wraparound  Severity : medium  CvssScore : 8.8  CVE : [CVE-2022-42898]
Title : Integer Overflow or Wraparound  Severity : low  CvssScore : 7.5  CVE : [CVE-2018-5709]
Title : Integer Overflow or Wraparound  Severity : medium  CvssScore : 8.8  CVE : [CVE-2022-42898]
Title : Integer Overflow or Wraparound  Severity : low  CvssScore : 7.5  CVE : [CVE-2018-5709]
Title : Integer Overflow or Wraparound  Severity : medium  CvssScore : 8.8  CVE : [CVE-2022-42898]
Title : Integer Overflow or Wraparound  Severity : low  CvssScore : 7.5  CVE : [CVE-2018-5709]
Title : Integer Overflow or Wraparound  Severity : medium  CvssScore : 8.8  CVE : [CVE-2022-42898]
Title : Integer Overflow or Wraparound  Severity : low  CvssScore : 7.5  CVE : [CVE-2018-5709]
Title : Integer Overflow or Wraparound  Severity : medium  CvssScore : 8.8  CVE : [CVE-2022-42898]
Title : Integer Overflow or Wraparound  Severity : low  CvssScore : 7.5  CVE : [CVE-2018-5709]
Title : Integer Overflow or Wraparound  Severity : medium  CvssScore : 8.8  CVE : [CVE-2022-42898]
Title : Integer Overflow or Wraparound  Severity : low  CvssScore : 7.5  CVE : [CVE-2018-5709]
Title : Integer Overflow or Wraparound  Severity : medium  CvssScore : 8.8  CVE : [CVE-2022-42898]
Title : Integer Overflow or Wraparound  Severity : low  CvssScore : 7.5  CVE : [CVE-2018-5709]
Title : Integer Overflow or Wraparound  Severity : medium  CvssScore : 8.8  CVE : [CVE-2022-42898]
Title : Out-of-bounds Read  Severity : low  CvssScore : 7.1  CVE : [CVE-2022-29458]
Title : Out-of-bounds Read  Severity : low  CvssScore : 7.1  CVE : [CVE-2022-29458]
Title : Out-of-bounds Read  Severity : low  CvssScore : 7.1  CVE : [CVE-2022-29458]
Title : Out-of-bounds Read  Severity : low  CvssScore : 7.1  CVE : [CVE-2022-29458]
Title : Out-of-bounds Read  Severity : low  CvssScore : 7.1  CVE : [CVE-2022-29458]
Title : Out-of-bounds Read  Severity : low  CvssScore : 7.1  CVE : [CVE-2022-29458]
Title : Out-of-bounds Read  Severity : low  CvssScore : 7.1  CVE : [CVE-2022-29458]
Title : Out-of-bounds Read  Severity : low  CvssScore : 7.1  CVE : [CVE-2022-29458]
Title : Out-of-bounds Read  Severity : low  CvssScore : 7.1  CVE : [CVE-2022-29458]
Title : Out-of-bounds Read  Severity : low  CvssScore : 7.1  CVE : [CVE-2022-29458]
Title : Out-of-bounds Read  Severity : low  CvssScore : 7.1  CVE : [CVE-2022-29458]
Title : Out-of-bounds Read  Severity : low  CvssScore : 7.1  CVE : [CVE-2022-29458]
Title : Out-of-bounds Read  Severity : low  CvssScore : 7.1  CVE : [CVE-2022-29458]
Title : Improper Locking  Severity : low  CvssScore : 7.5  CVE : [CVE-2022-3996]
Title : Improper Locking  Severity : low  CvssScore : 7.5  CVE : [CVE-2022-3996]
Title : Uncontrolled Recursion  Severity : low  CvssScore : 7.5  CVE : [CVE-2017-11164]
Title : Uncontrolled Recursion  Severity : low  CvssScore : 7.5  CVE : [CVE-2017-11164]
Title : Off-by-one Error  Severity : medium  CvssScore : 5.5  CVE : [CVE-2022-3821]
Title : Off-by-one Error  Severity : medium  CvssScore : 5.5  CVE : [CVE-2022-3821]
Title : Off-by-one Error  Severity : medium  CvssScore : 5.5  CVE : [CVE-2022-3821]
Title : Off-by-one Error  Severity : medium  CvssScore : 5.5  CVE : [CVE-2022-3821]
Title : Off-by-one Error  Severity : medium  CvssScore : 5.5  CVE : [CVE-2022-3821]
Title : Off-by-one Error  Severity : medium  CvssScore : 5.5  CVE : [CVE-2022-3821]
Title : Off-by-one Error  Severity : medium  CvssScore : 5.5  CVE : [CVE-2022-3821]
Title : Off-by-one Error  Severity : medium  CvssScore : 5.5  CVE : [CVE-2022-3821]
Title : Off-by-one Error  Severity : medium  CvssScore : 5.5  CVE : [CVE-2022-3821]
372.60000000000014
Converting docker image to tar file: ubuntu.tar
2023/01/12 16:58:30 API error (404): 404 page not found

AESKey: [51 55 49 56 56 49 51 54 49 55 50 50 49 50 50 57 50 51 50 56 49 50 49 51 57 56 51 55 56 51 50 50]
Adding enimage on IPFS
added QmRDVgeRxegkm4GpKx6ZJWZYLihA8xdc1kN3k9LhEPHbB8
Adding ubuntu_cvelist on IPFS
added QmeRnQ292XU1J6me6FT2Wg9DSW1FAuWWrHFzgyfF29M2eG
Register data to blockchain
tx sent: 0x255054a5759c52d80284d001edfe3af7e2f89b41dc9723d9e5c5bcc51fb2e2d4
Register data to blockchain done
Docker image upload time Total cost : 51.4653186s
```

## 在区块链上分享镜像

## 从区块链上下载镜像
