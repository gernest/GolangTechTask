package GolangTechTask

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/buffup/GolangTechTask/api"
	"github.com/urfave/cli"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Config struct {
	Port     int
	Endpoint string
	Region   string
	Memory   bool
	Trace    bool
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
		cli.StringFlag{
			Name:   "region,r",
			Usage:  "aws region",
			EnvVar: "VOTING_AWS_REGION",
			Value:  "local",
		},
		cli.StringFlag{
			Name:   "endpoint,e",
			Usage:  "dynamodb endpoint",
			EnvVar: "VOTING_DYNAMODB_ENDPOINT",
			Value:  "http://localhost:8000",
		},
		cli.BoolFlag{
			Name:   "mem,m",
			Usage:  "Uses an in memory storage",
			EnvVar: "VOTING_MEMORY_STORE",
		},
		cli.BoolFlag{
			Name:   "trace,t",
			Usage:  "Enable open tracing, the traces will be exported to stdout",
			EnvVar: "VOTING_TRACE",
		},
	}
	a.Action = command
	return a
}

func command(ctx *cli.Context) error {
	c := &Config{
		Port:     ctx.GlobalInt("port"),
		Endpoint: ctx.GlobalString("endpoint"),
		Region:   ctx.GlobalString("region"),
		Memory:   ctx.GlobalBool("mem"),
		Trace:    ctx.GlobalBool("trace"),
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
	return run(context.Background(), c)
}

func run(ctx context.Context, c *Config) error {
	m := Logger.Named("main")
	m.Info("Opening storage")
	store, err := NewStore(c)
	if err != nil {
		return err
	}
	store.Clear()
	m.Info("Setting up open telemetry")
	tp, err := CreateOpenTelemetry(c)
	if err != nil {
		return err
	}
	m.Info("Starting listener for votable service", zap.Int("port", c.Port))
	ls, err := net.Listen("tcp", fmt.Sprintf(":%d", c.Port))
	if err != nil {
		return err
	}
	defer tp.Shutdown(context.Background())

	return serve(ctx, store, c, ls)
}

func serve(ctx context.Context, store Store, c *Config, ls net.Listener) error {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
		grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()),
	)
	api.RegisterVotingServiceServer(s, &Server{
		store: store,
	})
	return s.Serve(ls)
}
