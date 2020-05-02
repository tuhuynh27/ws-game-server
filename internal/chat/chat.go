package chat

const (
	CollectionName = "chat"
)

type Chat struct {
	User		string `json:"user" bson:"user"`
	Message     string `json:"message" bson:"message"`
	Time		int64  `json:"time" bson:"time"`
}
