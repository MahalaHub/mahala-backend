package hangout

import (
	"github.com/google/uuid"
)

type Activity struct {
	ID          int      `json:"id"`
	Description string   `json:"description"`
	Location    Location `json:"location"`
	Tags        []string `json:"tags"`
	Media       []Media  `json:"media"`
}

type Location struct {
	Name        string      `json:"name"`
	Coordinates Coordinates `json:"coordinates"`
}

type Coordinates struct {
	Lat  string `json:"lat"`
	Long string `json:"long"`
}

type Media struct {
	ID  uuid.UUID `json:"id"`
	URL string    `json:"url"`
}

type ActivityService struct {
	repo Repository
}

type Repository interface {
	AddActivity(activity Activity) (Activity, error)
	GetActivities(filter SearchFilter) ([]Activity, error)
	GetActivity(activityID int) (Activity, error)
	TagActivity(activityID int, tags []string) error
}

type SearchFilter struct {
	ByCoordinates string   `json:"byCoordinates"`
	ByTags        []string `json:"byTags"`
}

func NewActivityService(repo Repository) ActivityService {
	return ActivityService{
		repo: repo,
	}
}

func (svc ActivityService) CreateActivity(activity Activity) (Activity, error) {
	return svc.repo.AddActivity(activity)
}

func (svc ActivityService) SearchActivities(filter SearchFilter) ([]Activity, error) {
	return svc.repo.GetActivities(filter)
}

func (svc ActivityService) GetActivityDetails(activityID int) (Activity, error) {
	return svc.repo.GetActivity(activityID)
}

func (svc ActivityService) TagActivity(activityID int, tags []string) error {
	return svc.repo.TagActivity(activityID, tags)
}
