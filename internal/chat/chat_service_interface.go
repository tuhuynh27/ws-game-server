package chat

type IService interface {
	ListService() ([]*Chat, error)
	InsertOneService(*Chat) (string, error)
}
