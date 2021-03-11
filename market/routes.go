/*
 * App market
 *
 * API version: 0.0.1
 * Contact: support@peraMIC.io
 */

package market

import (
	"github.com/peramic/utils"
)

// Routes returns all of the api route for the Controller
var Routes = utils.Routes{

	utils.Route{
		Name:        "GetAvailableApps",
		Method:      "GET",
		Pattern:     "/rest/apps/{namespace}",
		HandlerFunc: GetAvailableApps,
	},
	utils.Route{
		Name:        "GetInstalledApps",
		Method:      "GET",
		Pattern:     "/rest/apps/{namespace}/installed",
		HandlerFunc: GetInstalledApps,
	},
	utils.Route{
		Name:        "GetUpdates",
		Method:      "GET",
		Pattern:     "/rest/apps/{namespace}/updates",
		HandlerFunc: GetUpdates,
	},
	utils.Route{
		Name:        "InstallApp",
		Method:      "POST",
		Pattern:     "/rest/apps/{namespace}",
		HandlerFunc: InstallApp,
	},
	utils.Route{
		Name:        "UpgradeApps",
		Method:      "PUT",
		Pattern:     "/rest/apps/{namespace}",
		HandlerFunc: UpgradeApps,
	},
	utils.Route{
		Name:        "UpdateApp",
		Method:      "PUT",
		Pattern:     "/rest/apps/{namespace}/{name}",
		HandlerFunc: UpdateApp,
	},
	utils.Route{
		Name:        "UninstallApp",
		Method:      "DELETE",
		Pattern:     "/rest/apps/{namespace}/{name}",
		HandlerFunc: UninstallApp,
	},
	utils.Route{
		Name:        "Installfile",
		Method:      "PUT",
		Pattern:     "/rest/apps/{namespace}/add/{filename}",
		HandlerFunc: installFile,
	},
}
