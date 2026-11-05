package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/AtlasRegap/gator/internal/config"
	"github.com/AtlasRegap/gator/internal/database"
	_ "github.com/lib/pq" // registers the "postgres" driver; never called directly
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	dbQueries := database.New(db)

	programState := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cmds := commands{registeredCommands: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
        cmds.register("users", handlerListUsers)
        cmds.register("agg", handlerAgg)
        cmds.register("addfeed", handlerAddFeed)
        cmds.register("feeds", handlerListFeeds)
        cmds.register("follow", handlerFollow)
        cmds.register("following", handlerListFeedFollows)
        cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
        cmds.register("browse", middlewareLoggedIn(handlerBrowse))

	if len(os.Args) < 2 {
		log.Fatal("not enough arguments: expected a command")
	}
	cmd := command{Name: os.Args[1], Args: os.Args[2:]}

	if err := cmds.run(programState, cmd); err != nil {
		log.Fatal(err)
	}
}
