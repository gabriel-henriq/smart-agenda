package paseto

import (
	"github.com/gabriel-henriq/smart-agenda/db"
	"github.com/gabriel-henriq/smart-agenda/token"
	"github.com/gabriel-henriq/smart-agenda/util"
	"github.com/gin-gonic/gin"
)

type IToken interface {
	SetupTokenRoute(routerGroup *gin.RouterGroup)
}

type Token struct {
	db         db.Store
	tokenMaker token.Maker
	config     util.Config
}

func NewToken(db db.Store, config util.Config) IToken {
	return Token{
		db: db,
	}
}

func (t Token) SetupTokenRoute(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/token", t.renewAccessToken)
}
