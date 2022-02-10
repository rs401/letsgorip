package middlewares

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/rs401/letsgorip/authapi/tokenutils"
)

func TestAuthMiddleware(t *testing.T) {
	// HandlerFunc for 'next'
	myNext := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {})

	// Good token
	tokens, err := tokenutils.CreateToken(1)
	if err != nil {
		t.Errorf("Error creating tokens: %v", err)
	}

	// Bad token
	claims := &tokenutils.Claims{UserId: 1}
	claims.ExpiresAt = time.Now().Add(-24 * time.Hour).Unix()
	bt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	btStr, err := bt.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		t.Errorf("Error signing token: %v\n", err)
	}

	// rw's
	rwBad := httptest.NewRecorder()
	rwGood := httptest.NewRecorder()

	// Tests
	tests := []struct {
		name    string
		token   string
		rw      http.ResponseWriter
		r       *http.Request
		wantErr bool
	}{
		{
			name:    "Bad token",
			token:   fmt.Sprintf("Bearer %s", btStr),
			rw:      rwBad,
			r:       httptest.NewRequest(http.MethodGet, "/api/", nil),
			wantErr: true,
		},
		{
			name:    "Good token",
			token:   fmt.Sprintf("Bearer %s", tokens.AccessToken),
			rw:      rwGood,
			r:       httptest.NewRequest(http.MethodGet, "/api/", nil),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.Header.Set("Authorization", tt.token)
			AuthMiddleware(myNext).ServeHTTP(tt.rw, tt.r)
			if (tt.rw.(*httptest.ResponseRecorder).Code != http.StatusOK) != tt.wantErr {
				t.Errorf("Bad status code: %v\n", tt.rw.(*httptest.ResponseRecorder).Code)
			}
		})
	}

}
