package helper

import (
	"encoding/json"
	"go-app/domain"
	"net/http"
	"os"
	"strconv"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gofiber/fiber/v2"
)

type Client struct {
	Client *memcache.Client
	Auth  Auth
}

func NewMemcached() (*Client, error) {
	client := memcache.New(os.Getenv("MEMCACHED_HOST"))

	if err := client.Ping(); err != nil {
		return nil, err
	}
	return &Client{
		Client: client,
	}, nil
}

func (cl *Client) VerifyCache(c *fiber.Ctx) error {
	user := cl.Auth.GetUser(c)
	idStr := strconv.FormatUint(uint64(user.ID), 10)
    val, err := cl.Client.Get(idStr)
    if err != nil {
        return c.Next()
    }
	data := domain.User{}
    err = json.Unmarshal(val.Value, &data)
	if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
		"message": "error unmarshalling data",
	})}
    return c.Status(http.StatusOK).JSON(&fiber.Map{"message": data})
}

