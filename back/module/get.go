package module

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/juanesech/topo/constants"
	db "github.com/juanesech/topo/database"
	"github.com/juanesech/topo/utils"
)

func Get(ctx *gin.Context) {
	var modulesFromDB []*Module
	var loadedModule *Module
	var module Module
	ctx.Header("Access-Control-Allow-Origin", "*")

	session, sessionErr := db.Client.OpenSession(constants.DBName)
	utils.CheckError(sessionErr)
	defer session.Close()

	query := session.QueryCollectionForType(reflect.TypeOf(&Module{})).WhereEquals("Name", ctx.Param("name"))
	utils.CheckError(query.GetResults(&modulesFromDB))

	if len(modulesFromDB) != 0 {
		module.ID = modulesFromDB[0].ID
		session.Load(&loadedModule, module.ID)
		module = *loadedModule
		ctx.JSON(http.StatusOK, module)
	} else {
		ctx.String(http.StatusNotFound, fmt.Sprintf("Module %s not found", ctx.Param("name")))
	}
}
