package templates

import (
	"fmt"
	"strings"
)

func ModuleServiceGoSimple(name string) string {
	titleName := strings.Title(name)
	return fmt.Sprintf(`package %s

import (
	"fmt"
)

// Service handles business logic for %s
type Service struct {
	repo *Repository
}

// New%sService creates a new service instance
func New%sService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// GetAll retrieves all %s with pagination
func (s *Service) GetAll(page, pageSize int) (*PaginationResponse, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	items, total, err := s.repo.GetAll(page, pageSize)
	if err != nil {
		return nil, fmt.Errorf("failed to get %s: %%w", err)
	}

	totalPages := int(total) / pageSize
	if int(total)%%pageSize != 0 {
		totalPages++
	}

	return &PaginationResponse{
		Page:       page,
		PageSize:   pageSize,
		Total:      total,
		TotalPages: totalPages,
		Data:       items,
	}, nil
}

// GetByID retrieves a %s by ID
func (s *Service) GetByID(id uint) (*%s, error) {
	item, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get %s: %%w", err)
	}
	return item, nil
}

// Create creates a new %s
func (s *Service) Create(req *Create%sRequest) (*%s, error) {
	item := &%s{
		// Map request fields to entity
		// Name: req.Name,
	}

	if err := s.repo.Create(item); err != nil {
		return nil, fmt.Errorf("failed to create %s: %%w", err)
	}

	return item, nil
}

// Update updates an existing %s
func (s *Service) Update(id uint, req *Update%sRequest) (*%s, error) {
	item, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get %s: %%w", err)
	}

	// Update fields from request
	// item.Name = req.Name

	if err := s.repo.Update(item); err != nil {
		return nil, fmt.Errorf("failed to update %s: %%w", err)
	}

	return item, nil
}

// Delete deletes a %s by ID
func (s *Service) Delete(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		return fmt.Errorf("failed to delete %s: %%w", err)
	}
	return nil
}
`, name, name, titleName, titleName, name, name, name, titleName, name, name, titleName, titleName, titleName, name, name, titleName, titleName, name, name, name, name)
}