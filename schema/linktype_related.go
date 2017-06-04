package schema

import (
	"themis/models"
)

func createLinkTypeRelated() []models.LinkType {
	linkTypes := []models.LinkType {
		createLinkTypeRelatedStoryToTask(),
		createLinkTypeRelatedBugToTask(),
		createLinkTypeRelatedStoryToBug(),
	}
	return linkTypes
}

func createLinkTypeRelatedStoryToTask() models.LinkType {
	linkType := models.NewLinkType() 
	linkType.Name = "Related Link"
	linkType.Description = "The related relationship."
	linkType.Version = 0
	linkType.ForwardName = "is related to"
	linkType.ReverseName = "is related to"
	linkType.Topology = "graph"
	linkType.CategoryRef = "default"
	linkType.SourceWorkItemTypeRef = "story"
	linkType.TargetWorkItemTypeRef = "task"
	return *linkType
}

func createLinkTypeRelatedBugToTask() models.LinkType {
	linkType := models.NewLinkType() 
	linkType.Name = "Related Link"
	linkType.Description = "The related relationship."
	linkType.Version = 0
	linkType.ForwardName = "is related to"
	linkType.ReverseName = "is related to"
	linkType.Topology = "graph"
	linkType.CategoryRef = "default"
	linkType.SourceWorkItemTypeRef = "bug"
	linkType.TargetWorkItemTypeRef = "task"
	return *linkType
}

func createLinkTypeRelatedStoryToBug() models.LinkType {
	linkType := models.NewLinkType() 
	linkType.Name = "Related Link"
	linkType.Description = "The related relationship."
	linkType.Version = 0
	linkType.ForwardName = "is related to"
	linkType.ReverseName = "is related to"
	linkType.Topology = "graph"
	linkType.CategoryRef = "default"
	linkType.SourceWorkItemTypeRef = "story"
	linkType.TargetWorkItemTypeRef = "bug"
	return *linkType
}
