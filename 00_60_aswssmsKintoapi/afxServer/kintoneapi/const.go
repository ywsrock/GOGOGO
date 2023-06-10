package kintoneapi

const (
	// BaseURL          = "https://w1ylayksvqqk.cybozu.com"
	// ApiToken         = "STRUGo43oDKFK24uXy6bhVM4wmaKQyg5VXWyxIoB"
	// AppID        int = 5
	BaseURL      = "https://9xqesjnpoad9.cybozu.com"
	ApiToken     = "RqtKq1jmxXgyjMyZu4fvcejmul0QIYqJszM6dTdo"
	AppID    int = 8

	GuestSpaceID = ""

	// 1件のレコードを取得する	GET
	GET_RECODE_URL = "/k/v1/record.json"
	// 1件のレコードを登録する	POST
	ADD_RECODE_URL = "/k/v1/record.json"
	// 1件のレコードを更新する	PUT
	UPDATE_RECODE_URL = "/k/v1/record.json"
	// 複数のレコードを取得する	GET
	GET_RECODES_URL = "/k/v1/records.json"
	// 複数のレコードを登録する	POST
	ADD_RECODES_URL = "/k/v1/records.json"
	// 複数のレコードを更新する	PUT
	UPDATE_RECODES_URL = "/k/v1/records.json"
)
