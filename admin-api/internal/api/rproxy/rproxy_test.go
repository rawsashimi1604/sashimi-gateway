package rproxy

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"testing"

	gs "github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/gateway/service"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TODO: update tests for mock analytics tracker.

// mockTransport is a mocked object that mocks the http.RoundTripper, an interface http.Client uses to send the requests.
type mockTransport struct{}

func (t *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBufferString("OK")),
		Header:     make(http.Header),
	}, nil
}

// MockServiceGateway is a mocked object that implements the ServiceGateway interface
type MockServiceGateway struct {
	mock.Mock
}

func (m *MockServiceGateway) GetServiceByPath(path string) (models.Service, error) {
	args := m.Called(path)
	return args.Get(0).(models.Service), args.Error(1)
}

func (m *MockServiceGateway) GetServiceByTargetUrl(targetUrl string) (models.Service, error) {
	args := m.Called(targetUrl)
	return args.Get(0).(models.Service), args.Error(1)
}

func (m *MockServiceGateway) GetAllServices() ([]models.Service, error) {
	args := m.Called()
	return args.Get(0).([]models.Service), args.Error(1)
}

func (m *MockServiceGateway) RegisterService(service models.Service) (models.Service, error) {
	args := m.Called()
	return args.Get(0).(models.Service), args.Error(1)
}

func TestMatchServiceNotFound(t *testing.T) {
	mockServiceGateway := new(MockServiceGateway)
	rps := NewReverseProxy(mockServiceGateway, &mockTransport{})

	// Mocking the service not found error
	mockServiceGateway.On("GetServiceByPath", "nonexistent").Return(models.Service{}, gs.ErrServiceNotFound)

	service, err := rps.matchService("/nonexistent/path")
	assert.True(t, errors.Is(err, gs.ErrServiceNotFound))
	assert.Equal(t, models.Service{}, service)
}

func TestReverseProxyMiddleware(t *testing.T) {
	mockServiceGateway := new(MockServiceGateway)
	rps := NewReverseProxy(mockServiceGateway, &mockTransport{})

	// Create a test service and route
	testService := models.Service{
		Path:      "/testservice",
		TargetUrl: "http://backend-service.com",
		Routes:    []models.Route{{Path: "/"}, {Path: "/testroute"}},
	}

	t.Run("valid request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/testservice/testroute", bytes.NewBuffer([]byte{}))
		rr := httptest.NewRecorder()

		mockServiceGateway.On("GetServiceByPath", "testservice").Return(testService, nil)

		// Set up proxy transport before forwarding request
		origin, _ := url.Parse(testService.TargetUrl)
		proxy := httputil.NewSingleHostReverseProxy(origin)
		proxy.Transport = &mockTransport{}
		proxy.ErrorHandler = func(w http.ResponseWriter, req *http.Request, err error) {
			log.Info().Msgf("Error while proxying request: %v", err)
			http.Error(w, "Error while proxying request", http.StatusInternalServerError)
		}

		// Apply the middleware to a test route and invoke it
		rps.ReverseProxyMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})).ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("service not found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/invalidservice/testroute", bytes.NewBuffer([]byte{}))
		rr := httptest.NewRecorder()

		mockServiceGateway.On("GetServiceByPath", "invalidservice").Return(models.Service{}, gs.ErrServiceNotFound)

		origin, _ := url.Parse(testService.TargetUrl)
		proxy := httputil.NewSingleHostReverseProxy(origin)
		proxy.Transport = &mockTransport{}
		rps.ReverseProxyMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})).ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
	})

}
