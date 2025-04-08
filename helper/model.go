package helper

import (
	"farhanhr/go-restful-api/model/domain"
	"farhanhr/go-restful-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var CategoryResponses []web.CategoryResponse
	for _, category := range categories {
		CategoryResponses = append(CategoryResponses, ToCategoryResponse(category))
	}

	return CategoryResponses
}