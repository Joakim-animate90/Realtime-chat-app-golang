package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type Message struct {
	ID          int32                 `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time             `json:"createAt"`
	UpdatedAt   time.Time             `json:"updatedAt"`
	DeletedAt   soft_delete.DeletedAt `json:"deletedAt"`
	FromUserId  int32                 `json:"fromUserId" gorm:"index"`
	ToUserId    int32                 `json:"toUserId" gorm:"index;comment:'The ID sent to the end can be a user ID or a group ID'"`
	Content     string                `json:"content" gorm:"type:varchar(2500)"`
	MessageType int16                 `json:"messageType" gorm:"comment:'Message type: 1 single chat, 2 group chat'"`
	ContentType int16                 `json:"contentType" gorm:"comment:'Message content: 1 Text 2. Ordinary File 3. Picture 4. Audio 5. Video 6. Voice Chat 7. Video chat'"`
	Pic         string                `json:"pic" gorm:"type:text;comment:'Thumbnail"`
	Url         string                `json:"url" gorm:"type:varchar(350);comment:'File or picture address'"`
}
