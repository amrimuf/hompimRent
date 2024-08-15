package repositories

import (
	"context"
	"log"
	"time"

	"github.com/amrimuf/hompimRent/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ListingRepository struct {
	DB *pgxpool.Pool
}

func NewListingRepository(db *pgxpool.Pool) *ListingRepository {
	return &ListingRepository{DB: db}
}

func (r *ListingRepository) Create(listing *models.Listing) error {
	_, err := r.DB.Exec(context.Background(), "INSERT INTO listings (title, description, price, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)",
		listing.Title, listing.Description, listing.Price, time.Now(), time.Now())
	if err != nil {
		log.Printf("Error creating listing: %v", err)
	}
	return err
}

func (r *ListingRepository) GetAll() ([]models.Listing, error) {
	rows, err := r.DB.Query(context.Background(), "SELECT id, title, description, price, created_at, updated_at FROM listings")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listings []models.Listing
	for rows.Next() {
		var listing models.Listing
		if err := rows.Scan(&listing.ID, &listing.Title, &listing.Description, &listing.Price, &listing.CreatedAt, &listing.UpdatedAt); err != nil {
			return nil, err
		}
		listings = append(listings, listing)
	}
	return listings, nil
}

func (r *ListingRepository) GetByID(id int) (*models.Listing, error) {
	var listing models.Listing
	err := r.DB.QueryRow(context.Background(), "SELECT id, title, description, price, created_at, updated_at FROM listings WHERE id=$1", id).
		Scan(&listing.ID, &listing.Title, &listing.Description, &listing.Price, &listing.CreatedAt, &listing.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &listing, nil
}

func (r *ListingRepository) Update(listing *models.Listing) error {
	_, err := r.DB.Exec(context.Background(), "UPDATE listings SET title=$1, description=$2, price=$3, updated_at=$4 WHERE id=$5",
		listing.Title, listing.Description, listing.Price, time.Now(), listing.ID)
	return err
}

func (r *ListingRepository) Delete(id int) error {
	_, err := r.DB.Exec(context.Background(), "DELETE FROM listings WHERE id=$1", id)
	return err
}
