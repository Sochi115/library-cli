package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/Sochi115/library-cli/db"
	"github.com/Sochi115/library-cli/save"
	"github.com/Sochi115/library-cli/search"
)

func main() {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	sqliteDb := db.ConnectToSqliteDb(dirname + "\\library_cli.db")
	sqliteDb.InitBookTable()
	defer sqliteDb.CloseDb()

	var listFlag string
	var id int

	bookFlag := &cli.StringFlag{
		Name:    "book",
		Aliases: []string{"b"},
		Usage:   "Specifies book name `BOOK`",
	}

	isbnFlag := &cli.StringFlag{
		Name:    "isbn",
		Aliases: []string{"i"},
		Usage:   "Specifies isbn number `ISBN`",
	}

	authorFlag := &cli.StringFlag{
		Name:    "author",
		Aliases: []string{"a"},
		Usage:   "Specifies author name `AUTHOR`",
	}

	worksFlag := &cli.StringFlag{
		Name:    "works",
		Aliases: []string{"w"},
		Usage:   "Specifies OpenLibrary's works ID string `WORKS`",
	}

	app := &cli.App{
		Name:  "library-cli",
		Usage: "Cli tool to manage and search for books",
		Commands: []*cli.Command{
			{
				Name:  "save",
				Usage: "Saves book",
				Flags: []cli.Flag{
					worksFlag,
					isbnFlag,
				},
				Action: func(ctx *cli.Context) error {
					works := ctx.String("works")
					isbn := ctx.String("isbn")

					if len(works) > 0 {
						book := save.HandleSaveBookByWorks(works)
						sqliteDb.AddBook(book)
						return nil
					}

					if len(isbn) > 0 {
						book := save.HandleSaveBookByIsbn(isbn)
						sqliteDb.AddBook(book)
						return nil
					}

					fmt.Fprintln(
						os.Stderr,
						"No valid command options detected\nSee `library-cli save --help` for a usage guide",
					)
					return nil
				},
			},
			{
				Name:  "search",
				Usage: "Retrieves data of book or author",
				Flags: []cli.Flag{
					bookFlag,
					authorFlag,
				},
				SkipFlagParsing: false,
				OnUsageError: func(cCtx *cli.Context, err error, isSubcommand bool) error {
					if isSubcommand {
						return err
					}
					fmt.Fprintf(cCtx.App.Writer, "Usage error")
					return err
				},
				Action: func(ctx *cli.Context) error {
					book := ctx.String("book")
					author := ctx.String("author")

					if len(book) > 0 {
						fmt.Println("Searching for titles:", book)
						search.GetBookDataByTitle(book)
						return nil
					}

					if len(author) > 0 {
						fmt.Println("Searching for authors:", author)
						search.GetAuthorData(author)
						return nil
					}

					fmt.Fprintln(
						os.Stderr,
						"No valid command options detected\nSee `library-cli search --help` for a usage guide",
					)
					return nil
				},
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "list",
				Value:       "all",
				Usage:       "List personal books within the library",
				Aliases:     []string{"l"},
				Destination: &listFlag,
			},
			&cli.IntFlag{
				Name:        "mark-read",
				Usage:       "Mark book with id: `ID` as read",
				Aliases:     []string{"mr"},
				Destination: &id,
			},
			&cli.IntFlag{
				Name:        "delete",
				Usage:       "Delete book with id: `ID` from the database",
				Aliases:     []string{"d"},
				Destination: &id,
			},
			&cli.IntFlag{
				Name:        "rate",
				Usage:       "Rate book with id: `ID`",
				Aliases:     []string{"r"},
				Destination: &id,
			},
		},
		Action: func(ctx *cli.Context) error {
			fmt.Println("Hello World")
			if ctx.String("list") == "all" {
				sqliteDb.FetchAll()
				return nil
			}
			return nil
		},
	}
	app.Suggest = true

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
