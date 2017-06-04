package schema

import (
	"themis/models"
)

func createLinkTypeBlocks() []models.LinkType {
	linkTypes := []models.LinkType {
		createLinkTypeBlocksBugToStory(),
		createLinkTypeBlocksBugToTask(),
		createLinkTypeBlocksTaskToStory(),
		createLinkTypeBlocksStoryToStory(),
	}
	return linkTypes
}

func createLinkTypeBlocksBugToStory() models.LinkType {
	linkType := models.NewLinkType() 
	linkType.Name = "Blocker Link"
	linkType.Description = "The blocker relationship."
	linkType.Version = 0
	linkType.ForwardName = "blocks"
	linkType.ReverseName = "is blocked by"
	linkType.Topology = "graph"
	linkType.CategoryRef = "default"
	linkType.SourceWorkItemTypeRef = "bug"
	linkType.TargetWorkItemTypeRef = "story"
	return *linkType
}

func createLinkTypeBlocksBugToTask() models.LinkType {
	linkType := models.NewLinkType() 
	linkType.Name = "Blocker Link"
	linkType.Description = "The blocker relationship."
	linkType.Version = 0
	linkType.ForwardName = "blocks"
	linkType.ReverseName = "is blocked by"
	linkType.Topology = "graph"
	linkType.CategoryRef = "default"
	linkType.SourceWorkItemTypeRef = "bug"
	linkType.TargetWorkItemTypeRef = "task"
	return *linkType
}

func createLinkTypeBlocksTaskToStory() models.LinkType {
	linkType := models.NewLinkType() 
	linkType.Name = "Blocker Link"
	linkType.Description = "The blocker relationship."
	linkType.Version = 0
	linkType.ForwardName = "blocks"
	linkType.ReverseName = "is blocked by"
	linkType.Topology = "graph"
	linkType.CategoryRef = "default"
	linkType.SourceWorkItemTypeRef = "task"
	linkType.TargetWorkItemTypeRef = "story"
	return *linkType
}

func createLinkTypeBlocksStoryToStory() models.LinkType {
	linkType := models.NewLinkType() 
	linkType.Name = "Blocker Link"
	linkType.Description = "The blocker relationship."
	linkType.Version = 0
	linkType.ForwardName = "blocks"
	linkType.ReverseName = "is blocked by"
	linkType.Topology = "graph"
	linkType.CategoryRef = "default"
	linkType.SourceWorkItemTypeRef = "story"
	linkType.TargetWorkItemTypeRef = "story"
	return *linkType
}
