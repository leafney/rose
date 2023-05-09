package rose

import (
	"golang.org/x/crypto/bcrypt"
)

// BcryptHash 通过Bcrypt生成哈希值
func BcryptHash(str string) (newStr string, err error) {
	newByte, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	return string(newByte), err
}

// BcryptCheck 通过Bcrypt校验哈希值
func BcryptCheck(hashedPwd, pwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd))
}
