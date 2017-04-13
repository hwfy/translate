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
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Eval using Baidu translation of a single character, key is to make the error
// prompt more detailed,such as you want to translate a table in the field ,
// then the key is the table name
//   From 		To  		Translation direction
//-------------------------------------------------------------------------------
//   auto     		auto    		Automatic Identification
//   zh       		en      		Chinese -> English
//   zh       		jp      		Chinese -> Japanese
//   en       		zh      		English -> Chinese
//   jp       		zh      		Japanese -> Chinese
func Eval(from, to, query, key string) string {
	url := "http://fanyi.baidu.com/v2transapi?from=" + from + "&to=" + to + "&query=" + query

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
		return query
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return query
	}

	j2 := make(map[string]interface{})

	if err = json.Unmarshal(body, &j2); err != nil {
		log.Println(key+" ", err)
		return query
	}
	if j2["error"] != nil {
		log.Printf("%s_%s form %s to %s error:%v", key, query, from, to, j2["error"])
		return query
	}
	result := j2["trans_result"].(map[string]interface{})
	datas := result["data"].([]interface{})
	reply := datas[0].(map[string]interface{})

	return reply["dst"].(string)
}

// Evals using Baidu translation of multiple characters, key is to make the error
// prompt more detailed,such as you want to translate a table in the field ,then
// the key is the table name
//   From 		To  		Translation direction
//-------------------------------------------------------------------------------
//   auto     		auto    		Automatic Identification
//   zh       		en      		Chinese -> English
//   zh       		jp      		Chinese -> Japanese
//   en       		zh      		English -> Chinese
//   jp       		zh      		Japanese -> Chinese
func Evals(from, to string, querys []string, key string) []string {
	url := "http://fanyi.baidu.com/v2transapi?from=" + from + "&to=" + to + "&query=" + strings.Join(querys, "%0A")

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
		return querys
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return querys
	}
	j2 := make(map[string]interface{})

	if err = json.Unmarshal(body, &j2); err != nil {
		log.Println(key+" ", err)
		return querys
	}
	if j2["error"] != nil {
		log.Printf("%s_%s form %s to %s error:%v", key, querys, from, to, j2["error"])
		return querys
	}
	if j2["trans_result"] == nil {
		log.Println("result null")
		return querys
	}
	result := j2["trans_result"].(map[string]interface{})
	datas := result["data"].([]interface{})
	replys := make([]string, 0, len(datas))

	for _, data := range datas {
		reply := data.(map[string]interface{})
		replys = append(replys, reply["dst"].(string))
	}
	return replys
}
