package serializer

import "golang.org/x/crypto/bcrypt"

// EncryptPassword
// @Func: 加密
func EncryptPassword(password string, level int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), level)
	if err != nil {
		// TODO:日志输出
		panic("加密失败！！！")
	}
	passwordEncrypt := string(bytes)
	return passwordEncrypt, err
}

// CheckPassword
// @Func：校验密码
func CheckPassword(encryptPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(encryptPassword), []byte(password))
	if err != nil {
		// TODO:日志输出
		return err
	}
	return nil
}

// ComparePassword
// @Func: 比较密码
func ComparePassword(encryptPassword string, password string) bool {
	err := CheckPassword(encryptPassword, password)
	return err == nil
}
