package main

// book request
type CreateBookFormat struct {
	Name     string `json:"id" binding:"required"`
	Category uint   `json:"category" binding:"required"`
}

type UpdateBookFormat struct {
	ID       uint   `json:"id" binding:"required"`
	Name     string `json:"name"`
	Category uint   `json:"category"`
}

type RemoveBookFormat struct {
	ID uint `json:"id"`
}

// book response
type GetBookListResFormat struct {
	Books []Book `json:"books"`
}

type GetOneBookResFormat struct {
	Book Book `json:"book"`
}

type UpdateBookResFormat struct {
	Book Book `json:"book"`
}

type RemoveBookResFormat struct {
	Book Book `json:"book"`
}

// category request
type CreateCategoryFormat struct {
	Name string `json:"name" binding:"required"`
}

type UpdateCategoryFormat struct {
	// ID   uint   `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type RemoveCategoryFormat struct {
	ID uint `json:"id" binding:"required"`
}

// category response
type GetCategoryResFormat struct {
	// ID   uint   `json:"category_id"`
	// Name string `json:"name"`
	Category Category `json:"category"`
}

type GetCategoryListResFormat struct {
	Category []Category `json:"categories"`
}

type CreateCategoryResFormat struct {
	Message  string   `json:"message"`
	Category Category `json:"category"`
}

type UpdateCategoryResFormat struct {
	Message  string   `json:"message"`
	Category Category `json:"category"`
}

type RemoveCategoryResFormat struct {
	Message string `json:"message"`
	// Category Category `json:"category"`
}
