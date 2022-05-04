package api

import (
	"IM-chat/model"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var Manage = NewManageClient()

func HandleWebsocketConnect(c *gin.Context) {
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		}}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}

	id := c.Param("id")
	userId, _ := strconv.Atoi(id)
	cUser := GetUser(userId)

	//实例化新连接的用户
	client := &Client{
		User: cUser,
		conn: conn,
	}

	Manage.Register <- client

}

type ManageClient struct {
	OnlineClients map[int]*Client
	Register      chan *Client

	Lock sync.Mutex
}

func NewManageClient() *ManageClient {

	mm := &ManageClient{
		OnlineClients: make(map[int]*Client),
		Register:      make(chan *Client, 100),
		Lock:          sync.Mutex{},
	}

	return mm
}

func (m *ManageClient) Run() {
	// go m.receive()

	m.receive()
}

func (m *ManageClient) receive() {

	for {
		select {
		case client := <-Manage.Register:

			m.broudcastOnline(client.User)
			m.sendUserList(client.conn)

			m.Lock.Lock()
			Manage.OnlineClients[client.User.Id] = client
			m.Lock.Unlock()

			log.Println("收到一个连接::-------", client.conn.RemoteAddr(), client.User.Id)
			log.Println("当前连接数量::", len(m.OnlineClients))

			go m._receive(client)
		default:
			time.Sleep(time.Millisecond * 10)
		}
	}
}

func (m *ManageClient) _receive(client *Client) {

	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error:
			log.Println("websocket-----runtime error:", err)
		default:
		}
	}()

	// 接收消息 或者 断开连接
	for {
		_, message, err := client.conn.ReadMessage()

		if err == io.EOF || websocket.IsCloseError(err, 1005, 1001, 1002, 1003, 1004) {
			log.Println("当前连接数量::", len(m.OnlineClients))
			log.Printf("[%v]-[%v]断开了websocket连接-----", client.conn.RemoteAddr(), client.User.Id)
			delete(m.OnlineClients, client.User.Id)

			log.Println("当前连接数量::", len(m.OnlineClients))

			m.broudcastOffline(client.User)

			client.conn.Close()
			return
		}

		// 解析发送的消息内容
		log.Printf("[%v]-[%v]发送了:%v", client.conn.RemoteAddr(), client.User.Id, string(message))

		var req = WebsocketRequest{}
		var res = WebsocketResponse{ID: req.ID}
		err = json.Unmarshal(message, &req)
		if err != nil {
			res.Type = WS_TYPE_PARAMS_INVALID
			m._sendJSON(client.conn, res)
			continue
		}

		switch req.Type {
		case WS_TYPE_HEARTBEAT:
			res.Type = WS_TYPE_HEARTBEAT
			m._sendJSON(client.conn, res)
		case WS_TYPE_CHAT_USER_TO_USER:
			m.handleChat1(&req)
		case WS_TYPE_CHAT_USER_TO_PUBLIC:
			m.handleChat2(&req)

		}

	}
}

// SendUserList 先构建系统消息——获取当前连接的所有用户信息
func (m *ManageClient) sendUserList(conn *websocket.Conn) {

	users := make([]*User, 0)

	m.Lock.Lock()
	defer m.Lock.Unlock()

	for _, val := range m.OnlineClients {
		users = append(users, val.User)
	}

	m._sendJSON(conn, WebsocketResponse{Type: WS_TYPE_ONLINE_USERS, Data: users})
}

func (m *ManageClient) broudcastOnline(user *User) {
	m.Lock.Lock()
	defer m.Lock.Unlock()

	res := WebsocketResponse{
		Type: WS_TYPE_USER_ONLINE,
		Data: user,
	}
	for _, val := range m.OnlineClients {
		go m._sendJSON(val.conn, res)
	}
}

func (m *ManageClient) broudcastOffline(user *User) {
	m.Lock.Lock()
	defer m.Lock.Unlock()

	res := WebsocketResponse{
		Type: WS_TYPE_USER_OFFLINE,
		Data: user,
	}

	for _, val := range m.OnlineClients {
		if val.User.Id == user.Id {
			continue
		}

		go m._sendJSON(val.conn, res)
	}
}

// handleChat2 广播消息
func (m *ManageClient) handleChat2(req *WebsocketRequest) {

	m.Lock.Lock()
	from := m.OnlineClients[req.FromUserId]
	m.Lock.Unlock()

	// 回复
	go m._sendJSON(from.conn, WebsocketResponse{Type: WS_TYPE_NO_REPLY, ID: req.ID})

	m.Lock.Lock()

	msg := WebsocketResponse{
		Type:       WS_TYPE_CHAT_USER_TO_PUBLIC,
		Data:       req.Text,
		FromUserId: from.User.Id,
		FromUser:   from.User,
		Text:       req.Text,
	}
	for _, v := range m.OnlineClients {
		go m._sendJSON(v.conn, msg)
	}
	m.Lock.Unlock()

}

// handleChat1 用户发给用户
func (m *ManageClient) handleChat1(req *WebsocketRequest) {

	m.Lock.Lock()
	from := m.OnlineClients[req.FromUserId]
	to := m.OnlineClients[req.ToUserId]
	m.Lock.Unlock()

	// save to db
	model.AddInfo(&model.Info{
		FromUserId: req.FromUserId,
		ToUserId:   req.ToUserId,
		Message:    req.Text,
	})

	go m._sendJSON(from.conn, WebsocketResponse{Type: WS_TYPE_NO_REPLY})

	req.FromUser = from.User
	go m._sendJSON(to.conn, req)
}

func (w *ManageClient) _sendJSON(c *websocket.Conn, v interface{}) {

	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error:
			log.Println("websocket-----runtime--error:", err)
		default:
		}
	}()

	if err := c.WriteJSON(v); err != nil {
		log.Println(err)
	}

}
