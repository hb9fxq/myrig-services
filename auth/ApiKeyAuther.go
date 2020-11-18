package auth

import (
	"github.com/hb9fxq/myrig-services/globals"
	"log"
)

func IsValidApiKeyForUser(appCtx *globals.ApplicationContext, userId string) {

	log.Printf("auth lookup for user %s", userId)
}
