package resources

import (
	"errors"
	"net/http"

	"github.com/manyminds/api2go"
	"gopkg.in/mgo.v2/bson"

	"themis/utils"
	"themis/models"
	"themis/database"
)

// WorkItemTypeResource for api2go routes.
type WorkItemTypeResource struct {
	WorkItemTypeStorage *database.WorkItemTypeStorage
	WorkItemStorage *database.WorkItemStorage
}

func (c WorkItemTypeResource) getFilterFromRequest(r api2go.Request) (bson.M, error) {
	var filter bson.M
	// Getting reference context
	// TODO: find a more elegant way, maybe using function literals.
	sourceContext, sourceContextID, thisContext := utils.ParseContext(r)
	switch sourceContext {
		case models.WorkItemName:
			workItem, err := c.WorkItemStorage.GetOne(bson.ObjectIdHex(sourceContextID))
			if (err != nil) {
				return nil, err
			}
			if thisContext == "baseType" {
				filter = bson.M{"_id": workItem.BaseTypeID}
			}
		default:
			// build standard filter expression
			filter = (utils.BuildDbFilterFromRequest(r)).(bson.M)
	}
	return filter, nil
}

// FindAll WorkItemTypes.
func (c WorkItemTypeResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	// build filter expression
	filter, err := c.getFilterFromRequest(r)
	if err != nil {
		return &api2go.Response{}, err
	}
	workItemTypes, _ := c.WorkItemTypeStorage.GetAll(filter)
	return &api2go.Response{Res: workItemTypes}, nil
}

// PaginatedFindAll can be used to load users in chunks.
// Possible success status code 200.
func (c WorkItemTypeResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {

	// build filter expression
	filter, err := c.getFilterFromRequest(r)
	if err != nil {
		return 0, &api2go.Response{}, err
	}

	// parse out offset and limit
	queryOffset, queryLimit, err := utils.ParsePaging(r)
	if err!=nil {
		return 0, &api2go.Response{}, err
	}

	// get the paged data from storage
	result, err := c.WorkItemTypeStorage.GetAllPaged(filter, queryOffset, queryLimit)
	if err!=nil {
		return 0, &api2go.Response{}, err
	}

	// get total count for paging
	allCount, err := c.WorkItemTypeStorage.GetAllCount(filter)
	if err!=nil {
		return 0, &api2go.Response{}, err
	}

	// return everything
	return uint(allCount), &api2go.Response{Res: result}, nil
}

// FindOne WorkItemType.
// Possible success status code 200
func (c WorkItemTypeResource) FindOne(id string, r api2go.Request) (api2go.Responder, error) {
	utils.DebugLog.Printf("Received FindOne with ID %s.", id)
	res, err := c.WorkItemTypeStorage.GetOne(bson.ObjectIdHex(id))
	return &api2go.Response{Res: res}, err
}

// Create a new WorkItemType.
// Possible status codes are:
// - 201 Created: Resource was created and needs to be returned
// - 202 Accepted: Processing is delayed, return nothing
// - 204 No Content: Resource created with a client generated ID, and no fields were modified by
//   the server
func (c WorkItemTypeResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	workItemType, ok := obj.(models.WorkItemType)
	if !ok {
		return &api2go.Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}
	id, _ := c.WorkItemTypeStorage.Insert(workItemType)
	workItemType.ID = id
	return &api2go.Response{Res: workItemType, Code: http.StatusCreated}, nil
}

// Delete a WorkItemType.
// Possible status codes are:
// - 200 OK: Deletion was a success, returns meta information, currently not implemented! Do not use this
// - 202 Accepted: Processing is delayed, return nothing
// - 204 No Content: Deletion was successful, return nothing
func (c WorkItemTypeResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.WorkItemTypeStorage.Delete(bson.ObjectIdHex(id))
	return &api2go.Response{Code: http.StatusOK}, err
}

// Update a WorkItemType.
// Possible status codes are:
// - 200 OK: Update successful, however some field(s) were changed, returns updates source
// - 202 Accepted: Processing is delayed, return nothing
// - 204 No Content: Update was successful, no fields were changed by the server, return nothing
func (c WorkItemTypeResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	workItemType, ok := obj.(models.WorkItemType)
	if !ok {
		return &api2go.Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}
	err := c.WorkItemTypeStorage.Update(workItemType)
	return &api2go.Response{Res: workItemType, Code: http.StatusNoContent}, err
}
