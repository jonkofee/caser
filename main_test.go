package main

import "testing"

func TestUpperCamelToSnake(t *testing.T) {
	result := CamelToSnake("UpdatedAt")
	if result != "updated_at" {
		t.Error("Incorrect UpperCamelCase to snake_case, got ", result)
	}
}

func TestLowerCamelToSnake(t *testing.T) {
	result := CamelToSnake("updatedAt")
	if CamelToSnake("updatedAt") != "updated_at" {
		t.Error("Incorrect lowerCamelCase to snake_case, got ", result)
	}
}

func TestUpperCamelToSnakeWithNumber(t *testing.T) {
	result := CamelToSnake("UpdatedAt420")
	if CamelToSnake("UpdatedAt420") != "updated_at_420" {
		t.Error("Incorrect UpperCamelCase to snake_case with number in end, got ", result)
	}
}

func TestLowerCamelToSnakeWithNumber(t *testing.T) {
	result := CamelToSnake("updatedAt420")
	if CamelToSnake("updatedAt420") != "updated_at_420" {
		t.Error("Incorrect lowerCamelCase to snake_case with number in end, got ", result)
	}
}

func TestUpperCamelToSnakeStartWithNumber(t *testing.T) {
	result := CamelToSnake("15UpdatedAt420")
	if CamelToSnake("15UpdatedAt420") != "15_updated_at_420" {
		t.Error("Incorrect UpperCamelCase to snake_case with number in begin, got ", result)
	}
}

func TestLowerCamelToSnakeStartWithNumber(t *testing.T) {
	result := CamelToSnake("15updatedAt420")
	if CamelToSnake("15updatedAt420") != "15_updated_at_420" {
		t.Error("Incorrect lowerCamelCase to snake_case with number in begin, got ", result)
	}
}

func TestUpperCamelToSnakeWithNumberInMiddle(t *testing.T) {
	result := CamelToSnake("15UpdatedAt420")
	if CamelToSnake("15UpdatedAt420") != "15_updated_at_420" {
		t.Error("Incorrect UpperCamelCase to snake_case with number in middle, got ", result)
	}
}

func TestLowerCamelToSnakeWithNumberInMiddle(t *testing.T) {
	result := CamelToSnake("15updated8At420")
	if CamelToSnake("15updated8At420") != "15_updated_8_at_420" {
		t.Error("Incorrect lowerCamelCase to snake_case with number in middle, got ", result)
	}
}
