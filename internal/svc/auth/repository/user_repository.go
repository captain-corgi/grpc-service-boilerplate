package repository

type UserRepository interface {
	CreateUser()

	FindOneUser()
	FindAllUser()

	UpdateUser()
	BulkUpdateUser()

	DeleteUser()
	DeleteUserPermanently()
}
