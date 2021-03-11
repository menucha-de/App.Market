/*
 * App market
 *
 * API version: 0.0.1
 * Contact: support@peraMIC.io
 */

package market

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// GetAvailableApps - Returns a list of available apps
func GetAvailableApps(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ns := params["namespace"]
	if err := checkNamespace(ns); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	apps, err := GetAvailableApps2(ns)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	installable := []App{}
	for _, v := range apps {
		installable = append(installable, v)
	}

	err = json.NewEncoder(w).Encode(installable)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetInstalledApps - Returns a list of installed apps
func GetInstalledApps(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ns := params["namespace"]
	if err := checkNamespace(ns); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	installed, err := getInstalledApps(ns)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Cache-Control", "no-cache")
	err = json.NewEncoder(w).Encode(installed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// GetUpdates - Get apps with available updates
func GetUpdates(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ns := params["namespace"]
	if err := checkNamespace(ns); err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	updates, err := getUpdates(ns, r.Header.Get("username"), r.Header.Get("password"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Cache-Control", "no-cache")
	err = json.NewEncoder(w).Encode(updates)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// InstallApp - Installs the specified app
func InstallApp(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ns := params["namespace"]

	if err := checkNamespace(ns); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var app App
	if err := json.NewDecoder(r.Body).Decode(&app); err != nil {
		log.Error(err.Error())
		http.Error(w, "Update failed. Please see log for details.", http.StatusBadRequest)
		return
	}

	if ns == "system" && app.Name != "vpn" {
		http.Error(w, "Installation of system apps not allowed", http.StatusForbidden)
		return
	}

	err := installApp(ns, app)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// UpgradeApps - Upgrade the specified list of apps
func UpgradeApps(w http.ResponseWriter, r *http.Request) {
	var err error
	params := mux.Vars(r)
	ns := params["namespace"]
	if err := checkNamespace(ns); err != nil {
		http.Error(w, "Upgrade failed. Failed to check namespace. Please see log for details.", http.StatusInternalServerError)
		return
	}

	var apps []App
	if ns != "system" {
		if err := json.NewDecoder(r.Body).Decode(&apps); err != nil {
			log.Error(err.Error())
			http.Error(w, "Upgrade failed. Please see log for details.", http.StatusBadRequest)
			return
		}
	}

	// err := upgradeApps(ns, apps)
	if err != nil {
		http.Error(w, "Upgrade failed. Please see log for details.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// UninstallApp - Uninstalls the specified app
func UninstallApp(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	ns := params["namespace"]
	if err := checkNamespace(ns); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err := UninstallApp2(ns, name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// UpdateApp - Updates the specified app
func UpdateApp(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	ns := params["namespace"]
	if err := checkNamespace(ns); err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	app := App{}
	if err := json.NewDecoder(r.Body).Decode(&app); err != nil {
		log.Error(err.Error())
		http.Error(w, "Update failed. Invalid app "+name+". Please see log for details.", http.StatusBadRequest)
		return
	}

	err := updateApp(ns, app)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func installFile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ns := params["namespace"]

	if err := checkNamespace(ns); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if ns == "system" {
		http.Error(w, "Installation of system apps not allowed", http.StatusForbidden)
		return
	}

	filename := params["filename"]
	filePath := "/opt/" + filename
	f, err := os.Create(filePath)
	if err != nil {
		log.WithError(err).Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()
	_, err = io.Copy(f, r.Body)
	if err != nil {
		log.WithError(err).Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = installFromfile(ns, filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
func checkNamespace(ns string) error {
	var namespaces [2]string
	err := client.Call("Service.GetNamespaces", 0, &namespaces)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	for _, v := range namespaces {
		if v == ns {
			return nil
		}
	}

	err = errors.New("Unknown namespace \"" + ns + "\"")
	log.Error(err.Error())
	return err
}
