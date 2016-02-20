package main

import (
	"testing"

	"github.com/javierprovecho/ci-example/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"github.com/javierprovecho/ci-example/Godeps/_workspace/src/github.com/stretchr/testify/assert"
)

func TestCatsAndDogs(t *testing.T) {
	if cats && dogs {
		t.Error("there are cats with dogs in the scene!")
	}
}

func TestCreateRouter(t *testing.T) {
	router := createRouter()
	list := router.Routes()

	assert.Len(t, list, 1)

	assertRoutePresent(t, list, gin.RouteInfo{
		Method: "GET",
		Path:   "/",
	})
}

func assertRoutePresent(t *testing.T, gotRoutes gin.RoutesInfo, wantRoute gin.RouteInfo) {
	for _, gotRoute := range gotRoutes {
		if gotRoute.Path == wantRoute.Path && gotRoute.Method == wantRoute.Method {
			assert.Regexp(t, wantRoute.Path, gotRoute.Path)
			return
		}
	}
	t.Errorf("route not found: %v", wantRoute)
}
