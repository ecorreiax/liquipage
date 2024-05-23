# Liquipage
Liquipage is a tiny Go package that generates a single static HTML page from Markdown files that can be hosted on Github pages for small documentation. It is meant to be easy, simple, and fast.

## Getting started
To run Liquipage CLI using a Docker container:

```bash
docker run --rm -it liquipage build --dir docs --output docs --title Liquipage
```
Flags can be passed to change some default behaviors. If no flags are provided, `--dir` and `--out` will both use the current directory where the command is being executed.

| Flag    | Description                                              |
| :-------| :------------------------------------------------------- |
| --dir   | Source of the markdown files                             |
| --out   | Output where the HTML should be generated                |
| --title | Title that goes in the HTML title tag, header and footer |

## License
This project is under the [MIT License](https://github.com/ecorreiax/liquipage/blob/main/LICENSE)