package main

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func main(){
	//Initializing Browser Context (if headless mode is not disabled this doesn't work)
	ctx, cancel := chromedp.NewExecAllocator(
		context.Background(),
		append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless", false))...
	)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	//var data string
	var data []*cdp.Node
	var data_status bool

	log.Println("Initializing Browser...")

	err := chromedp.Run(ctx,
		chromedp.Navigate("http://web.whatsapp.com/"),
		chromedp.WaitEnabled("._25pwu", chromedp.ByQuery),
		//chromedp.InnerHTML("._2UwZ_", &data, chromedp.ByQuery),
		chromedp.Nodes("._2UwZ_", &data),
		)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(data))
	log.Println(data_status)
}