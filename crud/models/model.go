package models

type (
	Book struct {
		Id     string `json:"id"`
		Name   string `json:"name"`
		Author string `json:"author"`
	}
	User struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Books []Book `json:"books"`
	}
	DeleteBook struct {
		Id string `json:"id"`
	}
	Successfull struct {
		Success string `json:"success"`
		Err     error  `json:"err"`
	}
)
