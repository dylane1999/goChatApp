package types

type ChatMessage struct {
	Username string`json:"username"`
	Text string`json:"text"`
	ChatroomId string `json:"chatroomId"`
}