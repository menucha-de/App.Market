/*
 * App market
 *
 * API version: 0.0.1
 * Contact: support@peraMIC.io
 */

package market

import (
	"net/http"
	"time"
)

// Logger ...
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Tracef(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
