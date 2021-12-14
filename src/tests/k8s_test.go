package tests

import (
	"fmt"
	"k8s.io/apimachinery/pkg/api/resource"
	"testing"
)

func TestResourceQuantity(t *testing.T) {

	fmt.Println(resource.ParseQuantity("15Mi"))
}
