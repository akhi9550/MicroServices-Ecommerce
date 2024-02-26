package server

import (
	"grpc-api-gateway/pkg/api/handler"
	"log"

	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(adminHandler *handler.AdminHandler, productHandler *handler.ProductHandler, userHandler *handler.UserHandler, cartHandler *handler.CartHandler, orderhandler *handler.OrderHandler) *ServerHTTP {
	router := gin.New()

	router.Use(gin.Logger())

	//admin side
	router.POST("/admin/login", adminHandler.LoginHandler)

	//user side
	router.POST("/user/signup", userHandler.UserSignup)
	router.POST("/user/login", userHandler.Userlogin)

	//product side
	router.GET("/user/home", productHandler.ShowAllProducts)
	router.POST("/admin/product/add", productHandler.AddProducts)
	router.DELETE("/admin/product/delete", productHandler.DeleteProducts)
	router.PUT("/admin/product/update", productHandler.UpdateProduct)

	//cart side
	router.POST("/user/cart", cartHandler.AddToCart)
	router.GET("/user/cart", cartHandler.GetCart)

	//order
	router.POST("user/cart/order", orderhandler.OrderItemsFromCart)
	router.GET("user/orders", orderhandler.GetOrderDetails)

	return &ServerHTTP{engine: router}
}

func (s *ServerHTTP) Start() {
	log.Printf("starting server on :3000")
	err := s.engine.Run(":3000")
	if err != nil {
		log.Printf("error while starting the server")
	}
}
