package schema

import (
	"themis/models"
)

func createWorkItemTypeTask() models.WorkItemType {
	workItemType := models.NewWorkItemType()
	workItemType.RefID = "task"
	workItemType.Name = "Task"
	workItemType.Description = "A story task."
	workItemType.Version = 0
	workItemType.Icon = "fa fa-bolt"
	workItemType.Fields = map[string]models.WorkItemTypeField{
		"system.area": models.WorkItemTypeField{
			Description: "The area to which the work item belongs",
			Label:       "Area",
			Required:    false,
			Type: models.WorkItemTypeFieldDescriptor{
				Kind: "area",
			},
		},
		"system.assignees": models.WorkItemTypeField{
			Description: "The users that are assigned to the work item",
			Label:       "Assignees",
			Required:    false,
			Type: models.WorkItemTypeFieldDescriptor{
				ComponentType: "user",
				Kind:          "list",
			},
		},
		"system.created_at": models.WorkItemTypeField{
			Description: "The date and time when the work item was created",
			Label:       "Created at",
			Required:    false,
			Type: models.WorkItemTypeFieldDescriptor{
				Kind: "instant",
			},
		},
		"system.updated_at": models.WorkItemTypeField{
			Description: "The date and time when the work item was last updated",
			Label:       "Updated at",
			Required:    false,
			Type: models.WorkItemTypeFieldDescriptor{
				Kind: "instant",
			},
		},
		"system.creator": models.WorkItemTypeField{
			Description: "The user that created the work item",
			Label:       "Creator",
			Required:    true,
			Type: models.WorkItemTypeFieldDescriptor{
				Kind: "user",
			},
		},
		"system.description": models.WorkItemTypeField{
			Description: "A descriptive text of the work item",
			Label:       "Description",
			Required:    false,
			Type: models.WorkItemTypeFieldDescriptor{
				Kind: "markup",
			},
		},
		"system.iteration": models.WorkItemTypeField{
			Description: "The iteration to which the work item belongs",
			Label:       "Iteration",
			Required:    false,
			Type: models.WorkItemTypeFieldDescriptor{
				Kind: "iteration",
			},
		},
		"system.order": models.WorkItemTypeField{
			Description: "Execution Order of the workitem",
			Label:       "Execution Order",
			Required:    false,
			Type: models.WorkItemTypeFieldDescriptor{
				Kind: "float",
			},
		},
		"system.state": models.WorkItemTypeField{
			Description: "The state of the work item",
			Label:       "State",
			Required:    true,
			Type: models.WorkItemTypeFieldDescriptor{
				BaseType: "string",
				Kind:     "enum",
				Values: []string{
					"new",
					"open",
					"in progress",
					"resolved",
					"closed",
				},
			},
		},
		"system.title": models.WorkItemTypeField{
			Description: "The title text of the work item",
			Label:       "Title",
			Required:    true,
			Type: models.WorkItemTypeFieldDescriptor{
				Kind: "string",
			},
		},
	}
	return *workItemType
}
