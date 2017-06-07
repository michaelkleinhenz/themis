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

// SpaceResource for api2go routes.
type SpaceResource struct {
	SpaceStorage *database.SpaceStorage
	WorkItemTypeStorage *database.WorkItemTypeStorage
	AreaStorage *database.AreaStorage

	WorkItemStorage *database.WorkItemStorage
	IterationStorage *database.IterationStorage
	LinkCategoryStorage *database.LinkCategoryStorage
	LinkStorage *database.LinkStorage
	LinkTypeStorage *database.LinkTypeStorage
}

func (c SpaceResource) getFilterFromRequest(r api2go.Request) (bson.M, error) {
	var filter bson.M
	// Getting reference context
	// TODO: find a more elegant way, maybe using function literals.
	sourceContext, sourceContextID, thisContext := utils.ParseContext(r)
	switch sourceContext {
		case models.WorkItemTypeName:
			entity, err := c.WorkItemTypeStorage.GetOne(bson.ObjectIdHex(sourceContextID))
			if (err != nil) {
				return nil, err
			}
			// TODO find out why "spaces" on the WorkItemType and "space" on the area.
			if thisContext == "space" || thisContext == "spaces" {
				filter = bson.M{"_id": entity.SpaceID}
			}
		case models.AreaName:
			entity, err := c.AreaStorage.GetOne(bson.ObjectIdHex(sourceContextID))
			if (err != nil) {
				return nil, err
			}
			if thisContext == "space" || thisContext == "spaces" {
				filter = bson.M{"_id": entity.SpaceID}
			}
		case models.WorkItemName:
			entity, err := c.WorkItemStorage.GetOne(bson.ObjectIdHex(sourceContextID))
			if (err != nil) {
				return nil, err
			}
			if thisContext == "space" || thisContext == "spaces" {
				filter = bson.M{"_id": entity.SpaceID}
			}
		case models.IterationName:
			entity, err := c.IterationStorage.GetOne(bson.ObjectIdHex(sourceContextID))
			if (err != nil) {
				return nil, err
			}
			if thisContext == "space" || thisContext == "spaces" {
				filter = bson.M{"_id": entity.SpaceID}
			}
		case models.LinkCategoryName:
			entity, err := c.LinkCategoryStorage.GetOne(bson.ObjectIdHex(sourceContextID))
			if (err != nil) {
				return nil, err
			}
			if thisContext == "space" || thisContext == "spaces" {
				filter = bson.M{"_id": entity.SpaceID}
			}
		case models.LinkTypeName:
			entity, err := c.LinkTypeStorage.GetOne(bson.ObjectIdHex(sourceContextID))
			if (err != nil) {
				return nil, err
			}
			if thisContext == "space" || thisContext == "spaces" {
				filter = bson.M{"_id": entity.SpaceID}
			}
		case models.LinkName:
			entity, err := c.LinkStorage.GetOne(bson.ObjectIdHex(sourceContextID))
			if (err != nil) {
				return nil, err
			}
			if thisContext == "space" || thisContext == "spaces" {
				filter = bson.M{"_id": entity.SpaceID}
			}
		default:
			// build standard filter expression
			filter = (utils.BuildDbFilterFromRequest(r)).(bson.M)
	}
	return filter, nil
}

// FindAll Spaces.
func (c SpaceResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	// build filter expression
	filter, err := c.getFilterFromRequest(r)
	if err != nil {
		return api2go.Response{}, err
	}
	spaces, err := c.SpaceStorage.GetAll(filter)
	if err != nil {
		return api2go.Response{}, err
	}
	return &api2go.Response{Res: spaces}, nil
}

// PaginatedFindAll can be used to load users in chunks.
// Possible success status code 200.
func (c SpaceResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {

	// build filter expression
	filter, err := c.getFilterFromRequest(r)
	if err != nil {
		return 0, api2go.Response{}, err
	}

	// parse out offset and limit
	queryOffset, queryLimit, err := utils.ParsePaging(r)
	if err!=nil {
		return 0, &api2go.Response{}, err
	}

	// get the paged data from storage
	result, err := c.SpaceStorage.GetAllPaged(filter, queryOffset, queryLimit)
	if err!=nil {
		return 0, &api2go.Response{}, err
	}

	// get total count for paging
	allCount, err := c.SpaceStorage.GetAllCount(filter)
	if err!=nil {
		return 0, &api2go.Response{}, err
	}

	// return everything
	return uint(allCount), &api2go.Response{Res: result}, nil
}

// FindOne Space.
// Possible success status code 200
func (c SpaceResource) FindOne(id string, r api2go.Request) (api2go.Responder, error) {
	utils.DebugLog.Printf("Received FindOne with ID %s.", id)
	res, err := c.SpaceStorage.GetOne(bson.ObjectIdHex(id))
	return &api2go.Response{Res: res}, err
}

// Create a new Space.
// Possible status codes are:
// - 201 Created: Resource was created and needs to be returned
// - 202 Accepted: Processing is delayed, return nothing
// - 204 No Content: Resource created with a client generated ID, and no fields were modified by
//   the server
func (c SpaceResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	space, ok := obj.(models.Space)
	if !ok {
		return &api2go.Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}
	id, _ := c.SpaceStorage.Insert(space)
	space.ID = id
	return &api2go.Response{Res: space, Code: http.StatusCreated}, nil
}

// Delete a Space.
// Possible status codes are:
// - 200 OK: Deletion was a success, returns meta information, currently not implemented! Do not use this
// - 202 Accepted: Processing is delayed, return nothing
// - 204 No Content: Deletion was successful, return nothing
func (c SpaceResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.SpaceStorage.Delete(bson.ObjectIdHex(id))
	return &api2go.Response{Code: http.StatusOK}, err
}

// Update a Space.
// Possible status codes are:
// - 200 OK: Update successful, however some field(s) were changed, returns updates source
// - 202 Accepted: Processing is delayed, return nothing
// - 204 No Content: Update was successful, no fields were changed by the server, return nothing
func (c SpaceResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	space, ok := obj.(models.Space)
	if !ok {
		return &api2go.Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}
	err := c.SpaceStorage.Update(space)
	return &api2go.Response{Res: space, Code: http.StatusNoContent}, err
}
