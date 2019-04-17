package main

import (
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"mather/internal/api"
	"mather/internal/client"
	"mather/internal/commands"
	"os"
)

const (
	hostnameKey = "HOSTNAME"
	portKey     = "PORT"
	apiKeyKey   = "API_KEY"
)

func main() {
	loadConfig()
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", viper.GetString(hostnameKey), viper.GetInt(portKey)), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	rpcClient := api.NewMatherClient(conn)
	remoteClient := &client.RemoteClient{
		RPCClient: rpcClient,
		APIKey:    viper.GetString(apiKeyKey),
	}

	command := commands.NewMatherCommand(remoteClient)
	if err := command.Execute(); err != nil {
		log.Fatalf("%v\n", err)
	}
}

func loadConfig() {
	// TODO: Could customise these as flags
	viper.AddConfigPath(".")
	viper.SetConfigName("cfg")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv() // Can customise config via Env
	_ = viper.BindEnv(hostnameKey, portKey, apiKeyKey)
	if err := viper.ReadInConfig(); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			break // If we are in the env var only case
		default:
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
