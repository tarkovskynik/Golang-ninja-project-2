package repository

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tarkovskynik/Golang-ninja-project-2/internal/domain"
)

func createUser(t *testing.T, user *domain.User) {
	err := TestUserRepository.Create(Ctx, user)
	require.NoError(t, err)
}

func TestUser_Create(t *testing.T) {
	user := &domain.User{
		Email:     randomString(5) + "@email.com",
		Password:  randomString(10),
		Token:     randomString(15),
		CreatedAt: time.Now(),
	}

	createUser(t, user)
}

func TestUser_GetByEmailAndPassword(t *testing.T) {
	user := &domain.User{
		Email:     randomString(5) + "@email.com",
		Password:  randomString(10),
		Token:     randomString(15),
		CreatedAt: time.Now(),
	}

	createUser(t, user)

	findUser, err := TestUserRepository.GetByEmailAndPassword(Ctx, user.Email, user.Password)
	require.NoError(t, err)
	require.NotEmpty(t, findUser)

	require.Equal(t, user.Email, findUser.Email)
	require.Equal(t, user.Token, findUser.Token)
	require.Equal(t, user.Password, findUser.Password)
	//require.Equal(t, user.CreatedAt, findUser.CreatedAt) // TODO fix check time variable
}

func TestUser_DeleteByEmail(t *testing.T) {
	user := &domain.User{
		Email:     randomString(5) + "@email.com",
		Password:  randomString(10),
		Token:     randomString(15),
		CreatedAt: time.Now(),
	}

	createUser(t, user)

	err := TestUserRepository.DeleteByEmail(Ctx, user.Email)
	require.NoError(t, err)

	findUser, err := TestUserRepository.GetByEmailAndPassword(Ctx, user.Email, user.Password)
	fmt.Println()
	require.Empty(t, findUser)
}

func TestUser_UpdateByEmail(t *testing.T) {
	user := &domain.User{
		Email:     randomString(5) + "@email.com",
		Password:  randomString(10),
		Token:     randomString(15),
		CreatedAt: time.Now(),
	}

	createUser(t, user)

	updateUser := &domain.User{
		Email:     user.Email,
		Password:  randomString(10),
		Token:     randomString(15),
		CreatedAt: time.Now(),
	}

	err := TestUserRepository.UpdateByEmail(Ctx, updateUser.Email, updateUser)
	require.NoError(t, err)

	findUser, err := TestUserRepository.GetByEmailAndPassword(Ctx, updateUser.Email, updateUser.Password)
	require.NoError(t, err)

	require.Equal(t, updateUser.Email, findUser.Email)
	require.Equal(t, updateUser.Token, findUser.Token)
	require.Equal(t, updateUser.Password, findUser.Password)
	//require.Equal(t, updateUser.CreatedAt, findUser.CreatedAt) // TODO fix check time variable
}

func randomString(n int) string {
	str := "abcdefghijklmnopqrstuvwxyz"
	ret := strings.Builder{}
	for i := 0; i < n; i++ {
		ret.WriteByte(str[rand.Intn(n)])
	}
	return ret.String()
}
