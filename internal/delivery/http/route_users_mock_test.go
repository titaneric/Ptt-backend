package http

import (
	"context"
	"time"

	"github.com/Ptt-official-app/go-bbs"
)

func (usecase *MockUsecase) GetUserByID(ctx context.Context, userID string) (bbs.UserRecord, error) {
	result := NewMockUserRecord(userID)
	return result, nil
}

func (usecase *MockUsecase) GetUserFavorites(ctx context.Context, userID string) ([]interface{}, error) {
	result := []interface{}{
		map[string]interface{}{
			"type":     "board",
			"board_id": "test_board_001",
		},
	}
	return result, nil
}

func (usecase *MockUsecase) GetUserInformation(ctx context.Context, userID string) (map[string]interface{}, error) {
	result := map[string]interface{}{
		"user_id": "id",
	}
	return result, nil
}

func (usecase *MockUsecase) GetUserPreferences(ctx context.Context, userID string) (map[string]string, error) {
	result := map[string]string{
		"favorite_no_highlight": "false",
	}

	return result, nil
}

func (usecase *MockUsecase) GetUserArticles(ctx context.Context, userID string) ([]interface{}, error) {
	return nil, nil
}

func (usecase *MockUsecase) GetUserComments(ctx context.Context, userID string) ([]interface{}, error) {
	result := []interface{}{
		map[string]interface{}{
			"board_id": "SYSOP",
		},
	}
	return result, nil
}

type MockUserRecord struct {
	userID string
}

// NewMockUserRecord generates fake user record for developing
func NewMockUserRecord(userID string) *MockUserRecord { return &MockUserRecord{userID: userID} }
func (u *MockUserRecord) UserID() string              { return u.userID }

// HashedPassword return user hashed password, it only for debug,
// If you want to check is user password correct, please use
// VerifyPassword insteaded.
func (u *MockUserRecord) HashedPassword() string { return "" }

// VerifyPassword will check user's password is OK. it will return null
// when OK and error when there are something wrong
func (u *MockUserRecord) VerifyPassword(password string) error { return nil }

// Nickname return a string for user's nickname, this string may change
// depend on user's mood, return empty string if this bbs system do not support
func (u *MockUserRecord) Nickname() string { return "" }

// RealName return a string for user's real name, this string may not be changed
// return empty string if this bbs system do not support
func (u *MockUserRecord) RealName() string { return "" }

// NumLoginDays return how many days this have been login since account created.
func (u *MockUserRecord) NumLoginDays() int { return 0 }

// NumPosts return how many posts this user has posted.
func (u *MockUserRecord) NumPosts() int { return 0 }

// Money return the money this user have.
func (u *MockUserRecord) Money() int { return 0 }

// LastLogin return last login time of user
func (u *MockUserRecord) LastLogin() time.Time { return time.Now() }

// LastHost return last login host of user, it is IPv4 address usually, but it
// could be domain name or IPv6 address.
func (u *MockUserRecord) LastHost() string { return "" }
