package errors

type UserAlreadyExist struct {
}

func (m *UserAlreadyExist) Error() string {
	return "User already exist"
}

var (
	UserAlreadyExistError = &UserAlreadyExist{}
)
