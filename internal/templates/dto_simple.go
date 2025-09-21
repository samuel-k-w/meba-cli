package templates

import (
	"fmt"
	"strings"
)

func ModuleDtoGoSimple(name string) string {
	titleName := strings.Title(name)
	return fmt.Sprintf(`package %s

import "time"

// Create%sRequest represents the request to create a %s
type Create%sRequest struct {
	// Add your create request fields here
	// Name        string `+"`json:\"name\" validate:\"required,min=2,max=100\"`"+`
	// Description string `+"`json:\"description\" validate:\"max=500\"`"+`
}

// Update%sRequest represents the request to update a %s
type Update%sRequest struct {
	// Add your update request fields here
	// Name        string `+"`json:\"name\" validate:\"required,min=2,max=100\"`"+`
	// Description string `+"`json:\"description\" validate:\"max=500\"`"+`
}

// %sResponse represents the response for %s operations
type %sResponse struct {
	ID        uint      `+"`json:\"id\"`"+`
	CreatedAt time.Time `+"`json:\"created_at\"`"+`
	UpdatedAt time.Time `+"`json:\"updated_at\"`"+`
	
	// Add your response fields here
	// Name        string `+"`json:\"name\"`"+`
	// Description string `+"`json:\"description\"`"+`
	// Status      string `+"`json:\"status\"`"+`
}

// PaginationRequest provides common pagination parameters
type PaginationRequest struct {
	Page     int `+"`json:\"page\" form:\"page\" validate:\"min=1\"`"+`
	PageSize int `+"`json:\"page_size\" form:\"page_size\" validate:\"min=1,max=100\"`"+`
}

// PaginationResponse provides common pagination response
type PaginationResponse struct {
	Page       int         `+"`json:\"page\"`"+`
	PageSize   int         `+"`json:\"page_size\"`"+`
	Total      int64       `+"`json:\"total\"`"+`
	TotalPages int         `+"`json:\"total_pages\"`"+`
	Data       interface{} `+"`json:\"data\"`"+`
}
`, name, titleName, name, titleName, titleName, name, titleName, titleName, name, titleName)
}