package controller

import (
	"fmt"
	"testing"

	"github.com/zytzjx/warehouse/models"
)

func TestFindAll(t *testing.T) {
	models.ConnectDatabase()
	var devices []models.Handsets
	models.DB.Find(&devices)
	fmt.Print(devices)
}
