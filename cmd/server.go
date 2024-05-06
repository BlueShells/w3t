package cmd

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"io"
	"mos/docs"
	"net/http"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		listen, _ := cmd.Flags().GetString("listen")
		startHttpServer(listen)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.PersistentFlags().StringP("listen", "l", "127.0.0.1:9280", "ip:port to listen on")
}

func startHttpServer(listen string) {
	r := setupRouter()

	r.Run(listen)
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")

	v1.GET("/ping", pingHandler)
	v1.GET("/jsonrpc", jsonrpcGetHandler)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}

// @BasePath /api/v1
// Ping test
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} pong
// @Router /ping [get]
func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}

// @BasePath /api/v1
// @Summary JSONRPC over HTTP Proxy
// @Schemes
// @Description
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} jsonrpc result
// @Router /jsonrpc [get]
// @Param target query string true "Target url of jsonrpc server"
// @Param method query string true "JSONRPC methods" Enums(eth_blockNumber, eth_web3_clientversion, optimism_rollupConfig, optimism_syncStatus, admin_nodeInfo, admin_peers)
// @Param id query int false "Request ID"
func jsonrpcGetHandler(c *gin.Context) {
	q := c.Request.URL.Query()
	target := ""
	method := ""
	id := "1"

	methods, ok := q["method"]
	if !ok {
		c.JSON(http.StatusBadRequest, "arg method not exist")
	}

	method = methods[0]

	targets, ok := q["target"]
	if !ok {
		c.JSON(http.StatusBadRequest, "arg target not exist")
	}

	target = targets[0]

	ids, ok := q["id"]

	if ok {
		id = ids[0]
	}

	var jsonData = []byte(`{
		"id": ` + id + `,
		"jsonrpc": "2.0",
		"method": "` + method + `",
		"params": []
	}`)

	resp, err := http.Post(target, "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		c.JSON(http.StatusInternalServerError, "RPC response error: "+err.Error())
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "RPC response error: "+err.Error())
	}
	c.JSON(http.StatusOK, json.RawMessage(body))
}
