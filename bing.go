package translate

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type BingReply struct {
	XMLName xml.Name `xml:"string"`
	Text    string   `xml:",innerxml"`
}

//A4D660A48A6A97CCA791C34935E4C02BBB1BEC1C
const BINGURL = "http://api.microsofttranslator.com/v2/Http.svc/Translate?appId=AFC76A66CF4F434ED080D245C30CF1E71C22959C"

// Detect Automatically detect language,
// ja 	-> ja_JP
// ms 	-> ms_MY
// id 	-> in_ID
// zh-CHS	-> zh_CN
// zh-CHT	-> zh_TW
func Detect(query string) string {
	url := "http://api.microsofttranslator.com/V2/Http.svc/Detect?appId=AFC76A66CF4F434ED080D245C30CF1E71C22959C"

	resp, err := http.Get(url + "&text=" + query)
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
	var reply BingReply

	if err = xml.Unmarshal(data, &reply); err != nil {
		log.Println(err)
		return ""
	}
	switch reply.Text {
	case "ja":
		return "ja_JP"
	case "ms":
		return "ms_MY"
	case "id":
		return "in_ID"
	case "zh-CHS":
		return "zh_CN"
	case "zh-CHT":
		return "zh_TW"
	default:
		return reply.Text
	}
}

// Bing using bing translation of a single character
//
//   From 		To  		Translation direction
//-------------------------------------------------------------------------------
//   zh-CN       	en      		Simplified Chinese -> English
//   zh-CN       	zh-TW      	Simplified Chinese -> traditional Chinese
//   zh-CN       	ja-JP      		Simplified Chinese -> Japanese
func Bing(from, to, query string) string {
	query = strings.Replace(query, " ", "%20", -1)
	to = strings.Replace(to, "_", "-", 1)

	resp, err := http.Get(BINGURL + "&from=" + from + "&to=" + to + "&text=" + query)
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
	var reply BingReply

	if err = xml.Unmarshal(data, &reply); err != nil {
		log.Println(err)
		return ""
	}
	return reply.Text
}

// Bings using bing translation of multiple characters
//
//   From 		To  		Translation direction
//-------------------------------------------------------------------------------
//   zh-CN       	en      		Simplified Chinese -> English
//   zh-CN       	zh-TW      	Simplified Chinese -> traditional Chinese
//   zh-CN       	ja-JP      		Simplified Chinese -> Japanese
func Bings(from, to string, querys []string) []string {
	query := strings.Join(querys, "%0A")
	query = strings.Replace(query, " ", "%20", -1)

	resp, err := http.Get(BINGURL + "&from=" + from + "&to=" + to + "&text=" + query)
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
	var reply BingReply

	if err = xml.Unmarshal(data, &reply); err != nil {
		log.Println(err)
		return nil
	}
	result := make([]string, 0)

	if strings.Contains(reply.Text, "\n") {
		result = strings.Split(reply.Text, "\n")
	} else {
		result = strings.Split(reply.Text, " ")
	}
	result[0] = strings.TrimRight(result[0], ",")

	return result
}

// ToEnglishByBing bing translated into english
func ToEnglishByBing(query string) string {
	query = strings.Trim(query, " ")
	query = strings.Replace(query, " ", "%20", -1)

	prefix, suffix := getNumberStringPosition(query)

	resp, err := http.Get(BINGURL + "&from=&to=en&text=" + query[prefix:suffix])
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
	var reply BingReply

	if err = xml.Unmarshal(data, &reply); err != nil {
		log.Println(err)
		return ""
	}
	result := query[:prefix] + " " + reply.Text + " " + query[suffix:]

	return strings.Trim(result, " ")
}

// ToTraditionalByBing bing translated into Chinese traditional
func ToTraditionalByBing(query string) string {
	query = strings.Replace(query, " ", "%20", -1)

	resp, err := http.Get(BINGURL + "&from=&to=zh-TW&text=" + query)
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
	var reply BingReply

	if err = xml.Unmarshal(data, &reply); err != nil {
		log.Println(err)
		return ""
	}
	return reply.Text
}

// ToSimplifiedByBing bing translated into Chinese Simplified
func ToSimplifiedByBing(query string) string {
	query = strings.Replace(query, " ", "%20", -1)

	resp, err := http.Get(BINGURL + "&from=&to=zh-CN&text=" + query)
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
	var reply BingReply

	if err = xml.Unmarshal(data, &reply); err != nil {
		log.Println(err)
		return ""
	}
	return reply.Text
}

//  	   # 取得acess token的两个参数，常量
//     ACCESS_TOKEN_REQ_SCOPE = "http://api.microsofttranslator.com"
//     ACCESS_TOKEN_REQ_GRANT_TYPE = "client_credentials"

//     # POST取得ACESS TOKEN的URL
//     ACCESS_TOKEN_REQ_URL = "https://datamarket.accesscontrol.windows.net/v2/OAuth2-13"
//     # GET方法得到翻译的结果，只得到一个结果，估计这个最常用
//     TRANSLATE_REQ_URL = "http://api.microsofttranslator.com/V2/Http.svc/Translate"
//     # POST取得翻译结果的结果的URL,这个是一次可以取回多个翻译结果
//     GET_TRANSLATE_REQ_URL = "http://api.microsofttranslator.com/V2/Http.svc/GetTranslations"
//     # 检测语句的语言
//     DETECT_REQ_URL = "http://api.microsofttranslator.com/V2/Http.svc/Detect"
//     # 增加翻译的URL
//     ADD_TRANSLATION_URL = "http://api.microsofttranslator.com/V2/Http.svc/AddTranslation"
//     # 发音的请求
//     SPEAK_REQ_URL = "http://api.microsofttranslator.com/V2/Http.svc/Speak"
