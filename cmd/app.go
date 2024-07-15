package main

import (
	"fmt"
	"lets-blog/config"
	"lets-blog/routes"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type App struct {
	config *config.Config
	router http.Handler
	ddb    *dynamodb.DynamoDB
}

func NewApp(config *config.Config) *App {
	awsSession := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,

		Config: aws.Config{
			Region:   &config.AwsConfig.Region,
			Endpoint: &config.DynamoEndpoint,
			Credentials: credentials.NewStaticCredentials(
				config.AwsConfig.AccessKey,
				config.AwsConfig.SecretKey,
				"",
			),
		},
	}))

	app := &App{
		config: config,
		ddb:    dynamodb.New(awsSession),
	}

	app.loadRoutes()

	return app
}

func (app *App) Start() error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", app.config.ServerPort),
		Handler: app.router,
	}

	_, err := app.ddb.ListTables(&dynamodb.ListTablesInput{Limit: aws.Int64(1)})
	if err != nil {
		return fmt.Errorf("failed to connect to dynamodb: %w", err)
	}

	fmt.Println("Starting server at port: ", app.config.ServerPort)

	ch := make(chan error, 1)
	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err)
		}

		close(ch)
	}()

	err = <-ch
	return err
}

func (app *App) loadRoutes() {
	app.router = routes.NewRoutes(app.ddb)
}
