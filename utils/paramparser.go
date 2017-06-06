package utils

import (
	"strconv"

	"github.com/manyminds/api2go"
)

func GetPathID(r api2go.Request, typeStr string) (string, bool) {
	spacesID, ok := r.QueryParams[typeStr]
	if ok {
		spaceID := spacesID[0]
		return spaceID, true
	}
	return "", false
}

func GetSpaceID(r api2go.Request) (string, bool) {
	return GetPathID(r, "spacesID")
}

func ParsePaging(r api2go.Request) (int, int, error) {
	var number, size, offset, limit	string

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