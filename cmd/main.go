package main

// https://www.youtube.com/watch?v=Q9hl2oSo0i0&list=PLbTTxxr-hMmyFAvyn7DeOgNRN8BQdjFm8&index=2
import (
	"log"

	"github.com/seephp/todo"
	"github.com/seephp/todo/pkg/handler"
	"github.com/seephp/todo/pkg/service"
	"github.com/seephp/todo/pkg/repository"
)

func main() {

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

srv := new(todo.Server)
if err := srv.Run("8000", handler.InitRoutes()); err!= nil {
	log.Fatalf("error occured wile running http server: %s", err.Error())
}

}



