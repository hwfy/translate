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

//type (
//	BaiduReply struct {
//		Trans_result transResult
//		Error        int
//	}
//	transResult struct {
//		Data []data
//	}
//	data struct {
//		Dst string
//	}
//)
type (
	BaiduReply struct {
		From   string
		To     string
		Domain string `json:"-"`
		Type   int    `json:"-"`
		Status int    `json:"-"`
		Data   []data
		Error  int
	}
	data struct {
		Dst        string
		PrefixWrap int `json:"-"`
		Src        string
		Relation   []string        `json:"-"`
		Result     [][]interface{} `json:"-"`
	}
)

//const BAIDUURL = "http://fanyi.baidu.com/v2transapi?"
const BAIDUURL = "http://fanyi.baidu.com/transapi?"

// Baidu using Baidu translation of a single character, key is to make the error
// prompt more detailed,such as you want to translate a table in the field ,
// then the key is the table name
//   From 		To  		Translation direction
//-------------------------------------------------------------------------------
//   auto     		auto    		Automatic Identification
//   zh       		en      		Simplified Chinese -> English
//   zh       		cht      		Simplified Chinese -> traditional Chinese
//   zh       		jp      		Simplified Chinese -> Japanese
func Baidu(from, to, query, key string) string {
	query = strings.Replace(query, " ", "%20", -1)

	resp, err := http.Get(BAIDUURL + "from=" + from + "&to=" + to + "&query=" + query)
	if err != nil {
		log.Fatalln(err)
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return ""
	}
	var reply BaiduReply

	if err = json.Unmarshal(body, &reply); err != nil {
		log.Println(key+" ", err)
		return ""
	}
	if reply.Error != 0 {
		log.Printf("%s %s form %s to %s error code: %v", key, query, from, to, reply.Error)
		return ""
	}
	datas := reply.Data
	if len(datas) == 0 {
		return ""
	}
	return strings.Replace(datas[0].Dst, "?", " no", -1)
}

// Baidus using Baidu translation of multiple characters, key is to make the error
// prompt more detailed,such as you want to translate a table in the field ,then
// the key is the table name
//   From 		To  		Translation direction
//-------------------------------------------------------------------------------
//   auto     		auto    		Automatic Identification
//   zh       		en      		Simplified Chinese -> English
//   zh       		cht      		Simplified Chinese -> traditional Chinese
//   zh       		jp      		Simplified Chinese -> Japanese
func Baidus(from, to string, querys []string, key string) (results []string) {
	query := strings.Join(querys, "%0A")
	query = strings.Replace(query, " ", "%20", -1)

	resp, err := http.Get(BAIDUURL + "from=" + from + "&to=" + to + "&query=" + query)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil
	}
	var reply BaiduReply

	if err = json.Unmarshal(body, &reply); err != nil {
		log.Println(key+" ", err)
		return nil
	}
	if reply.Error != 0 {
		log.Printf("%s %s form %s to %s error code: %v", key, querys, from, to, reply.Error)
		return nil
	}
	for _, data := range reply.Data {
		result := strings.Replace(data.Dst, "?", " no", -1)
		results = append(results, result)
	}
	return
}
