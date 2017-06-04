package fixtures

import (
	"themis/models"
)

func createLinkTypeChild() models.LinkType {
	linkType := models.NewLinkType()
	linkType.Name = "Child Link"
	linkType.Description = "The parent-child relationship."
	linkType.Version = 0
	linkType.ForwardName = "is child of"
	linkType.ReverseName = "is parent of"
	linkType.Topology = "graph"
	return *linkType
}
