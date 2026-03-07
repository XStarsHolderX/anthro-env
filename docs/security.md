## Security


这个项目服务的是“频繁切换配置”的真实场景。  
切换越频繁，就越容易把密钥散落在各种文件里。

因此 anthro-env 在设计时同时考虑了 **方便性和安全性**：

默认将 Token 存储在 macOS Keychain，而不是写入 profile 文件。

This tool is built for workflows where configuration switching happens frequently.

Frequent switching can easily lead to secrets being scattered across different files.

To avoid this, anthro-env is designed with both convenience and security in mind.

Tokens are stored in the macOS Keychain by default instead of being written to profile files.

### Security Design

- `ANTHROPIC_AUTH_TOKEN` 存储在 macOS Keychain  
- profile 文件只包含非敏感配置  
- `anthro-env env` 通过 shell hook 导出环境变量

  
- `ANTHROPIC_AUTH_TOKEN` is stored in macOS Keychain.
- Profile files contain non-secret metadata only.
- `anthro-env env` exports variables to the current shell through hook scripts.
- If old profiles contain plaintext token, run `anthro-env migrate-tokens`.
