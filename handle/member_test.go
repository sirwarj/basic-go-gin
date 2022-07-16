package handle

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sirwarj/go-api-gin/models"
)

func TestAllData(t *testing.T) {
	m := NewMember()

	result := m.AllData()

	var expect []models.Member

	assert.Equal(t, expect, result)
}

func TestAddData(t *testing.T) {
	m := NewMember()

	data := models.Member{
		Name:  "John",
		Phone: "0987654321",
	}

	m.AddData(data)

	result := m.AllData()

	var expect []models.Member
	expect = append(expect, data)

	assert.Equal(t, expect, result)

}
