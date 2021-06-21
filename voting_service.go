package GolangTechTask

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/buffup/GolangTechTask/api"
	"github.com/urfave/cli"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Config struct {
	Port int
}

func App() *cli.App {
	a := cli.NewApp()
	a.Name = "Voting Service"
	a.Usage = "Simple gRPC service for voting"
	a.Flags = []cli.Flag{
		cli.IntFlag{
			Name:   "port,p",
			Usage:  "Port to bind the service",
			Value:  8080,
			EnvVar: "VOTING_SERVICE_PORT",
		},
		cli.StringFlag{
			Name:   "config,c",
			Usage:  "Path to the configuration file",
			EnvVar: "VOTING_CONFIG_FILE",
		},
	}
	a.Action = command
	return a
}

func command(ctx *cli.Context) error {
	c := &Config{
		Port: ctx.GlobalInt("port"),
	}
	if config := ctx.GlobalString("config"); config != "" {
		f, err := os.Open(config)
		if err != nil {
			return err
		}
		if err := json.NewDecoder(f).Decode(c); err != nil {
			f.Close()
			return err
		}
		f.Close()
	}
	return run(context.Background(), &Config{})
}

func run(ctx context.Context, c *Config) error {
	m := Logger.Named("main")
	m.Info("Starting listener for votable service", zap.Int("port", c.Port))
	ls, err := net.Listen("tcp", fmt.Sprintf(":%d", c.Port))
	if err != nil {
		return err
	}
	store, err := NewStore(c)
	if err != nil {
		return err
	}
	return serve(ctx, store, c, ls)
}

func serve(ctx context.Context, store Store, c *Config, ls net.Listener) error {
	s := grpc.NewServer()
	api.RegisterVotingServiceServer(s, &Server{
		store: store,
	})
	return s.Serve(ls)
}
