package spentcalories

import (
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

	parsedData := strings.Split(data, ",")
	if len(parsedData) != 3 {
		return 0, "", 0, fmt.Errorf("incorrect Input")
	}

	steps, err := strconv.Atoi(parsedData[0])
	if err != nil {
		return 0, "", 0, err
	} else if steps <= 0 {
		return 0, "", 0, fmt.Errorf("количество шагов должно быть больше 0")
	}

	walkingTime, err := time.ParseDuration(parsedData[1])
	if err != nil {
		return 0, "", 0, err
	} else if walkingTime <= 0 {
		return 0, "", 0, fmt.Errorf("время должно быть больше 0")
	}

	return steps, parsedData[1], walkingTime, nil
}

func distance(steps int, height float64) float64 {

	stepLength := height * stepLengthCoefficient
	distance := stepLength * float64(steps) / mInKm

	return distance
}
func meanSpeed(steps int, height float64, duration time.Duration) float64 {

	if duration <= 0 {
		return 0
	}

	distance := distance(steps, height)

	return distance * duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {

	steps, trainingType, walkingTime, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	switch trainingType {
	case "Бег":
		distance, meanSpeed := distance(steps, height), meanSpeed(steps, height, walkingTime)
		spentCalories, err := RunningSpentCalories(steps, weight, height, walkingTime)
		if err != nil {
			return "", err
		}
		outputFormat := "Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n"
		output := fmt.Sprintf(outputFormat, trainingType, walkingTime, distance, meanSpeed, spentCalories)

		return output, nil
	case "Ходьба":
		distance, meanSpeed := distance(steps, height), meanSpeed(steps, height, walkingTime)
		spentCalories, err := WalkingSpentCalories(steps, weight, height, walkingTime)
		if err != nil {
			return "", err
		}
		outputFormat := "Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n"
		output := fmt.Sprintf(outputFormat, trainingType, walkingTime, distance, meanSpeed, spentCalories)

		return output, nil
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("incorrect input")
	}

	meanSpeed := meanSpeed(steps, height, duration)

	spentCalories := weight * meanSpeed * duration.Minutes() / minInH

	return spentCalories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("incorrect input")
	}

	meanSpeed := meanSpeed(steps, height, duration)

	spentCalories := weight * meanSpeed * duration.Minutes() / minInH * walkingCaloriesCoefficient

	return spentCalories, nil
}
