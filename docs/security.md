# Security model

## 背景

这个项目服务的是“频繁切换配置”的真实场景。切换越频繁，越容易把密钥散落在各种文件里。
所以 `anthro-env` 把“方便”和“安全”一起考虑：默认把 token 放进 Keychain，而不是明文写到 profile 文件。

- `ANTHROPIC_AUTH_TOKEN` is stored in macOS Keychain.
- Profile files contain non-secret metadata only.
- `anthro-env env` exports variables to the current shell through hook scripts.
