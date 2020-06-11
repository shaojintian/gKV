# gKV

A lightweight k-v store based on golang and takes example by redis(C code version).

<!-- PROJECT SHIELDS -->

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]
<a title="Build Status" target="_blank" href="https://travis-ci.com/panjf2000/gnet"><img src="https://img.shields.io/travis/com/panjf2000/gnet?style=flat-square&logo=appveyor"></a>
<a title="Codecov" target="_blank" href="https://codecov.io/gh/panjf2000/gnet"><img src="https://img.shields.io/codecov/c/github/panjf2000/gnet?style=flat-square&logo=appveyor"></a>
<a title="Supported Platforms" target="_blank" href="https://github.com/panjf2000/gnet"><img src="https://img.shields.io/badge/platform-Linux%20%7C%20macOS%20%7C%20Windows-549688?style=flat-square&logo=appveyor"></a>
<a title="Require Go Version" target="_blank" href="https://github.com/panjf2000/gnet"><img src="https://img.shields.io/badge/go-%3E%3D1.9-30dff3?style=flat-square&logo=appveyor"></a>
<a title="Release" target="_blank" href="https://github.com/panjf2000/gnet/releases"><img src="https://img.shields.io/github/release/panjf2000/gnet.svg?color=161823&style=flat-square&logo=appveyor"></a>
<br/>
<a title="" target="_blank" href="https://golangci.com/r/github.com/panjf2000/gnet"><img src="https://golangci.com/badges/github.com/panjf2000/gnet.svg"></a>
<a title="Doc for gnet" target="_blank" href="https://gowalker.org/github.com/panjf2000/gnet?lang=zh-CN"><img src="https://img.shields.io/badge/api-reference-8d4bbb.svg?style=flat-square&logo=appveyor"></a>
<a title="gnet on Sourcegraph" target="_blank" href="https://sourcegraph.com/github.com/panjf2000/gnet?badge"><img src="https://sourcegraph.com/github.com/panjf2000/gnet/-/badge.svg?style=flat-square"></a>
<a title="Mentioned in Awesome Go" target="_blank" href="https://github.com/avelino/awesome-go#networking"><img src="https://awesome.re/mentioned-badge-flat.svg"></a>

<!-- PROJECT LOGO -->
<br />

<p align="center">
  <a href="https://github.com/shaojintian/gKV/">
    <img src="docs/images/logo.png" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">Best-README-Template</h3>
  <p align="center">
    An awesome README template to jumpstart your projects!
    <br />
    <a href="https://github.com/shaojintian/gKV"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/shaojintian/gKV"> View Demo</a>
    ·
    <a href="https://github.com/shaojintian/gKV/issues">Report Bug</a>
    ·
    <a href="https://github.com/shaojintian/gKV/issues">Request Feature</a>
  </p>

</p>

 本篇README.md面向开发者 

# 🚀 功能

- [x] 支持数据类型：string,zlist,zset,zsortedSet
- [x] Reactor 模式
- [x] 持久化机制：支持RDB和AOF
- [x] 支持集群
- [x] 支持超时失效
- [x] API名称完美复刻redis，降低学习成本
- [x] 优雅退出：client/server接受sigxxx信号结束，并执行持久化存盘
- [x] 支持订阅发布


# 🏃‍ 进度
1. 完成set,get,del,lpush,llen,lrange,append
2. 正在做： RDB and AOF

# 🤜 难点
1. 传输数据make([]byte,len)==[0,0,0,0,0,...,0]  需要截取[:n]否则会有冗余的0
2. conn,err := netListen.Accept()只接受一次连接
3. 优雅接受signal：起一个goroutine，因为监听signal不能阻塞主goroutine，让它在新的goroutine中阻塞等待，一但有信号再经过cpu调度处理signal
# 🌐 展望
1. 可以考虑之后用共享内存/整理内存算法提高内存利用率，因为不整理会有内存碎片(golang内存管理不佳)


## 目录

- [上手指南](#上手指南)
  - [开发前的配置要求](#开发前的配置要求)
  - [安装步骤](#安装步骤)
- [文件目录说明](#文件目录说明)
- [开发的架构](#开发的架构)
- [部署](#部署)
- [使用到的框架](#使用到的框架)
- [核心设计💡](#核心设计)
- [性能测试📊](#性能测试)
- [贡献者](#贡献者)
  - [如何参与开源项目](#如何参与开源项目)
- [版本控制](#版本控制)
- [作者](#作者)
- [赞助❤](#赞助)
- [鸣谢](#鸣谢)

### 上手指南

请将所有链接中的“shaojintian/gKV”改为“your_github_name/your_repository”



###### 开发前的配置要求

1. golang >=1.12.5
2. xxxxx x.x.x

###### **安装步骤**

1. Get a free API Key at [https://example.com](https://example.com)
2. Clone the repo

```sh
git clone https://github.com/shaojintian/gKV.git
```
3. Start server && client
```go
go run server_main.go
go run client_main.go
```
run each command in different terminal

4. How to terminate client/server

```bash
press on Ctrl+C to close client/server gracefully.
```

### 文件目录说明

eg:


### 开发的架构 

请阅读[ARCHITECTURE.md](https://github.com/shaojintian/gKV/blob/master/ARCHITECTURE.md) 查阅为该项目的架构。

### 部署

暂无

### 使用到的框架

- [xxxxxxx](https://getbootstrap.com)
- [xxxxxxx](https://jquery.com)
- [xxxxxxx](https://laravel.com)


## 💡核心设计

1. xxxxx
2. xxxxx
3. xxxxx

## 📊性能测试

 1. xxxxx
 2. xxxxx


### 贡献者

请阅读**CONTRIBUTING.md** 查阅为该项目做出贡献的开发者。

#### 如何参与开源项目

贡献使开源社区成为一个学习、激励和创造的绝佳场所。你所作的任何贡献都是**非常感谢**的。

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



### 版本控制

该项目使用Git进行版本管理。您可以在repository参看当前可用版本。

### 作者

E-mail: sjt@hnu.edu.cn

知乎:[笃行er](https://www.zhihu.com/people/sjt_ai/activities)  &ensp; qq:1075803623    

 *您也可以在贡献者名单中参看所有参与该项目的开发者。*

### 版权说明

该项目签署了MIT 授权许可，详情请参阅 [LICENSE.txt](https://github.com/shaojintian/gKV/blob/master/LICENSE.txt)

### 鸣谢

- [GitHub Emoji Cheat Sheet](https://www.webpagefx.com/tools/emoji-cheat-sheet)
- [Img Shields](https://shields.io)
- [Choose an Open Source License](https://choosealicense.com)
- [GitHub Pages](https://pages.github.com)
- [Animate.css](https://daneden.github.io/animate.css)
- [xxxxxxxxxxxxxx](https://connoratherton.com/loaders)

### 赞助

If you like this project and want to sponsor the author, you can reward the author using Wechat or Alipay by scanning the following QR code.

<figure class="half">
  <img src="docs/images/reward_wechat.png" width="200" height="260"/>
  <img src="docs/images/reward_alipay.png" width="200" height="260"/>
</figure>
<!-- links -->

[your-project-path]: shaojintian/gKV
[contributors-shield]: https://img.shields.io/github/contributors/shaojintian/gKV.svg?style=flat-square
[contributors-url]: https://github.com/shaojintian/gKV/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/shaojintian/gKV.svg?style=flat-square
[forks-url]: https://github.com/shaojintian/gKV/network/members
[stars-shield]: https://img.shields.io/github/stars/shaojintian/gKV.svg?style=flat-square
[stars-url]: https://github.com/shaojintian/gKV/stargazers
[issues-shield]: https://img.shields.io/github/issues/shaojintian/gKV.svg?style=flat-square
[issues-url]: https://img.shields.io/github/issues/shaojintian/gKV.svg
[license-shield]: https://img.shields.io/github/license/shaojintian/gKV.svg?style=flat-square
[license-url]: https://github.com/shaojintian/gKV/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=flat-square&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/shaojintian