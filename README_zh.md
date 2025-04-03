<div align="center">
<a href="https://github.com/Astronaut-X-X/etcdm/"><img src="build/appicon.png" width="120"/></a>
</div>
<h1 align="center">ETCDM</h1>
<h4 align="center"><a href="https://github.com/Astronaut-X-X/etcdm/blob/main/README.md">English</a> | <strong>简体中文</strong></h4>
<div align="center">

[![License](https://img.shields.io/github/license/Astronaut-X-X/etcdm)](https://github.com/Astronaut-X-X/etcdm/blob/main/LICENSE)
[![GitHub release](https://img.shields.io/github/release/Astronaut-X-X/etcdm)](https://github.com/Astronaut-X-X/etcdm/releases)
![GitHub All Releases](https://img.shields.io/github/downloads/Astronaut-X-X/etcdm/total)
[![GitHub stars](https://img.shields.io/github/stars/Astronaut-X-X/etcdm)](https://github.com/Astronaut-X-X/etcdm/stargazers)
[![GitHub forks](https://img.shields.io/github/forks/Astronaut-X-X/etcdm)](https://github.com/Astronaut-X-X/etcdm/fork)

<strong>ETCDM 是一个现代化的轻量级跨平台 ETCD 桌面管理工具，支持 Mac、Windows 和 Linux。</strong>
</div>

<picture>
 <source media="(prefers-color-scheme: dark)" srcset="docs/screenshots/dark_en.jpg">
 <source media="(prefers-color-scheme: light)" srcset="docs/screenshots/light_en.jpg">
 <img alt="screenshot" src="docs/screenshots/dark_en.jpg">
</picture>

<picture>
 <source media="(prefers-color-scheme: dark)" srcset="docs/screenshots/dark_en2.jpg">
 <source media="(prefers-color-scheme: light)" srcset="docs/screenshots/light_en2.jpg">
 <img alt="screenshot" src="docs/screenshots/dark_en2.jpg">
</picture>

## 特性

* 超轻量级，基于 Webview2 构建，无内嵌浏览器（感谢 [Wails](https://github.com/wailsapp/wails)）
* 提供视觉友好的用户界面，支持明暗主题（感谢 [Naive UI](https://github.com/tusen-ai/naive-ui) 和 [IconPark](https://iconpark.oceanengine.com)）
* 多语言支持（[需要更多语言？点击这里贡献](.github/CONTRIBUTING.md)）
* 更好的连接管理：支持 SSL/HTTP 代理/SOCKS5 代理
* 可视化的键值操作，支持键值对的增删改查
* 支持多种数据查看格式
* 使用列表分段加载，轻松处理数百万个键
* 命令操作历史日志列表
* 提供命令行模式
* 支持数据导入/导出
* 支持连接配置导入/导出

## 安装

可以从[这里](https://github.com/Astronaut-X-X/etcdm/releases)免费下载。

> 如果在 macOS 上安装后无法打开，请执行以下命令后重新打开：
> ```shell
> sudo xattr -d com.apple.quarantine /Applications/ETCDM.app
> ```

## 构建指南

### 环境要求

* Go（最新版本）
* Node.js >= 16
* NPM >= 9

### 安装 Wails

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest