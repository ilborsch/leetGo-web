
# LeetGo: tiny Gin, GORM, gRPC, html templating example web-app 


Welcome to LeetGo, a tiny web-service where people can enhance their programming skills by solving problems and reading articles related to the fields needed in orded to succeed.
LeetGo is the environment where people have freedom of creating their own tags, topics, problems and share their thoughts in articles.
The project is built using Go language leveraging functionality of its Gin, GORM frameworks and templating libraries.
The project is an example of gRPC client project as well.

<p align="center">
    <img style="width: 200px;" src="https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png" alt="logo">
</p>

**Note**: The project is created exclusively in educational purposes so will not be maintained properly. I don't plan on making further updates.

## Run Locally

Clone the project

```bash
  git clone https://github.com/ilborsch/leetGo-web
```

Go to the project directory

```bash
  cd leetGo-web
```

Install Go dependencies

```bash
  go get
```

Start the server manually

```bash
   go run ./cmd/webapp/main.go --config=./config/local.yaml
```

Or start the server using make
```bash
   make run
```





## Authors

- [@ilborsch](https://www.github.com/ilborsch)


## License

[MIT](https://choosealicense.com/licenses/mit/)