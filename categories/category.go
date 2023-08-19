package categories

type Category struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"imageURL"`
	Ordering    int    `json:"ordering"`
	Active      bool   `json:"active"`
}

type Repository interface {
	GetAll() ([]Category, error)
}

func GetAll(repo Repository) ([]Category, error) {
	return repo.GetAll()
}
