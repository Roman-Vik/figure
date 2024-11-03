package decimal

import (
	"fmt"
	"strconv"
)

func DecimalSum(a, b string) (string, error) {
	first, err := strconv.ParseFloat(a, 32)
	if err != nil {
		return "", fmt.Errorf("ошибка преобразования первого числа: %w", err)
	}

	second, err := strconv.ParseFloat(b, 32)
	if err != nil {
		return "", fmt.Errorf("ошибка преобразования второго числа: %w", err)
	}

	// Преобразуем сумму обратно в строку с нужным форматом
	return fmt.Sprintf("%.2f", first+second), nil // Выводим с двумя знаками после запятой
}

func DecimalSubtract(a, b string) (string, error) {
	first, err := strconv.ParseFloat(a, 32)
	if err != nil {
		return "", err
	}

	second, err := strconv.ParseFloat(b, 32)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%.2f", first-second), nil
}

func DecimalMultiply(a, b string) (string, error) {
	first, err := strconv.ParseFloat(a, 32)
	if err != nil {
		return "", err
	}

	second, err := strconv.ParseFloat(b, 32)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%.2f", first*second), nil
}

func DecimalDivide(a, b string) (string, error) {
	first, err := strconv.ParseFloat(a, 64)
	fmt.Println(first)
	if err != nil {
		return "", err
	}
	second, err := strconv.ParseFloat(b, 64)
	fmt.Println(second)
	if err != nil {
		return "", err
	}

	// Проверка деления на ноль
	if second == 0.0 || first == 0.0 {
		return "", fmt.Errorf("на ноль делить нельзя")
	}
	return fmt.Sprintf("%.2f", first/second), nil
}

func DecimalRound(a string, precision int32) (string, error) {
	first, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return "", err
	}

	// Задаем динамическое количество знаков после запятой
	format := fmt.Sprintf("%%.%df", precision) // Создаем строку формата динамически

	// Форматируем число
	r := fmt.Sprintf(format, first) // Используем динамически созданный формат
	return r, nil                   // Правильная подстановка переменной r
}

func DecimalGreaterThan(a, b string) (bool, error) {
	first, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return false, fmt.Errorf("ошибка преобразования первого числа: %w", err)
	}

	second, err := strconv.ParseFloat(b, 64)
	if err != nil {
		return false, fmt.Errorf("ошибка преобразования второго числа: %w", err)
	}

	return first > second, nil
}
func DecimalLessThan(a, b string) (bool, error) {
	first, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return false, fmt.Errorf("ошибка преобразования первого числа: %w", err)
	}

	second, err := strconv.ParseFloat(b, 64)
	if err != nil {
		return false, fmt.Errorf("ошибка преобразования второго числа: %w", err)
	}

	return first < second, nil
}

func DecimalEqual(a, b string) (bool, error) {
	first, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return false, fmt.Errorf("ошибка преобразования первого числа: %w", err)
	}

	second, err := strconv.ParseFloat(b, 64)
	if err != nil {
		return false, fmt.Errorf("ошибка преобразования второго числа: %w", err)
	}

	return first == second, nil
}

