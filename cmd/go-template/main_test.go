package main

import (
	"testing"
)

func TestGetSensorValue(t *testing.T) {
	expected := exampleSensorValue
	actual := getSensorValue("sensor1")

	t.Errorf("expected %v, but got %v", expected, actual)
	if actual != expected {
		t.Errorf("expected %v, but got %v", expected, actual)
	}
}
