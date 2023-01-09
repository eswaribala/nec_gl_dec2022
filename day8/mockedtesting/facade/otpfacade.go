package facade

type OTPFacade interface {
	GetOTP(seed int) int
}
type standardOTP struct{}
