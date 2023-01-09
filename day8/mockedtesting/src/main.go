package src

import "necdec2022/day8/mockedtesting/facade"

func GenerateOTP(seed int, otpFacade facade.OTPFacade) int {

	return otpFacade.GetOTP(seed)
}
