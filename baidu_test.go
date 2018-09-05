//Copyright (c) 2017, hwfy

//Permission to use, copy, modify, and/or distribute this software for any
//purpose with or without fee is hereby granted, provided that the above
//copyright notice and this permission notice appear in all copies.

//THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
//WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
//MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
//ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
//WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
//ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
//OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

package translate

import (
	"strings"
	"testing"
)

//   zh       		en      		Chinese -> English
//   zh       		jp      		Chinese -> Japanese
//   en       		zh      		English -> Chinese
//   jp       		zh      		Japanese -> Chinese

func TestTranslate(t *testing.T) {
	query := "感谢"
	querys := []string{"感谢", "全球"}

	result := Baidu("zh", "cht", query, "")
	if result != "感謝" {
		t.Errorf("The ideal result 感謝, but practical %s", result)
	}
	result = Baidu("zh", "en", query, "")
	if result != "Thank" {
		t.Errorf("The ideal result Thank, but practical %s", result)
	}
	result = Baidu("zh", "jp", query, "")
	if result != "ありがとうございます" {
		t.Errorf("The ideal result ありがとうございます, but practical %s", result)
	}

	results := Baidus("zh", "cht", querys, "")
	word := strings.Join(results, ",")
	if word != "感謝,全球" {
		t.Errorf("The ideal result 感謝,全球, but practical %s", word)
	}
	results = Baidus("zh", "en", querys, "")
	word = strings.Join(results, ",")
	if word != "Thank,Global" {
		t.Errorf("The ideal result Thank,Global, but practical %s", word)
	}
	results = Baidus("zh", "jp", querys, "")
	word = strings.Join(results, ",")
	if word != "ありがとうございます,グローバル" {
		t.Errorf("The ideal result ありがとうございます,グローバル, but practical %s", word)
	}
}
