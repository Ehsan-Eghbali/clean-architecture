package main

import (
	"log"
	"net/http"

	deliveryHttp "clean-architecture-go/internal/delivery/http" // اسم مستعار برای پکیج داخلی
	"clean-architecture-go/internal/repository/memory"
	"clean-architecture-go/internal/usecase"
)

func main() {
	// ساخت ریپازیتوری در حافظه
	userRepo := memory.NewUserRepositoryMemory()

	// ساخت یوزکیس
	userUC := usecase.NewUserUseCase(userRepo)

	// ساخت هندلر
	userHandler := deliveryHttp.NewUserHandler(userUC)

	// ساخت روتر
	router := deliveryHttp.NewRouter(userHandler)

	log.Println("Server is running on port :8080")
	// از پکیج استاندارد net/http برای ران کردن سرور استفاده می‌کنیم
	log.Fatal(http.ListenAndServe(":8082z", router))
}
