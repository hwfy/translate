package translate

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type (
	GoogleReply struct {
		Sentences []sentence
		Src       string //源语言标识
		Ld_result ld_result
	}

	sentence struct {
		Trans   string //目标文本
		Orig    string //源文本
		Backend int
	}

	ld_result struct {
		Srclangs             []string //源语言标识
		Srclangs_confidences []float64
		Extended_srclangs    []string //源语言标识
	}
)

//&dj=1&source=icon&oe=UTF-8
const GOOGLEURL = "http://translate.google.cn/translate_a/single?client=gtx&dt=t&dj=1&ie=UTF-8"

// Srclang get the language identification,return to the default identity Lang when the translation fails
func Srclang(query string) string {
	client := getHttpClient(0, 0)

	resp, err := client.Get(GOOGLEURL + "&sl=auto&tl=zh_TW&q=" + query)
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
	var reply GoogleReply

	if err = json.Unmarshal(data, &reply); err != nil {
		log.Println(err)
		return ""
	}
	langs := reply.Ld_result.Extended_srclangs
	if len(langs) == 0 {
		return ""
	}
	if langs[0] == "ja" {
		return "ja-JP"
	}
	return langs[0]
}

// Google using Google translation of a single character
//
//   From 		To  		Translation direction
//-------------------------------------------------------------------------------
//   auto     		auto    		Automatic Identification
//   zh-CN       	en      		Simplified Chinese -> English
//   zh-CN       	zh-TW      	Simplified Chinese -> traditional Chinese
//   zh-CN       	ja-JP      		Simplified Chinese -> Japanese
func Google(from, to, query string) string {
	client := getHttpClient(0, 0)
	query = strings.Replace(query, " ", "%20", -1)
	prefix, suffix := getNumberStringPosition(query)

	resp, err := client.Get(GOOGLEURL + "&sl=" + from + "&tl=" + to + "&q=" + query[prefix:suffix])
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
	var reply GoogleReply

	if err = json.Unmarshal(data, &reply); err != nil {
		log.Println(err)
		return ""
	}
	result := query[:prefix] + " " + reply.Sentences[0].Trans + " " + query[suffix:]

	return strings.Trim(result, " ")
}

// Googles using Google translation of multiple characters
//
//   From 		To  		Translation direction
//-------------------------------------------------------------------------------
//   auto     		auto    		Automatic Identification
//   zh-CN       	en      		Simplified Chinese -> English
//   zh-CN       	zh-TW      	Simplified Chinese -> traditional Chinese
//   zh-CN       	ja-JP      		Simplified Chinese -> Japanese
func Googles(from, to string, querys []string) (results []string) {
	query := strings.Join(querys, "%0A")
	query = strings.Replace(query, " ", "%20", -1)

	resp, err := http.Get(GOOGLEURL + "&sl=" + from + "&tl=" + to + "&q=" + query)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil
	}
	var reply GoogleReply

	if err = json.Unmarshal(data, &reply); err != nil {
		log.Println(err)
		return nil
	}
	for _, v := range reply.Sentences {
		results = append(results, strings.TrimSuffix(v.Trans, "\n"))
	}
	return
}

// ToEnglish Google translated into english
func ToEnglish(query string) string {
	//为了翻译句子编码空格为%20
	query = strings.Replace(query, " ", "%20", -1)

	prefix, suffix := getNumberStringPosition(query)

	//"上班时间1"和"下班时间1"都会翻译成"Working time 1",因此去除字符串前后的数值
	resp, err := http.Get(GOOGLEURL + "&sl=auto&tl=en&q=" + query[prefix:suffix])
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
	var reply GoogleReply

	if err = json.Unmarshal(data, &reply); err != nil {
		log.Println(err)
		return ""
	}
	result := query[:prefix] + " " + reply.Sentences[0].Trans + " " + query[suffix:]

	return strings.Trim(result, " ")
}

// ToTraditional Google translated into Chinese traditional
func ToTraditional(query string) string {
	query = strings.Replace(query, " ", "%20", -1)

	resp, err := http.Get(GOOGLEURL + "&sl=auto&tl=zh-TW&q=" + query)
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
	var reply GoogleReply

	if err = json.Unmarshal(data, &reply); err != nil {
		log.Println(err)
		return ""
	}
	return reply.Sentences[0].Trans
}

// ToSimplified Google translated into Chinese Simplified
func ToSimplified(query string) string {
	query = strings.Replace(query, " ", "%20", -1)

	resp, err := http.Get(GOOGLEURL + "&sl=auto&tl=zh-CN&q=" + query)
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
	var reply GoogleReply

	if err = json.Unmarshal(data, &reply); err != nil {
		log.Println(err)
		return ""
	}
	return reply.Sentences[0].Trans
}

// getNumberStringPosition 获取字符串中前面数值的末尾位置和后面数值的起始位置
// 例如 str= "012下班时间345", 返回 prifix= 3,suffix= 15, 其中中文占3个位置
func getNumberStringPosition(str string) (prefix, suffix int) {
	for i, v := range str {
		if v > '9' || v < '0' {
			prefix = i
			break
		}
	}
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] > '9' || str[i] < '0' {
			suffix = i + 1
			break
		}
	}
	return
}
