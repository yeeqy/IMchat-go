package api

import (
	"github.com/gorilla/websocket"
)

type User struct {
	Id        int    `json:"id" form:"id"`
	Username  string `json:"username" form:"username"`
	UpdatedAt string `json:"updated_at" form:"updated_at"`
	Avatar    string `json:"avatar" form:"avatar"`
}

type UserResp struct {
	Flag  bool        `json:"flag"`
	Data  interface{} `json:"data"`
	Msg   string      `json:"message"`
	Error string      `json:"error"`
}
type UserData struct {
	*User `json:"user"`
	Token string `json:"token"`
}

type MessageForSelect struct {
	FromUserId int `json:"from_user_id"`
	ToUserId   int `json:"to_user_id"`
	PageNum    int `json:"page_num"`
	PageSize   int `json:"page_size"`
}

type BroadCast struct {
	IsSystem bool    `json:"is_system"`
	Message  *Client `json:"message"`
}

type Client struct {
	User *User `json:"fromUser"`
	conn *websocket.Conn
	//manage *ManageClient
}

type WebsocketType string

const (
	WS_TYPE_HEARTBEAT    WebsocketType = "heartbeat"
	WS_TYPE_ONLINE_USERS WebsocketType = "onlineUsers"

	WS_TYPE_USER_ONLINE  WebsocketType = "userOnline"
	WS_TYPE_USER_OFFLINE WebsocketType = "userOffline"

	WS_TYPE_CHAT_USER_TO_USER   WebsocketType = "chatUserToUser"
	WS_TYPE_CHAT_USER_TO_PUBLIC WebsocketType = "chatUserToPublic"
	// 群组...

	WS_TYPE_PARAMS_INVALID WebsocketType = "paramsInvalid"
	WS_TYPE_SERVER_ERROR   WebsocketType = "error"
	WS_TYPE_NO_REPLY       WebsocketType = "noreply"
)

type WebsocketRequest struct {
	ID   int           `json:"id"`
	Type WebsocketType `json:"type"`
	Data interface{}   `json:"data"`

	FromUserId int    `json:"fromUserId"`
	FromUser   *User  `json:"fromUser"`
	Text       string `json:"text"`
	ToUserId   int    `json:"toUserId"`
	PageNum    int    `json:"page_num"`
	PageSize   int    `json:"page_size"`
}

type WebsocketResponse struct {
	ID   int           `json:"id"`
	Type WebsocketType `json:"type"`
	Data interface{}   `json:"data"`

	FromUserId int    `json:"fromUserId"`
	FromUser   *User  `json:"fromUser"`
	Text       string `json:"text"`
	ToUserId   int    `json:"toUserId"`
}
