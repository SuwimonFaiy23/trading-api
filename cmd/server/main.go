package main

import (
	"log"
	"trading-api/internal/config"
	"trading-api/internal/db"
	"trading-api/internal/handler"
	"trading-api/internal/repository"
	"trading-api/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	db, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("cannot connect database: %v", err)
	}

	r := gin.Default()

	userHandler := handler.NewUserHandler(handler.UserHandlerDependencies{
		UserService: service.NewUserService(service.UserServiceDependencies{
			UserRepository: repository.NewUserRepository(repository.UserDependencies{Database: db}),
		}),
	})

	productHandler := handler.NewProductHandler(handler.ProductHandlerDependencies{
		ProductService: service.NewProductService(service.ProductServiceDependencies{
			ProductRepository: repository.NewProductRepository(repository.ProductDependencies{
				Database: db,
			}),
		}),
	})

	orderHandler := handler.NewOrderHandler(handler.OrderHandlerDependencies{
		OrderService: service.NewOrderService(service.OrderServiceDependencies{
			OrderRepository: repository.NewOrderRepository(repository.OrderDependencies{
				Database: db,
			}),
			UserRepository: repository.NewUserRepository(
				repository.UserDependencies{
					Database: db,
				},
			),
			ProductRepository: repository.NewProductRepository(
				repository.ProductDependencies{
					Database: db,
				},
			),
		}),
	})

	commissionHandler := handler.NewCommissionHandler(handler.CommissionHandlerDependencies{
		CommissionService: service.NewCommissionService(service.CommissionServiceDependencies{
			CommissionRepository: repository.NewCommissionRepository(repository.CommissionDependencies{
				Database: db,
			}),
		}),
	})

	affiliateHandler := handler.NewAffiliateHandler(handler.AffiliateHandlerDependencies{
		AffiliateService: service.NewAffiliateService(service.AffiliateServiceDependencies{
			AffiliateRepository: repository.NewAffiliateRepository(repository.AffiliateDependencies{
				Database: db,
			}),
			UserRepository: repository.NewUserRepository(
				repository.UserDependencies{
					Database: db,
				},
			),
		}),
	})

	users := r.Group("/user")
	users.POST("", userHandler.CreateUser)
	users.PATCH("/:id", userHandler.UpdateUser)
	users.PATCH("/add/balance/:id", userHandler.AddBalanceUser)
	users.PATCH("/deduct/balance/:id", userHandler.DeductBalanceUser)
	users.GET("/:id", userHandler.GetUserByID)
	users.GET("/all", userHandler.GetUsers)

	product := r.Group("/product")
	product.POST("", productHandler.CreateProduct)
	product.GET("/:id", productHandler.GetProductByID)
	product.GET("/list", productHandler.GetListProduct)

	order := r.Group("/order")
	order.POST("", orderHandler.CreateOrder)
	order.GET("/:id", orderHandler.GetOrderByID)

	commission := r.Group("/commission")
	commission.GET("/:id", commissionHandler.GetCommissionByID)
	commission.GET("/list", commissionHandler.GetListCommission)

	affiliate := r.Group("/affiliate")
	affiliate.POST("", affiliateHandler.CreateAffiliate)
	affiliate.GET("/:id", affiliateHandler.GetAffiliateByID)
	affiliate.GET("/list", affiliateHandler.GetListAffiliate)

	r.Run(":8080")

	log.Println("🚀 Server starting...")
}
