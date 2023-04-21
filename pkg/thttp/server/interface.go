package server

import "github.com/gofiber/fiber/v2"

type IHandler interface {
	GetRouter() fiber.Router
	GetMappingRoutes() map[string]*RouteMapping
}
