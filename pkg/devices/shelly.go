package devices

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/url"
	"strconv"

	"github.com/hashicorp/mdns"
	shelly "github.com/jcodybaker/go-shelly"
	"github.com/mongoose-os/mos/common/mgrpc"
	"github.com/rs/zerolog/log"
)

const (
	defaultShellyMDNSServiceName = "_shelly._tcp"
)

func NewShellyDiscover(opts ...ShellyOption) Discoverer {
	return &shellyDiscover{}
}

type ShellyOption func(d *shellyDiscover)

type shellyDiscover struct{}

type shellyDevice struct {
	name     string
	addr     string
	port     int
	switchID int
}

func (d *shellyDiscover) Discover(ctx context.Context) ([]Device, error) {
	serviceEntries := make(chan *mdns.ServiceEntry)
	q := &mdns.QueryParam{
		Service: defaultShellyMDNSServiceName,
		Entries: serviceEntries,
	}
	go func() {
		defer close(serviceEntries)
		if err := mdns.Query(q); err != nil {
			log.Error().Err(err).Msg("querying mDNS for shelly devices")
		}
	}()
	var outDevices []Device
	for se := range serviceEntries {
		d, err := serviceEntryToShellyDevices(ctx, se)
		if err != nil {
			log.Error().Err(err).Msg("converting shelly entry to devices")
		}
		if d == nil {
			continue
		}
		outDevices = append(outDevices, d...)
	}
	return outDevices, nil
}

func serviceEntryToShellyDevices(ctx context.Context, se *mdns.ServiceEntry) ([]Device, error) {
	addr := se.AddrV4
	if addr.IsUnspecified() {
		addr = se.AddrV6
	}
	if addr.IsUnspecified() {
		return nil, errors.New("device had no addresses")
	}
	rpcAddr := (&url.URL{
		Scheme: "http",
		Host:   net.JoinHostPort(addr.String(), strconv.Itoa(se.Port)),
		Path:   "/rpc",
	}).String()

	c, err := mgrpc.New(ctx, rpcAddr, mgrpc.UseHTTPPost())
	if err != nil {
		return nil, fmt.Errorf("establishing rpc channel: %w", err)
	}
	defer c.Disconnect(ctx)

	var devices []Device
	for i := 0; true; i++ {
		req := &shelly.SwitchGetConfigRequest{
			ID: i,
		}
		resp, _, err := req.Do(ctx, c)
		if err != nil {
			var statusErr *shelly.BadStatusError
			if !errors.As(err, &statusErr) || statusErr.Status != shelly.StatusIDNotFound {
				return nil, err
			}
			break // error is a NOT found, meaning this switch doesn't exist. stop.
		}
		name := se.Name
		if resp.Name != nil {
			name = *resp.Name
		}
		devices = append(devices, &shellyDevice{
			name:     name,
			addr:     rpcAddr,
			port:     se.Port,
			switchID: i,
		})
	}
	return devices, nil
}

func (d *shellyDevice) GetName() string {
	return d.name
}

func (d *shellyDevice) GetID() string {
	return fmt.Sprintf("%s-%d", d.addr, d.switchID)
}

func (d *shellyDevice) GetStatus() Status {
	return Status{}
}
