package model

// AiPostResponse AI算法传回来的数据
type AiPostResponse struct {
	Result string
	Data   *[]byte
}
