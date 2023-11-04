package services

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"go.uber.org/mock/gomock"
	"job-portal-api/internal/models"
	"job-portal-api/internal/repository"
	"reflect"
	"testing"
)

func TestStore_CreateUser(t *testing.T) {

	type args struct {
		ctx context.Context
		nu  models.NewUser
	}

	tests := []struct {
		name             string
		args             args
		want             models.User
		wantErr          bool
		mockRepoResponse func() (models.User, error)
	}{
		{name: "error from database",
			args: args{
				ctx: context.Background(),
				nu: models.NewUser{
					Name:     "satyam",
					Email:    "satyam@gmail.com",
					Password: "password123",
				},
			},
			want:    models.User{},
			wantErr: true,
			mockRepoResponse: func() (models.User, error) {
				return models.User{}, errors.New("error in database")
			},
		},
		{
			name: "ok",
			args: args{
				ctx: context.Background(),
				nu: models.NewUser{
					Name:     "satyam",
					Email:    "satyam@gmail.com",
					Password: "password123",
				},
			},
			want: models.User{

				Name:         "satyam",
				Email:        "satyam@gmail.com",
				PasswordHash: "euubfbu3rbvub",
			},
			wantErr: false,
			mockRepoResponse: func() (models.User, error) {
				return models.User{
					Name:         "satyam",
					Email:        "satyam@gmail.com",
					PasswordHash: "euubfbu3rbvub",
				}, nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, err := NewStore(mockRepo)
			if err != nil {
				log.Print(err)
				return
			}
			got, err := s.CreateUser(tt.args.ctx, tt.args.nu)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Store.CreateUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_Authenticate(t *testing.T) {
	type args struct {
		ctx      context.Context
		email    string
		password string
	}
	tests := []struct {
		name             string
		args             args
		want             jwt.RegisteredClaims
		wantErr          bool
		mockRepoResponse func() (jwt.RegisteredClaims, error)
	}{
		{
			name: "Error",
			args: args{
				ctx:      context.Background(),
				email:    "satyam@gmail.com",
				password: "satyam",
			},
			want:    jwt.RegisteredClaims{},
			wantErr: true,
			mockRepoResponse: func() (jwt.RegisteredClaims, error) {
				return jwt.RegisteredClaims{}, errors.New("error in token")
			},
		},
		{
			name: "Ok",
			args: args{
				ctx:      context.Background(),
				email:    "satyam@gmail.com",
				password: "satyam",
			},
			want: jwt.RegisteredClaims{
				Issuer:  "satyam",
				Subject: "1",
				ID:      "2",
			},
			wantErr: false,
			mockRepoResponse: func() (jwt.RegisteredClaims, error) {
				return jwt.RegisteredClaims{
					Issuer:  "satyam",
					Subject: "1",
					ID:      "2",
				}, nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().CheckEmail(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()
			}

			s, err := NewStore(mockRepo)
			if err != nil {
				t.Fatalf("error creating Store: %v", err)
			}

			got, err := s.Authenticate(tt.args.ctx, tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Authenticate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Authenticate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
