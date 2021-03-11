/*
 * App market
 *
 * API version: 0.0.1
 * Contact: support@peraMIC.io
 */

package market

import (
	"github.com/peramic/App.Containerd/go/containers"
)

// App - App information
type App struct {
	containers.Container

	// App description
	Description string `json:"description,omitempty"`

	// Name of the app icon
	Icon string `json:"icon,omitempty"`
}
