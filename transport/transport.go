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

type Transport struct {
	ID        int       `json:"id"`
	BeginsAt  time.Time `json:"beginsAt"`
	ReturnsAt time.Time `json:"returnsAt"`
	Cargo     Cargo     `json:"cargo"`
	Carrier   Carrier   `json:"carrier"`
	Notes     string    `json:"notes"`
	Status    Status    `json:"status"`
}

type Cargo struct {
	NumberOfPeople int      `json:"numberOfPeople"`
	Materials      []string `json:"materials"`
}

type Carrier struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}

type Service struct {
	repo Repository
}

type Repository interface {
	AddTransport(transport Transport) (Transport, error)
	UpdateStatus(transportID int, newStatus Status) error
}

func NewService(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

func (svc Service) RequestNew(transport Transport) (Transport, error) {
	transport.Status = Created
	return svc.repo.AddTransport(transport)
}

func (svc Service) AnnounceNew(transport Transport) (Transport, error) {
	transport.Status = Created
	return svc.repo.AddTransport(transport)
}

func (svc Service) Accept(transportID int) error {
	return svc.repo.UpdateStatus(transportID, Accepted)
}

func (svc Service) Reject(transportID int) error {
	return svc.repo.UpdateStatus(transportID, Rejected)
}

func (svc Service) Close(transportID int) error {
	return svc.repo.UpdateStatus(transportID, Closed)
}
