package paseto

import (
	"github.com/gin-gonic/gin"

	"github.com/gabriel-henriq/smart-agenda/db"
	"github.com/gabriel-henriq/smart-agenda/token"
	"github.com/gabriel-henriq/smart-agenda/util"
)

type IToken interface {
	SetupTokenRoute(routerGroup *gin.RouterGroup)
}

type Token struct {
	db         db.Store
	tokenMaker token.Maker
	config     util.Config
}

func NewToken(db db.Store) IToken {
	return Token{
		db: db,
	}
}

func (t Token) SetupTokenRoute(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/token", t.renewAccessToken)
}
