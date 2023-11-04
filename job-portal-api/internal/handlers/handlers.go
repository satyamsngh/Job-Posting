package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"job-portal-api/internal/auth"
	"job-portal-api/internal/middleware"
	"job-portal-api/internal/repository"
	"job-portal-api/internal/services"
	"net/http"
	"time"
)

func API(a *auth.Auth, c repository.UserRepo) *gin.Engine {
	r := gin.New()

	m, err := middlewares.NewMid(a)
	if err != nil {
		log.Error().Err(err).Msg("Error setting up middlewares")
		return nil
	}

	ms, err := services.NewStore(c)
	if err != nil {
		log.Error().Err(err).Msg("Error setting up services")
		return nil
	}

	h := handler{
		s: ms,
		a: a,
	}

	r.Use(m.Log(), gin.Recovery())
	r.GET("/api/check", m.Authenticate(check))
	r.POST("/api/register", h.Register)
	r.POST("/api/login", h.Login)
	r.POST("/api/listcompanies", m.Authenticate(h.AddCompanies))
	r.GET("/api/viewcompanies", m.Authenticate(h.ViewCompanies))
	r.GET("/api/companies/:companyID", m.Authenticate(h.ViewCompaniesById))
	r.POST("/companies/:companyID/jobs", m.Authenticate(h.CreateJob))
	r.GET("api/companies/:companyID/list-jobs", m.Authenticate(h.ListJobs))
	r.GET("api/jobs", m.Authenticate(h.AllJobs))
	r.GET("/api/jobs/:jobID", m.Authenticate(h.JobsByID))

	return r
}

func check(c *gin.Context) {
	time.Sleep(time.Second * 3)
	select {
	case <-c.Request.Context().Done():
		err := c.Request.Context().Err()
		log.Error().Err(err).Msg("Request cancelled")
		c.Error(err)
		return
	default:
		c.JSON(http.StatusOK, gin.H{"msg": "statusOk"})
	}
}
