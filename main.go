package main

import (
	"fmt"
	"encoding/json"
	"net/url"
)

const jsonStr string = `
{
  "name": "日本",
  "prefectures": [
    {
      "name": "東京都",
      "capital": "東京",
      "population": 13482040
    },
    {
      "name": "埼玉県",
      "capital": "さいたま市",
      "population": 7249287
    },
    {
      "name": "神奈川県",
      "capital": "横浜市",
      "population": 9116252
    }
  ]
}
`

type Country struct {
	Name string `json:"name"`
	Prefectures []Prefecture `json:"prefectures"`
}

type Prefecture struct {
	Name string    `json:"name"`
}

func main() {
	var jsonBytes []byte = ([]byte)(jsonStr)
	var data Country

	if err := json.Unmarshal(jsonBytes, &data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return
	}

	v := url.Values{}
	v.Set("name", "Ava")
	v.Add("friend", "Jess")
	v.Add("friend", `{"value": "hello"}`)
	v.Add("friend", "Zoe")
	aa := v.Encode()

	aaa, _ := url.Parse(aa)

	fmt.Println(aa)
	fmt.Println(aaa.RawQuery)
	fmt.Println(data.Prefectures[0].Name)
	fmt.Println(data.Prefectures[1].Name)
}
