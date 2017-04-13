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
	query := "��л"

	result := translate.Eval("zh", "cht", query, "")
	fmt.Println(result)

	result = translate.Eval("zh", "en", query, "")
	fmt.Println(result)

	result = translate.Eval("zh", "jp", query, "")
	fmt.Println(result)

	querys := []string{"���", "ȫ��"}

	results := translate.Evals("zh", "cht", querys, "")
	fmt.Println(results)

	results = translate.Evals("zh", "en", querys, "")
	fmt.Println(results)

	results = translate.Evals("zh", "jp", querys, "")
	fmt.Println(results)
}
// OutPut:
// ���x
// Thank
// ���꤬�Ȥ��������ޤ�
// [��� ȫ��]
// [Hello. Global]
// [����ˤ��� ����`�Х�]
```