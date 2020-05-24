package auth

import (
	"github.com/krippendorf/myrig-services/globals"
	"log"
)

func IsValidApiKeyForUser(appCtx *globals.ApplicationContext, userId string) {

	log.Printf("auth lookup for user %s", userId)
}
