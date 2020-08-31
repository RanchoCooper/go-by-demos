package service

// Service handle business
type Service struct {
	m MailSender
	u UserRepo
	h Handler
}

type Handle struct {
}

func (H *Handle) DoA() {

}

func (H *Handle) DoB() {

}

type Handler interface {
	DoA()
	DoB()
}

func (s *Service) UserSignUp() {
	s.u.AddUser()
	s.m.Send()
	s.h.DoA()
}

func NewService(m MailSender, u UserRepo) *Service {
	return &Service{m: m, u: u, h: &Handle{}}
}
