package views

// R - ...
type R struct {
	Status    string      `json:"status"`
	ErrorCode int         `json:"error_code"`
	ErrorNote string      `json:"error_note"`
	Data      interface{} `json:"data"`
}

type Response struct {
	RequestID string `json:"requestId"`
	Code      string `json:"code"`
	Message   string `json:"message"`
}

type Respdata struct {
	HTTPStatus   string   `json:"httpStatus"`
	ErrorCode    string   `json:"errorCode"`
	ErrorMessage string   `json:"errorMessage"`
	Response     Response `json:"response"`
}

type ResponseToIABS struct {
	Type     string   `json:"msgtype"`
	ID       string   `json:"msgid"`
	Corrid   string   `json:"msgcorrid"`
	Date     string   `json:"msgdate"`
	Source   string   `json:"msgsource"`
	Respcode string   `json:"msgrespcode"`
	Respdata Respdata `json:"msgrespdata"`
}
