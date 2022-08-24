package main

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
)

func main() {
	//This is the code to actually see the browser
	//ctx, cancel := chromedp.NewExecAllocator(context.Background(), append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless", false))...)
	//defer cancel()

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var res string

	//The Query options may be the key
	err := chromedp.Run(ctx, 
		chromedp.Navigate("https://web.whatsapp.com/"),
		chromedp.WaitEnabled("._25pwu", chromedp.ByQuery),
		chromedp.OuterHTML("html", &res, chromedp.ByQuery),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
}