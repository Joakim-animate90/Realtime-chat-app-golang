package server

import (
	"realtime-chat-app/config"
	"realtime-chat-app/internal/service"
	"realtime-chat-app/pkg/common/constant"
	"realtime-chat-app/pkg/common/util"
	"realtime-chat-app/pkg/global/log"
	"realtime-chat-app/pkg/protocol"

	"encoding/base64"
	"io/ioutil"
	"strings"
	"sync"

	"github.com/gogo/protobuf/proto"
	"github.com/google/uuid"
)

var MyServer = NewServer()

type Server struct {
	Clients   map[string]*Client
	mutex     *sync.Mutex
	Broadcast chan []byte
	Register  chan *Client
	Ungister  chan *Client
}

func NewServer() *Server {
	return &Server{
		mutex:     &sync.Mutex{},
		Clients:   make(map[string]*Client),
		Broadcast: make(chan []byte),
		Register:  make(chan *Client),
		Ungister:  make(chan *Client),
	}
}

// Consumption of the message in KAFKA, and then put it directly in Go Channel for uniform consumption
func ConsumerKafkaMsg(data []byte) {
	MyServer.Broadcast <- data
}

func (s *Server) Start() {
	log.Logger.Info("start server", log.Any("start server", "start server..."))
	for {
		select {
		case conn := <-s.Register:
			log.Logger.Info("login", log.Any("login", "new user login in"+conn.Name))
			s.Clients[conn.Name] = conn
			msg := &protocol.Message{
				From:    "System",
				To:      conn.Name,
				Content: "welcome!",
			}
			protoMsg, _ := proto.Marshal(msg)
			conn.Send <- protoMsg

		case conn := <-s.Ungister:
			log.Logger.Info("loginout", log.Any("loginout", conn.Name))
			if _, ok := s.Clients[conn.Name]; ok {
				close(conn.Send)
				delete(s.Clients, conn.Name)
			}

		case message := <-s.Broadcast:
			msg := &protocol.Message{}
			proto.Unmarshal(message, msg)

			if msg.To != "" {
				// General message, such as text messages, video file messages, etc.
				if msg.ContentType >= constant.TEXT && msg.ContentType <= constant.VIDEO {
					// Save the message will only be saved on one end of the Socket to prevent distributed deployment.
					_, exits := s.Clients[msg.From]
					if exits {
						saveMessage(msg)
					}

					if msg.MessageType == constant.MESSAGE_TYPE_USER {
						client, ok := s.Clients[msg.To]
						if ok {
							msgByte, err := proto.Marshal(msg)
							if err == nil {
								client.Send <- msgByte
							}
						}
					} else if msg.MessageType == constant.MESSAGE_TYPE_GROUP {
						sendGroupMessage(msg, s)
					}
				} else {
					// voice calls, video calls, etc., only support single people chatting, do not support group chat
					// Do not save the file, forward it directly
					client, ok := s.Clients[msg.To]
					if ok {
						client.Send <- message
					}
				}

			} else {
				// No corresponding acceptance personnel to broadcast
				for id, conn := range s.Clients {
					log.Logger.Info("allUser", log.Any("allUser", id))

					select {
					case conn.Send <- message:
					default:
						close(conn.Send)
						delete(s.Clients, conn.Name)
					}
				}
			}
		}
	}
}

// Send to the group message, you need to check all the personnel of the group to send in turn
func sendGroupMessage(msg *protocol.Message, s *Server) {
	// The message sent to the group to find all users in the group to send
	users := service.GroupService.GetUserIdByGroupUuid(msg.To)
	for _, user := range users {
		if user.Uuid == msg.From {
			continue
		}

		client, ok := s.Clients[user.Uuid]
		if !ok {
			continue
		}

		fromUserDetails := service.UserService.GetUserDetails(msg.From)
		// Because when sending group chats, From is an individual, to is a group chat with UUID.So when returning the message, modify FORM to a group chat UUID and unify the single chat
		msgSend := protocol.Message{
			Avatar:       fromUserDetails.Avatar,
			FromUsername: msg.FromUsername,
			From:         msg.To,
			To:           msg.From,
			Content:      msg.Content,
			ContentType:  msg.ContentType,
			Type:         msg.Type,
			MessageType:  msg.MessageType,
			Url:          msg.Url,
		}

		msgByte, err := proto.Marshal(&msgSend)
		if err == nil {
			client.Send <- msgByte
		}
	}
}

// Save the message, if it is stored directly, if it is file, voice and other messages, after saving the file, save the corresponding file path
func saveMessage(message *protocol.Message) {
	// If you upload the BASE64 string file, the analysis file is stored
	if message.ContentType == 2 {
		url := uuid.New().String() + ".png"
		index := strings.Index(message.Content, "base64")
		index += 7

		content := message.Content
		content = content[index:]

		dataBuffer, dataErr := base64.StdEncoding.DecodeString(content)
		if dataErr != nil {
			log.Logger.Error("transfer base64 to file error", log.String("transfer base64 to file error", dataErr.Error()))
			return
		}
		err := ioutil.WriteFile(config.GetConfig().StaticPath.FilePath+url, dataBuffer, 0666)
		if err != nil {
			log.Logger.Error("write file error", log.String("write file error", err.Error()))
			return
		}
		message.Url = url
		message.Content = ""
	} else if message.ContentType == 3 {
		// Ordinary file binary upload
		fileSuffix := util.GetFileType(message.File)
		nullStr := ""
		if nullStr == fileSuffix {
			fileSuffix = strings.ToLower(message.FileSuffix)
		}
		contentType := util.GetContentTypeBySuffix(fileSuffix)
		url := uuid.New().String() + "." + fileSuffix
		err := ioutil.WriteFile(config.GetConfig().StaticPath.FilePath+url, message.File, 0666)
		if err != nil {
			log.Logger.Error("write file error", log.String("write file error", err.Error()))
			return
		}
		message.Url = url
		message.File = nil
		message.ContentType = contentType
	}

	service.MessageService.SaveMessage(*message)
}
