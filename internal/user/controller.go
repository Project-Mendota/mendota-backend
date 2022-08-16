package user

type Controller struct {
	service Service
}

func NewUserController(service Service) Controller {
	return Controller{service: service}
}
