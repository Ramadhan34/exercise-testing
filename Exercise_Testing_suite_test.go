package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestExerciseTesting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ExerciseTesting Suite")
}
