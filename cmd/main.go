package main

import "github.com/seephp/todo"
import "log"

func main() {
srv := new(todo.Server)
if err := srv.Run("8000"); err!= nil {
	log.Fatalf("error occured wile running http server: %s", err.Error())
}

}



