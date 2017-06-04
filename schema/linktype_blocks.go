package schema

import (
	"themis/models"
)

func createLinkTypeBlocks() models.LinkType {
	linkType := models.NewLinkType()
	linkType.Name = "Blocker Link"
	linkType.Description = "The blocker relationship."
	linkType.Version = 0
	linkType.ForwardName = "blocks"
	linkType.ReverseName = "is blocked by"
	linkType.Topology = "graph"
	return *linkType
}
