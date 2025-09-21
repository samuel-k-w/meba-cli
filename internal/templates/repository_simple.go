package templates

import (
	"fmt"
	"strings"
)

func ModuleRepositoryGoSimple(name string) string {
	titleName := strings.Title(name)
	return fmt.Sprintf(`package %s

import (
	"gorm.io/gorm"
)

// Repository handles data access for %s
type Repository struct {
	db *gorm.DB
}

// New%sRepository creates a new repository instance
func New%sRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// GetAll retrieves all %s with pagination
func (r *Repository) GetAll(page, pageSize int) ([]*%s, int64, error) {
	var items []*%s
	var total int64

	offset := (page - 1) * pageSize

	if err := r.db.Model(&%s{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Offset(offset).Limit(pageSize).Find(&items).Error; err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

// GetByID retrieves a %s by ID
func (r *Repository) GetByID(id uint) (*%s, error) {
	var item %s
	if err := r.db.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

// Create creates a new %s
func (r *Repository) Create(item *%s) error {
	return r.db.Create(item).Error
}

// Update updates an existing %s
func (r *Repository) Update(item *%s) error {
	return r.db.Save(item).Error
}

// Delete deletes a %s by ID
func (r *Repository) Delete(id uint) error {
	return r.db.Delete(&%s{}, id).Error
}
`, name, name, titleName, titleName, name, titleName, titleName, titleName, name, titleName, titleName, name, titleName, name, titleName, name, titleName)
}