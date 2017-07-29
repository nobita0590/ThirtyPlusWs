package models

type (
	Category struct {
		Id                  int
		CategoryName        string
		PrettyUrl           string
		IsActive            bool
		Stt                 int
		Description         string
		ParentId            int
		IsEnd               bool
		IsAdvisory          bool
	}

	CategoryModel  struct {
		MainModel
	}
	CategoryFilter struct {

	}
)
