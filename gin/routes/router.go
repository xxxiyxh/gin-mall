package routes

import (
	api "gin/api/v1"
	"gin/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("api/v1")
	{
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		v1.GET("carousels", api.ListCarousel)

		v1.GET("products", api.ListProduct)
		v1.GET("products/：id", api.ShowProduct)
		v1.GET("imgs/：id", api.ListProductImg)
		v1.GET("categories", api.ListCategory)

		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.PUT("user", api.UserUpdate)
			authed.POST("user/sending-email", api.SendEmail)
			authed.POST("user/valid-email", api.ValidEmail)
			authed.POST("money", api.ShowMoney)

			authed.POST("product", api.CreateProduct)
			authed.POST("product", api.SearchProduct)

			authed.GET("favorite", api.ListFavorite)
			authed.POST("favorite", api.CreateFavorite)
			authed.DELETE("favorite/:id", api.DeleteFavorite)

			authed.POST("address", api.CreateAddress)
			authed.GET("address/:id", api.ShowAddress)
			authed.GET("address", api.ListAddress)
			authed.PUT("address/:id", api.UpdateAddress)
			authed.DELETE("address/:id", api.DeleteAddress)

			authed.POST("carts", api.CreateCart)
			authed.GET("carts", api.ListCart)
			authed.PUT("carts/:id", api.UpdateCart)
			authed.DELETE("carts/:id", api.DeleteCart)

			authed.POST("orders", api.CreateOrder)
			authed.GET("orders", api.ListOrder)
			authed.PUT("orders/:id", api.ShowOrder)
			authed.DELETE("orders/:id", api.DeleteOrder)

			authed.POST("paydown", api.OrderPay)

		}
	}
	return r
}
