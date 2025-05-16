package services

import (
	"errors"
	"log"

	"go-gin-project/config"
	"go-gin-project/models"

	"gorm.io/gorm"
)

// UserService provides user-related services
type UserService struct {
	db *gorm.DB
}

// NewUserService creates a new UserService
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

// CreateInitialAdminUser creates the initial admin user if it doesn't exist.
// The admin username is fixed as "admin" and the password is read from config.
func (s *UserService) CreateInitialAdminUser() error {
	adminUsername := "admin"
	var existingUser models.User
	if err := s.db.Where("username = ?", adminUsername).First(&existingUser).Error; err == nil {
		log.Printf("Admin user '%s' already exists.", adminUsername)
		return nil // Admin user already exists
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("Error checking for admin user: %v", err)
		return err // Other database error
	}

	// Admin user does not exist, create it
	adminPassword := config.AppConfig.AdminPassword
	if adminPassword == "" {
		log.Println("Warning: ADMIN_PASSWORD is not set in config. Using a default insecure password for admin.")
		adminPassword = "admin123" // Fallback, though InitConfig should set a default
	}

	adminUser := models.User{
		Username: adminUsername,
		Password: adminPassword, // Password will be hashed before saving
	}

	if err := adminUser.HashPassword(); err != nil {
		log.Printf("Error hashing admin password: %v", err)
		return err
	}

	if err := s.db.Create(&adminUser).Error; err != nil {
		log.Printf("Error creating admin user '%s': %v", adminUsername, err)
		return err
	}

	log.Printf("Admin user '%s' created successfully.", adminUsername)
	return nil
}

// GetUserByUsername retrieves a user by their username.
func (s *UserService) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

