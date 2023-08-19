package transport

import "time"

const (
	Car   Vehicle = 1
	Bus   Vehicle = 2
	Plane Vehicle = 3
)

type Vehicle int

type Transport struct {
	ID        int       `json:"id"`
	BeginsAt  time.Time `json:"beginsAt"`
	ReturnsAt time.Time `json:"returnsAt"`
	Cargo     Cargo     `json:"cargo"`
	Carrier   Carrier   `json:"carrier"`
	Notes     string    `json:"notes"`
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

func RequestNew(transport Transport) (Transport, error) {
	return transport, nil
}

func AnnounceNew(transport Transport) (Transport, error) {
	return transport, nil
}

func Accept(transportID int) error {
	return nil
}

func Reject(transportID int) error {
	return nil
}

func Close(transportID int) error {
	return nil
}

func SuggestChange(transport Transport) error {
	return nil
}
