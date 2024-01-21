package models

type News struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type NewsId struct {
	Id string `json:"id"`
}

type CreateNews struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateNews struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type GetListNewsRequest struct {
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
	Title  string `json:"title"`
}

type GetListNewsResponse struct {
	Count int32  `json:"count"`
	News  []News `json:"news"`
}
