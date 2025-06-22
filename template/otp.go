package template

type IOtp interface {
	GenRandomOTP(int) string
	SaveOTPCache(string)
	GetMessage(string) string
	SendNotification(string) error
}

type Otp struct {
	IOtp IOtp
}

func (o *Otp) GenAndSendOTP(otpLength int) error {
	otp := o.IOtp.GenRandomOTP(otpLength)
	o.IOtp.SaveOTPCache(otp)
	message := o.IOtp.GetMessage(otp)
	err := o.IOtp.SendNotification(message)
	if err != nil {
		return err
	}
	return nil
}
