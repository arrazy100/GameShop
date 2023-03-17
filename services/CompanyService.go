package services

import (
	"main/models"

	"gorm.io/gorm"
)

type CompanyService struct {
	db *gorm.DB
}

func (repo *CompanyService) NewCompanyService(db *gorm.DB) CompanyService {
	return CompanyService{
		db: db,
	}
}

func (repo *CompanyService) FindAll() ([]models.CompanyDetails, error) {
	var companies []models.CompanyDetails

	// get all companies
	err := repo.db.Model(&models.Company{}).Find(&companies).Error

	if err != nil {
		return nil, err
	}

	return companies, nil
}

func (repo *CompanyService) Create(company models.CompanyInput) (*models.CompanyDetails, error) {
	var newCompany models.Company

	// create new company
	newCompany.Name = company.Name
	newCompany.State = company.State
	newCompany.City = company.City
	newCompany.Country = company.Country

	err := repo.db.Model(&models.Company{}).Create(&newCompany).Error

	if err != nil {
		return &models.CompanyDetails{}, err
	}

	return &models.CompanyDetails{
		CompanyID: newCompany.CompanyID,
		Name:      newCompany.Name,
		State:     newCompany.State,
		City:      newCompany.City,
		Country:   newCompany.Country,
		CreatedAt: newCompany.CreatedAt,
		UpdatedAt: newCompany.UpdatedAt,
		IsActive:  newCompany.IsActive,
	}, nil
}
