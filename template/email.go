package template

import "fmt"

type Email struct {
	Otp
}

func (e *Email) GenRandomOTP(i int) string {
	randomOTP := "1234"
	fmt.Printf("Email: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (e *Email) SaveOTPCache(otp string) {
	fmt.Printf("Email: saving otp %s to cache\n", otp)
}

func (e *Email) GetMessage(otp string) string {
	return "Email OTP for login is " + otp
}

func (e *Email) SendNotification(message string) error {
	fmt.Printf("Email: sending email: %s\n", message)
	return nil
}
