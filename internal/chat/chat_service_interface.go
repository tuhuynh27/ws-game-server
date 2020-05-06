package chat

type IService interface {
	ListService() ([]*Chat, error)
	InsertOneService(newChat Chat) (string, error)
}
