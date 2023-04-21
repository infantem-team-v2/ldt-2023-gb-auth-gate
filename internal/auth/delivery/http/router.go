package http

func (ah *AuthHandler) SetRoutes() {
	ah.router.Post("sign/up", ah.SignUp())
	ah.router.Post("sign/in", ah.SignIn())
	ah.router.Post("email/validate", ah.ValidateEmail())
	ah.router.Get("/", ah.VendorAuth())

}
