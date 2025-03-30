package spentcalories

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                            = 0.65  // средняя длина шага.
	mInKm                              = 1000  // количество метров в километре.
	minInH                             = 60    // количество минут в часе.
	runningCaloriesMeanSpeedMultiplier = 18.0  // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 20.0  // среднее количество сжигаемых калорий при беге.
	walkingCaloriesWeightMultiplier    = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier       = 0.029 // множитель роста.
)

var (
	errConversion = errors.New("conversion error")
	errZeroSteps  = errors.New("steps count is zero")
	errLessData   = errors.New("data is not full")
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// ваш код ниже
	parsed := strings.Split(data, ",")
	if len(parsed) != 3 {
		return 0, "", 0, fmt.Errorf("error: %w", errLessData)
	}
	steps, err := strconv.Atoi(parsed[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("%w: %w", errConversion, err)
	}
	if steps <= 0 {
		return 0, "", 0, fmt.Errorf("error: %w", errZeroSteps)
	}

	activity := parsed[1]

	duration, err := time.ParseDuration(parsed[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("%w: %w", errConversion, err)
	}

	return steps, activity, duration, nil
}

// distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий (число шагов при ходьбе и беге).
func distance(steps int) float64 {
	// ваш код ниже
	dist := (float64(steps) * lenStep) / mInKm
	return dist
}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий(число шагов при ходьбе и беге).
// duration time.Duration — длительность тренировки.
func meanSpeed(steps int, duration time.Duration) float64 {
	// ваш код ниже
	if duration <= 0 {
		return 0
	}
	dist := distance(steps)

	speed := dist / duration.Hours()

	return speed
}

// RunningSpentCalories возвращает количество потраченных колорий при беге.
//
// Параметры:
//
// steps int - количество шагов.
// weight float64 — вес пользователя.
// duration time.Duration — длительность тренировки.
func RunningSpentCalories(steps int, weight float64, duration time.Duration) float64 {
	// ваш код здесь
	speed := meanSpeed(steps, duration)

	calories := ((runningCaloriesMeanSpeedMultiplier * speed) - runningCaloriesMeanSpeedShift) * weight

	return calories
}

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
//
// Параметры:
//
// steps int - количество шагов.
// duration time.Duration — длительность тренировки.
// weight float64 — вес пользователя.
// height float64 — рост пользователя.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) float64 {
	// ваш код здесь
	speed := meanSpeed(steps, duration)

	calories := ((walkingCaloriesWeightMultiplier * weight) + (speed*speed/height)*walkingSpeedHeightMultiplier) * duration.Hours() * minInH

	return calories
}

// ShowTrainingInfo возвращает строку с информацией о тренировке.
//
// Параметры:
//
// data string - строка с данными.
// weight, height float64 — вес и рост пользователя.
func TrainingInfo(data string, weight, height float64) string {
	// ваш код ниже
	//	parsed := strings.Split(data, ",")
	steps, trainingType, duration, err := parseTraining(data)
	if err != nil {
		return fmt.Sprintf("%v\n", err)
	}
	/*	steps, err := strconv.Atoi(parsed[0])
		if err != nil {
			return fmt.Sprint(err)
		}

		trainingType := parsed[1]

		duration, err := time.ParseDuration(parsed[2])
		if err != nil {
			return fmt.Sprint(err)
		}
	*/
	dist := distance(steps)
	speed := meanSpeed(steps, duration)

	var calories float64
	switch trainingType {
	case "Бег":
		calories = RunningSpentCalories(steps, weight, duration)
	case "Ходьба":
		calories = WalkingSpentCalories(steps, weight, height, duration)
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration.Hours(), dist, speed, calories)
}
