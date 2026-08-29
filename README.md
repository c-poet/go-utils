# go-utils

`go-utils` 是一组按职责拆分的 Go 通用工具包。

## 安装

```bash
go get github.com/c-poet/go-utils
```

## 模块

| 模块 | 导入路径 | 功能 |
| --- | --- | --- |
| `cipher` | `github.com/c-poet/go-utils/cipher` | AES-CBC 加密、解密及 Base64 字符串辅助方法 |
| `json` | `github.com/c-poet/go-utils/json` | 通过 JSON 在对象、Map 与结构体间转换 |
| `os` | `github.com/c-poet/go-utils/os` | Windows、macOS、Linux 运行环境判断 |
| `strings` | `github.com/c-poet/go-utils/strings` | 按指定分隔符或逗号拆分字符串，可保留或过滤空字符串 |
| `types` | `github.com/c-poet/go-utils/types` | 反射类型判断 |

## 示例

```go
package main

import (
	"fmt"

	jsonutil "github.com/c-poet/go-utils/json"
	"github.com/c-poet/go-utils/os"
	"github.com/c-poet/go-utils/types"
)

type User struct {
	Name string `json:"name"`
}

func main() {
	user := User{Name: "Alice"}
	values := jsonutil.ConvertToMap(user)

	var decoded User
	_ = jsonutil.ConvertToStruct(values, &decoded)

	fmt.Println(types.IsStruct(user))
	fmt.Println(os.IsWindows())
}
```

## 开发

请使用以下命令格式化并运行全部测试：

```bash
gofmt -w <changed-go-files>
go test ./...
```

`cipher` 中的 AES-CBC 工具为兼容既有接口而保留了密钥派生 IV 的行为。新协议或新数据存储建议使用带随机 nonce 的认证加密方案。
