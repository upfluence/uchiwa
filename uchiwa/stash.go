package uchiwa

import "github.com/upfluence/uchiwa/uchiwa/logger"

type stash struct {
	Dc      string                 `json:"dc"`
	Path    string                 `json:"path"`
	Content map[string]interface{} `json:"content"`
	Expire  int32                  `json:"expire,omitempty"`
}

// PostStash send a POST request to the /stashes endpoint in order to create a stash
func (u *Uchiwa) PostStash(data stash) error {
	api, err := getAPI(u.Datacenters, data.Dc)
	if err != nil {
		logger.Warning(err)
		return err
	}

	_, err = api.CreateStash(data)
	if err != nil {
		logger.Warning(err)
		return err
	}

	return nil
}

// DeleteStash send a DELETE request to the /stashes/*path* endpoint in order to delete a stash
func (u *Uchiwa) DeleteStash(data stash) error {
	api, err := getAPI(u.Datacenters, data.Dc)
	if err != nil {
		logger.Warning(err)
		return err
	}

	err = api.DeleteStash(data.Path)
	if err != nil {
		logger.Warning(err)
		return err
	}

	return nil
}
