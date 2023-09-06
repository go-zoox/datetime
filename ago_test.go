package datetime

import (
	"testing"
	"time"

	"github.com/go-zoox/testify"
)

func TestAgoJust(t *testing.T) {
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
