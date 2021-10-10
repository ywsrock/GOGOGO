package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	body := strings.NewReader("sEcho=66&iColumns=9&sColumns=%2C%2C%2C%2C%2C%2C%2C%2C&iDisplayStart=0&iDisplayLength=10&mDataProp_0=toggle&mDataProp_1=select&mDataProp_2=medical_center_name&mDataProp_3=street_address&mDataProp_4=vaccine_name&mDataProp_5=reservation_reception&mDataProp_6=select_explanation&mDataProp_7=medical_center_cd&mDataProp_8=reservation_site_cd&_token=2QpK5PHE1FHYig8EHKWlZ7WkJJatNSD7Ky7Ccxn3&position=1&search=0&search_medical_institution_name=&search_medical_institution_name_kana=&search_ward_name=&search_street_address=&search_date_limit=2021%2F09%2F18&start_count=0&first_load=1&medical_saerch_flg=1&free_saerch_flg=1&reservations_number=1")
	req, err := http.NewRequest("POST", "https://v-yoyaku.jp/searchmedicalinstitutions/search_medical_institutions", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Sec-Ch-Ua", "\" Not;A Brand\";v=\"99\", \"Google Chrome\";v=\"91\", \"Chromium\";v=\"91\"")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Origin", "https://v-yoyaku.jp")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://v-yoyaku.jp/reservation/calendar")
	req.Header.Set("Accept-Language", "ja,en-US;q=0.9,en;q=0.8,zh-CN;q=0.7,zh;q=0.6")
	req.Header.Set("Cookie", "ck_local_key=eyJpdiI6IlpFREp2cklUMmlyb3JCSVU2MjZaNUE9PSIsInZhbHVlIjoiNUhDcEp1WlFEUWV3RlluK2JmS1lXQT09IiwibWFjIjoiNjQzYjM4ZjgzODM0MjE3NWJmZmFjODU2MWE3MTEyNTk3NDM2NjYxZmJjNTllYzljYzUwNGVhODlmMTEyMzZlYiJ9; ck_ward_key=eyJpdiI6IngrVVBrbm8rbGZwTTFuMWUrWHl1bWc9PSIsInZhbHVlIjoieGJ2c25jYlovN3lkVDhPTnRFbWt2UT09IiwibWFjIjoiMWFjODliMzQ1NDIzM2YwYzE5ZDQ0MTFjYjE2NWU4OWRjOTU1ZDBmZGEzODQyMGI4YWFmM2IxNmZkYTUyNWM1YSJ9; ck_ward_name=eyJpdiI6InZTQ1kwRkxpOVJrbVlTUmZJOGVKTFE9PSIsInZhbHVlIjoib2R5cWMwZXJEait2aHdSMG1pYjFvUT09IiwibWFjIjoiZTE3MDI0N2ExMDIwYjUxMjZjNWJmYWRhYzQ3MmNiZjc0ZDcxNGUyOWZkNmZlOWM4ZTQyOTQ5NTAxMTE4MmRjYiJ9; ck_auth_key=eyJpdiI6IlJVRWlxeVFuNTk3R2k1VzdQRDZOUUE9PSIsInZhbHVlIjoidWRHZ3dIc2NFbEJQQ3pEL2pOdFBiQT09IiwibWFjIjoiM2ZiMzZmYTQzMjdiMzUyZWQ1OWI2ZjlhZWM4YWM3OGEzZTU2MjdjNjM2OGMyMmM5YmEzZTc1N2U3NGNlYTJiMCJ9; AWSALB=k5syyNY6I5A9D/Vi+/gsgZMCQifVV3dYl8a4ctasHmZkyZF4mOuQrc2tpK5ijZyPSMDOuGVIUlM0LxGI6JmUbS24c+EeV0Ug2fRQi8r0fkXC5vBFnae1XExqM0QH; AWSALBCORS=k5syyNY6I5A9D/Vi+/gsgZMCQifVV3dYl8a4ctasHmZkyZF4mOuQrc2tpK5ijZyPSMDOuGVIUlM0LxGI6JmUbS24c+EeV0Ug2fRQi8r0fkXC5vBFnae1XExqM0QH; XSRF-TOKEN=eyJpdiI6IkJZbUx2REVlZG9HckNSVTEwamh6WFE9PSIsInZhbHVlIjoiOTVKb082MmFQT1hSWHg3eGdBR3hlTDZFQ2tVeS80L3dpUEx0cFFKRnVyZmJrelA1WE9MSWQ3eE0zVi9lSTF5MSIsIm1hYyI6IjdhZDUxMGYwYjU0YmRiYzMxOTAxYmY4ZGE5OTM0ZGRhODMwMDkzMDlmMWMwZjEyZjFlZTMxNzcwZmExZjU4YjIifQ%3D%3D; _session=eyJpdiI6InZ5cWlNRk83NnhXaVhCTFVMTmxlREE9PSIsInZhbHVlIjoia1Iwdk5VQWlvbUdqblBLdzZvaDRPTTlUTGxqSnJHODd4a0hobERmYTJiNUVnYW1Rd1QvckFMME5GcFJVVmF1TCIsIm1hYyI6IjlmZmU2ZWExMzA1ZmNlNGZkZmNhOWQ1MzI5NzNhYzdjNzc4NjVkMTY5YWMxNDBiYTg1NTlhYmFhMzI4MTJjNzQifQ%3D%3D")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	bs, _ := ioutil.ReadAll(resp.Body)
	decoded, error := url.QueryUnescape(string(body))
	fmt.Println(string(bs))
	defer resp.Body.Close()
}
