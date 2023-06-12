package main

import (
	"github.com/spf13/cobra"

	"github.com/kebe7jun/miragedebug/config"
	"github.com/kebe7jun/miragedebug/internal/apps"
	"github.com/kebe7jun/miragedebug/internal/servers"
	"github.com/kebe7jun/miragedebug/pkg/log"
)

func main() {
	httpAddr := ""
	grpcAddr := ""
	kubeconfig := ""
	root := &cobra.Command{
		Use: "mirage-debug-server",
		Args: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			log.SetDebug()
			config.SetKubeconfig(kubeconfig)
			grpcServer := servers.NewGRPCServer(grpcAddr, apps.RegisterGRPCRoutes)
			go grpcServer.Run()
			gwServer := servers.NewGatewayServer("mirage debug server", httpAddr, grpcAddr, apps.RegisterHTTPRoutes())
			return gwServer.Run()
		},
	}
	root.PersistentFlags().StringVarP(&httpAddr, "http-addr", "", ":38080", "HTTP listen address.")
	root.PersistentFlags().StringVarP(&grpcAddr, "grpc-addr", "", ":38081", "GRPC listen address.")
	root.PersistentFlags().StringVarP(&kubeconfig, "kubeconfig", "k", "~/.kube/config", "Kubeconfig file path.")
	root.Execute()
}
