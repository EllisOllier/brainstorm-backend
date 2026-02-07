package page

type Page struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"Content"`
}

type PageService struct {
	pageRepository *PageRepository
}

func NewPageService(givenPageRepository *PageRepository) *PageService {
	return &PageService{
		pageRepository: givenPageRepository,
	}
}
