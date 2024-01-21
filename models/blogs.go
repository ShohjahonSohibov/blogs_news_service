package models

type Blog struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type BlogId struct {
	Id string `json:"id"`
}

type CreateBlog struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}


type UpdateBlog struct {
	Id        string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type GetListBlogsRequest struct {
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
	Title  string `json:"title"`
}

type GetListBlogsResponse struct {
	Count  int32  `json:"count"`
	Blogs  []Blog `json:"blogs"`
}
