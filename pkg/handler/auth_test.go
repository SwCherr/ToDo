package handler

import (
	"app"
	mock_service "app/mocks/service"
	"app/pkg/service"
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
)

func TestHandler_SignUp(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *mock_service.MockAuthorization, user app.User)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            app.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"email": "email", "password": "qwerty"}`,
			inputUser: app.User{
				Email:    "email",
				Password: "qwerty",
			},
			mockBehavior: func(r *mock_service.MockAuthorization, user app.User) {
				r.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:                 "Wrong Input",
			inputBody:            `{"email": "email"}`,
			inputUser:            app.User{},
			mockBehavior:         func(r *mock_service.MockAuthorization, user app.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Service Error",
			inputBody: `{"email": "email", "password": "qwerty"}`,
			inputUser: app.User{
				Email:    "email",
				Password: "qwerty",
			},
			mockBehavior: func(r *mock_service.MockAuthorization, user app.User) {
				r.EXPECT().CreateUser(user).Return(0, errors.New("error on server's side"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"error on server's side"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockAuthorization(c)
			test.mockBehavior(repo, test.inputUser)

			services := &service.Service{Authorization: repo}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.POST("/signup", handler.SignUp)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/signup",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_GetPareTokens(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *mock_service.MockAuthorization, input app.Sesion)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            app.Sesion
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"id": 1, "guid": "1", "refreshToken": "mafmaf"}`,
			inputUser: app.Sesion{
				UserID:       1,
				GUID:         "1",
				RefreshToken: "mafmaf",
			},
			mockBehavior: func(r *mock_service.MockAuthorization, input app.Sesion) {
				r.EXPECT().GetPareToken(input).Return("", "", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"accessToken":"","refreshToken":""}`,
		},
		{
			name:                 "Wrong Input",
			inputBody:            `{"email": "email"}`,
			inputUser:            app.Sesion{},
			mockBehavior:         func(r *mock_service.MockAuthorization, input app.Sesion) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Service Error",
			inputBody: `{"id": 1, "guid": "1", "refreshToken": "mafmaf"}`,
			inputUser: app.Sesion{
				UserID:       1,
				GUID:         "1",
				RefreshToken: "mafmaf",
			},
			mockBehavior: func(r *mock_service.MockAuthorization, input app.Sesion) {
				r.EXPECT().GetPareToken(input).Return("", "", errors.New("error on server's side"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"error on server's side"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockAuthorization(c)
			test.mockBehavior(repo, test.inputUser)

			services := &service.Service{Authorization: repo}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.POST("/get", handler.GetPareTokens)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/get",
				bytes.NewBufferString(test.inputBody))

			// set IP
			req.RemoteAddr = ""

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_RefreshToken(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *mock_service.MockAuthorization, input app.Sesion)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            app.Sesion
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"id": 1, "guid": "1", "refreshToken": "mafmaf"}`,
			inputUser: app.Sesion{
				UserID:       1,
				GUID:         "1",
				RefreshToken: "mafmaf",
			},
			mockBehavior: func(r *mock_service.MockAuthorization, input app.Sesion) {
				r.EXPECT().RefreshToken(input).Return("", "", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"accessToken":"","refreshToken":""}`,
		},
		{
			name:                 "Wrong Input",
			inputBody:            `{"email": "email"}`,
			inputUser:            app.Sesion{},
			mockBehavior:         func(r *mock_service.MockAuthorization, input app.Sesion) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Service Error",
			inputBody: `{"id": 1, "guid": "1", "refreshToken": "mafmaf"}`,
			inputUser: app.Sesion{
				UserID:       1,
				GUID:         "1",
				RefreshToken: "mafmaf",
			},
			mockBehavior: func(r *mock_service.MockAuthorization, input app.Sesion) {
				r.EXPECT().RefreshToken(input).Return("", "", errors.New("error on server's side"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"error on server's side"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockAuthorization(c)
			test.mockBehavior(repo, test.inputUser)

			services := &service.Service{Authorization: repo}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.POST("/refresh", handler.RefreshToken)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/refresh",
				bytes.NewBufferString(test.inputBody))

			// set IP
			req.RemoteAddr = ""

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
