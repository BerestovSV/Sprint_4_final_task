package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-4-sprint-final/internal/spentcalories"
)

var (
	StepLength    = 0.65 // длина шага в метрах
	ErrConversion = errors.New("conversion error")
	ErrZeroSteps  = errors.New("steps count is zero")
	ErrLessData   = errors.New("data is not full")
)

func parsePackage(data string) (int, time.Duration, error) {
	// ваш код ниже
	parsed := strings.Split(data, ",")

	if len(parsed) != 2 {
		return 0, 0, fmt.Errorf("error: %w", ErrLessData)
	}

	steps, err := strconv.Atoi(parsed[0])
	if err != nil {
		return 0, 0, fmt.Errorf("%w: %w", ErrConversion, err)
	}
	if steps <= 0 {
		return 0, 0, fmt.Errorf("error: %w", ErrZeroSteps)
	}

	duration, err := time.ParseDuration(parsed[1])
	if err != nil {
		return 0, 0, fmt.Errorf("%w: %w", ErrConversion, err)
	}

	return steps, duration, nil
}

// DayActionInfo обрабатывает входящий пакет, который передаётся в
// виде строки в параметре data. Параметр storage содержит пакеты за текущий день.
// Если время пакета относится к новым суткам, storage предварительно
// очищается.
// Если пакет валидный, он добавляется в слайс storage, который возвращает
// функция. Если пакет невалидный, storage возвращается без изменений.
func DayActionInfo(data string, weight, height float64) string {
	// ваш код ниже
	steps, duration, err := parsePackage(data)
	if err != nil {
		return fmt.Sprintf("%v\n", err)
	}

	if steps <= 0 {
		return " \n"
	}

	distance := (float64(steps) * StepLength) / 1000

	calories := spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distance, calories)
}
