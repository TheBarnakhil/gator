package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/TheBarnakhil/gator/internal/config"
	"github.com/TheBarnakhil/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	db, err := sql.Open("postgres", cfg.DbURL)
	dbQueries := database.New(db)

	if err != nil {
		log.Fatalf("Error reading config file : %v", err)
	}
	s := &state{
		cfg: &cfg,
		db:  dbQueries,
	}
	cmds := &commands{
		cmds: map[string]func(*state, command) error{},
	}
	cmds.register("login", handleLogin)
	cmds.register("register", handleRegister)
	args := os.Args
	if len(args) < 2 {
		log.Fatal("At least two arguments expected!")
	}
	cmd_name, cmd_args := args[1], args[2:]
	cmd := command{
		name: cmd_name,
		args: cmd_args,
	}
	if err := cmds.run(s, cmd); err != nil {
		log.Fatal(err)
	}
}
