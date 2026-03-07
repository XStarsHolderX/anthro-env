# anthro-env

[English](README.md) | 中文

[![Release](https://img.shields.io/github/v/release/kelaocai/anthro-env)](https://github.com/kelaocai/anthro-env/releases)
[![Homebrew Tap](https://img.shields.io/badge/Homebrew-kelaocai%2Fhomebrew--tap-blue)](https://github.com/kelaocai/homebrew-tap)

`anthro-env` 是一个面向 macOS 的 Claude Code / Anthropic 环境变量配置切换 CLI。
它可以让你在多个兼容 Anthropic 协议的网关与模型之间一键切换，并把 token 存在 macOS Keychain。

## 快速链接

- English README: [README.md](README.md)
- 中文 README: [README.zh.md](README.zh.md)
- 快速上手（3分钟）: [docs/快速上手.md](docs/快速上手.md)
- 完整中文手册: [docs/项目中文手册.md](docs/项目中文手册.md)

## Homebrew 安装（推荐）

```bash
brew tap kelaocai/homebrew-tap
brew install anthro-env
```

普通用户安装不需要本地 Go 环境。

## 30 秒开始

```bash
anthro-env init
source ~/.zshrc
anthro-env menu
```

## 常用命令

```bash
anthro-env init
anthro-env menu
anthro-env add <name>
anthro-env use <name>
anthro-env ls
anthro-env current
anthro-env rm <name>
anthro-env doctor
anthro-env -v
```

## 配置示例（已脱敏）

以下示例基于真实可用配置，API Key 已脱敏。  
注意：在 `anthro-env` 中，推荐把 token 存到 Keychain。

### 示例 1：bailian-kimi-k2.5

```bash
ANTHROPIC_AUTH_TOKEN=sk-********
ANTHROPIC_BASE_URL=https://coding.dashscope.aliyuncs.com/apps/anthropic
CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC=1
ANTHROPIC_MODEL=kimi-k2.5
ANTHROPIC_SMALL_FAST_MODEL=kimi-k2.5
ANTHROPIC_DEFAULT_SONNET_MODEL=kimi-k2.5
ANTHROPIC_DEFAULT_OPUS_MODEL=kimi-k2.5
ANTHROPIC_DEFAULT_HAIKU_MODEL=kimi-k2.5
```

### 示例 2：MiniMax-M2.5

```bash
ANTHROPIC_AUTH_TOKEN=sk-********
ANTHROPIC_BASE_URL=https://api.minimax.io/anthropic
CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC=1
ANTHROPIC_MODEL=MiniMax-M2.5
ANTHROPIC_SMALL_FAST_MODEL=MiniMax-M2.5
ANTHROPIC_DEFAULT_SONNET_MODEL=MiniMax-M2.5
ANTHROPIC_DEFAULT_OPUS_MODEL=MiniMax-M2.5
ANTHROPIC_DEFAULT_HAIKU_MODEL=MiniMax-M2.5
```

## 源码编译（给极客用户）

如果你希望本地编译或参与开发：

```bash
git clone https://github.com/kelaocai/anthro-env.git
cd anthro-env
go test ./...
go build -o ./bin/anthro-env ./cmd/anthro-env
./bin/anthro-env --help
```

## License

MIT
