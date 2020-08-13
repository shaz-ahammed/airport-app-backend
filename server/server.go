package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"golang-app-starter/certs"
	"golang-app-starter/config"

	"crypto/tls"
	"errors"
	"net/http"
	"time"
)

type AppServer struct {
	router *gin.Engine
}

var isServerConfigured = false

func NewServer(rtr *gin.Engine) *AppServer {
	srv := &AppServer{
		router: rtr,
	}
	srv.additionalRouterConfiguration()
	srv.setupRoutesAndMiddleware()
	isServerConfigured = true
	return srv
}

func (srv *AppServer) RunServer() error {
	if !isServerConfigured {
		return errors.New("server has not been configured. Call NewServer() first")
	}

	if config.EnableTls {
		srv := configureServerWithTLS(srv.router)
		log.Info().Msg("Starting server with TLS enabled")
		return srv.ListenAndServeTLS(config.CertPath, config.KeyPath)
	} else {
		srv := configureServerWithTLSDisabled(srv.router)
		log.Info().Msg("Starting server with TLS disabled")
		return srv.ListenAndServe()
	}
}

func (srv *AppServer) additionalRouterConfiguration() {
	log.Info().Msg("Configuring GIN router")
	srv.router.HandleMethodNotAllowed = true
}

func configureServerWithTLS(handler http.Handler) *http.Server {
	// Generate self-signed certificate is necessary
	if config.GenerateSelfSignedCert {
		certs.GenerateSelfSignedCertAndKey(config.CertPath, config.KeyPath, config.CertValidityDays)
	}

	// TLS Config - https://ssl-config.mozilla.org/#server=go&version=1.14.4&config=modern&hsts=false&guideline=5.4
	// Supports min Firefox 63, Android 10.0, Chrome 70, Edge 75, Java 11, OpenSSL 1.1.1, Opera 57, and Safari 12.1
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS13,
	}

	srv := &http.Server{
		Handler:           handler,
		Addr:              config.ServerAddress,
		ReadHeaderTimeout: time.Duration(config.ServerReadHeaderTimeout) * time.Second,
		ReadTimeout:       time.Duration(config.ServerReadTimeout) * time.Second,
		WriteTimeout:      time.Duration(config.ServerWriteTimeout) * time.Second,
		IdleTimeout:       time.Duration(config.ServerIdleTimeout) * time.Second,
		TLSConfig:         tlsConfig,
	}

	return srv
}

func configureServerWithTLSDisabled(handler http.Handler) *http.Server {
	srv := &http.Server{
		Handler:           handler,
		Addr:              config.ServerAddress,
		ReadHeaderTimeout: time.Duration(config.ServerReadHeaderTimeout) * time.Second,
		ReadTimeout:       time.Duration(config.ServerReadTimeout) * time.Second,
		WriteTimeout:      time.Duration(config.ServerWriteTimeout) * time.Second,
		IdleTimeout:       time.Duration(config.ServerIdleTimeout) * time.Second,
	}

	return srv
}
