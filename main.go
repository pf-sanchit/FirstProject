package main

import (
	"context"
	v1 "firstProject/service/v1"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	////TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	//// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	//fmt.Println("Enter your name...")
	//var fname, lname string
	//_, err := fmt.Scanln(&fname, &lname)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	panic(err)
	//}
	//fmt.Println("Hello and welcome,", fname, lname, "!")
	//
	//for i := 1; i <= 5; i++ {
	//	//TIP <p>To start your debugging session, right-click your code in the editor and select the Debug option.</p> <p>We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
	//	// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>.</p>
	//	fmt.Printf("100/i = %f, where i = %d\n", 100/float64(i), i)
	//}

	var (
		httpAddr = flag.String("http", ":8080", "http listen address")
	)
	flag.Parse()
	ctx := context.Background()

	service := v1.NewService()
	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	endpoints := v1.Endpoints{
		HelloEndpoint: v1.MakeHelloEndpoint(service),
	}

	go func() {
		log.Println("Service is listening on port", *httpAddr)
		handler := v1.MyHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	log.Fatal(<-errChan)
}
