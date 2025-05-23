package cookie

import (
	"errors"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetCookie(c *fiber.Ctx, name string) (string, error) {
	cookie := c.Cookies(name)
	if cookie == "" {
		return "", errors.New("cookie is empty")
	}
	return cookie, nil
}

func ClearCookie(c *fiber.Ctx, name, domain string) {
	cookie := &fiber.Cookie{
		Name:     name,
		Value:    "",
		Path:     "/",
		Secure:   true,
		HTTPOnly: true,
		Domain:   domain,
		Expires:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
	}

	if strings.HasPrefix(c.Hostname(), "localhost") {
		cookie.Domain = "localhost"
	}

	c.Cookie(cookie)
}


type CookieData struct {
	Name    string
	Value   string
	Expires time.Time
	Domain  string
}


func SetCookie(c *fiber.Ctx, params CookieData) {
	cookie := &fiber.Cookie{
		Name:     params.Name,
		Value:    params.Value,
		Path:     "/",
		Secure:   true,
		HTTPOnly: true,
		Domain:   params.Domain,
		Expires:  params.Expires,
	}

	if strings.HasPrefix(c.Hostname(), "localhost") {
		cookie.Domain = "localhost"
	}

	c.Cookie(cookie)
}
