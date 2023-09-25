package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/koba1108/sls-golang-lambda-provided.al2/src/config"
)

var ginLambda *ginadapter.GinLambda

func init() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		conf, err := config.LoadConfigFromENV()
		if err != nil {
			c.JSON(200, gin.H{
				"error": err.Error(),
			})
			return
		}
		serverConf, err := config.LoadServerConfigFromENV()
		if err != nil {
			c.JSON(200, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"conf":       conf,
			"serverConf": serverConf,
		})
	})
	ginLambda = ginadapter.New(r)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
