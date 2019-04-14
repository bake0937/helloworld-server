package main

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"proton"

	"server"

	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Version and Revision are replaced when building.
// To set specific version, edit Makefile.
var (
	Version  = "0.0.1"
	Revision = "xxx"

	Name  = "helloworld-server"
	Usage = "Notification service for clients"
)

func main() {
	// l := logger.NewLogger("")
	// defer l.Close()
	// l.ReplaceGlobal()

	app := cli.NewApp()
	app.Version = fmt.Sprintf("%s (%s)", Version, Revision)
	app.Name = Name
	app.Usage = Usage

	var (
		// serverCrtPath string
		// serverKeyPath string
		binding string
		// fluentdAddr   string
		// fluentdTag    string
		// configPath    string
		// insecureFlag  bool
		healthBinding string
	)

	app.Flags = []cli.Flag{
		// cli.StringFlag{
		// 	Name:        "key",
		// 	Usage:       "Path to certification private key",
		// 	Destination: &serverKeyPath,
		// },
		// cli.StringFlag{
		// 	Name:        "cert",
		// 	Usage:       "Path to certification",
		// 	Destination: &serverCrtPath,
		// },
		// cli.BoolFlag{
		// 	Name:        "insecure",
		// 	Usage:       "Running without TLS",
		// 	Destination: &insecureFlag,
		// },
		cli.StringFlag{
			Name:        "binding, b",
			Usage:       "Server binding address",
			Value:       "127.0.0.1:50051",
			Destination: &binding,
		},
		// cli.StringFlag{
		// 	Name:        "health",
		// 	Usage:       "Health check binding",
		// 	Value:       "127.0.0.1:9999",
		// 	Destination: &healthBinding,
		// },
		// cli.StringFlag{
		// 	Name:        "fluentd",
		// 	Usage:       "Address to fluentd (like 127.0.0.1:24224)",
		// 	Value:       "",
		// 	Destination: &fluentdAddr,
		// },
		// cli.StringFlag{
		// 	Name:        "fluentdTag",
		// 	Usage:       "Fluentd access log tag",
		// 	Value:       fmt.Sprintf("log.access.%s", Name),
		// 	Destination: &fluentdTag,
		// },
		// cli.StringFlag{
		// 	Name:        "config, c",
		// 	Usage:       "Path to configuration file",
		// 	Value:       "config.yml",
		// 	Destination: &configPath,
		// },
	}

	app.Action = func(c *cli.Context) error {
		// var fluentLogger *logger.FluentdClient
		// if fluentdAddr != "" {
		// 	fluentLogger = logger.NewFluentdClient(
		// 		fluentdAddr,
		// 		5*time.Second,
		// 		4*1024,
		// 	)

		// 	defer fluentLogger.Close()
		// }

		// cfg, err := config.NewConfig(configPath)
		// if err != nil {
		// 	logger.Fatal(err.Error())
		// }

		// db, err := md.NewDB(
		// 	cfg.Db.User,
		// 	cfg.Db.Password,
		// 	cfg.Db.Host,
		// 	cfg.Db.Port,
		// 	cfg.Db.DBName,
		// )
		// if err != nil {
		// 	logger.Fatal(err.Error())
		// }
		// defer db.Close()
		// m := model.NewModel(db)

		opts := []grpc.ServerOption{}

		// logger.Infow("Starting gRPC server",
		// 	"binding", binding,
		// 	"insecure", insecureFlag,
		// 	"key", serverKeyPath,
		// 	"cert", serverCrtPath,
		// 	"fluentdAddr", fluentdAddr,
		// 	"fluentdTag", fluentdTag,
		// 	"healthBinding", healthBinding,
		// )

		listener, err := net.Listen("tcp", binding)
		if err != nil {
			return cli.NewExitError(err, 1)
		}

		svr := grpc.NewServer(opts...)

		proton.RegisterGreeterServer(
			svr,
			server.NewServer(),
		)
		reflection.Register(svr)

		go func() {
			http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte(Revision))
			})
			http.ListenAndServe(healthBinding, nil)
		}()

		if err := svr.Serve(listener); err != nil {
			return cli.NewExitError(err, 1)
		}

		return nil
	}

	app.Run(os.Args)
}
