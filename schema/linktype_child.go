package schema

import (
	"themis/models"
)

func createLinkTypeChild() []models.LinkType {
	linkTypes := []models.LinkType {
		createLinkTypeChildTaskToStory(),
		createLinkTypeChildStoryToBug(),
	}
	return linkTypes
}

func createLinkTypeChildTaskToStory() models.LinkType {
	linkType := models.NewLinkType() 
	linkType.Name = "Child Link"
	linkType.Description = "The parent-child relationship."
	linkType.Version = 0
	linkType.ForwardName = "is child of"
	linkType.ReverseName = "is parent of"
	linkType.Topology = "graph"
	linkType.CategoryRef = "default"
	linkType.SourceWorkItemTypeRef = "task"
	linkType.TargetWorkItemTypeRef = "story"
	return *linkType
}

func createLinkTypeChildStoryToBug() models.LinkType {
	linkType := models.NewLinkType() 
	linkType.Name = "Child Link"
	linkType.Description = "The parent-child relationship."
	linkType.Version = 0
	linkType.ForwardName = "is child of"
	linkType.ReverseName = "is parent of"
	linkType.Topology = "graph"
	linkType.CategoryRef = "default"
	linkType.SourceWorkItemTypeRef = "story"
	linkType.TargetWorkItemTypeRef = "bug"
	return *linkType
}
