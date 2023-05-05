package services

type IUserService interface {
	GetProfile()
}

type userService struct {
}

func NewUserService() IUserService {
	return &userService{}
}

func (*userService) GetProfile() {

}
