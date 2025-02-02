/*
 * Simple Inventory API
 *
 * This is a simple API
 *
 * API version: 1.0.0
 * Contact: you@your-company.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type InventoryItem struct {
	Id string `json:"id"`
	Name string `json:"name"`
	ReleaseDate time.Time `json:"releaseDate"`
	Manufacturer *Manufacturer `json:"manufacturer"`
}
