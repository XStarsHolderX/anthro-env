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
anthro-env edit <name>
anthro-env use <name>
anthro-env ls
anthro-env current
anthro-env rm <name>
anthro-env doctor
anthro-env -v
```

如果你之前的 profile 文件里有明文 `ANTHROPIC_AUTH_TOKEN`：

```bash
anthro-env migrate-tokens
```

### 详细用法：`edit`

```bash
anthro-env edit <name>
```

交互规则：
- `ANTHROPIC_BASE_URL`：直接回车 = 保留原值
- `ANTHROPIC_MODEL`：直接回车 = 保留；输入 `-` = 清空（走网关默认模型）
- `ANTHROPIC_AUTH_TOKEN`：直接回车 = 保留 Keychain 当前值；输入 `-` = 从 Keychain 删除；输入新值 = 覆盖

示例：

```bash
anthro-env edit ai-router
anthro-env use ai-router
anthro-env doctor
```

### 详细用法：`migrate-tokens`

```bash
anthro-env migrate-tokens
```

这个命令会：
- 读取 profile 文件中的明文 `ANTHROPIC_AUTH_TOKEN`
- 写入对应 profile 的 macOS Keychain
- 从 profile 文件删除明文 token
- 输出迁移统计（`migrated` / `skipped`）

建议迁移后执行：

```bash
anthro-env doctor
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

### 示例 3：ai-router（走网关默认模型路由）

```bash
ANTHROPIC_AUTH_TOKEN=sk-********
ANTHROPIC_BASE_URL=https://ai-router.plugins-world.cn
CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC=1
# 这里故意不设置 ANTHROPIC_MODEL
# 具体模型由网关侧默认路由策略决定
```

## 安全与优先级规则

- Token 推荐存放在 macOS Keychain。
- profile 文件建议只放非敏感配置。
- Token 生效优先级：`Keychain > .env`。  
  如果两边都存在，以 Keychain 为准。

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

发布机制与流水线说明：见 [docs/release.md](docs/release.md)
