# translate
Call Baidu Translate api to translate one or more characters

# Installation
>go get github.com/hwfy/translate

# Example
```go
package main

import (
	"cy/CloudERP/common/middlewares/translate"
	"fmt"
)

func main() {
	query := "感谢"

	result := translate.Eval("zh", "cht", query, "")
	fmt.Println(result)

	result = translate.Eval("zh", "en", query, "")
	fmt.Println(result)

	result = translate.Eval("zh", "jp", query, "")
	fmt.Println(result)

	querys := []string{"你好", "全球"}

	results := translate.Evals("zh", "cht", querys, "")
	fmt.Println(results)

	results = translate.Evals("zh", "en", querys, "")
	fmt.Println(results)

	results = translate.Evals("zh", "jp", querys, "")
	fmt.Println(results)
}
// OutPut:
// 感謝
// Thank
// ありがとうございます
// [你好 全球]
// [Hello. Global]
// [こんにちは グローバル]
```