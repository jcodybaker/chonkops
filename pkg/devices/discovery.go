package devices

import "context"

type Devices []Device

func Run(ctx context.Context) error {
	<-ctx.Done()
	return nil
}
