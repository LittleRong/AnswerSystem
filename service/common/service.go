package common

type service struct{}

func (this *service) init() {
	DatabaseInit()
}
