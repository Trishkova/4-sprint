package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	if data == "" {
		return 0, 0, errors.New("empty input")
	}

	arrData := strings.Split(data, ",")
	if len(arrData) != 2 {
		return 0, 0, errors.New("wrong input")
	}

	steps, err := strconv.Atoi(arrData[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid steps format: %w", err)
	}

	if steps <= 0 {
		return 0, 0, errors.New("steps must be more 0")
	}

	err = nil
	duration, err := time.ParseDuration(arrData[1])
	if duration <= 0 {
		return 0, 0, errors.New("duration must be more 0")
	}
	if err != nil {
		return 0, 0, err
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}

	if steps <= 0 {
		return ""
	}

	distance := (float64(steps) * stepLength) / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	if err != nil {
		log.Println(err)
		return ""
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distance, calories)
}
