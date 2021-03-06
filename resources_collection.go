package corbel

import (
	"fmt"
	"net/http"
)

// AddToCollection add the required struct formated as json to the desired collection
// resource must have exported variables and optionally its representation as JSON.
func (r *ResourcesService) AddToCollection(collectionName string, resource interface{}) (string, error) {
	var (
		req *http.Request
		err error
	)

	req, err = r.client.NewRequest("POST", "resources", fmt.Sprintf("/v1.0/resource/%s", collectionName), resource)
	return returnErrorHTTPSimple(r.client, req, err, 201)
}

// UpdateInCollection updates the required struct formated as json to the desired collection
// resource must have exported variables and optionally its representation as JSON.
func (r *ResourcesService) UpdateInCollection(collectionName, id string, resource interface{}) error {
	var (
		req *http.Request
		err error
	)

	req, err = r.client.NewRequest("PUT", "resources", fmt.Sprintf("/v1.0/resource/%s/%s", collectionName, id), resource)
	_, err = returnErrorHTTPSimple(r.client, req, err, 204)
	return err
}

// SearchCollection gets the desired objects in base of a search query
func (r *ResourcesService) SearchCollection(collectionName string) *Search {
	return NewSearch(r.client, "resources", fmt.Sprintf("/v1.0/resource/%s", collectionName))
}

// GetFromCollection gets the desired object from the collection by id
func (r *ResourcesService) GetFromCollection(collectionName, id string, resource interface{}) error {
	var (
		req *http.Request
		err error
	)

	req, err = r.client.NewRequest("GET", "resources", fmt.Sprintf("/v1.0/resource/%s/%s", collectionName, id), nil)
	_, err = returnErrorHTTPInterface(r.client, req, err, resource, 200)
	return err
}

// GetFromRelationDefinition gets the desired object from the collection by id
func (r *ResourcesService) GetFromRelationDefinition(id string, resource interface{}) error {
	var (
		req *http.Request
		err error
	)

	req, err = r.client.NewRequest("GET", "resources", fmt.Sprintf("/v1.0/resource/%s", id), nil)
	_, err = returnErrorHTTPInterface(r.client, req, err, resource, 200)
	return err
}

// DeleteFromCollection deletes the desired resource from the platform by id
func (r *ResourcesService) DeleteFromCollection(collectionName, id string) error {

	var (
		req *http.Request
		err error
	)

	req, err = r.client.NewRequest("DELETE", "resources", fmt.Sprintf("/v1.0/resource/%s/%s", collectionName, id), nil)
	_, err = returnErrorHTTPSimple(r.client, req, err, 204)
	return err
}

// UpdateResourceACL updates the acl of the associated resource. ACL entries will be added if they were not previously
// there or modified otherwise. Any entries previously added but not passed will be removed.
func (r *ResourcesService) UpdateResourceACL(collectionName, id string, acl interface{}) error {
	req, err := r.client.NewRequest("PUT", "resources", fmt.Sprintf("/v1.0/resource/%s/%s", collectionName, id), acl)
	req.Header.Set("Accept", "application/corbel.acl+json")
	_, err = returnErrorHTTPSimple(r.client, req, err, 204)
	return err
}
