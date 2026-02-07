package project

type Project struct {
	ID          int    `json"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ProjectService struct {
	projectRepository *ProjectRepository
}

func NewProjectService(givenProjectRepository *ProjectRepository) *ProjectService {
	return &ProjectService{
		projectRepository: givenProjectRepository,
	}
}
