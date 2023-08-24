package transport

import "time"

const (
	Car   Vehicle = 1
	Bus   Vehicle = 2
	Plane Vehicle = 3
)

const (
	Created  Status = 1
	Accepted Status = 2
	Rejected Status = 3
	Closed   Status = 4
)

type Vehicle int
type Status int

type Ride struct {
	ID        int       `json:"id"`
	BeginsAt  time.Time `json:"beginsAt"`
	ReturnsAt time.Time `json:"returnsAt"`
	Cargo     Cargo     `json:"cargo"`
	Rider     Rider     `json:"rider"`
	Notes     string    `json:"notes"`
	Status    Status    `json:"status"`
}

type Cargo struct {
	NumberOfPeople int      `json:"numberOfPeople"`
	Materials      []string `json:"materials"`
}

type Rider struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}

type Service struct {
	repo Repository
}

type Repository interface {
	AddRide(ride Ride) (Ride, error)
	UpdateRideStatus(rideID int, newStatus Status) error
}

func NewService(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

func (svc Service) RequestNewRide(ride Ride) (Ride, error) {
	ride.Status = Created
	return svc.repo.AddRide(ride)
}

func (svc Service) AnnounceNewRide(ride Ride) (Ride, error) {
	ride.Status = Created
	return svc.repo.AddRide(ride)
}

func (svc Service) AcceptRide(rideID int) error {
	return svc.repo.UpdateRideStatus(rideID, Accepted)
}

func (svc Service) RejectRide(rideID int) error {
	return svc.repo.UpdateRideStatus(rideID, Rejected)
}

func (svc Service) CloseRide(rideID int) error {
	return svc.repo.UpdateRideStatus(rideID, Closed)
}
