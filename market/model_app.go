/*
 * App market
 *
 * API version: 1.0.0
 * Contact: info@menucha.de
 */

package market

import (
	"github.com/menucha-de/art/art/containers"
)

// App - App information
type App struct {
	containers.Container

	// App description
	Description string `json:"description,omitempty"`

	// Name of the app icon
	Icon string `json:"icon,omitempty"`
}
