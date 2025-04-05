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
 <source media="(prefers-color-scheme: dark)" srcset="docs/screenshots/dark_en.jpg">
 <source media="(prefers-color-scheme: light)" srcset="docs/screenshots/light_en.jpg">
 <img alt="screenshot" src="docs/screenshots/dark_en.jpg">
</picture>

<picture>
 <source media="(prefers-color-scheme: dark)" srcset="docs/screenshots/dark_en2.jpg">
 <source media="(prefers-color-scheme: light)" srcset="docs/screenshots/light_en2.jpg">
 <img alt="screenshot" src="docs/screenshots/dark_en2.jpg">
</picture>

## Features

* Super lightweight, built on Webview2, without embedded browsers (Thanks to [Wails](https://github.com/wailsapp/wails))
* Provides visually and user-friendly UI, light and dark themes (Thanks to [Naive UI](https://github.com/tusen-ai/naive-ui) and [IconPark](https://iconpark.oceanengine.com))
* Visualize key value operations, CRUD support for key-value pairs

## TODO

* Multi-language support
* Support multiple data viewing formats 
* Support import/export data
* Support import/export connection profiles

## Installation

Available to download for free from [here](https://github.com/Astronaut-X-X/etcdm/releases).

## Build Guidelines

### Prerequisites

* Go >= 1.23.0
* Node.js >= 16
* NPM >= 9

### Install Wails

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

### Build Project
```bash
wails build
```

### Notinges

If you get a "damaged and can't be opened" error after installing on macOS, run the following command.

```bash
 sudo xattr -d com.apple.quarantine /Applications/etcdm.app
```