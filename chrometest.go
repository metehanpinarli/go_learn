package main

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
)

func main() {
	//chrome örneği oluşturalım
	ctx, cancel := chromedp.NewExecAllocator(
		context.Background(),
		append(
			chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("headless", false),
		)...,
	)
	//headless: false ayarlayarak pencerenin görünmesini istedik

	//chrome nesnesini defer ile kapatmayı unutmuyoruz
	defer cancel()

	//yeni durum oluşturuyoruz
	ctx, cancel = chromedp.NewContext(ctx)

	//aynı şekilde defer ile penceremizide kapatıyoruz
	defer cancel()

	//Twitter isminin kaydedileceği değişkeni oluşturalım
	var twitterName string

	//chromedp.Run() içerisinde tarayıcıda yapılacak işlemleri yazıyoruz.
	err := chromedp.Run(ctx, //önce durumu (hangi pencere) olacağını belirtiyoruz

		//tarayıcının gitmsini istediğimiz adresi yazalım
		chromedp.Navigate(`https://kaanksc.com/posts/webview-statik-uygulama-ornegi3/`),

		//css seçici ile belirttiğimiz elementin yüklenmesini bekleyelim
		chromedp.WaitVisible(`.single__contents > p:nth-child(16) > a:nth-child(1)`, chromedp.ByQuery),

		//Tıklanılacak nesneyi yine css seçici ile belirtelim
		chromedp.Click(`.single__contents > p:nth-child(16) > a:nth-child(1)`, chromedp.ByQuery),
		//Bu işlemden sonra twitter'a gidecek

		//Twitter profilinde adın gösterildiği yeri css seçici ile beklemesini istedik
		chromedp.WaitVisible(`div.r-1b6yd1w:nth-child(1) > span:nth-child(1)`, chromedp.ByQuery),

		//belirttiğimiz css seçicisi ile elementin içindeki yazıyı twitterName değişkenine atayalım
		chromedp.Text(`div.r-1b6yd1w:nth-child(1) > span:nth-child(1)`, &twitterName),

		//burdan sonra tarayıcı penceresi kapanacak
	)

	//hata kontrolü yapaım
	if err != nil {
		log.Fatal(err)
	}

	//son olarak twitterName içindeki değişkeni ekrana bastıralım
	log.Printf("Twitter İsim:%s\n", twitterName)
}
