# FaucetPadPorter

多线程，高速`Hyperos Pad`包移植工具

usage:

```shell
go build
./faucetpadporter  -base 原包.zip -port 移植包.zip
```

鉴于这个项目过于不要脸，以后部分代码随缘更新，哈哈



| 系统/硬件要求 | 适配情况 |
|--------------|---------|
| Linux (amd64) | 良好    |
| Windows      | 未测试   |

**注意：** 以上适配情况基于以下条件：
- 系统线程数大于8
- 对于 Linux 只有amd64平台