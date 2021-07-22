package test

import (
	"github.com/jonkofee/caser"
	"testing"
)

func TestUpperCamelToSnake(t *testing.T) {
	result := caser.CamelToSnake("UpdatedAt")
	if result != "updated_at" {
		t.Error("Incorrect UpperCamelCase to snake_case, got ", result)
	}
}

func TestLowerCamelToSnake(t *testing.T) {
	result := caser.CamelToSnake("updatedAt")
	if result != "updated_at" {
		t.Error("Incorrect lowerCamelCase to snake_case, got ", result)
	}
}

func TestUpperCamelToSnakeWithNumber(t *testing.T) {
	result := caser.CamelToSnake("UpdatedAt420")
	if result != "updated_at_420" {
		t.Error("Incorrect UpperCamelCase to snake_case with number in end, got ", result)
	}
}

func TestLowerCamelToSnakeWithNumber(t *testing.T) {
	result := caser.CamelToSnake("updatedAt420")
	if result != "updated_at_420" {
		t.Error("Incorrect lowerCamelCase to snake_case with number in end, got ", result)
	}
}

func TestUpperCamelToSnakeStartWithNumber(t *testing.T) {
	result := caser.CamelToSnake("15UpdatedAt420")
	if result != "15_updated_at_420" {
		t.Error("Incorrect UpperCamelCase to snake_case with number in begin, got ", result)
	}
}

func TestLowerCamelToSnakeStartWithNumber(t *testing.T) {
	result := caser.CamelToSnake("15updatedAt420")
	if result != "15updated_at_420" {
		t.Error("Incorrect lowerCamelCase to snake_case with number in begin, got ", result)
	}
}

func TestUpperCamelToSnakeWithNumberInMiddle(t *testing.T) {
	result := caser.CamelToSnake("15UpdatedAt420")
	if result != "15_updated_at_420" {
		t.Error("Incorrect UpperCamelCase to snake_case with number in middle, got ", result)
	}
}

func TestLowerCamelToSnakeWithNumberInMiddle(t *testing.T) {
	result := caser.CamelToSnake("15updated8At420")
	if result != "15updated_8_at_420" {
		t.Error("Incorrect lowerCamelCase to snake_case with number in middle, got ", result)
	}
}

func TestLowerCamelToSnakeSingleChar(t *testing.T) {
	result := caser.CamelToSnake("a")
	if result != "a" {
		t.Error("Incorrect lowerCamelCase to snake_case with single char, got ", result)
	}
}

func TestUpperCamelToSnakeSingleChar(t *testing.T) {
	result := caser.CamelToSnake("A")
	if result != "a" {
		t.Error("Incorrect lowerCamelCase to snake_case with single char, got ", result)
	}
}

func TestLowerCamelToSnakeWithUpperWord(t *testing.T) {
	result := caser.CamelToSnake("reserveUSD")
	if result != "reserve_usd" {
		t.Error("Incorrect lowerCamelCase to snake_case with upper word, got ", result)
	}
}

func TestUpperCamelToSnakeWithUpperWord(t *testing.T) {
	result := caser.CamelToSnake("ReserveUSD")
	if result != "reserve_usd" {
		t.Error("Incorrect lowerCamelCase to snake_case with upper word, got ", result)
	}
}
