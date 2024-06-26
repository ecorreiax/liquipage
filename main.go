package main

import (
	"errors"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type Config struct {
	Dir  string
	Out  string
	File string
}

type Layout struct {
	Title   string
	Content []template.HTML
}

var AppConfig = Config{
	Dir:  "./",
	Out:  "./",
	File: "",
}

var AppLayout = Layout{
	Title:   "Liquipage",
	Content: []template.HTML{},
}

func main() {
	app := cli.NewApp()
	app.Name = "liquipage"
	app.Usage = "Tiny static site generator from markdown files."
	app.Version = "0.0.0"
	app.UsageText = "liquipage --dir path/to/docs"
	app.Authors = []*cli.Author{
		{
			Name:  "Eduardo Correia",
			Email: "ecorreiax@gmail.com",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "Generates a static page from markdown files in directory",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "dir",
					Value:       "./",
					Aliases:     []string{"d"},
					Usage:       "directory containing the markdown files",
					Destination: &AppConfig.Dir,
				},
				&cli.StringFlag{
					Name:        "out",
					Value:       "./",
					Aliases:     []string{"o"},
					Usage:       "output directory for the html file",
					Destination: &AppConfig.Out,
				},
				&cli.StringFlag{
					Name:        "file",
					Aliases:     []string{"f"},
					Usage:       "a single file",
					Destination: &AppConfig.File,
				},
				&cli.StringFlag{
					Name:        "title",
					Aliases:     []string{"t"},
					Usage:       "title for you static page",
					Destination: &AppLayout.Title,
				},
			},
			Action: func(c *cli.Context) error {
				if c.NArg() > 0 {
					return errors.New("too many args where given")
				}

				paths, err := getMDFiles()
				if err != nil {
					return err
				}

				mds := []string{}

				for _, v := range paths {
					content, err := getContentFromMDFile(v)
					if err != nil {
						return err
					}

					if len(content) == 0 {
						continue
					}

					mds = append(mds, content)
				}

				err = generateHTMLFile(mds)
				if err != nil {
					return err
				}

				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func getMDFiles() ([]string, error) {
	paths := []string{}

	if len(AppConfig.File) > 0 {
		paths = append(paths, AppConfig.File)
		return paths, nil
	}

	err := filepath.Walk(AppConfig.Dir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".md" {
			paths = append(paths, path)
		}
		return err
	})
	if err != nil {
		return nil, err
	}

	if len(paths) == 0 {
		return nil, errors.New("no markdown files found")
	}

	return paths, nil
}

func getContentFromMDFile(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	content := string(b)

	return content, nil
}

func generateHTMLFile(mds []string) error {
	for _, v := range mds {
		md := []byte(v)
		html := convertMDToHTML(md)
		AppLayout.Content = append(AppLayout.Content, template.HTML(html))
	}

	tmpl := template.Must(template.ParseFiles("layout.html"))

	if AppConfig.Out != "./" {
		if _, err := os.Stat(AppConfig.Out); os.IsNotExist(err) {
			err := os.MkdirAll(AppConfig.Out, 0755)
			if err != nil {
				return err
			}
		}
	}

	filepath := filepath.Join(AppConfig.Out, "index.html")
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = tmpl.Execute(file, AppLayout)
	if err != nil {
		return err
	}
	return nil
}

func convertMDToHTML(md []byte) []byte {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}
