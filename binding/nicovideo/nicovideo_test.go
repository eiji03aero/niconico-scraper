package nicovideo

import (
	"testing"
	"time"

	"github.com/go-rod/rod"
)

func Test_Demo(t *testing.T) {
	cred := NicoVideoCredential{Email: "aj07light@gmail.com", Password: "Aeropro2"}
	svc := NewNicoVideoSvc(cred)

	page := rod.New().MustConnect().MustPage("https://google.com")
	svc.Login(page)

	page.Navigate("https://www.nicovideo.jp/watch/sm40590922")
	page.MustWaitLoad()

	videoStartButton := page.MustElementByJS(`() => document.querySelector('button.VideoStartButton')`)
	videoStartButton.MustClick()

	time.Sleep(3000 * time.Millisecond)

	page.MustScreenshot("scsho.png")

	var adSkipButton *rod.Element
	adSkipButton = page.MustElementByJS(`() => document.querySelector('button.VideoAdSkippableSkipButton')`)
	adSkipButton.MustClick()
	time.Sleep(5000 * time.Millisecond)

	page.MustScreenshot("scsho.png")
}
