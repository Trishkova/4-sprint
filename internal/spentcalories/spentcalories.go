package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	if data == "" {
		return 0, "", 0, errors.New("empty input")
	}

	arrData := strings.Split(data, ",")
	if len(arrData) != 3 {
		return 0, "", 0, errors.New("wrong input")
	}

	steps, err := strconv.Atoi(arrData[0])
	if err != nil {
		return 0, "", 0, err
	}

	if steps <= 0 {
		return 0, "", 0, errors.New("steps must be more 0")
	}

	err = nil
	duration, err := time.ParseDuration(arrData[2])
	if duration <= 0 {
		return 0, "", 0, errors.New("duration must be more 0")
	}
	if err != nil {
		return 0, "", 0, err
	}

	return steps, arrData[1], duration, nil
}

func distance(steps int, height float64) float64 {
	lengthStep := stepLengthCoefficient * height
	distance := (lengthStep * float64(steps)) / mInKm

	return distance
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	distance := distance(steps, height)

	hours := duration.Hours()
	avrSpeed := distance / hours

	return avrSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	if data == "" {
		log.Println("empty input")
		return "", errors.New("empty input")
	}

	arrData := strings.Split(data, ",")
	if len(arrData) < 3 {
		log.Println("wrong input")
		return "", errors.New("wrong input")
	}

	steps, err := strconv.Atoi(arrData[0])
	if err != nil {
		log.Println(err)
		return "", err
	}

	if steps <= 0 {
		log.Println("steps must be more 0")
		return "", errors.New("steps must be more 0")
	}

	trainingType := arrData[1]
	if trainingType == "" {
		log.Println("traning type cant be empty")
		return "", errors.New("traning type cant be empty")
	}

	err = nil
	duration, err := time.ParseDuration(arrData[2])
	if duration <= 0 {
		log.Println("duration must be more 0")
		return "", errors.New("duration must be more 0")
	}

	if err != nil {
		log.Println(err)
		return "", err
	}

	var calories float64
	switch trainingType {
	case "Бег":
		calories, _ = RunningSpentCalories(steps, weight, height, duration)
	case "Ходьба":
		calories, _ = WalkingSpentCalories(steps, weight, height, duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	distance := distance(steps, height)
	speed := meanSpeed(steps, height, duration)
	time := duration.Hours()

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, time, distance, speed, calories), nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("empty input")
	}

	avrSpeed := meanSpeed(steps, height, duration)
	calories := (duration.Minutes() * weight * avrSpeed) / minInH

	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("wrong input")
	}

	avrSpeed := meanSpeed(steps, height, duration)
	calories := ((duration.Minutes() * weight * avrSpeed) / minInH) * walkingCaloriesCoefficient

	return calories, nil

}
