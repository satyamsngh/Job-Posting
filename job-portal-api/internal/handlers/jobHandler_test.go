package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
	"job-portal-api/internal/auth"
	middlewares "job-portal-api/internal/middleware"
	"job-portal-api/internal/models"
	"job-portal-api/internal/services"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestViewCompanies(t *testing.T) {
	// Sets the Gin router mode to test.
	gin.SetMode(gin.TestMode)
	fakeClaims := jwt.RegisteredClaims{
		Subject: "1",
	}
	// MockUser struct initialization
	mockCompanies := []models.Companies{
		{
			Model: gorm.Model{
				ID:        1,
				CreatedAt: time.Date(2006, 1, 1, 1, 1, 1, 1, time.UTC),
				UpdatedAt: time.Date(2006, 1, 1, 1, 1, 1, 1, time.UTC),
			},
			CompanyName: "infy",
			FoundedYear: 2019,
			Location:    "banglore",
			UserId:      1,
			Address:     "blndr",
		},
		// Add more sample companies if needed.
	}
	// Define the list of test cases
	testCases := []struct {
		name              string                        // Name of the test case
		expectedStatus    int                           // Expected status of the response
		expectedResponse  string                        // Expected response body
		expectedCompanies []models.Companies            // Expected user after signup
		mockService       func(m *services.MockService) // Mock service function
	}{
		{
			name:              "OK",
			expectedStatus:    200,
			expectedResponse:  `{"companies list":[{"ID":1,"CreatedAt":"2006-01-01T01:01:01.000000001Z","UpdatedAt":"2006-01-01T01:01:01.000000001Z","DeletedAt":null,"company_name":"infy","founded_year":2019,"location":"banglore","user_id":1,"address":"blndr"}]}`,
			expectedCompanies: mockCompanies,
			mockService: func(m *services.MockService) {
				m.EXPECT().ViewCompanies(gomock.Any(), gomock.Any()).Times(1).
					Return(mockCompanies, nil)
			},
		},
		{
			name:              "Error",
			expectedStatus:    400,
			expectedResponse:  `{"msg":"problem in viewing company"}`,
			expectedCompanies: nil,
			mockService: func(m *services.MockService) {
				m.EXPECT().ViewCompanies(gomock.Any(), gomock.Any()).Times(1).
					Return(nil, errors.New("internal server error"))
			},
		},
	}

	// Start a loop over `testCases` array where each element is represented by `tc`.
	for _, tc := range testCases {
		// Run a new test with the `tc.name` as its identifier.
		t.Run(tc.name, func(t *testing.T) {
			// Create a new Gomock controller.
			ctrl := gomock.NewController(t)

			// Create a mock Inventory using the Gomock controller.
			mockS := services.NewMockService(ctrl)

			// Apply the mock to the user service.
			tc.mockService(mockS)

			// Create a new instance of `models.Service` with the mock service.
			ms := services.NewStore(mockS)

			// Create a new context. This is typically passed between functions
			// carrying deadline, cancellation signals, and other request-scoped values.
			ctx := context.Background()
			// Insert the TraceId into the context.
			ctx = context.WithValue(ctx, auth.Key, fakeClaims)
			// Create a fake TraceID. This would typically be used for request tracing.
			traceID := "fake-trace-id"
			// Insert the TraceId into the context.
			ctx = context.WithValue(ctx, middlewares.TraceIdKey, traceID)

			// Create a new Gin router.
			router := gin.New()

			// Create a new handler which uses the service model.
			h := handler{s: ms}

			// Register an endpoint and its handler with the router.
			router.GET("/api/companies", h.ViewCompanies)

			// Create a new HTTP GET request to "/api/companies".
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/api/companies", nil)
			// If the request creation fails, raise an error and stop the test.
			require.NoError(t, err)

			// Create a new HTTP Response Recorder. This is used to capture the HTTP response for analysis.
			resp := httptest.NewRecorder()

			// Pass the HTTP request to the router. This effectively "performs" the request and gets the associated response.
			router.ServeHTTP(resp, req)

			// Assert the returned HTTP status code is as expected.
			require.Equal(t, tc.expectedStatus, resp.Code)

			// Assert the response matches the expected response.
			require.Equal(t, tc.expectedResponse, string(resp.Body.Bytes()))
		})
	}
}

func TestViewCompaniesById(t *testing.T) {
	// Sets the Gin router mode to test.
	gin.SetMode(gin.TestMode)
	fakeClaims := jwt.RegisteredClaims{
		Subject: "1",
	}

	// MockUser struct initialization
	mockCompanies := []models.Companies{
		{
			Model: gorm.Model{
				ID:        1,
				CreatedAt: time.Date(2006, 1, 1, 1, 1, 1, 1, time.UTC),
				UpdatedAt: time.Date(2006, 1, 1, 1, 1, 1, 1, time.UTC),
			},
			CompanyName: "infy",
			FoundedYear: 2019,
			Location:    "banglore",
			UserId:      1,
			Address:     "blndr",
		},
		// Add more sample companies if needed.
	}
	// Define the list of test cases
	testCases := []struct {
		name              string // Name of the test case
		body              any
		expectedStatus    int // Expected status of the response
		expectedResponse  string
		expectedcompanies []models.Companies            // Expected response body 		// Expected user after signup
		mockService       func(m *services.MockService) // Mock service function
	}{
		{
			name:           "OK",
			body:           mockCompanies,
			expectedStatus: 200,

			expectedResponse: `[{"ID":1,"CreatedAt":"2006-01-01T01:01:01.000000001Z","UpdatedAt":"2006-01-01T01:01:01.000000001Z","DeletedAt":null,"company_name":"infy","founded_year":2019,"location":"banglore","user_id":1,"address":"blndr"}]`,
			mockService: func(m *services.MockService) {

				m.EXPECT().ViewCompaniesById(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1).
					Return(mockCompanies, nil)

			},
		},
		{
			name: "Error - Company Not Found",
			body: []models.Companies{
				{
					Model: gorm.Model{
						ID: 5,

						CreatedAt: time.Date(2006, 1, 1, 1, 1, 1, 1, time.UTC),
						UpdatedAt: time.Date(2006, 1, 1, 1, 1, 1, 1, time.UTC),
					},
					CompanyName: "infy",
					FoundedYear: 2019,
					Location:    "banglore",
					UserId:      1,
					Address:     "blndr",
				},
				// Add more sample companies if needed.
			},
			expectedStatus:   400,
			expectedResponse: `{"msg":"problem in fetching company details"}`,
			mockService: func(m *services.MockService) {
				m.EXPECT().ViewCompaniesById(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1).Return(nil, errors.New(""))
			},
		},
	}

	// Start a loop over `testCases` array where each element is represented by `tc`.
	for _, tc := range testCases {
		// Run a new test with the `tc.name` as its identifier.
		t.Run(tc.name, func(t *testing.T) {
			// Create a new Gomock controller.
			ctrl := gomock.NewController(t)

			// Create a mock Inventory using the Gomock controller.
			mockS := services.NewMockService(ctrl)

			// Apply the mock to the user service.
			tc.mockService(mockS)

			// Create a new instance of `models.Service` with the mock service.
			ms := services.NewStore(mockS)

			// Create a new context. This is typically passed between functions
			// carrying deadline, cancellation signals, and other request-scoped values.
			ctx := context.Background()
			// Insert the TraceId into the context.
			ctx = context.WithValue(ctx, auth.Key, fakeClaims)
			// Create a fake TraceID. This would typically be used for request tracing.
			traceID := "fake-trace-id"
			// Insert the TraceId into the context.
			ctx = context.WithValue(ctx, middlewares.TraceIdKey, traceID)

			// Create a new Gin router.
			router := gin.New()

			// Create a new handler which uses the service model.
			h := handler{s: ms}

			// Register an endpoint and its handler with the router.
			router.GET("/api/companies/:companyID", h.ViewCompaniesById)

			// Create a new HTTP POST equest to "/signup".
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/api/companies/1", nil)
			// If the request creation fails, raise an error and stop the test.
			require.NoError(t, err)

			// Create a new HTTP Response Recorder. This is used to capture the HTTP response for analysis.
			resp := httptest.NewRecorder()

			// Pass the HTTP request to the router. This effectively "performs" the request and gets the associated response.
			router.ServeHTTP(resp, req)

			// Assert the returned HTTP status code is as expected.
			require.Equal(t, tc.expectedStatus, resp.Code)

			// Assert the response matches the expected response.
			require.Equal(t, tc.expectedResponse, string(resp.Body.Bytes()))
		})
	}
}

func TestHandler_CreateJob(t *testing.T) {
	// Sets the Gin router mode to test.
	gin.SetMode(gin.TestMode)
	fakeClaims := jwt.RegisteredClaims{
		Subject: "1",
	}

	// Define the input data for creating a job
	jobData := models.Job{
		Title:       "Software Engineer",
		Description: "Senior",
		CompanyID:   1,
	}

	// Define the list of test cases
	testCases := []struct {
		name             string                        // Name of the test case
		expectedStatus   int                           // Expected status of the response
		expectedResponse string                        // Expected response body
		mockService      func(m *services.MockService) // Mock service function
	}{
		{
			name:           "OK",
			expectedStatus: 201,
			// You can adjust the expected response based on your application's actual response format.
			expectedResponse: `{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"title":"Software Engineer","description":"Senior","company_id":1}`,
			// Function for mocking service.
			// This simulates CreateJob service and its return value.
			mockService: func(m *services.MockService) {
				m.EXPECT().CreateJob(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).
					Return(jobData, nil)
			},
		},
	}

	// Start a loop over `testCases` array where each element is represented by `tc`.
	for _, tc := range testCases {
		// Run a new test with the `tc.name` as its identifier.
		t.Run(tc.name, func(t *testing.T) {
			// Create a new Gomock controller.
			ctrl := gomock.NewController(t)

			// Create a mock Inventory using the Gomock controller.
			mockS := services.NewMockService(ctrl)

			// Apply the mock to the user service.
			tc.mockService(mockS)

			// Create a new instance of `models.Service` with the mock service.

			ms := services.NewStore(mockS)

			// Create a new context. This is typically passed between functions
			// carrying deadline, cancellation signals, and other request-scoped values.
			ctx := context.Background()
			// Insert the TraceId into the context.
			ctx = context.WithValue(ctx, auth.Key, fakeClaims)
			// Create a fake TraceID. This would typically be used for request tracing.
			traceID := "fake-trace-id"
			// Insert the TraceId into the context.
			ctx = context.WithValue(ctx, middlewares.TraceIdKey, traceID)

			// Create a new Gin router.
			router := gin.New()

			// Create a new handler which uses the service model.
			h := handler{s: ms}

			// Register an endpoint and its handler with the router.
			router.POST("/companies/:companyID/jobs", h.CreateJob)

			// Serialize jobData to JSON and create a request body
			reqBody, _ := json.Marshal(jobData)

			// Create a new HTTP POST request to "/createjob".
			req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/companies/1/jobs", bytes.NewReader(reqBody))
			// If the request creation fails, raise an error and stop the test.
			require.NoError(t, err)

			// Create a new HTTP Response Recorder. This is used to capture the HTTP response for analysis.
			resp := httptest.NewRecorder()

			// Pass the HTTP request to the router. This effectively "performs" the request and gets the associated response.
			router.ServeHTTP(resp, req)

			// Assert the returned HTTP status code is as expected.
			require.Equal(t, tc.expectedStatus, resp.Code)

			// Assert the response matches the expected response.
			require.Equal(t, tc.expectedResponse, string(resp.Body.Bytes()))
		})
	}
}
func TestAllJobs(t *testing.T) {
	// Sets the Gin router mode to test.
	gin.SetMode(gin.TestMode)
	fakeClaims := jwt.RegisteredClaims{
		Subject: "1",
	}
	// MockUser struct initialization
	mockJob := []models.Job{
		{
			Model: gorm.Model{
				ID:        1,
				CreatedAt: time.Date(2006, 1, 1, 1, 1, 1, 1, time.UTC),
				UpdatedAt: time.Date(2006, 1, 1, 1, 1, 1, 1, time.UTC),
			},
			Title:       "Software Engineer",
			Description: "Senior",
			CompanyID:   1,
		},
		// Add more sample companies if needed.
	}
	// Define the list of test cases
	testCases := []struct {
		name              string                        // Name of the test case
		expectedStatus    int                           // Expected status of the response
		expectedResponse  string                        // Expected response body
		expectedCompanies []models.Job                  // Expected user after signup
		mockService       func(m *services.MockService) // Mock service function
	}{
		{
			name:              "OK",
			expectedStatus:    200,
			expectedResponse:  `[{"ID":1,"CreatedAt":"2006-01-01T01:01:01.000000001Z","UpdatedAt":"2006-01-01T01:01:01.000000001Z","DeletedAt":null,"title":"Software Engineer","description":"Senior","company_id":1}]`,
			expectedCompanies: mockJob,
			mockService: func(m *services.MockService) {

				m.EXPECT().AllJob(gomock.Any(), gomock.Any()).Times(1).
					Return(mockJob, nil)
			},
		},
	}

	// Start a loop over `testCases` array where each element is represented by `tc`.
	for _, tc := range testCases {
		// Run a new test with the `tc.name` as its identifier.
		t.Run(tc.name, func(t *testing.T) {
			// Create a new Gomock controller.
			ctrl := gomock.NewController(t)

			// Create a mock Inventory using the Gomock controller.
			mockS := services.NewMockService(ctrl)

			// Apply the mock to the user service.
			tc.mockService(mockS)

			// Create a new instance of `models.Service` with the mock service.
			ms := services.NewStore(mockS)

			// Create a new context. This is typically passed between functions
			// carrying deadline, cancellation signals, and other request-scoped values.
			ctx := context.Background()
			// Insert the TraceId into the context.
			ctx = context.WithValue(ctx, auth.Key, fakeClaims)
			// Create a fake TraceID. This would typically be used for request tracing.
			traceID := "fake-trace-id"
			// Insert the TraceId into the context.
			ctx = context.WithValue(ctx, middlewares.TraceIdKey, traceID)

			// Create a new Gin router.
			router := gin.New()

			// Create a new handler which uses the service model.
			h := handler{s: ms}

			// Register an endpoint and its handler with the router.
			router.GET("/jobs", h.AllJobs)

			// Create a new HTTP POST request to "/signup".
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/jobs", nil)
			// If the request creation fails, raise an error and stop the test.
			require.NoError(t, err)

			// Create a new HTTP Response Recorder. This is used to capture the HTTP response for analysis.
			resp := httptest.NewRecorder()

			// Pass the HTTP request to the router. This effectively "performs" the request and gets the associated response.
			router.ServeHTTP(resp, req)

			// Assert the returned HTTP status code is as expected.
			require.Equal(t, tc.expectedStatus, resp.Code)

			// Assert the response matches the expected response.
			require.Equal(t, tc.expectedResponse, string(resp.Body.Bytes()))
		})
	}
}
func TestJobsById(t *testing.T) {
	// Sets the Gin router mode to test.
	gin.SetMode(gin.TestMode)
	fakeClaims := jwt.RegisteredClaims{
		Subject: "1",
	}
	// MockUser struct initialization
	mockJobs := []models.Job{
		{
			Model: gorm.Model{
				ID:        1,
				CreatedAt: time.Date(2006, 1, 1, 1, 1, 1, 1, time.UTC),
				UpdatedAt: time.Date(2006, 1, 1, 1, 1, 1, 1, time.UTC),
			},
			Title:       "Software Engineer",
			Description: "Senior",
			CompanyID:   1,
		},
		// Add more sample companies if needed.
	}
	// Define the list of test cases
	testCases := []struct {
		name             string                        // Name of the test case
		expectedStatus   int                           // Expected status of the response
		expectedResponse string                        // Expected response body 		// Expected user after signup
		mockService      func(m *services.MockService) // Mock service function
	}{
		{
			name:             "OK",
			expectedStatus:   200,
			expectedResponse: `[{"ID":1,"CreatedAt":"2006-01-01T01:01:01.000000001Z","UpdatedAt":"2006-01-01T01:01:01.000000001Z","DeletedAt":null,"title":"Software Engineer","description":"Senior","company_id":1}]`,
			mockService: func(m *services.MockService) {

				m.EXPECT().JobsById(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1).
					Return(mockJobs, nil)

			},
		},
	}

	// Start a loop over `testCases` array where each element is represented by `tc`.
	for _, tc := range testCases {
		// Run a new test with the `tc.name` as its identifier.
		t.Run(tc.name, func(t *testing.T) {
			// Create a new Gomock controller.
			ctrl := gomock.NewController(t)

			// Create a mock Inventory using the Gomock controller.
			mockS := services.NewMockService(ctrl)

			// Apply the mock to the user service.
			tc.mockService(mockS)

			// Create a new instance of `models.Service` with the mock service.
			ms := services.NewStore(mockS)

			// Create a new context. This is typically passed between functions
			// carrying deadline, cancellation signals, and other request-scoped values.
			ctx := context.Background()
			// Insert the TraceId into the context.
			ctx = context.WithValue(ctx, auth.Key, fakeClaims)
			// Create a fake TraceID. This would typically be used for request tracing.
			traceID := "fake-trace-id"
			// Insert the TraceId into the context.
			ctx = context.WithValue(ctx, middlewares.TraceIdKey, traceID)

			// Create a new Gin router.
			router := gin.New()

			// Create a new handler which uses the service model.
			h := handler{s: ms}

			// Register an endpoint and its handler with the router.
			router.GET("/api/jobs/:jobID", h.JobsByID)

			// Create a new HTTP POST equest to "/signup".
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/api/jobs/1", nil)
			// If the request creation fails, raise an error and stop the test.
			require.NoError(t, err)

			// Create a new HTTP Response Recorder. This is used to capture the HTTP response for analysis.
			resp := httptest.NewRecorder()

			// Pass the HTTP request to the router. This effectively "performs" the request and gets the associated response.
			router.ServeHTTP(resp, req)

			// Assert the returned HTTP status code is as expected.
			require.Equal(t, tc.expectedStatus, resp.Code)

			// Assert the response matches the expected response.
			require.Equal(t, tc.expectedResponse, string(resp.Body.Bytes()))
		})
	}
}
func TestHandler_CreateCompany(t *testing.T) {
	// Sets the Gin router mode to test.
	gin.SetMode(gin.TestMode)
	fakeClaims := jwt.RegisteredClaims{
		Subject: "1",
	}

	// Define the input data for creating a job
	jobData := models.Companies{
		CompanyName: "infy",
		FoundedYear: 2019,
		Location:    "banglore",
		UserId:      1,
		Address:     "blndr",
	}

	// Define the list of test cases
	testCases := []struct {
		name             string                        // Name of the test case
		expectedStatus   int                           // Expected status of the response
		expectedResponse string                        // Expected response body
		mockService      func(m *services.MockService) // Mock service function
	}{
		{
			name:           "OK",
			expectedStatus: 200,
			// You can adjust the expected response based on your application's actual response format.
			expectedResponse: `{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"company_name":"infy","founded_year":2019,"location":"banglore","user_id":1,"address":"blndr"}`,
			// Function for mocking service.
			// This simulates CreateJob service and its return value.
			mockService: func(m *services.MockService) {
				m.EXPECT().CreatCompanies(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).
					Return(jobData, nil)
			},
		},
	}

	// Start a loop over `testCases` array where each element is represented by `tc`.
	for _, tc := range testCases {
		// Run a new test with the `tc.name` as its identifier.
		t.Run(tc.name, func(t *testing.T) {
			// Create a new Gomock controller.
			ctrl := gomock.NewController(t)

			// Create a mock Inventory using the Gomock controller.
			mockS := services.NewMockService(ctrl)

			// Apply the mock to the user service.
			tc.mockService(mockS)

			// Create a new instance of `models.Service` with the mock service.

			ms := services.NewStore(mockS)

			// Create a new context. This is typically passed between functions
			// carrying deadline, cancellation signals, and other request-scoped values.
			ctx := context.Background()
			// Insert the TraceId into the context.
			ctx = context.WithValue(ctx, auth.Key, fakeClaims)
			// Create a fake TraceID. This would typically be used for request tracing.
			traceID := "fake-trace-id"
			// Insert the TraceId into the context.
			ctx = context.WithValue(ctx, middlewares.TraceIdKey, traceID)

			// Create a new Gin router.
			router := gin.New()

			// Create a new handler which uses the service model.
			h := handler{s: ms}

			// Register an endpoint and its handler with the router.
			router.POST("/api/companies", h.AddCompanies)

			// Serialize jobData to JSON and create a request body
			reqBody, _ := json.Marshal(jobData)

			// Create a new HTTP POST request to "/createjob".
			req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/api/companies", bytes.NewReader(reqBody))
			// If the request creation fails, raise an error and stop the test.
			require.NoError(t, err)

			// Create a new HTTP Response Recorder. This is used to capture the HTTP response for analysis.
			resp := httptest.NewRecorder()

			// Pass the HTTP request to the router. This effectively "performs" the request and gets the associated response.
			router.ServeHTTP(resp, req)

			// Assert the returned HTTP status code is as expected.
			require.Equal(t, tc.expectedStatus, resp.Code)

			// Assert the response matches the expected response.
			require.Equal(t, tc.expectedResponse, string(resp.Body.Bytes()))
		})
	}
}
func TestViewJobById(t *testing.T) {
	// Sets the Gin router mode to test.
	gin.SetMode(gin.TestMode)
	fakeClaims := jwt.RegisteredClaims{
		Subject: "1",
	}
	// MockUser struct initialization
	mockCompanies := []models.Job{
		{
			Model: gorm.Model{
				ID:        1,
				CreatedAt: time.Date(2006, 1, 1, 1, 1, 1, 1, time.UTC),
				UpdatedAt: time.Date(2006, 1, 1, 1, 1, 1, 1, time.UTC),
			},
			Title:       "Software Engineer",
			Description: "Senior",
			CompanyID:   1,
		},
		// Add more sample companies if needed.
	}
	// Define the list of test cases
	testCases := []struct {
		name             string                        // Name of the test case
		expectedStatus   int                           // Expected status of the response
		expectedResponse string                        // Expected response body 		// Expected user after signup
		mockService      func(m *services.MockService) // Mock service function
	}{
		{
			name:             "OK",
			expectedStatus:   200,
			expectedResponse: `[{"ID":1,"CreatedAt":"2006-01-01T01:01:01.000000001Z","UpdatedAt":"2006-01-01T01:01:01.000000001Z","DeletedAt":null,"title":"Software Engineer","description":"Senior","company_id":1}]`,
			mockService: func(m *services.MockService) {

				m.EXPECT().ListJobs(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1).
					Return(mockCompanies, nil)

			},
		},
	}

	// Start a loop over `testCases` array where each element is represented by `tc`.
	for _, tc := range testCases {
		// Run a new test with the `tc.name` as its identifier.
		t.Run(tc.name, func(t *testing.T) {
			// Create a new Gomock controller.
			ctrl := gomock.NewController(t)

			// Create a mock Inventory using the Gomock controller.
			mockS := services.NewMockService(ctrl)

			// Apply the mock to the user service.
			tc.mockService(mockS)

			// Create a new instance of `models.Service` with the mock service.
			ms := services.NewStore(mockS)

			// Create a new context. This is typically passed between functions
			// carrying deadline, cancellation signals, and other request-scoped values.
			ctx := context.Background()
			// Insert the TraceId into the context.
			ctx = context.WithValue(ctx, auth.Key, fakeClaims)
			// Create a fake TraceID. This would typically be used for request tracing.
			traceID := "fake-trace-id"
			// Insert the TraceId into the context.
			ctx = context.WithValue(ctx, middlewares.TraceIdKey, traceID)

			// Create a new Gin router.
			router := gin.New()

			// Create a new handler which uses the service model.
			h := handler{s: ms}

			// Register an endpoint and its handler with the router.
			router.GET("/api/companies/:companyID/list-jobs", h.ListJobs)

			// Create a new HTTP POST equest to "/signup".
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/api/companies/1/list-jobs", nil)
			// If the request creation fails, raise an error and stop the test.
			require.NoError(t, err)

			// Create a new HTTP Response Recorder. This is used to capture the HTTP response for analysis.
			resp := httptest.NewRecorder()

			// Pass the HTTP request to the router. This effectively "performs" the request and gets the associated response.
			router.ServeHTTP(resp, req)

			// Assert the returned HTTP status code is as expected.
			require.Equal(t, tc.expectedStatus, resp.Code)

			// Assert the response matches the expected response.
			require.Equal(t, tc.expectedResponse, string(resp.Body.Bytes()))
		})
	}
}
