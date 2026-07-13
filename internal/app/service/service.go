package service

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) TransactionList() any {
	return 0
}
