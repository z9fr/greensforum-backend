package events

import "github.com/z9fr/greensforum-backend/internal/utils"

func (s *Service) GetAllEvents() []Event {
	var events []Event
	s.DB.Debug().Order("created_at DESC").Find(&events)
	return events
}

func (s *Service) GeteventbySlug(slug string) *Event {
	// slug

	var event *Event
	// only get the accepted posts
	s.DB.Debug().Order("created_at DESC").Where("slug = ?", slug).First(&event)
	return event
}

func (s *Service) IsEventSlugExist(slug string) bool {
	var exists bool
	if err := s.DB.Debug().Model(&Event{}).
		Select("count(*) > 0").
		Where("slug = ?", slug).
		Find(&exists).
		Error; err != nil {
		utils.LogWarn(err)
	}

	return exists
}
