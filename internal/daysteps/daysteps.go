package daysteps

import (
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

	parsedData := strings.Split(data, ",")
	if len(parsedData) != 2 {
		return 0, 0, fmt.Errorf("incorrect Input")
	}

	steps, err := strconv.Atoi(parsedData[0])
	if err != nil {
		return 0, 0, err
	} else if steps <= 0 {
		return 0, 0, fmt.Errorf("количество шагов должно быть больше 0")
	}

	walkingTime, err := time.ParseDuration(parsedData[1])
	if err != nil {
		return 0, 0, err
	} else if walkingTime <= 0 {
		return 0, 0, fmt.Errorf("время должно быть больше 0")
	}

	return steps, walkingTime, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, walkingTime, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}

	distance := float64(steps) * stepLength / mInKm

	walkingSpentCalories, err := spentcalories.WalkingSpentCalories(steps, weight, height, walkingTime)
	if err != nil {
		log.Println(err)
		return ""
	}

	outputFormat := "Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n"
	output := fmt.Sprintf(outputFormat, steps, distance, walkingSpentCalories)

	return output
}
