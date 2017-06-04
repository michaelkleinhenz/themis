package schema

import (
	"themis/models"
)

func createLinkTypeRelated() models.LinkType {
	linkType := models.NewLinkType()
	linkType.Name = "Related Link"
	linkType.Description = "The related relationship."
	linkType.Version = 0
	linkType.ForwardName = "is related to"
	linkType.ReverseName = "is related to"
	linkType.Topology = "graph"
	return *linkType
}
