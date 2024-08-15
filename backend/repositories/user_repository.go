package repositories

import (
	"context"

	"github.com/amrimuf/hompimRent/models"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepository struct {
    DB *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
    return &UserRepository{DB: db}
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
    var user models.User
    err := r.DB.QueryRow(context.Background(), "SELECT id, username, email, password_hash, role, created_at, updated_at FROM users WHERE email=$1", email).
        Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.Role, &user.CreatedAt, &user.UpdatedAt)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserRepository) Create(user *models.User) error {
    _, err := r.DB.Exec(context.Background(), `
        INSERT INTO users (id, username, email, password_hash, role, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7)`,
        user.ID, user.Username, user.Email, user.PasswordHash, user.Role, user.CreatedAt, user.UpdatedAt)
    return err
}

func (r *UserRepository) GetAll() ([]models.User, error) {
    rows, err := r.DB.Query(context.Background(), "SELECT id, username, email, role, created_at, updated_at FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []models.User
    for rows.Next() {
        var user models.User
        if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt); err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    return users, nil
}

func (r *UserRepository) GetByID(id uuid.UUID) (*models.User, error) {
    var user models.User
    err := r.DB.QueryRow(context.Background(), `
        SELECT id, username, email, password_hash, role, created_at, updated_at
        FROM users WHERE id = $1`, id).Scan(
        &user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.Role, &user.CreatedAt, &user.UpdatedAt)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserRepository) Update(user *models.User) error {
    _, err := r.DB.Exec(context.Background(), `
        UPDATE users SET username = $1, email = $2, password_hash = $3, role = $4, updated_at = $5
        WHERE id = $6`,
        user.Username, user.Email, user.PasswordHash, user.Role, user.UpdatedAt, user.ID)
    return err
}

func (r *UserRepository) Delete(id uuid.UUID) error {
    _, err := r.DB.Exec(context.Background(), `DELETE FROM users WHERE id = $1`, id)
    return err
}