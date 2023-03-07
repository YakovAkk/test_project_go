package service

type Sender interface {
	Send()
}

type Service struct {
	Sender
}

func NewService() *Service {
	return &Service{
		Sender: NewSenderService(),
	}
}
