package course

import (
	"github.com/go-co-op/gocron"
	"math/rand"
	"time"
	"trading/services/db"
)

var min = 0.1

func generateCurrentCourseHistory() {
	goods := db.Goods()
	for _, good := range goods {
		value := newValue(good.CurrentCourse)
		good.CurrentCourse = value
		db.SaveGood(good)
	}
}

func generateCourseHistory() {
	goods := db.Goods()
	for _, good := range goods {
		db.SaveGoodDataPoint(good)
	}
}

func Init() {
	s := gocron.NewScheduler(time.UTC)

	_, err := s.Every(15).Seconds().Do(generateCurrentCourseHistory)
	if err != nil {
		panic(err)
	}
	_, err = s.Every(5).Minutes().Do(generateCourseHistory)
	if err != nil {
		panic(err)
	}
	s.StartAsync()
}

func newValue(oldValue float64) float64 {
	if oldValue == 0 {
		return 100 * rand.Float64()
	}
	var value = oldValue + rand.Float64() - 0.5

	if value < min {
		return min
	}
	if value < min*10 {
		return value * 2
	}
	return value
}
