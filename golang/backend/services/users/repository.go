package user

type UserRepository struct {
    // DB connection, etc.
}

func NewUserRepository() *UserRepository {
    return &UserRepository{}
}
