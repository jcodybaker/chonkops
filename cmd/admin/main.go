package main

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/jcodybaker/chonkops/pkg/admin"
	"github.com/jcodybaker/chonkops/pkg/devices"
)

const defaultAddr = ":8080"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx, signalStop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer signalStop()

	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	addr := defaultAddr
	if port := viper.GetInt("port"); port != 0 {
		addr = fmt.Sprintf(":%d", port)
	}

	var assetFS fs.FS
	var assetPath string
	var err error
	if assetPath = viper.GetString("ASSET_PATH"); assetPath == "" {
		assetPath, err = os.Getwd()
		if err != nil {
			log.Fatal().Err(err).Msg("finding current dir")
		}
	}
	assetFS = os.DirFS(assetPath)

	a, err := admin.NewServer(
		admin.WithAssetFileSystem(assetFS),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("initializing server")
	}

	s := http.Server{
		Addr:    addr,
		Handler: a,
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer cancel()
		log.Info().Str("addr", s.Addr).Msg("starting http server")
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error().Err(err).Msg("running server")
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		log.Info().Msg("shutting down; 5s countdown")
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer shutdownCancel()
		s.Shutdown(shutdownCtx)
		log.Info().Msg("http shutdown complete")
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		devices.Run(ctx)
		log.Info().Msg("device shutdown complete")
	}()
	wg.Wait()
}
