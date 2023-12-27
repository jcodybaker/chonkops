package devices

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
)

type Devices []Device

func Run(ctx context.Context) error {
	t := time.NewTicker(5 * time.Second)
	shelly := NewShellyDiscover()
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-t.C:
			devices, err := shelly.Discover(ctx)
			if err != nil {
				log.Warn().Err(err).Msg("discovering shelly devices")
			}
			for _, d := range devices {
				log.Info().Str("name", d.GetName()).Str("id", d.GetID()).Msg("got device")
			}
		}
	}
}
