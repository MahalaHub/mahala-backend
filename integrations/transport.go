package integrations

import (
	"github.com/mahalahub/mahala/internal/redis"
	"github.com/mahalahub/mahala/transport"
)

type TransportRepository struct {
	redisClient *redis.Client
}

func NewTransportRepository(redisClient *redis.Client) *TransportRepository {
	return &TransportRepository{
		redisClient: redisClient,
	}
}

func (repo TransportRepository) AddRide(ride transport.Ride) (transport.Ride, error) {
	return transport.Ride{}, nil
}

func (repo TransportRepository) UpdateRideStatus(rideID int, newStatus transport.Status) error {
	return nil
}
