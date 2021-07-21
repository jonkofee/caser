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
	if caser.CamelToSnake("updatedAt") != "updated_at" {
		t.Error("Incorrect lowerCamelCase to snake_case, got ", result)
	}
}

func TestUpperCamelToSnakeWithNumber(t *testing.T) {
	result := caser.CamelToSnake("UpdatedAt420")
	if caser.CamelToSnake("UpdatedAt420") != "updated_at_420" {
		t.Error("Incorrect UpperCamelCase to snake_case with number in end, got ", result)
	}
}

func TestLowerCamelToSnakeWithNumber(t *testing.T) {
	result := caser.CamelToSnake("updatedAt420")
	if caser.CamelToSnake("updatedAt420") != "updated_at_420" {
		t.Error("Incorrect lowerCamelCase to snake_case with number in end, got ", result)
	}
}

func TestUpperCamelToSnakeStartWithNumber(t *testing.T) {
	result := caser.CamelToSnake("15UpdatedAt420")
	if caser.CamelToSnake("15UpdatedAt420") != "15_updated_at_420" {
		t.Error("Incorrect UpperCamelCase to snake_case with number in begin, got ", result)
	}
}

func TestLowerCamelToSnakeStartWithNumber(t *testing.T) {
	result := caser.CamelToSnake("15updatedAt420")
	if caser.CamelToSnake("15updatedAt420") != "15_updated_at_420" {
		t.Error("Incorrect lowerCamelCase to snake_case with number in begin, got ", result)
	}
}

func TestUpperCamelToSnakeWithNumberInMiddle(t *testing.T) {
	result := caser.CamelToSnake("15UpdatedAt420")
	if caser.CamelToSnake("15UpdatedAt420") != "15_updated_at_420" {
		t.Error("Incorrect UpperCamelCase to snake_case with number in middle, got ", result)
	}
}

func TestLowerCamelToSnakeWithNumberInMiddle(t *testing.T) {
	result := caser.CamelToSnake("15updated8At420")
	if caser.CamelToSnake("15updated8At420") != "15_updated_8_at_420" {
		t.Error("Incorrect lowerCamelCase to snake_case with number in middle, got ", result)
	}
}
