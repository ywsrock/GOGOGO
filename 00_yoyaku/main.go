// Command click2 is a chromedp example demonstrating how to use a selector to
// click on an element.
package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/target"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {

	quit := make(chan os.Signal)
	signal.Notify(quit,os.Interrupt)

	//headless　表示設定
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false), // headless=false に変更
		chromedp.Flag("disable-gpu", false),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-extensions", false),
		chromedp.Flag("hide-scrollbars", false),
		chromedp.Flag("mute-audio", false),
	)
	//create allocator
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...) // cancel() を呼ばないように変更

	// create chrome instance by allocator
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf)) // cancel() を呼ばないように変更

	defer cancel()

	//wati new target
	ch := chromedp.WaitNewTarget(ctx, func(info *target.Info) bool {
		fmt.Println(info.URL)
		return info.URL != ""
	})

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 1*time.Hour)
	defer cancel()

	// navigate to a page, wait for an element, click
	var example string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.kds-net.jp/`),
		chromedp.ScrollIntoView(`#container > div:nth-child(3) > footer`),
		//予約クリック
		chromedp.Click(`#global-nav1 > li:nth-child(6)`, chromedp.NodeVisible),
		chromedp.WaitVisible(`#container > div:nth-child(3) > div > div.span9.flex-priority1 > div:nth-child(3) > div > div > div`),
		chromedp.Click(`#container > div:nth-child(3) > div > div.span9.flex-priority1 > div:nth-child(2) > div > div > div > div:nth-child(1) > div > p > u > a`),
	)

	if err != nil {
		log.Fatal(err)
	}

	newCtx, cancel := chromedp.NewContext(ctx, chromedp.WithTargetID(<-ch))
	defer cancel()

	//var urlstr string
	err1 := chromedp.Run(newCtx,
		//chromedp.Location(&urlstr),
		chromedp.WaitVisible(`#p01aForm_login`),
		chromedp.SendKeys(`#p01aForm_b_studentId`,"51395"),

	)

	if err1 != nil {
		log.Fatal(err1)
	}

	time.Sleep(2 * time.Second)
	err1 = chromedp.Run(newCtx,
		chromedp.SendKeys(`#p01aForm_b_password`,"yws13474987768"),
	)

	time.Sleep(2 * time.Second)
		err1 = chromedp.Run(newCtx,
		chromedp.Click(`#p01aForm_login`),
		chromedp.WaitVisible("body > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(3) > td"),
		chromedp.Text(`body > table > tbody > tr:nth-child(2) > td > table > tbody > tr:nth-child(1) > td > font > b`,&example),

		)

	log.Printf("Go's time.After example:\n%s", example)
	// 受信まで、待機
	<-quit

}
