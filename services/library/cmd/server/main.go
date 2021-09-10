package server

import (
	"github.com/mohammadne/bookman/library/internal/configs"
	"github.com/spf13/cobra"
)

const (
	use   = "server"
	short = "run server"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{Use: use, Short: short, Run: main}

	envFlag := "set config environment, default is dev"
	cmd.Flags().StringP("env", "e", "", envFlag)

	return cmd
}

func main(cmd *cobra.Command, args []string) {
	env := cmd.Flag("env").Value.String()
	config := configs.Server(env)
	_ = config
	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// db := database.NewMysqlDatabase(server.Config.Database)

	// // repositories
	// booksRepository := books.NewRepository(db, server.Logger)

	// // usecases
	// booksUsecase := books.NewUsecase(booksRepository)

	// rest := rest.New(server.Config.Rest)
	// v1Group := rest.Instance().Group("/api/v1")

	// booksHandler := books_rest.NewHandler(booksUsecase)
	// booksHandler.Route(v1Group.Group("/books"))

	// go func(s web.Server) {
	// 	server.Logger.Info("start serving rest server")
	// 	if err := s.Serve(nil); err != nil {
	// 		server.Logger.Panic("server failed to start", logger.Error(err))
	// 	}
	// }(rest)

	// <-quit
	// server.Logger.Info("Server Exited Properly")
}
