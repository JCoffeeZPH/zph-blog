package session

type KeyType string

func (key KeyType) ToString() string {
	return string(key)
}

const (
	KeyUserId KeyType = "key_user_id"
	KeyUsername KeyType = "key_username"
	KeyNickname KeyType = "key_nickname"
	KeyAvatar KeyType = "key_avatar"
	KeyEmail KeyType = "key_email"
	KeyRole KeyType = "key_role"
)
