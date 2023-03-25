package users

import (
	"context"
	"errors"

	log "github.com/sirupsen/logrus"
)

var (
	ErrFetchingUser   = errors.New("could not fetch User by email")
	ErrUpdatingUser   = errors.New("could not update User ")
	ErrNoUserFound    = errors.New("no User  found")
	ErrDeletingUser   = errors.New("could not delete User ")
	ErrNotImplemented = errors.New("not implemented")
)

// User - defines our user structure
type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserStore - defines the interface we need our user storage
// layer to implement
type UserStore interface {
	GetUser(ctx context.Context, email string) (User, error)
	CreateUser(ctx context.Context, in User) (User, error)
	UpdateUser(ctx context.Context, in User) (User, error)
	DeleteUser(ctx context.Context, email string) error
}

type Service struct {
	Store UserStore
}

func NewService(store UserStore) *Service {
	return &Service{
		Store: store,
	}
}

// GetUser - retrieves user data by their email from the database
func (s *Service) GetUser(ctx context.Context, email string) (User, error) {
	// calls store passing in the context
	in, err := s.Store.GetUser(ctx, email)
	if err != nil {
		log.Errorf("an error occured fetching the user: %s", err.Error())
		return User{}, ErrFetchingUser
	}
	return in, nil
}

// CreateUser - adds a new user to the database
func (s *Service) CreateUser(ctx context.Context, in User) (User, error) {
	in, err := s.Store.CreateUser(ctx, in)
	if err != nil {
		log.Errorf("an error occurred adding the user: %s", err.Error())
	}
	return in, nil
}

// UpdateUser - updates a user data with new user info
func (s *Service) UpdateUser(
	ctx context.Context, in User,
) (User, error) {
	in, err := s.Store.UpdateUser(ctx, in)
	if err != nil {
		log.Errorf("an error occurred updating the user: %s", err.Error())
	}
	return in, nil
}

// DeleteUser - deletes a user from the database by email
func (s *Service) DeleteUser(ctx context.Context, email string) error {
	return s.Store.DeleteUser(ctx, email)
}
