package http

func (ah *AuthHandler) SetRoutes() {
	ah.Router.Get("/", ah.VendorAuth())
}
