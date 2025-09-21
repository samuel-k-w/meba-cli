package templates

import (
	"fmt"
	"strings"
)

func ModuleGo(name string) string {
	titleName := strings.Title(name)
	return fmt.Sprintf(`package %s

import (
	"github.com/google/wire"
)

// Module is the wire set for the %s module
var Module = wire.NewSet(
	New%sService,
	New%sRepository,
	New%sHandlers,
)
`, name, name, titleName, titleName, titleName)
}

func ModuleHandlersGo(name string) string {
	titleName := strings.Title(name)
	return fmt.Sprintf(`package %s

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handlers handles HTTP requests for %s
type Handlers struct {
	service *Service
}

// New%sHandlers creates a new handlers instance
func New%sHandlers(service *Service) *Handlers {
	return &Handlers{
		service: service,
	}
}

// SetupRoutes configures routes for %s module
func (h *Handlers) SetupRoutes(r *gin.RouterGroup) {
	%sGroup := r.Group("/%s")
	{
		%sGroup.GET("", h.GetAll)
		%sGroup.GET("/:id", h.GetByID)
		%sGroup.POST("", h.Create)
		%sGroup.PUT("/:id", h.Update)
		%sGroup.DELETE("/:id", h.Delete)
	}
}

// GetAll godoc
// @Summary Get all %s
// @Description Get all %s with pagination
// @Tags %s
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Success 200 {object} PaginationResponse
// @Router /%s [get]
func (h *Handlers) GetAll(c *gin.Context) {
	var req PaginationRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid query parameters",
			"error":   err.Error(),
		})
		return
	}

	result, err := h.service.GetAll(req.Page, req.PageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to get %s",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    result,
	})
}

// GetByID godoc
// @Summary Get %s by ID
// @Description Get a single %s by ID
// @Tags %s
// @Accept json
// @Produce json
// @Param id path int true "%s ID"
// @Success 200 {object} %s
// @Router /%s/{id} [get]
func (h *Handlers) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid ID",
		})
		return
	}

	result, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "%s not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    result,
	})
}

// Create godoc
// @Summary Create %s
// @Description Create a new %s
// @Tags %s
// @Accept json
// @Produce json
// @Param %s body Create%sRequest true "%s data"
// @Success 201 {object} %s
// @Router /%s [post]
func (h *Handlers) Create(c *gin.Context) {
	var req Create%sRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	result, err := h.service.Create(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to create %s",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    result,
	})
}

// Update godoc
// @Summary Update %s
// @Description Update an existing %s
// @Tags %s
// @Accept json
// @Produce json
// @Param id path int true "%s ID"
// @Param %s body Update%sRequest true "%s data"
// @Success 200 {object} %s
// @Router /%s/{id} [put]
func (h *Handlers) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid ID",
		})
		return
	}

	var req Update%sRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	result, err := h.service.Update(uint(id), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to update %s",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    result,
	})
}

// Delete godoc
// @Summary Delete %s
// @Description Delete a %s by ID
// @Tags %s
// @Accept json
// @Produce json
// @Param id path int true "%s ID"
// @Success 200 {object} BaseResponse
// @Router /%s/{id} [delete]
func (h *Handlers) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid ID",
		})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to delete %s",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "%s deleted successfully",
	})
}
`, name, name, titleName, titleName, name, name, name, name, name, name, name, name, name, name, name, name, name, titleName, name, titleName, name, name, titleName, titleName, name, titleName, name, titleName, name, name, titleName, name, titleName, titleName, name, titleName, name, titleName, name, name, titleName, name, titleName, name, titleName)
}

func ModuleServiceGo(name string) string {
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

	%ss, total, err := s.repo.GetAll(page, pageSize)
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
		Data:       %ss,
	}, nil
}

// GetByID retrieves a %s by ID
func (s *Service) GetByID(id uint) (*%s, error) {
	%s, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get %s: %%w", err)
	}
	return %s, nil
}

// Create creates a new %s
func (s *Service) Create(req *Create%sRequest) (*%s, error) {
	%s := &%s{
		// Map request fields to entity
		// Name: req.Name,
	}

	if err := s.repo.Create(%s); err != nil {
		return nil, fmt.Errorf("failed to create %s: %%w", err)
	}

	return %s, nil
}

// Update updates an existing %s
func (s *Service) Update(id uint, req *Update%sRequest) (*%s, error) {
	%s, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get %s: %%w", err)
	}

	// Update fields from request
	// %s.Name = req.Name

	if err := s.repo.Update(%s); err != nil {
		return nil, fmt.Errorf("failed to update %s: %%w", err)
	}

	return %s, nil
}

// Delete deletes a %s by ID
func (s *Service) Delete(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		return fmt.Errorf("failed to delete %s: %%w", err)
	}
	return nil
}
`, name, name, titleName, titleName, name, name, name, name, titleName, name, name, name, titleName, titleName, name, titleName, name, name, name, name, name, titleName, titleName, name, name, name, name, name, name, name, name)
}

func ModuleRepositoryGo(name string) string {
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
	var %ss []*%s
	var total int64

	offset := (page - 1) * pageSize

	if err := r.db.Model(&%s{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Offset(offset).Limit(pageSize).Find(&%ss).Error; err != nil {
		return nil, 0, err
	}

	return %ss, total, nil
}

// GetByID retrieves a %s by ID
func (r *Repository) GetByID(id uint) (*%s, error) {
	var %s %s
	if err := r.db.First(&%s, id).Error; err != nil {
		return nil, err
	}
	return &%s, nil
}

// Create creates a new %s
func (r *Repository) Create(%s *%s) error {
	return r.db.Create(%s).Error
}

// Update updates an existing %s
func (r *Repository) Update(%s *%s) error {
	return r.db.Save(%s).Error
}

// Delete deletes a %s by ID
func (r *Repository) Delete(id uint) error {
	return r.db.Delete(&%s{}, id).Error
}

// GetByField retrieves %s by a specific field (example)
func (r *Repository) GetByField(field, value string) ([]*%s, error) {
	var %ss []*%s
	if err := r.db.Where(field+" = ?", value).Find(&%ss).Error; err != nil {
		return nil, err
	}
	return %ss, nil
}
`, name, name, titleName, titleName, name, titleName, name, titleName, titleName, name, titleName, name, titleName, name, titleName, titleName, name, name, titleName, name, titleName, name, titleName, name, titleName, name, titleName, titleName, name, titleName, name)
}

func ModuleEntityGo(name string) string {
	titleName := strings.Title(name)
	return fmt.Sprintf(`package %s

import (
	"time"
	"gorm.io/gorm"
)

// %s represents the %s entity
type %s struct {
	ID        uint           ` + "`json:\"id\" gorm:\"primarykey\"`" + `
	CreatedAt time.Time      ` + "`json:\"created_at\"`" + `
	UpdatedAt time.Time      ` + "`json:\"updated_at\"`" + `
	DeletedAt gorm.DeletedAt ` + "`json:\"deleted_at\" gorm:\"index\"`" + `
	
	// Add your %s fields here
	// Name        string ` + "`json:\"name\" gorm:\"not null\" validate:\"required\"`" + `
	// Description string ` + "`json:\"description\"`" + `
	// Status      string ` + "`json:\"status\" gorm:\"default:active\"`" + `
}

// TableName returns the table name for %s
func (%s) TableName() string {
	return "%ss"
}

// BeforeCreate hook
func (%s *%s) BeforeCreate(tx *gorm.DB) error {
	// Add any pre-creation logic here
	return nil
}

// BeforeUpdate hook
func (%s *%s) BeforeUpdate(tx *gorm.DB) error {
	// Add any pre-update logic here
	return nil
}
`, name, titleName, name, titleName, name, titleName, name, name, titleName, name, titleName, name, titleName)
}

func ModuleDtoGo(name string) string {
	titleName := strings.Title(name)
	return fmt.Sprintf(`package %s

import "time"

// Create%sRequest represents the request to create a %s
type Create%sRequest struct {
	// Add your create request fields here
	// Name        string ` + "`json:\"name\" validate:\"required,min=2,max=100\"`" + `
	// Description string ` + "`json:\"description\" validate:\"max=500\"`" + `
}

// Update%sRequest represents the request to update a %s
type Update%sRequest struct {
	// Add your update request fields here
	// Name        string ` + "`json:\"name\" validate:\"required,min=2,max=100\"`" + `
	// Description string ` + "`json:\"description\" validate:\"max=500\"`" + `
}

// %sResponse represents the response for %s operations
type %sResponse struct {
	ID        uint      ` + "`json:\"id\"`" + `
	CreatedAt time.Time ` + "`json:\"created_at\"`" + `
	UpdatedAt time.Time ` + "`json:\"updated_at\"`" + `
	
	// Add your response fields here
	// Name        string ` + "`json:\"name\"`" + `
	// Description string ` + "`json:\"description\"`" + `
	// Status      string ` + "`json:\"status\"`" + `
}

// %sListResponse represents the response for listing %s
type %sListResponse struct {
	%ss []*%sResponse ` + "`json:\"%ss\"`" + `
	Total int64              ` + "`json:\"total\"`" + `
}

// PaginationRequest provides common pagination parameters
type PaginationRequest struct {
	Page     int ` + "`json:\"page\" form:\"page\" validate:\"min=1\"`" + `
	PageSize int ` + "`json:\"page_size\" form:\"page_size\" validate:\"min=1,max=100\"`" + `
}

// PaginationResponse provides common pagination response
type PaginationResponse struct {
	Page       int         ` + "`json:\"page\"`" + `
	PageSize   int         ` + "`json:\"page_size\"`" + `
	Total      int64       ` + "`json:\"total\"`" + `
	TotalPages int         ` + "`json:\"total_pages\"`" + `
	Data       interface{} ` + "`json:\"data\"`" + `
}

// BaseResponse provides common response structure
type BaseResponse struct {
	Success bool        ` + "`json:\"success\"`" + `
	Message string      ` + "`json:\"message\"`" + `
	Data    interface{} ` + "`json:\"data,omitempty\"`" + `
	Error   string      ` + "`json:\"error,omitempty\"`" + `
}
`, name, titleName, name, titleName, titleName, name, titleName, titleName, name, titleName, titleName, name, titleName, titleName, titleName, name)
}