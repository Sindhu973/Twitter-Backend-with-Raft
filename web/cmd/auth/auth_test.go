package auth

import(
	"testing"
)

func TestRegister(t *testing.T) {
	user := "testuser"
	password := "password"

	err = auth.Register(user, password)
	if err != nil {
		t.Error(err)
	}
}

