package datetime

import (
	"testing"
	"time"

	"github.com/go-zoox/testify"
)

func TestAgo(t *testing.T) {
	testify.Equal(t, "just now", Now().Ago())
	testify.Equal(t, "5 second ago", Now().Add(-5*time.Second).Ago())
	testify.Equal(t, "5 minute ago", Now().Add(-5*time.Minute).Ago())
	testify.Equal(t, "5 hour ago", Now().Add(-5*time.Hour).Ago())
	testify.Equal(t, "5 day ago", Now().Add(-5*24*time.Hour).Ago())
	testify.Equal(t, "3 week ago", Now().Add(-3*7*24*time.Hour).Ago())
	testify.Equal(t, "5 month ago", Now().Add(-5*30*24*time.Hour).Ago())
	testify.Equal(t, "5 year ago", Now().Add(-5*365*24*time.Hour).Ago())
	testify.Equal(t, "50 year ago", Now().Add(-50*365*24*time.Hour).Ago())
}

func TestAgoWithI18n(t *testing.T) {
	SetLocales(I18nLocales{
		"zh-CN": I18nTranslations{
			"just now":      "刚刚",
			"%d second ago": "%d秒前",
			"%d minute ago": "%d分钟前",
			"%d hour ago":   "%d小时前",
			"%d day ago":    "%d天前",
			"%d week ago":   "%d周前",
			"%d month ago":  "%d月前",
			"%d year ago":   "%d年前",
		},
	})
	SetLanguage("zh-CN")

	testify.Equal(t, "刚刚", Now().Ago())
	testify.Equal(t, "5秒前", Now().Add(-5*time.Second).Ago())
	testify.Equal(t, "5分钟前", Now().Add(-5*time.Minute).Ago())
	testify.Equal(t, "5小时前", Now().Add(-5*time.Hour).Ago())
	testify.Equal(t, "5天前", Now().Add(-5*24*time.Hour).Ago())
	testify.Equal(t, "3周前", Now().Add(-3*7*24*time.Hour).Ago())
	testify.Equal(t, "5月前", Now().Add(-5*30*24*time.Hour).Ago())
	testify.Equal(t, "5年前", Now().Add(-5*365*24*time.Hour).Ago())
	testify.Equal(t, "50年前", Now().Add(-50*365*24*time.Hour).Ago())
}
