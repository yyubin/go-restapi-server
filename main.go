package main

import (
	"admin-server/pkg/configs"
	middleware "admin-server/pkg/middlewares"
	"admin-server/pkg/routes"
	"admin-server/pkg/utils"

	_ "admin-server/docs"

	"github.com/gofiber/fiber/v2"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:5001
// @BasePath /
func main() {
	// // consul 설정 및 등록
	// config := api.DefaultConfig()
	// client, err := api.NewClient(config)

	// serviceID := "admin-server"
	// serviceName := "callisto-admin-server"
	// serviceAddress := "127.0.0.1"
	// servicePort := 3000

	// registration := new(api.AgentServiceRegistration)
	// registration.ID = serviceID
	// registration.Name = serviceName
	// registration.Address = serviceAddress
	// registration.Port = servicePort

	// agent := client.Agent()
	// err = agent.ServiceRegister(registration)
	// if err != nil {
	// 	log.Fatal("Falied to register service to Consul")
	// }

	// fiberConfig := configs.FiberConfig()

	// app := fiber.New(fiberConfig)
	// routes.PublicRoutes(app)
	// routes.PrivateRoutes(app)

	// // http 서버 실행
	// go func() {
	// 	utils.StartServerWithGracefulShutdown(app)
	// }()

	// // kafka 구독
	// fmt.Println("kafka on")
	// topic := "example-topic"
	// partition := 0

	// conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	// if err != nil {
	// 	log.Fatal("failed to dial leader", err)
	// }

	// defer conn.Close()

	// reader := kafka.NewReader(kafka.ReaderConfig{
	// 	Brokers: []string{"localhost:9092"},
	// 	Topic:   topic,
	// })

	// var wg sync.WaitGroup
	// wg.Add(1)

	// // 별도 고루틴에서 Kafka Consumer 실행
	// go func() {
	// 	defer wg.Done()

	// 	for {
	// 		m, err := reader.ReadMessage(context.Background())
	// 		if err != nil {
	// 			log.Fatal("failed to read message: ", err)
	// 		}

	// 		fmt.Printf("received message: %s\n", string(m.Value))
	// 	}
	// }()

	// // 프로그램 종료를 기다림
	// wg.Wait()

	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app)

	// Routes.
	routes.SwaggerRoute(app)  // Register a route for API Docs (Swagger).
	routes.PublicRoutes(app)  // Register a public routes for app.
	routes.PrivateRoutes(app) // Register a private routes for app.

	// Start server (with graceful shutdown).
	utils.StartServer(app)
}
