package services

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"go.uber.org/mock/gomock"
	"job-portal-api/internal/models"
	"job-portal-api/internal/repository"
	"reflect"
	"testing"
)

func TestStore_CreatCompanies(t *testing.T) {

	type args struct {
		ctx    context.Context
		nc     models.NewComapanies
		UserID uint
	}
	tests := []struct {
		name             string
		args             args
		want             models.Companies
		wantErr          bool
		mockRepoResponse func() (models.Companies, error)
	}{
		{name: "error from database",
			args: args{
				ctx: context.Background(),
				nc: models.NewComapanies{
					CompanyName: "google",
					FoundedYear: 2019,
					Location:    "banglore",
					Address:     "blndr",
				},
				UserID: 1,
			},
			want:    models.Companies{},
			wantErr: true,
			mockRepoResponse: func() (models.Companies, error) {
				return models.Companies{}, errors.New("error in data base")

			},
		},
		{
			name: "ok",
			args: args{
				ctx: context.Background(),
				nc: models.NewComapanies{
					CompanyName: "google",
					FoundedYear: 2019,
					Location:    "banglore",
					Address:     "blndr",
				},
				UserID: 1,
			},
			want: models.Companies{
				CompanyName: "google",
				FoundedYear: 2019,
				Location:    "banglore",
				UserId:      1,
				Address:     "blndr",
			},
			wantErr: false,
			mockRepoResponse: func() (models.Companies, error) {
				return models.Companies{
					CompanyName: "google",
					FoundedYear: 2019,
					Location:    "banglore",
					UserId:      1,
					Address:     "blndr",
				}, nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().CreateCompany(gomock.Any(), gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, err := NewStore(mockRepo)
			if err != nil {
				log.Print(err)
				return
			}
			got, err := s.CreatCompanies(tt.args.ctx, tt.args.nc, tt.args.UserID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.ViewJobById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreatCompanies() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_ViewCompanies(t *testing.T) {

	type args struct {
		ctx       context.Context
		companyID string
	}
	tests := []struct {
		name           string
		args           args
		want           []models.Companies
		wantErr        bool
		mockNewService func() ([]models.Companies, error)
	}{
		{
			name: "error from database",
			args: args{
				ctx:       context.Background(),
				companyID: "1",
			},
			want:    nil,
			wantErr: true,
			mockNewService: func() ([]models.Companies, error) {
				return nil, errors.New("data base error")
			},
		},
		{
			name: "OK",
			args: args{
				ctx:       context.Background(),
				companyID: "1",
			},
			want: models.Companies{
				CompanyName: "slk",
				FoundedYear: 2011,
				Location:    "banglore",
				UserId:      1,
				Address:     "",
				Jobs:        nil,
			},
			wantErr: true,
			mockNewService: func() ([]models.Companies, error) {
				return nil, errors.New("data base error")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mock)
			if tt.mockNewService != nil {
				mockRepo.EXPECT().ViewCompanies(gomock.Any()).Return(tt.mockNewService()).AnyTimes()
			}
			s, err := NewStore(mockRepo)
			if err != nil {
				log.Print(err)
				return
			}

			got, err := s.ViewCompanies(tt.args.ctx, tt.args.companyID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ViewCompanies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ViewCompanies() got = %v, want %v", got, tt.want)
			}
		})
	}
}
