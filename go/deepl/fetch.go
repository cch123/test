package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var tpl = `
{
    "jsonrpc": "2.0",
    "method": "LMT_handle_jobs",
    "params": {
        "jobs": [
            {
                "kind": "default",
                "raw_en_sentence": "%v",
                "raw_en_context_before": [],
                "raw_en_context_after": [],
                "preferred_num_beams": 4
            }
        ],
        "lang": {
            "user_preferred_langs": [
                "EN",
                "ZH"
            ],
            "source_lang_computed": "ZH",
            "target_lang": "EN"
        },
        "priority": 1,
        "commonJobParams": {
            "regionalVariant": "en-US",
            "formality": null
        },
        "timestamp": 1713714284739
    },
    "id": 84460011
}
`

var content = ``

func main() {
	urll := "https://www2.deepl.com/jsonrpc"
	method := "POST"

	//payload := strings.NewReader(fmt.Sprintf(tpl, url.QueryEscape(content)))
	payload := strings.NewReader(fmt.Sprintf(tpl, content))

	client := &http.Client{}
	req, err := http.NewRequest(method, urll, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("authority", "www2.deepl.com")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36")
	req.Header.Add("dnt", "1")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("accept", "*/*")
	req.Header.Add("origin", "https://www.deepl.com")
	req.Header.Add("sec-fetch-site", "same-site")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("referer", "https://www.deepl.com/")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,zh-TW;q=0.8")
	req.Header.Add("Cookie", "LMTBID=v2|4b9f8665-2396-4bb9-a58a-d4a56daa0f95|475507a102f2ce0664045d468b16e4bc")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
