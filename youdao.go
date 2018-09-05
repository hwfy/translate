package translate

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//	ZH_CN2EN 中文　»　英语
//	ZH_CN2JA 中文　»　日语
//	ZH_CN2KR 中文　»　韩语
//	ZH_CN2FR 中文　»　法语
//	ZH_CN2RU 中文　»　俄语
//	ZH_CN2SP 中文　»　西语
//	EN2ZH_CN 英语　»　中文
//	JA2ZH_CN 日语　»　中文
//	KR2ZH_CN 韩语　»　中文
//	FR2ZH_CN 法语　»　中文
//	RU2ZH_CN 俄语　»　中文
//	SP2ZH_CN 西语　»　中文

type (
	YoudaoReply struct {
		Type            string
		ErrorCode       int
		ElapsedTime     int
		TranslateResult [][]translateResult
	}
	translateResult struct {
		Src string
		Tgt string
	}
)

const YOUDAOURL = "http://fanyi.youdao.com/translate?&doctype=json"

func Youdao(from, to, query string) string {
	type_ := from + "2" + to

	query = strings.Replace(query, " ", "%20", -1)

	resp, err := http.Get(YOUDAOURL + "&type=" + type_ + "&i=" + query)
	if err != nil {
		log.Fatalln(err)
		return ""
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return ""
	}
	var reply YoudaoReply

	if err = json.Unmarshal(data, &reply); err != nil {
		log.Println(err)
		return ""
	}
	for _, r := range reply.TranslateResult {
		for _, result := range r {
			return result.Tgt
		}
	}
	return query
}
