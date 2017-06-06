package client

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"

	"github.com/manyminds/api2go/jsonapi"

	"themis/models"
)

func doGETRequest(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	resultBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return resultBody, nil
}

func doPOSTRequest(url string, bodyType string, jsonStr string) ([]byte, error) {
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonStr)))
	// req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", bodyType)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
    return nil, err		
	}
	defer resp.Body.Close()
	resultBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return resultBody, nil
}

func doPATCH(url string, bodyType string, jsonStr string) ([]byte, error) {
	req, _ := http.NewRequest("PATCH", url, bytes.NewBuffer([]byte(jsonStr)))
	// req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", bodyType)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
    return nil, err		
	}
	defer resp.Body.Close()
	resultBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return resultBody, nil
}

func doDELETE(url string) ([]byte, error) {
	req, _ := http.NewRequest("DELETE", url, nil)
	// req.Header.Set("X-Custom-Header", "myvalue")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
    return nil, err		
	}
	defer resp.Body.Close()
	resultBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return resultBody, nil
}

// GetSpace retrieves a space from a given API URL for a given SpaceID.
func GetSpace(apiURL string, spaceID string) (*models.Space, map[string]interface{}, error) {
	result, err := doGETRequest(apiURL + "/api/spaces/" + spaceID)
	if err != nil {
		return nil, nil, err
	}
	var msgMapTemplate interface{}
	err = json.Unmarshal([]byte(result), &msgMapTemplate)
	if err != nil {
		return nil, nil, err
	}	
	msgMap := msgMapTemplate.(map[string]interface{})
	var space models.Space
	err = jsonapi.Unmarshal(result, &space)
	if err != nil {
		return nil, nil, err
	}	
	return &space, msgMap, nil
}

// GetWorkItem retrieves a space from a given API URL for a given SpaceID.
func GetWorkItem(apiURL string, spaceID string, workItemID string) (*models.WorkItem, map[string]interface{}, error) {
	result, err := doGETRequest(apiURL + "/api/spaces/" + spaceID + "/workitems/" + workItemID)
	if err != nil {
		return nil, nil, err
	}
	var msgMapTemplate interface{}
	err = json.Unmarshal([]byte(result), &msgMapTemplate)
	if err != nil {
		return nil, nil, err
	}	
	msgMap := msgMapTemplate.(map[string]interface{})
	var workItem models.WorkItem
	err = jsonapi.Unmarshal(result, &workItem)
	if err != nil {
		return nil, nil, err
	}	
	return &workItem, msgMap, nil
}

