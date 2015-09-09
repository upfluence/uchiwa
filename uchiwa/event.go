package uchiwa

import "github.com/upfluence/uchiwa/uchiwa/logger"
import "github.com/upfluence/uchiwa/uchiwa/daemon"

// ResolveEvent send a POST request to the /resolve endpoint in order to resolve an event
func (u *Uchiwa) ResolveEvent(data interface{}) error {
	api, m, err := daemon.FindDcFromInterface(data, u.Datacenters)
	_, err = api.ResolveEvent(m["payload"])
	if err != nil {
		logger.Warning(err)
		return err
	}
	return nil
}
