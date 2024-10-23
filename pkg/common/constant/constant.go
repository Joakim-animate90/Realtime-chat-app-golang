package constant

const (
	HEAT_BEAT = "heatbeat"
	PONG      = "pong"

	// Message type, single chat or group chat
	MESSAGE_TYPE_USER  = 1
	MESSAGE_TYPE_GROUP = 2

	// Message content type
	TEXT         = 1
	FILE         = 2
	IMAGE        = 3
	AUDIO        = 4
	VIDEO        = 5
	AUDIO_ONLINE = 6
	VIDEO_ONLINE = 7

	// Message queue type
	GO_CHANNEL = "gochannel"
	KAFKA      = "kafka"
)
