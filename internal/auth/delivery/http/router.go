package http

func (ah *AuthHandler) SetRoutes() {
	ah.router.Get("/", ah.VendorAuth())

}
