package service

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/chromedp/chromedp"
)

type BaseService struct {
}

var BaseSvr *BaseService

func init() {
	BaseSvr = &BaseService{}
}

func (s *BaseService) FetchHtml(url string) (string, error) {
	return s.FetchHtmlBase(url)
	// return s.FetchHtmlWithChrome(url)
}

func (s *BaseService) FetchHtmlBase(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (s *BaseService) FetchHtmlWithChrome(url string) (string, error) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
		chromedp.Flag("start-maximized", true),
		chromedp.Flag(
			"user-data-dir",
			"C:\\assets\\chromedata",
		),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(
		context.Background(),
		opts...,
	)

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	var html string

	err := chromedp.Run(ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			js := `
					Object.defineProperty(navigator, 'webdriver', {
						get: () => undefined,
					});

					window.navigator.chrome = {
						runtime: {},
					};

					Object.defineProperty(navigator, 'languages', {
						get: () => ['zh-CN', 'zh'],
					});

					Object.defineProperty(navigator, 'plugins', {
						get: () => [1, 2, 3, 4, 5],
					});
				`
			return chromedp.Evaluate(js, nil).Do(ctx)
		}),

		// 打开网页
		chromedp.Navigate(url),
		// 等待页面加载
		chromedp.Sleep(20*time.Second),
		// 获取整个HTML
		chromedp.OuterHTML("html", &html),
	)

	if err != nil {
		return "", err
	}

	return html, nil
}
