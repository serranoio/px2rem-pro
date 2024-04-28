package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func createConfig(factor float64, doNotInclude string) config {

	return config{
		conversionFactor: factor,
		doNotInclude:     doNotInclude,
	}
}

func TestMain(t *testing.T) {
	charmInterface(config{
		conversionFactor: 5.0,
		doNotInclude:     "padding",
	})
}

func TestConvertPxToRem(t *testing.T) {
	line := "padding: 20px 20px 20px 20px;"

	config := createConfig(.5, "padding")

	newLine := convertPxToRem(config, line)

	assert.Equal(t, "padding: 10.0rem 10.0rem 10.0rem 10.0rem;", newLine)
}
