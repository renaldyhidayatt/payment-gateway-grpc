package test

import (
	"MamangRust/paymentgatewaygrpc/internal/handler/api"
	mock_pb "MamangRust/paymentgatewaygrpc/internal/pb/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHandleHello(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockClient := mock_pb.NewMockAuthServiceClient(ctrl)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/auth/hello", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := api.NewHandlerAuth(mockClient, e, nil)

	err := handler.HandleHello(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Hello", rec.Body.String())
}
