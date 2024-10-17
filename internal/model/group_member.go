package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type GroupMember struct {
	ID        int32                 `json:"id" gorm:"primarykey"`
	CreatedAt time.Time             `json:"createAt"`
	UpdatedAt time.Time             `json:"updatedAt"`
	DeletedAt soft_delete.DeletedAt `json:"deletedAt"`
	UserId    int32                 `json:"userId" gorm:"index;comment:'User ID'"`
	GroupId   int32                 `json:"groupId" gorm:"index;comment:'Group ID'"`
	Nickname  string                `json:"nickname" gorm:"type:varchar(350);comment:'Nick name"`
	Mute      int16                 `json:"mute" gorm:"comment:'Whether to prohibit'"`
}
