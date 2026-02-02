package user

import "context"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ListUsers(ctx context.Context) ([]User, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) CreateUser(
	ctx context.Context,
	username, displayName, passwordHash string,
) (*User, error) {
	u := &User{
		Username:     username,
		DisplayName:  &displayName,
		PasswordHash: passwordHash,
	}

	if err := s.repo.Create(ctx, u); err != nil {
		return nil, err
	}

	return u, nil
}
