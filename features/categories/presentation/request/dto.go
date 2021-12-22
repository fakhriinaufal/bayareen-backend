package request

import "bayareen-backend/features/categories"

type Category struct {
	Name   string `json:"name"`
	ImgUrl string `json:"img_url"`
}

func (c *Category) ToCore() categories.Core {
	return categories.Core{
		Name:   c.Name,
		ImgUrl: c.ImgUrl,
	}
}
