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

// UserResource for api2go routes.
type UserResource struct {
	UserStorage *database.UserStorage
	SpaceStorage *database.SpaceStorage
}

func (c UserResource) getFilterFromRequest(r api2go.Request) (bson.M, error) {
	var filter bson.M
	// Getting reference context
	sourceContext, sourceContextID, thisContext := utils.ParseContext(r)
	// context: space/owned-by, returns "spaces", "owned-by", spaceID
	// TODO: we may want to replace that NewXX().GetCollectionName() with constants.
	if (sourceContext == models.NewSpace().GetCollectionName()) {
		space, err := c.SpaceStorage.GetOne(bson.ObjectIdHex(sourceContextID))
		if (err != nil) {
			return nil, err
		}
		if thisContext == "owned-by" {
			filter = bson.M{"_id": space.OwnerID}
		}
	} else {
		// build standard filter expression
		filter = (utils.BuildDbFilterFromRequest(r)).(bson.M)
	}
	return filter, nil
}

// FindAll Users.
func (c UserResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	// get filter expression
	filter, err := c.getFilterFromRequest(r)
	if err != nil {
		return api2go.Response{}, err
	}

	// do query
	users, err := c.UserStorage.GetAll(filter)
	if err != nil {
		return api2go.Response{}, err
	}
	return &api2go.Response{Res: users}, nil
}

// PaginatedFindAll can be used to load users in chunks.
// Possible success status code 200.
func (c UserResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {

	// get filter expression
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
	result, err := c.UserStorage.GetAllPaged(filter, queryOffset, queryLimit)
	if err!=nil {
		return 0, &api2go.Response{}, err
	}

	// get total count for paging
	allCount, err := c.UserStorage.GetAllCount(filter)
	if err!=nil {
		return 0, &api2go.Response{}, err
	}

	// return everything
	return uint(allCount), &api2go.Response{Res: result}, nil
}

// FindOne User.
// Possible success status code 200
func (c UserResource) FindOne(id string, r api2go.Request) (api2go.Responder, error) {
	utils.DebugLog.Printf("Received FindOne with ID %s.", id)
	res, err := c.UserStorage.GetOne(bson.ObjectIdHex(id))
	return &api2go.Response{Res: res}, err
}

// Create a new User.
// Possible status codes are:
// - 201 Created: Resource was created and needs to be returned
// - 202 Accepted: Processing is delayed, return nothing
// - 204 No Content: Resource created with a client generated ID, and no fields were modified by
//   the server
func (c UserResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	user, ok := obj.(models.User)
	if !ok {
		return &api2go.Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}
	id, _ := c.UserStorage.Insert(user)
	user.ID = id
	return &api2go.Response{Res: user, Code: http.StatusCreated}, nil
}

// Delete a User.
// Possible status codes are:
// - 200 OK: Deletion was a success, returns meta information, currently not implemented! Do not use this
// - 202 Accepted: Processing is delayed, return nothing
// - 204 No Content: Deletion was successful, return nothing
func (c UserResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.UserStorage.Delete(bson.ObjectIdHex(id))
	return &api2go.Response{Code: http.StatusOK}, err
}

// Update a User.
// Possible status codes are:
// - 200 OK: Update successful, however some field(s) were changed, returns updates source
// - 202 Accepted: Processing is delayed, return nothing
// - 204 No Content: Update was successful, no fields were changed by the server, return nothing
func (c UserResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	user, ok := obj.(models.User)
	if !ok {
		return &api2go.Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}
	err := c.UserStorage.Update(user)
	return &api2go.Response{Res: user, Code: http.StatusNoContent}, err
}
