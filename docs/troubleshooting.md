# Troubleshooting

## 背景

`anthro-env` 来自真实高频使用场景：配置切换越多，越容易遇到“明明切了但没生效”这类问题。
这份文档就是把最常见的问题和最快修复路径放在一起，减少来回排查时间。

## Profile not switching in current terminal

Run:

```bash
source ~/.zshrc
anthro-env doctor
```

## Keychain access failed

- Confirm keychain is unlocked.
- Re-run `anthro-env add <name>` and enter token again.
