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

// LinkCategoryResource for api2go routes.
type LinkCategoryResource struct {
	LinkCategoryStorage *database.LinkCategoryStorage
	LinkTypeStorage *database.LinkTypeStorage
}

func (c LinkCategoryResource) getFilterFromRequest(r api2go.Request) (bson.M, error) {
	var filter bson.M
	// Getting reference context
	sourceContext, sourceContextID, thisContext := utils.ParseContext(r)
	switch sourceContext {
		case models.LinkTypeName:
			linkType, err := c.LinkTypeStorage.GetOne(bson.ObjectIdHex(sourceContextID))
			if (err != nil) {
				return nil, err
			}
			if thisContext == "link_category" {
				filter = bson.M{"_id": linkType.LinkCategoryID}
			}
		default:
			// build standard filter expression
			filter = (utils.BuildDbFilterFromRequest(r)).(bson.M)
	}
	return filter, nil
}

// FindAll LinkCategorys.
func (c LinkCategoryResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	// build filter expression
	filter, err := c.getFilterFromRequest(r)
	if err != nil {
		return &api2go.Response{}, err
	}
	linkCategorys, _ := c.LinkCategoryStorage.GetAll(filter)
	return &api2go.Response{Res: linkCategorys}, nil
}

// PaginatedFindAll can be used to load users in chunks.
// Possible success status code 200.
func (c LinkCategoryResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {

	// build filter expression
	filter, err := c.getFilterFromRequest(r)
	if err!=nil {
		return 0, &api2go.Response{}, err
	}

	// parse out offset and limit
	queryOffset, queryLimit, err := utils.ParsePaging(r)
	if err!=nil {
		return 0, &api2go.Response{}, err
	}

	// get the paged data from storage
	result, err := c.LinkCategoryStorage.GetAllPaged(filter, queryOffset, queryLimit)
	if err!=nil {
		return 0, &api2go.Response{}, err
	}

	// get total count for paging
	allCount, err := c.LinkCategoryStorage.GetAllCount(filter)
	if err!=nil {
		return 0, &api2go.Response{}, err
	}

	// return everything
	return uint(allCount), &api2go.Response{Res: result}, nil
}

// FindOne LinkCategory.
// Possible success status code 200
func (c LinkCategoryResource) FindOne(id string, r api2go.Request) (api2go.Responder, error) {
	utils.DebugLog.Printf("Received FindOne with ID %s.", id)
	res, err := c.LinkCategoryStorage.GetOne(bson.ObjectIdHex(id))
	return &api2go.Response{Res: res}, err
}

// Create a new LinkCategory.
// Possible status codes are:
// - 201 Created: Resource was created and needs to be returned
// - 202 Accepted: Processing is delayed, return nothing
// - 204 No Content: Resource created with a client generated ID, and no fields were modified by
//   the server
func (c LinkCategoryResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	linkCategory, ok := obj.(models.LinkCategory)
	if !ok {
		return &api2go.Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}
	id, _ := c.LinkCategoryStorage.Insert(linkCategory)
	linkCategory.ID = id
	return &api2go.Response{Res: linkCategory, Code: http.StatusCreated}, nil
}

// Delete a LinkCategory.
// Possible status codes are:
// - 200 OK: Deletion was a success, returns meta information, currently not implemented! Do not use this
// - 202 Accepted: Processing is delayed, return nothing
// - 204 No Content: Deletion was successful, return nothing
func (c LinkCategoryResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.LinkCategoryStorage.Delete(bson.ObjectIdHex(id))
	return &api2go.Response{Code: http.StatusOK}, err
}

// Update a LinkCategory.
// Possible status codes are:
// - 200 OK: Update successful, however some field(s) were changed, returns updates source
// - 202 Accepted: Processing is delayed, return nothing
// - 204 No Content: Update was successful, no fields were changed by the server, return nothing
func (c LinkCategoryResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	linkCategory, ok := obj.(models.LinkCategory)
	if !ok {
		return &api2go.Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}
	err := c.LinkCategoryStorage.Update(linkCategory)
	return &api2go.Response{Res: linkCategory, Code: http.StatusNoContent}, err
}
