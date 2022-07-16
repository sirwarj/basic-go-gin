package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sirwarj/go-api-gin/controllers"
	"github.com/sirwarj/go-api-gin/handle"
	"github.com/sirwarj/go-api-gin/models"
	"github.com/stretchr/testify/assert"
)

func TestAddData(t *testing.T) {
	routeTest := gin.Default()

	m := handle.NewMember()

	routeTest.GET("/member", controllers.AllData(m))

	routeTest.POST("/member", controllers.AddData(m))

	t.Run("Get Data", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/member", nil)
		res := httptest.NewRecorder()

		routeTest.ServeHTTP(res, req)

		response := res.Result()
		body, _ := ioutil.ReadAll(response.Body)

		result := models.Member{}
		json.Unmarshal(body, &result)

		expect := models.Member{}

		assert.Equal(t, expect, result)
	})

	t.Run("Add Data", func(t *testing.T) {
		data := `{
			"name":"John",
			"phone":"0987654321"
		}`

		req := httptest.NewRequest("POST", "/member", strings.NewReader(data))
		res := httptest.NewRecorder()

		routeTest.ServeHTTP(res, req)

		response := res.Result()
		body, _ := ioutil.ReadAll(response.Body)

		result := models.Member{}
		json.Unmarshal(body, &result)

		expect := models.Member{
			Name:  "John",
			Phone: "0987654321",
		}

		assert.Equal(t, expect, result)
	})
}
