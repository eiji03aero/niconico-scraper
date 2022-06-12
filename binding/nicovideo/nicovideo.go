package nicovideo

import (
	"github.com/go-rod/rod"
)

type nicoVideoSvc struct {
	credential NicoVideoCredential
}

func NewNicoVideoSvc(cred NicoVideoCredential) *nicoVideoSvc {
	return &nicoVideoSvc{
		credential: cred,
	}
}

func (n *nicoVideoSvc) Login(page *rod.Page) {
	page.Navigate("https://account.nicovideo.jp/login")
	page.MustWaitLoad()

	mailInput := page.MustElementByJS(`() => document.querySelector('input#input__mailtel')`)
	mailInput.Input(n.credential.Email)

	passInput := page.MustElementByJS(`() => document.querySelector('input#input__password')`)
	passInput.Input(n.credential.Password)

	loginButton := page.MustElementByJS(`() => document.querySelector('input#login__submit')`)
	loginButton.MustClick()

	page.MustWaitLoad()
}
