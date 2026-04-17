package main

import (
	"fmt"
	httpservice "study/http_service"
)

func main() {
	/*ctx := context.Background()
	conn, err := sql.CreatConnection(ctx)
	if err != nil {
		panic(err)
	}
	if err := sql.CreateTable(*conn, ctx); err != nil {
		panic(err)
	}
	http.HandleFunc("/task", handlers.CreateTaskHandler(conn))
	fmt.Println("successfully!")
	if err := http.ListenAndServe(":9091", nil); err != nil {
		panic(err)
	}*/
	/*_, err := os.Create("out/newfile.txt")
	if err != nil {
		panic(err)
	}*/
	err := httpservice.StartHttpService()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Spoted")
	}
}
