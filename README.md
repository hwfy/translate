# translate
Call Translate api to translate one or more characters

# Installation
>go get github.com/hwfy/translate

# Example
```go
package main

import (
	"github.com/hwfy/translate"
	"fmt"
)

func main() {
	query := "��л"

	result := translate.Google("zh", "cht", query, "")
	fmt.Println(result)

	result = translate.Google("zh", "en", query, "")
	fmt.Println(result)

	result = translate.Google("zh", "jp", query, "")
	fmt.Println(result)

	querys := []string{"���", "ȫ��"}

	results := translate.Googles("zh", "cht", querys, "")
	fmt.Println(results)

	results = translate.Googles("zh", "en", querys, "")
	fmt.Println(results)

	results = translate.Googles("zh", "jp", querys, "")
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