package result

type Result struct {
	Code int         `bson:"code" json:"code"`
	Msg  string      `bson:"msg" json:"msg"`
	Data interface{} `bson:"data" json:"data"`
}

func GetSuccessJson() interface{} {
	return &Result{Code: 0, Msg: "ok"}
}

func GetSuccessJsonWithData(content interface{}) interface{} {
	return &Result{Code: 0, Msg: "ok", Data: content}
}

func GetFailJson(message string) interface{} {
	return &Result{Code: -1, Msg: message}
}
