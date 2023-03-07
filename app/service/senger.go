package service

type SenderService struct {
}

func (s *SenderService) Send() {
	return
}

func NewSenderService() *SenderService {
	return &SenderService{}
}
