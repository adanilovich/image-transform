package main

import (
	"log"
	"net"
	"net/http"

	"github.com/asfeather.com/internal/constant"
	"github.com/asfeather.com/internal/middleware"
	"github.com/asfeather.com/internal/router"
	"github.com/asfeather.com/internal/template"
	"github.com/asfeather.com/pkg/file_system"
	"github.com/spf13/pflag"

	"go.uber.org/zap"
)

const defaultConfigPath = "./build.json"

// dev flag and write logs to files
// add process of config
func main() {
	var configPath string
	pflag.StringVarP(&configPath, "config", "c", defaultConfigPath, "set config path")
	pflag.Parse()

	conf, err := newConf(configPath)
	if err != nil {
		log.Fatal(err)
	}

	logger, err := newZapLogger(conf)
	if err != nil {
		log.Fatal(err)
	}

	// Main server
	// Try to open port
	ln, err := net.Listen("tcp", conf.ServAddr)
	if err != nil {
		logger.Fatal("Service failed:", zap.Error(err))
	}

	// init middleware
	log := middleware.NewLogIntercept(logger)
	mon := middleware.NewMonitor(*logger)
	fs := file_system.NewFS() // dependency injection

	templateCollection, err := template.CompileTemplates(fs, constant.PATH_DIR_TEMPLATES)
	if err != nil {
		logger.Fatal("Load the templates:", zap.Error(err))
	}

	// adjust router
	r := router.New(fs, templateCollection)
	r.Use(log.AddToCtx)
	r.Use(mon.LogResponseTime)
	//	r.Use(middleware.AddResponseHeaders)
	//	r.Use(middleware.AddFileReader)

	logger.Info("Service started", zap.String("addr", conf.ServAddr))

	s := http.Server{
		Handler: r,
	}
	if err = s.Serve(ln); err != nil {
		logger.Fatal("Service failed: %v", zap.Error(err))
	}

}
