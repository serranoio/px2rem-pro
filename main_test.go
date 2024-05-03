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
		doNotInclude:     "padding, font-size",
	})
}

func TestCheckInclusion(t *testing.T) {
	config := createConfig(.1, "padding, font-size")

	begStr := "font-size: 16px;"
	str := checkInclusion(config, begStr)

	assert.Equal(t, begStr, str)
}

func TestConvertPxToRem(t *testing.T) {
	line := "padding: 20px 20px 20px 20px;"
	config := createConfig(.5, "padding")
	newLine := convertPxToRem(config, line)
	assert.Equal(t, "padding: 10.0rem 10.0rem 10.0rem 10.0rem;", newLine)

	line = "transform: translate(20px);"
	config = createConfig(.5, "padding")
	newLine = convertPxToRem(config, line)
	assert.Equal(t, "transform: translate(10.0rem);", newLine)

	line = "transform: translate(20px, 20px);"
	config = createConfig(.5, "padding")
	newLine = convertPxToRem(config, line)
	assert.Equal(t, "transform: translate(10.0rem, 10.0rem);", newLine)

	line = "box-shadow: 10px 10px 10px rgba(0, 0, 0, 0.5);"
	config = createConfig(.5, "padding")
	newLine = convertPxToRem(config, line)
	assert.Equal(t, "box-shadow: 5.0rem 5.0rem 5.0rem rgba(0, 0, 0, 0.5);", newLine)

	line = "outline: 2px solid blue;"
	config = createConfig(.5, "padding")
	newLine = convertPxToRem(config, line)
	assert.Equal(t, "outline: 1.0rem solid blue;", newLine)

	line = "height: 100px;"
	config = createConfig(.5, "padding")
	newLine = convertPxToRem(config, line)
	assert.Equal(t, "height: 50.0rem;", newLine)

	line = "width: 1000px;"
	config = createConfig(.5, "padding")
	newLine = convertPxToRem(config, line)
	assert.Equal(t, "width: 500.0rem;", newLine)

	line = "outline: 2px solid blue;"
	config = createConfig(.5, "padding")
	newLine = convertPxToRem(config, line)
	assert.Equal(t, "outline: 1.0rem solid blue;", newLine)
}

func TestParseContents(t *testing.T) {
	str := `
	.example {
		width: 200px;
		height: 100px;
		min-width: 50px;
		max-width: 300px;
		min-height: 50px;
		max-height: 200px;
		margin: 10px;
		margin-top: 20px;
		margin-right: 30px;
		margin-bottom: 40px;
		margin-left: 50px;
		padding: 5px;
		padding-top: 15px;
		padding-right: 25px;
		padding-bottom: 35px;
		padding-left: 45px;
		font-size: 16px;
		line-height: 20px;
		border: 1px solid black;
		border-width: 2px;
		border-top-width: 3px;
		border-right-width: 4px;
		border-bottom-width: 5px;
		border-left-width: 6px;
		border-radius: 10px;
		outline: 2px solid blue;
		outline-offset: 5px;
		box-shadow: 5px 5px 5px rgba(0, 0, 0, 0.5);
		text-shadow: 2px 2px 2px rgba(0, 0, 0, 0.5);
		inset: 10px;
		inset-block: 15px;
		inset-block-start: 20px;
		inset-block-end: 25px;
		inset-inline: 30px;
		inset-inline-start: 35px;
		inset-inline-end: 40px;
		top: 10px;
		right: 20px;
		bottom: 30px;
		left: 40px;
		grid-gap: 10px;
		grid-row-gap: 20px;
		grid-column-gap: 30px;
		gap: 10px;
		row-gap: 20px;
		column-gap: 30px;
		column-rule-width: 5px;
		flex-basis: 200px;
		flex-grow: 1;
		flex-shrink: 1;
		transform: translate(50px, 50px);
		translate: translate(50px, 50px);
		scroll-margin: 50px;
		scroll-margin-block: 60px;
		scroll-margin-block-start: 70px;
		scroll-margin-block-end: 80px;
		scroll-margin-inline: 90px;
		scroll-margin-inline-start: 100px;
		scroll-margin-inline-end: 110px;
		offset: 10px;
		offset-block: 20px;
		offset-block-start: 30px;
		offset-block-end: 40px;
		offset-inline: 50px;
		offset-inline-start: 60px;
		offset-inline-end: 70px;
	}`
	config := createConfig(.25, "")

	parseContents(config, str)
}
