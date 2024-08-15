package services

import (
	"github.com/amrimuf/hompimRent/models"
	"github.com/amrimuf/hompimRent/repositories"
)

type ListingService struct {
	Repo *repositories.ListingRepository
}

func NewListingService(repo *repositories.ListingRepository) *ListingService {
	return &ListingService{Repo: repo}
}

func (s *ListingService) CreateListing(listing *models.Listing) error {
	return s.Repo.Create(listing)
}

func (s *ListingService) GetAllListings() ([]models.Listing, error) {
	return s.Repo.GetAll()
}

func (s *ListingService) GetListingByID(id int) (*models.Listing, error) {
	return s.Repo.GetByID(id)
}

func (s *ListingService) UpdateListing(listing *models.Listing) error {
	return s.Repo.Update(listing)
}

func (s *ListingService) DeleteListing(id int) error {
	return s.Repo.Delete(id)
}
