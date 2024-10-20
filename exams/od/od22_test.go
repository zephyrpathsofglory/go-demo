package od

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	GoogleSite  = "www.google.com"
	TestSite    = "www.test.com"
	BaiduSite   = "www.baidu.com"
	YoutubeSite = "www.youtube.com"
)

const (
	W3Site = "w3.test.com"
	W4Site = "w4.test.com"
	W5Site = "w5.test.com"
	W6Site = "w6.test.com"
)

func TestBrowserHistorySys(t *testing.T) {
	{
		s := NewBrowserHistorySys(TestSite, 3)
		assert.Equal(t, 2, s.visit(W3Site))
		assert.Equal(t, 3, s.visit(W4Site))
		assert.Equal(t, W3Site, s.back())
		assert.Equal(t, 3, s.visit(TestSite))
		assert.Equal(t, 3, s.visit(W5Site))
		assert.Equal(t, 3, s.visit(W6Site))
	}

	{
		s := NewBrowserHistorySys(TestSite, 10)

		assert.Equal(t, 2, s.visit(GoogleSite))
		assert.Equal(t, TestSite, s.back())
		assert.Equal(t, GoogleSite, s.forward())
		assert.Equal(t, GoogleSite, s.forward())
		assert.Equal(t, 3, s.visit(BaiduSite))
		assert.Equal(t, 4, s.visit(YoutubeSite))
		assert.Equal(t, BaiduSite, s.back())
		assert.Equal(t, 4, s.visit(BaiduSite))
		assert.Equal(t, GoogleSite, s.back())
		assert.Equal(t, 3, s.visit("www.mail.com"))
	}
}
