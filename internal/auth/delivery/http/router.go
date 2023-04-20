package http

func (ah *AuthHandler) SetRoutes() {
	ah.Router.Get("/a", ah.VendorAuth())
	ah.Router.Post("sign/up", ah.)
}
