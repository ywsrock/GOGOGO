package handler

import (
	"afxServer/awssms"
	"afxServer/kintoneapi"
	"afxServer/pojo"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

//新規作成
func NewRequest(c *gin.Context) {
	var newRequestForm pojo.NewRequest
	if err := c.ShouldBind(&newRequestForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//ステータス変更：
	if newRequestForm.Status == "1" {
		newRequestForm.Status = "対応中"
	} else if newRequestForm.Status == "0" {
		newRequestForm.Status = "依頼中"
	} else if newRequestForm.Status == "2" {
		newRequestForm.Status = "対応完了"
	} else if newRequestForm.Status == "3" {
		newRequestForm.Status = "確認完了"
	}

	// Kintone
	k := kintoneapi.NewKintoneClient(kintoneapi.ADD_RECODE_URL)
	r := createRequestBody(newRequestForm)
	res, err := k.Create(r)
	if err != nil {
		c.JSON(300, gin.H{
			"message": err.Error(),
		})
	}

	//　SMS
	sendSMS(newRequestForm)

	c.JSON(200, gin.H{
		"message": "sucess",
		"result":  res,
	})
}

//更新　指定ID
func UpdateRequest(c *gin.Context) {
	var newRequestForm pojo.NewRequest
	if err := c.ShouldBind(&newRequestForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("----->  ", newRequestForm.RecordID)

	id, err := strconv.Atoi(newRequestForm.RecordID)
	if err != nil {
		c.JSON(300, gin.H{
			"message": err.Error(),
		})
		return
	}

	//ステータス変更：
	if newRequestForm.Status == "1" {
		newRequestForm.Status = "対応中"
	} else if newRequestForm.Status == "0" {
		newRequestForm.Status = "依頼中"
	} else if newRequestForm.Status == "2" {
		newRequestForm.Status = "対応完了"
	} else if newRequestForm.Status == "3" {
		newRequestForm.Status = "確認完了"
	}

	// Kintone
	k := kintoneapi.NewKintoneClient(kintoneapi.UPDATE_RECODE_URL)
	recod := kintoneapi.RecordUdp{
		RequestStatus: kintoneapi.CommonType{
			Value: newRequestForm.Status,
		},
		ReceptionContents: kintoneapi.CommonType{
			Value: newRequestForm.ReceptionContent,
		},
		RequestContents: kintoneapi.CommonType{
			Value: newRequestForm.RequestContent,
		},

		PhoneNo: kintoneapi.CommonType{
			Value: newRequestForm.PhoneNo,
		},
	}

	r := &kintoneapi.UpdateRecordRequest{
		App:    kintoneapi.AppID,
		ID:     id,
		Record: recod,
	}

	res, err := k.Update(r)
	if err != nil {
		c.JSON(300, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"result":  res,
	})
}

//1件取得
func GetHistory(c *gin.Context) {
	var newRequestForm pojo.NewRequest
	if err := c.ShouldBind(&newRequestForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(newRequestForm.RecordID)
	if err != nil {
		c.JSON(300, gin.H{
			"message": err.Error(),
		})
		return
	}

	k := kintoneapi.NewKintoneClient(kintoneapi.GET_RECODE_URL)
	r := &kintoneapi.GetRecordRequest{
		App: kintoneapi.AppID,
		ID:  id,
	}
	res, err := k.Query(r)

	if err != nil {
		c.JSON(300, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"message": "success",
		"result":  res,
	})
}

//全件取得
func ShowHistory(c *gin.Context) {
	k := kintoneapi.NewKintoneClient(kintoneapi.GET_RECODES_URL)
	r := &kintoneapi.GetRecordsRequest{
		App:        kintoneapi.AppID,
		TotalCount: true,
	}
	res, err := k.QueryAll(r)

	if err != nil {
		c.JSON(300, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"message": "success",
		"result":  res,
	})
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("%s%s\n", "エラー発生:", err)
		os.Exit(1)
	}
}

//新規作成Body
func createRequestBody(newRequestForm pojo.NewRequest) *kintoneapi.CreateRecordRequest {
	var cut_title string

	// Title=内容の左20文字
	if l := len([]byte(newRequestForm.RequestContent)); l <= 20 {
		cut_title = strings.Trim(string([]rune(newRequestForm.RequestContent)[0:l]), "\n")
	} else {
		cut_title = strings.Trim(string([]rune(newRequestForm.RequestContent)[0:20]), "\n")
	}

	r := &kintoneapi.CreateRecordRequest{
		App: kintoneapi.AppID,
		Record: kintoneapi.RecordCreate{
			RequestStatus: kintoneapi.CommonType{
				Value: newRequestForm.Status,
				// Value: "受付",
			},
			UpdateDate: kintoneapi.CommonType{
				// Value: "2023-03-29T05:49:00Z",
				Value: time.Now().UTC().Format("2006-01-02T15:04:05Z"),
			},
			PhoneNo: kintoneapi.CommonType{
				Value: newRequestForm.PhoneNo,
			},

			ReceptionContents: kintoneapi.CommonType{
				Value: newRequestForm.ReceptionContent,
			},

			UpdateUser: kintoneapi.CommonType{
				Value: kintoneapi.User{
					Code: "en.b@di-square.co.jp",
					Name: "YAN",
				},
			},

			RequestTitle: kintoneapi.CommonType{
				Value: cut_title,
			},

			RequestContents: kintoneapi.CommonType{
				Value: newRequestForm.RequestContent,
			},
			ReceptionUser: kintoneapi.CommonType{
				Value: []kintoneapi.User{
					{
						Code: "en.b@di-square.co.jp",
						Name: "YAN",
					},
				},
			},

			CreateUser: kintoneapi.CommonType{
				Value: kintoneapi.User{
					Code: "en.b@di-square.co.jp",
					Name: "YAN",
				},
			},

			CreateDate: kintoneapi.CommonType{
				// Value: "2023-03-29T05:49:00Z",
				Value: time.Now().UTC().Format("2006-01-02T15:04:05Z"),
			},

			Group: kintoneapi.CommonType{
				Value: []kintoneapi.Group{
					{
						Code: "Administrators",
						Name: "Administrators",
					},
				},
			},
		},
	}
	return r
}

//SMS送信
func sendSMS(newRequestForm pojo.NewRequest) {
	sendMsg := fmt.Sprintf("%s\n%s", newRequestForm.RequestContent, "https://dev.daptkb25b6ans.amplifyapp.com/")

	snsc := &awssms.SnsClient{}
	msgId, err := snsc.SendSms(sendMsg, newRequestForm.PhoneNo)
	checkError(err)

	fmt.Printf("%s->%s\n", newRequestForm.PhoneNo, msgId)
}
