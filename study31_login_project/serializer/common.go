package serializer

/*
	code : 0 成功,非0 失败
	msg: 消息
	data: 数据
	err : 错误信息

*/
type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Msg   interface{} `json:"msg"`
	Error string      `json:"err"`
}
