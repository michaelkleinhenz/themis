package resources

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/manyminds/api2go"

	"themis/models"
)

func NewRequest(queryParams map[string][]string, pagination map[string]string) api2go.Request {
	return api2go.Request {
		PlainRequest: nil,
		QueryParams: queryParams,
		Pagination: pagination,
		Header: nil,
		Context: nil,
	}
}

// WorkItemStorageMock mock for WorkItemStorage
type WorkItemStorageMock struct {}
// Insert mock.
func (mock WorkItemStorageMock) Insert(workItem models.WorkItem) (bson.ObjectId, error) {
	return bson.NewObjectId(), nil
}
// Update mock.
func (mock WorkItemStorageMock) Update(workItem models.WorkItem) error {
	return nil
}
// Delete mock.
func (mock WorkItemStorageMock) Delete(id bson.ObjectId) error {
	return nil
}
// GetOne mock.
func (mock WorkItemStorageMock) GetOne(id bson.ObjectId) (models.WorkItem, error) {
	return *models.NewWorkItem(), nil
}
// GetAll mock.
func (mock WorkItemStorageMock) GetAll(queryExpression interface{}) ([]models.WorkItem, error) {
	var entities []models.WorkItem
	entities = append(entities, *models.NewWorkItem())
	entities = append(entities, *models.NewWorkItem())
	return entities, nil
}
// GetAllChildIDs mock.
func (mock WorkItemStorageMock) GetAllChildIDs(id bson.ObjectId) ([]bson.ObjectId, error) {
	var entities []bson.ObjectId
	entities = append(entities, bson.NewObjectId())
	entities = append(entities, bson.NewObjectId())
	return entities, nil
}
// GetAllPaged mock.
func (mock WorkItemStorageMock) GetAllPaged(queryExpression interface{}, offset int, limit int) ([]models.WorkItem, error) {
	var entities []models.WorkItem
	entities = append(entities, *models.NewWorkItem())
	entities = append(entities, *models.NewWorkItem())
	return entities, nil
}
// GetAllCount mock.
func (mock WorkItemStorageMock) GetAllCount(queryExpression interface{}) (int, error) {
	return 42, nil
}
// NewDisplayID mock.
func (mock WorkItemStorageMock) NewDisplayID(spaceID string) (int, error) {
	return 1, nil
}

// IterationStorageMock mock for IterationStorage
type IterationStorageMock struct {}
// IsRoot mock.
func (mock IterationStorageMock) IsRoot(id bson.ObjectId) (bool, error) {
	return false, nil
}
// Insert mock.
func (mock IterationStorageMock) Insert(iteration models.Iteration) (bson.ObjectId, error) {
	return bson.NewObjectId(), nil
}
// Update mock.
func (mock IterationStorageMock) Update(iteration models.Iteration) error {
	return nil
}
// Delete mock.
func (mock IterationStorageMock) Delete(id bson.ObjectId) error {
	return nil
}
// GetOne mock.
func (mock IterationStorageMock) GetOne(id bson.ObjectId) (models.Iteration, error) {
	return *models.NewIteration(), nil
}
// GetAll mock.
func (mock IterationStorageMock) GetAll(queryExpression interface{}) ([]models.Iteration, error) {
	var entities []models.Iteration
	entities = append(entities, *models.NewIteration())
	entities = append(entities, *models.NewIteration())
	return entities, nil
}
// GetAllPaged mock.
func (mock IterationStorageMock) GetAllPaged(queryExpression interface{}, offset int, limit int) ([]models.Iteration, error) {
	var entities []models.Iteration
	entities = append(entities, *models.NewIteration())
	entities = append(entities, *models.NewIteration())
	return entities, nil
}
// GetAllCount mock.
func (mock IterationStorageMock) GetAllCount(queryExpression interface{}) (int, error) {
	return 42, nil
}
// NewDisplayID mock.
func (mock IterationStorageMock) NewDisplayID(spaceID string) (int, error) {
	return 1, nil
}
// GetParentPath mock.
func (mock IterationStorageMock) GetParentPath(id bson.ObjectId) (string, string, error) {
	return "/mock1_ID", "/mock1_NAME", nil
}
