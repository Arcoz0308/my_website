package discord

import (
	"encoding/json"
	"github.com/arcoz0308/arcoz0308.tech/handlers/config"
	"github.com/arcoz0308/arcoz0308.tech/handlers/logger"
	"github.com/arcoz0308/arcoz0308.tech/routes/api/v1/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"strings"
	"time"
)

var (
	i         = 0
	rateLimit map[string]rateLimitInfos
)

var (
	ApiURL   = "https://discord.com/api/v9/"
	AssetURL = "https://cdn.discordapp.com/"
)

// Flags public flag
var (
	Flags = map[string]int{
		"STAFF":                    1 << 0,
		"PARTNER":                  1 << 1,
		"HYPESQUAD":                1 << 2,
		"BUG_HUNTER_LEVEL_1":       1 << 3,
		"HYPESQUAD_ONLINE_HOUSE_1": 1 << 6,
		"HYPESQUAD_ONLINE_HOUSE_2": 1 << 7,
		"HYPESQUAD_ONLINE_HOUSE_3": 1 << 8,
		"PREMIUM_EARLY_SUPPORTER":  1 << 9,
		"TEAM_PSEUDO_USER":         1 << 10,
		"BUG_HUNTER_LEVEL_2":       1 << 14,
		"VERIFIED_BOT":             1 << 16,
		"VERIFIED_DEVELOPER":       1 << 17,
		"CERTIFIED_MODERATOR":      1 << 18,
		"BOT_HTTP_INTERACTIONS":    1 << 19,
	}
)

type rateLimitInfos struct {
}

type apiUser struct {
	Id            string  `json:"id"`
	Username      string  `json:"username"`
	Discriminator string  `json:"discriminator"`
	Avatar        *string `json:"avatar"`
	PublicFlags   int     `json:"public_flags"`
	Banner        *string `json:"banner"`
	BannerColor   string  `json:"banner_color"`
	accentColor   int     `json:"accent_color"`
}

type bUser struct {
	Id                  string   `json:"id"`
	Username            string   `json:"username"`
	Discriminator       string   `json:"discriminator"`
	Tag                 string   `json:"tag"`
	CreatedAtTimestamps int      `json:"created_at_timestamps"`
	CreatedAtDate       string   `json:"created_at_date"`
	Avatar              *string  `json:"avatar"`
	AvatarURL           string   `json:"avatar_url"`
	PublicFlags         int      `json:"public_flags"`
	PublicFlagsString   []string `json:"public_flags_string"`
	Banner              *string  `json:"banner"`
	BannerURL           *string  `json:"banner_url"`
	BannerColor         string   `json:"banner_color"`
	accentColor         int      `json:"accent_color"`
}

func New(a apiUser) bUser {
	b := bUser{
		Id:            a.Id,
		Username:      a.Username,
		Discriminator: a.Discriminator,
		Tag:           a.Username + "#" + a.Discriminator,
		Avatar:        a.Avatar,
		PublicFlags:   a.PublicFlags,
		Banner:        a.Banner,
		BannerColor:   a.BannerColor,
		accentColor:   a.accentColor,
	}

	// parse date
	snowFlake, _ := strconv.Atoi(b.Id)
	b.CreatedAtTimestamps = (snowFlake >> 22) + 1420070400000
	t := time.UnixMilli(int64(b.CreatedAtTimestamps))
	b.CreatedAtDate = t.String()

	// add avatar url
	if b.Avatar != nil && *b.Avatar != "" {

		// check if animated
		var format string
		if strings.HasPrefix(*b.Avatar, "a_") {
			format = ".gif"
		} else {
			format = ".png"
		}

		b.AvatarURL = AssetURL + "avatars/" + b.Id + "/" + *b.Avatar + format
	} else {
		d, _ := strconv.Atoi(b.Discriminator)
		b.AvatarURL = AssetURL + "embed/avatars/" + strconv.Itoa(d%5) + ".png"
	}

	// add public flag string
	for flag, by := range Flags {
		if hasFlag(by, b.PublicFlags) {
			b.PublicFlagsString = append(b.PublicFlagsString, flag)
		}
	}

	if b.Banner != nil && *b.Banner != "" {

		// check if animated
		var format string
		if strings.HasPrefix(*b.Banner, "a_") {
			format = ".gif"
		} else {
			format = ".png"
		}

		url := AssetURL + "banners/" + b.Id + "/" + *b.Banner + format
		b.BannerURL = &url
	}
	return b
}

func User(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, code := reqUser(id)

	switch code {
	case 0:
		user2 := New(user)
		return ctx.Status(fiber.StatusOK).JSON(user2)
	case 1:
		return ctx.Status(fiber.StatusNotFound).JSON(utils.ApiError{
			Code:    strconv.Itoa(fiber.StatusNotFound),
			Message: "user not found",
		})
	case 2:
		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.ApiError{
			Code:    strconv.Itoa(fiber.StatusInternalServerError),
			Message: "internal server error",
		})
	case 3:
		return ctx.Status(fiber.StatusBadGateway).JSON(utils.ApiError{
			Code:    strconv.Itoa(fiber.StatusBadGateway),
			Message: "something failed with the discord api",
		})
	default:
		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.ApiError{
			Code:    strconv.Itoa(fiber.StatusInternalServerError),
			Message: "internal server error",
		})
	}

}

// int declaration :
// 0 = success
// 1 = not found
// 2 = 500
// 3 = 502

func reqUser(id string) (apiUser, int) {

	// change to next token
	defer func() {
		if len(config.Discord.Token) != 1 {
			if i+1 == len(config.Discord.Token) {
				i = 0
			} else {
				i++
			}
		}
	}()

	c := fiber.AcquireAgent().Timeout(time.Second * 3)
	r := c.Request()
	r.Header.SetMethod("GET")
	r.Header.Set("Authorization", "Bot "+config.Discord.Token[i])
	r.SetRequestURI(ApiURL + "users/" + id)
	err := c.Parse()
	if err != nil {
		logger.AppErrorf("api", "error with discord request, error : %s, userid : %s", err.Error(), id)
		return apiUser{}, 2
	}
	code, b, err2 := c.Bytes()
	if len(err2) != 0 {
		return apiUser{}, 2
	}
	switch code {
	case 200:
		user := apiUser{}
		err = json.Unmarshal(b, &user)
		if err != nil {
			logger.AppErrorf("api", "error with unmarshal json for discord user, error : %s, user id : %s", err.Error(), id)
			return apiUser{}, 2
		}
		return user, 0
	case 404:
		return apiUser{}, 1
	case 401:
		logger.AppErrorf("api", "authentication with discord failed, token index : %d", i)
		return apiUser{}, 2
	case 500:
		return apiUser{}, 3
	default:
		logger.AppErrorf("api", "invalid status code response from discord, code : %d, response : %s", code, string(b))
		return apiUser{}, 2
	}

}

// hasFlag check if user has a flag
func hasFlag(flag int, flags int) bool {
	return flag&flags != 0
}
