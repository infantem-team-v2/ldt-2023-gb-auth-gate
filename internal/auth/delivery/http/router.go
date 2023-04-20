package http

func (ah *AuthHandler) SetRoutes() {
	ah.Router.Get("/a", ah.VendorAuth())
	ah.Router.Post("sign/up", ah.SignUp())
	ah.Router.Post("sign/in", ah.SignIn())
	ah.Router.Post("email/validate", ah.ValidateEmail())
}
