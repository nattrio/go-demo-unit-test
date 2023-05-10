package services_test

import (
	"testing"

	"github.com/nattrio/go-demo-unit-test/services"
)

func TestCheckGradeA(t *testing.T) {
	grade := services.CheckGrade(80)
	expected := "A"
	if grade != expected {
		t.Errorf("Expected %s, got %s", expected, grade)
	}
}

func TestCheckGradeB(t *testing.T) {
	grade := services.CheckGrade(70)
	expected := "B"
	if grade != expected {
		t.Errorf("Expected %s, got %s", expected, grade)
	}
}

func TestCheckGradeC(t *testing.T) {
	grade := services.CheckGrade(60)
	expected := "C"
	if grade != expected {
		t.Errorf("Expected %s, got %s", expected, grade)
	}
}
