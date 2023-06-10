package pojo

type NewRequest struct {
	UserName         string `form:"userName"`
	CustomerName     string `form:"customerName"`
	RequestContent   string `form:"requestContent"`
	Office           string `form:"office"`
	PhoneNo          string `form:"phoneNo"`
	Status           string `form:"status"`
	ReceptionContent string `form:"receptionContent"`
	RecordID         string `form:"recordID"`
}
