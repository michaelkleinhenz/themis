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

// WorkItemResource for api2go routes.
type WorkItemResource struct {
	WorkItemStorage *database.WorkItemStorage
}

func (c WorkItemResource) getFilterFromRequest(r api2go.Request) (bson.M, *utils.NestedEntityError) {
	var filter bson.M
	// Getting reference context
	sourceContext, sourceContextID, thisContext := utils.ParseContext(r)
	switch sourceContext {
		case models.WorkItemName:
			if thisContext == "children" {
				// we need the full children IDs here, so only query for the IDs, not the full documents!
				// (Otherwise, we might kill the service as child lists can get big).
				// Note that the empty selector automatically returns only the _id field.
				childrens, err := c.WorkItemStorage.GetAll(bson.M{"parent_workitem_id": bson.ObjectIdHex(sourceContextID)})
				if (err != nil) {
					return nil, &utils.NestedEntityError { InnerError: err, Code: 0 }
				}
				childIDs := make([]bson.ObjectId, 0)
				for _, children := range childrens {
					childIDs = append(childIDs, children.ID)
				}
				filter = bson.M{"_id": bson.M{"$in": childIDs}}
			}
		default:
			// build standard filter expression
			filter = utils.BuildDbFilterFromRequest(r)
	}
	return filter, nil
}

// FindAll WorkItems.
func (c WorkItemResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	utils.DebugLog.Printf("Received FindAll with params %s.", r.QueryParams)

	// build filter expression
	filter, nestedErr := c.getFilterFromRequest(r)
	if nestedErr != nil {
		return &api2go.Response{}, nestedErr.InnerError
	}

	// retrieve entities
	workItems, err := c.WorkItemStorage.GetAll(filter)	
	if err != nil {
		return &api2go.Response{}, err
	}
	return &api2go.Response{Res: workItems}, nil
}

// PaginatedFindAll can be used to load users in chunks
func (c WorkItemResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {

	// parse out offset and limit
	queryOffset, queryLimit, err := utils.ParsePaging(r)
	if err!=nil {
		return 0, &api2go.Response{}, err
	}

	// build filter expression
	filter, nestedErr := c.getFilterFromRequest(r)
	if nestedErr != nil {
		return 0, &api2go.Response{}, nestedErr.InnerError
	}

	// get the paged data from storage
	result, err := c.WorkItemStorage.GetAllPaged(filter, queryOffset, queryLimit)
	if err!=nil {
		return 0, &api2go.Response{}, err
	}

	// get total count for paging
	allCount, err := c.WorkItemStorage.GetAllCount(filter)
	if err!=nil {
		return 0, &api2go.Response{}, err
	}

	// return everything
	return uint(allCount), &api2go.Response{Res: result}, nil
}

// FindOne returns an object by its ID.
// Possible success status code 200
func (c WorkItemResource) FindOne(id string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.WorkItemStorage.GetOne(bson.ObjectIdHex(id))
	return &api2go.Response{Res: res}, err
}

// Create a new object. Newly created object/struct must be in Responder.
// Possible status codes are:
// - 201 Created: Resource was created and needs to be returned
// - 202 Accepted: Processing is delayed, return nothing
// - 204 No Content: Resource created with a client generated ID, and no fields were modified by
//   the server
func (c WorkItemResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	workItem, ok := obj.(models.WorkItem)
	if !ok {
		return &api2go.Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}
	id, _ := c.WorkItemStorage.Insert(workItem)
	workItem.ID = id
	return &api2go.Response{Res: workItem, Code: http.StatusCreated}, nil
}

// Delete an object
// Possible status codes are:
// - 200 OK: Deletion was a success, returns meta information, currently not implemented! Do not use this
// - 202 Accepted: Processing is delayed, return nothing
// - 204 No Content: Deletion was successful, return nothing
func (c WorkItemResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.WorkItemStorage.Delete(bson.ObjectIdHex(id))
	return &api2go.Response{Code: http.StatusOK}, err
}

// Update an object
// Possible status codes are:
// - 200 OK: Update successful, however some field(s) were changed, returns updates source
// - 202 Accepted: Processing is delayed, return nothing
// - 204 No Content: Update was successful, no fields were changed by the server, return nothing
func (c WorkItemResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	workItem, ok := obj.(models.WorkItem)
	if !ok {
		return &api2go.Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}
	err := c.WorkItemStorage.Update(workItem)
	return &api2go.Response{Res: workItem, Code: http.StatusNoContent}, err
}
