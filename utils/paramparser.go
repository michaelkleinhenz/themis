package utils

import (
	"strconv"
	"strings"

	"github.com/manyminds/api2go"
	"gopkg.in/mgo.v2/bson"
)

// ReplaceDotsToDollarsInAttributes replaces "." with "$" because MongoDB sucks.
func ReplaceDotsToDollarsInAttributes(attributes *map[string]string) {
	for k, v := range *attributes { 
		newK := strings.Replace(k, ".", "$", -1)
		(*attributes)[newK] = v
		delete(*attributes, k)
	}
}

// ReplaceDollarsToDotsInAttributes replaces "$" with "." because MongoDB sucks.
func ReplaceDollarsToDotsInAttributes(attributes *map[string]string) {
	for k, v := range *attributes { 
		newK := strings.Replace(k, "$", ".", -1)
		(*attributes)[newK] = v
		delete(*attributes, k)
	}
}

// ParseContext parses a possible subquery context, returning the parsed context keys.
// Example:
//   http://localhost:8080/api/workitemtypes/abc123/space has a QueryParams map:
//   map[workitemtypesID:[abc123] workitemtypesName:[space]]
// which will return
//   workitemtypes, abc123, space
func ParseContext(r api2go.Request) (sourceContext string, sourceContextID string, thisContext string) {
	for key, value := range r.QueryParams {
		if strings.HasSuffix(key, "ID") {
			sourceContextID = value[0]
			sourceContext = strings.Replace(key, "ID", "", -1)
		}
		if strings.HasSuffix(key, "Name") {
			thisContext = value[0]
		}
	}
	return sourceContext, sourceContextID, thisContext
}

// BuildDbFilterFromRequest builds the filter structure from the request.
func BuildDbFilterFromRequest(r api2go.Request) bson.M {
	var filter bson.M
	spaceID, ok := GetPathParam(r, "spacesID")
	if ok {
		filter = bson.M{"space_id": bson.ObjectIdHex(spaceID)}
	}
	workitemID, ok := GetPathParam(r, "workitemID")
	if ok {
		if filter == nil {
			filter = bson.M{}
		}
		filter["workitem_id"] = workitemID
	}
	workitemsID, ok := GetPathParam(r, "workitemsID")
	if ok {
		if filter == nil {
			filter = bson.M{}
		}
		filter["workitem_id"] = workitemsID
	}
	iterationsID, ok := GetPathParam(r, "iterationsID")
	if ok {
		if filter == nil {
			filter = bson.M{}
		}
		filter["iteration_id"] = iterationsID
	}
	areasID, ok := GetPathParam(r, "areasID")
	if ok {
		if filter == nil {
			filter = bson.M{}
		}
		filter["area_id"] = areasID
	}
	return filter
}

// ParsePaging parses the paging parameters of a request and returns them in a normalized version (start, limit).
func ParsePaging(r api2go.Request) (int, int, error) {
	var number, size, offset, limit string

	numberQuery, ok := r.QueryParams["page[number]"]
	if ok {
		number = numberQuery[0]
	}
	sizeQuery, ok := r.QueryParams["page[size]"]
	if ok {
		size = sizeQuery[0]
	}
	offsetQuery, ok := r.QueryParams["page[offset]"]
	if ok {
		offset = offsetQuery[0]
	}
	limitQuery, ok := r.QueryParams["page[limit]"]
	if ok {
		limit = limitQuery[0]
	}

	var resultStart int
	var resultLimit int
	if size != "" {
		// PAGE NUMBER AND SIZE MODE
		sizeI, err := strconv.Atoi(size)
		if err != nil {
			return -1, -1, err
		}
		numberI, err := strconv.Atoi(number)
		if err != nil {
			return -1, -1, err
		}
		resultStart = sizeI * (numberI - 1)
		resultLimit = sizeI
	} else {
		// PAGE OFFSET AND LIMIT MODE
		limitI, err := strconv.Atoi(limit)
		if err != nil {
			return -1, -1, err
		}
		offsetI, err := strconv.Atoi(offset)
		if err != nil {
			return -1, -1, err
		}
		resultStart = offsetI
		resultLimit = limitI
	}

	return resultStart, resultLimit, nil
}

// GetPathParam parses out a param from a QueryParam.
func GetPathParam(r api2go.Request, key string) (string, bool) {
	values, ok := r.QueryParams[key]
	if ok {
		value := values[0]
		return value, true
	}
	return "", false
}
