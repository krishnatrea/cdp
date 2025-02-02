package domain

type ProjectRepository interface {
	Create()
	GetByID()
	AssignRole()
	GetAllByUserId()
}
