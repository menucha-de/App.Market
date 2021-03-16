/*
 * App market
 *
 * API version: 1.0.0
 * Contact: info@menucha.de
 */

package market

import (
	"encoding/json"
	"errors"
	"net/rpc"
	"os"
	"strings"

	"github.com/menucha-de/art/art"
	"github.com/menucha-de/art/art/containers"
	"github.com/menucha-de/logging"
	"github.com/menucha-de/utils"
)

var log *logging.Logger = logging.GetLogger("market")

// Client ...
type Client struct {
	RPCClient *rpc.Client
}

var client utils.Client

const artRPCAdress = "art:8080"

var knownApps map[string](map[string]App)

var installed map[string]App

func init() {
	var file string
	var err error
	client.ServerAdress = artRPCAdress
	if err != nil {
		log.Error("Failed to establish connection ", err)
	}

	file = "apps.json"

	if err = readConfig(file); err != nil {
		log.Error("Failed to load config ", err)
	}

	installed = make(map[string]App)
}

// GetAvailableApps2 - Returns a list of available apps
func GetAvailableApps2(namespace string) (map[string]App, error) {
	if err := checkNamespace(namespace); err != nil {
		log.Error(err.Error())
	}

	var copy = make(map[string]App)
	for _, nsApp := range knownApps[namespace] {
		copy[nsApp.Name] = nsApp
	}

	installed, err := getInstalledApps(namespace)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	for _, n := range installed {
		delete(copy, n.Container.Name)
	}
	return copy, nil
}

func getInstalledApps(namespace string) ([]App, error) {
	apps := []App{}
	var cons []containers.Container
	err := client.Call("Service.GetContainers", namespace, &cons)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	a := knownApps[namespace]
	for _, c := range cons {
		app := a[c.Name]
		if app.Container.Name != "" {
			app.Id = c.Id
			if c.State != "" {
				app.State = c.State
			}
		} else {
			//c.Label = "Custom"
			app = App{
				Container:   c,
				Description: "Custom App",
				Icon:        "launch",
			}
		}
		apps = append(apps, app)
	}

	return apps, nil
}

func installApp(namespace string, app App) error {
	available, err := GetAvailableApps2(namespace)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	_, ok := available[app.Container.Name]

	if !ok {
		err := errors.New("Unknown container " + app.Container.Name)
		log.Error(err.Error())
		return err
	}
	request := art.ContainersRequest{
		Ns:         namespace,
		Containers: []containers.Container{app.Container},
	}

	var response int
	err = client.Call("Service.AddContainer", request, &response)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func installFromfile(namespace string, file string) error {
	c := containers.Container{Name: file, Label: file, State: "STARTED"}
	app := App{Container: c}

	request := art.ContainersRequest{
		Ns:         namespace,
		Containers: []containers.Container{app.Container},
	}

	var response int
	err := client.Call("Service.AddContainerFromFile", request, &response)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func updateApp(namespace string, app App) error {
	c := app.Container
	if strings.HasPrefix(app.State, "UPGRADING") {
		if app.State == "UPGRADINGSTARTED" {
			app.State = "UPGRADING"
			c.State = "STARTED"
		} else {
			app.State = "UPGRADING"
			c.State = "STOPPED"
		}
	} else if strings.HasPrefix(app.State, "RESETTING") {
		if app.State == "RESETTINGSTARTED" {
			app.State = "RESETTING"
			c.State = "STARTED"
		} else {
			app.State = "RESETTING"
			c.State = "STOPPED"
		}
	}
	cons := []containers.Container{}
	cons = append(cons, c)

	request := art.ContainersRequest{
		Ns:         namespace,
		Containers: cons,
	}

	var response int

	switch app.State {
	case "STARTING":
		err := client.Call("Service.StartContainer", request, &response)
		if err != nil {
			log.Error(err.Error())
			return err
		}
	case "STOPPING":
		err := client.Call("Service.StopContainer", request, &response)
		if err != nil {
			log.Error(err.Error())
			return err
		}
	case "UPGRADING":
		go client.Call("Service.Upgrade", request, &response)
	case "RESETTING":
		err := client.Call("Service.ResetContainer", request, &response)
		if err != nil {
			log.Error(err.Error())
			return err
		}
	}

	return nil
}

// UninstallApp2 - Uninstalls the specified app
func UninstallApp2(namespace string, name string) error {
	request := art.ContainersRequest{
		Ns: namespace,
		ID: name}
	var response int
	err := client.Call("Service.DeleteContainer", request, &response)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func getUpdates(namespace string, username string, password string) ([]App, error) {
	apps, err := getInstalledApps(namespace)
	if err != nil {
		return nil, err
	}

	var updates []App

	for _, app := range apps {
		app.User = username
		app.Passwd = password

		cons := []containers.Container{}
		cons = append(cons, app.Container)
		request := art.ContainersRequest{
			Ns:         namespace,
			Containers: cons}

		log.Debugf("Calling UpdateAvailable from art with request %+v\n", request)
		var response bool
		err := client.Call("Service.UpdateAvailable", request, &response)
		if err != nil {
			log.Error(err.Error())
			break
		}
		log.Debugf("Returning %t", response)
		if response {
			updates = append(updates, app)
		}
	}

	if err != nil {
		return nil, err
	}

	return updates, nil
}

func readConfig(file string) error {
	f, err := os.Open(file)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	defer f.Close()

	knownApps = make(map[string](map[string]App))
	if err := json.NewDecoder(f).Decode(&knownApps); err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}
