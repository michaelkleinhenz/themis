package database

// StorageBackends stores all the storage backends.
type StorageBackends struct {
	Space					*SpaceStorage
	WorkItem 			*WorkItemStorage
	WorkItemType	*WorkItemTypeStorage
	Area					*AreaStorage
	Comment				*CommentStorage
	Iteration			*IterationStorage
	LinkCategory	*LinkCategoryStorage
	Link					*LinkStorage
	LinkType			*LinkTypeStorage
	User					*UserStorage
}