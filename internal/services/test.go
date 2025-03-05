package test

import (
	"testing"
)

func IsTesting() bool {
	return testing.Short()
}
