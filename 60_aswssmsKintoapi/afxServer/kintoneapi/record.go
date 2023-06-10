package kintoneapi

// 新規
type RecordCreate struct {
	RequestStatus CommonType `json:"request_status"`
	ReceptionUser CommonType `json:"reception_user"`
	// Revision          CommonType `json:"$revision"`
	UpdateDate        CommonType `json:"update_date"`
	PhoneNo           CommonType `json:"phoneNo"`
	ReceptionContents CommonType `json:"reception_contents"`
	UpdateUser        CommonType `json:"update_user"`
	RequestTitle      CommonType `json:"request_title"`
	RequestContents   CommonType `json:"request_contents"`
	// RecordNo          CommonType `json:"record_no"`
	CreateUser CommonType `json:"create_user"`
	CreateDate CommonType `json:"create_date"`
	Group      CommonType `json:"group"`
	// ID                CommonType `json:"$id"`
}

//　検索
type RecordQuery struct {
	RequestStatus     CommonType `json:"request_status"`
	ReceptionUser     CommonType `json:"reception_user"`
	Revision          CommonType `json:"$revision"`
	UpdateDate        CommonType `json:"update_date"`
	PhoneNo           CommonType `json:"phoneNo"`
	ReceptionContents CommonType `json:"reception_contents"`
	UpdateUser        CommonType `json:"update_user"`
	RequestTitle      CommonType `json:"request_title"`
	RequestContents   CommonType `json:"request_contents"`
	RecordNo          CommonType `json:"record_no"`
	CreateUser        CommonType `json:"create_user"`
	CreateDate        CommonType `json:"create_date"`
	Group             CommonType `json:"group"`
	ID                CommonType `json:"$id"`
}

//　更新
type RecordUdp struct {
	RequestStatus CommonType `json:"request_status"`
	// UpdateDate        CommonType `json:"update_date"`
	ReceptionContents CommonType `json:"reception_contents"`
	// UpdateUser        CommonType `json:"update_user"`
	// Revision CommonType `json:"$revision"`
	// RequestTitle      CommonType `json:"request_title"`
	RequestContents CommonType `json:"request_contents"`
	PhoneNo         CommonType `json:"phoneNo"`
	// RecordNo CommonType `json:"record_no"`
}

type CommonType struct {
	// Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

type User struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Group struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// {
// 	"record": {
// 	  "request_status": {
// 		"type": "DROP_DOWN",
// 		"value": "受付"
// 	  },
// 	  "reception_user": {
// 		"type": "USER_SELECT",
// 		"value": [
// 		  {
// 			"code": "en.b@di-square.co.jp",
// 			"name": "YAN"
// 		  }
// 		]
// 	  },
// 	  "$revision": {
// 		"type": "__REVISION__",
// 		"value": "1"
// 	  },
// 	  "update_date": {
// 		"type": "UPDATED_TIME",
// 		"value": "2023-03-29T05:49:00Z"
// 	  },
// 	  "reception_contents": {
// 		"type": "MULTI_LINE_TEXT",
// 		"value": ""
// 	  },
// 	  "update_user": {
// 		"type": "MODIFIER",
// 		"value": {
// 		  "code": "en.b@di-square.co.jp",
// 		  "name": "YAN"
// 		}
// 	  },
// 	  "request_title": {
// 		"type": "SINGLE_LINE_TEXT",
// 		"value": "Test"
// 	  },
// 	  "request_contents": {
// 		"type": "MULTI_LINE_TEXT",
// 		"value": "問い合わせ内容"
// 	  },
// 	  "record_no": {
// 		"type": "RECORD_NUMBER",
// 		"value": "1"
// 	  },
// 	  "create_user": {
// 		"type": "CREATOR",
// 		"value": {
// 		  "code": "en.b@di-square.co.jp",
// 		  "name": "YAN"
// 		}
// 	  },
// 	  "create_date": {
// 		"type": "CREATED_TIME",
// 		"value": "2023-03-29T05:49:00Z"
// 	  },
// 	  "group": {
// 		"type": "GROUP_SELECT",
// 		"value": [
// 		  {
// 			"code": "Administrators",
// 			"name": "Administrators"
// 		  }
// 		]
// 	  },
// 	  "$id": {
// 		"type": "__ID__",
// 		"value": "1"
// 	  }
// 	}
//   }
