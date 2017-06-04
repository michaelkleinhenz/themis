package schema

import (
	"themis/models"
)

func createLinkCategoryDefault() models.LinkCategory {
	linkCategory := models.NewLinkCategory()
	linkCategory.Name = "Default Link Category"
	linkCategory.Description = "The default link category."
	linkCategory.Version = 0
	return *linkCategory
}
