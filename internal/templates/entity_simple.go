package templates

import (
	"fmt"
	"strings"
)

func ModuleEntityGoSimple(name string) string {
	titleName := strings.Title(name)
	return fmt.Sprintf(`package %s

import (
	"time"
	"gorm.io/gorm"
)

// %s represents the %s entity
type %s struct {
	ID        uint           `+"`json:\"id\" gorm:\"primarykey\"`"+`
	CreatedAt time.Time      `+"`json:\"created_at\"`"+`
	UpdatedAt time.Time      `+"`json:\"updated_at\"`"+`
	DeletedAt gorm.DeletedAt `+"`json:\"deleted_at\" gorm:\"index\"`"+`
	
	// Add your %s fields here
	// Name        string `+"`json:\"name\" gorm:\"not null\" validate:\"required\"`"+`
	// Description string `+"`json:\"description\"`"+`
	// Status      string `+"`json:\"status\" gorm:\"default:active\"`"+`
}

// TableName returns the table name for %s
func (e %s) TableName() string {
	return "%s"
}

// BeforeCreate hook
func (e *%s) BeforeCreate(tx *gorm.DB) error {
	// Add any pre-creation logic here
	return nil
}

// BeforeUpdate hook
func (e *%s) BeforeUpdate(tx *gorm.DB) error {
	// Add any pre-update logic here
	return nil
}
`, name, titleName, name, titleName, name, titleName, titleName, titleName, titleName, titleName)
}