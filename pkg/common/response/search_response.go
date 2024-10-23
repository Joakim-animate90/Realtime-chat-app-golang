package response

import "realtime-chat-app/internal/model"

type SearchResponse struct {
	User  model.User  `json:"user"`
	Group model.Group `json:"group"`
}
