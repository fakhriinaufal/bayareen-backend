package response

import (
	"bayareen-backend/features/categories"
	"time"
)

type Category struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	ImgUrl    string    `json:"img_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromCore(core *categories.Core) Category {
	return Category{
		Id:        core.Id,
		Name:      core.Name,
		ImgUrl:    core.ImgUrl,
		CreatedAt: core.CreatedAt,
		UpdatedAt: core.UpdatedAt,
	}
}

func FromCoreSlice(coreSlice []categories.Core) []Category {
	resp := []Category{}
	for _, val := range coreSlice {
		resp = append(resp, FromCore(&val))
	}
	return resp
}
