package service

import "sync"

// Rating contains the information of a laptop
type Rating struct {
	Count uint32  // the number of times the laptop is rated
	Sum   float64 // the sum of all rated scores
}

// RatingStore is an interface to store laptop ratings
type RatingStore interface {
	// Add adds a new laptop score to the store and return its rating
	Add(laptopId string, score float64) (*Rating, error)
}

type InMemoryRatingStore struct {
	mutex sync.RWMutex
	// key is laptop id, value is the Rating object
	rating map[string]*Rating
}

func (store *InMemoryRatingStore) Add(laptopId string, score float64) (*Rating, error) {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	rating := store.rating[laptopId]
	if rating == nil {
		rating = &Rating{
			Count: 1,
			Sum:   score,
		}
	} else {
		rating.Count++
		rating.Sum += score
	}

	store.rating[laptopId] = rating
	return rating, nil
}

func NewInMemoryRatingStore() *InMemoryRatingStore {
	return &InMemoryRatingStore{
		rating: make(map[string]*Rating),
	}
}
