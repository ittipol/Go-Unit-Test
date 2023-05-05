package repositories

type User struct {
	ID       uint
	Email    string
	Password string
	Name     string
}

type UserRepository interface {
	Login(email string, password string) (*User, error)
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (*userRepository) Login(email string, password string) (*User, error) {
	return &User{}, nil
}
