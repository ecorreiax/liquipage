# Liquipage
Liquipage is a tiny Go package that generates a single static HTML page from Markdown files that can be hosted on Github pages for small documentation. It is meant to be easy, simple, and fast.

## Getting started
To run Liquipage CLI using a Docker container:

```bash
docker run --rm -it liquipage --dir docs --output docs
```

You can pass a `--dir` flag indicating where all the Markdown files are stored and an `--output` flag indicating where the HTML files should be generated. 

If no flags are provided, both cases will use the current directory where the command is being executed.

## License
This project is under the [MIT License](https://github.com/ecorreiax/liquipage/blob/main/LICENSE)