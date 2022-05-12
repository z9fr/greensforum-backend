package events

import (
	"errors"

	"github.com/z9fr/greensforum-backend/internal/utils"
)

func (s *Service) CreateEvent(event Event) (Event, error) {

	event.Slug = utils.GenerateSlug(event.Name)
	// make sure slug is unique

	if s.IsEventSlugExist(event.Slug) {
		return Event{}, errors.New("Slug is already taken")
	}

	if event.Name == "" || event.Slug == "" {
		return Event{}, errors.New("Slug or Name can't be empty")
	}

	s.DB.Debug().Save(&event)
	return event, nil
}
