package tor

import (
	"context"
	"log"

	"github.com/cretz/bine/tor"
)

func NewTorService() *tor.OnionService {
	torInstance, err := tor.Start(context.Background(), nil)
	if err != nil {
		log.Fatalf("Unable to create Tor instance: %v", err)
	}
	onion, err := torInstance.Listen(context.Background(), &tor.ListenConf{
		Version3:    true,
		RemotePorts: []int{80},
	})
	if err != nil {
		log.Fatalf("Unable to start onion service: %v", err)
	}
	return onion
}
