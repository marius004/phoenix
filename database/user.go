package database

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/marius004/phoenix/internal/models"
)

// UserService implements services.UserService
type UserService struct {
	db     *sqlx.DB
	logger *log.Logger
}

func (s *UserService) GetById(ctx context.Context, id int) (*models.User, error) {
	var user models.User
	err := s.db.GetContext(ctx, &user, s.db.Rebind("SELECT * FROM users where id = ? LIMIT 1"), id)

	if err != nil {
		return nil, err
	}

	return &user, err
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	query := "SELECT * FROM users WHERE email = ? LIMIT 1"
	err := s.db.GetContext(ctx, &user, s.db.Rebind(query), email)

	if err != nil {
		return nil, err
	}

	return &user, err
}

func (s *UserService) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User

	query := "SELECT * FROM users WHERE username = ? LIMIT 1"
	err := s.db.GetContext(ctx, &user, s.db.Rebind(query), username)

	if err != nil {
		return nil, err
	}

	return &user, err
}

func (s *UserService) GetAll(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	err := s.db.Select(&users, "SELECT * FROM users")

	if err != nil {
		return nil, err
	}

	return users, err
}

func (s *UserService) Create(ctx context.Context, user *models.User) error {
	var id uint64

	query := "INSERT INTO users(username, email, password, bio, visible, is_admin, is_proposer, is_banned, verified_email) VALUES(?,?,?,?,?,?,?,?,?) RETURNING id"

	err := s.db.GetContext(ctx, &id, s.db.Rebind(query),
		user.Username, user.Email, user.Password,
		user.Bio, user.Visible, user.IsAdmin, user.IsProposer, user.IsBanned, user.VerifiedEmail)

	if err == nil {
		user.Id = id
	}

	return err
}

func (s *UserService) ExistsByUsername(ctx context.Context, username string) (bool, error) {
	var count int
	query := "SELECT COUNT(*) FROM users WHERE username = ?"
	err := s.db.GetContext(ctx, &count, s.db.Rebind(query), username)

	if err != nil {
		s.logger.Println(err)
		return false, nil
	}

	return count > 0, err
}

func (s *UserService) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var count int
	query := "SELECT COUNT(*) FROM users WHERE email = ?"
	err := s.db.GetContext(ctx, &count, s.db.Rebind(query), email)

	if err != nil {
		s.logger.Println(err)
		return false, nil
	}

	return count > 0, err
}

func (s *UserService) ExistsById(ctx context.Context, id int) (bool, error) {
	var count int
	query := "SELECT COUNT(*) FROM users WHERE id = ?"
	err := s.db.GetContext(ctx, &count, s.db.Rebind(query), id)

	if err != nil {
		s.logger.Println(err)
		return false, nil
	}

	return count > 0, err
}

func NewUserService(db *sqlx.DB, logger *log.Logger) *UserService {
	return &UserService{db, logger}
}
