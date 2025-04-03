<div align="center">
<a href="https://github.com/Astronaut-X-X/etcdm/"><img src="build/appicon.png" width="120"/></a>
</div>
<h1 align="center">ETCDM</h1>
<h4 align="center"><strong>English</strong> | <a href="https://github.com/Astronaut-X-X/etcdm/blob/main/README_zh.md">
简体中文</a></h4>
<div align="center">

[![License](https://img.shields.io/github/license/Astronaut-X-X/etcdm)](https://github.com/Astronaut-X-X/etcdm/blob/main/LICENSE)
[![GitHub release](https://img.shields.io/github/release/Astronaut-X-X/etcdm)](https://github.com/Astronaut-X-X/etcdm/releases)
![GitHub All Releases](https://img.shields.io/github/downloads/Astronaut-X-X/etcdm/total)
[![GitHub stars](https://img.shields.io/github/stars/Astronaut-X-X/etcdm)](https://github.com/Astronaut-X-X/etcdm/stargazers)
[![GitHub forks](https://img.shields.io/github/forks/Astronaut-X-X/etcdm)](https://github.com/Astronaut-X-X/etcdm/fork)

<strong>ETCDM is a modern lightweight cross-platform ETCD desktop manager available for Mac, Windows, and
Linux.</strong>
</div>

<picture>
 <source media="(prefers-color-scheme: dark)" srcset="screenshots/dark_en.png">
 <source media="(prefers-color-scheme: light)" srcset="screenshots/light_en.png">
 <img alt="screenshot" src="screenshots/dark_en.png">
</picture>

<picture>
 <source media="(prefers-color-scheme: dark)" srcset="screenshots/dark_en2.png">
 <source media="(prefers-color-scheme: light)" srcset="screenshots/light_en2.png">
 <img alt="screenshot" src="screenshots/dark_en2.png">
</picture>

## Feature

* Super lightweight, built on Webview2, without embedded browsers (Thanks
  to [Wails](https://github.com/wailsapp/wails)).
* Provides visually and user-friendly UI, light and dark themes (Thanks to [Naive UI](https://github.com/tusen-ai/naive-ui)
  and [IconPark](https://iconpark.oceanengine.com)).
* Multi-language support ([Need more languages ? Click here to contribute](.github/CONTRIBUTING.md)).
* Better connection management: supports SSL/HTTP proxy/SOCKS5 proxy.
* Visualize key value operations, CRUD support for key-value pairs.
* Support multiple data viewing format.
* Use list for segmented loading, making it easy to list millions of keys.
* Logs list for command operation history.
* Provides command-line mode.
* Support import/export data.
* Support import/export connection profile.

## Installation

Available to download for free from [here](https://github.com/Astronaut-X-X/etcdm/releases).

> If you can't open it after installation on macOS, exec the following command then reopen:
> ``` shell
>  sudo xattr -d com.apple.quarantine /Applications/ETCDM.app
> ```

## Build Guidelines

### Prerequisites

* Go (latest version)
* Node.js >= 16
* NPM >= 9

### Install Wails
