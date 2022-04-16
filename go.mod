module github.com/wrs-news/bfb-user-microservice

go 1.18

require (
	github.com/oklog/oklog v0.3.2
	github.com/spf13/cobra v1.4.0
	google.golang.org/grpc v1.45.0
)

require (
	github.com/twinj/uuid v1.0.0
	github.com/wrs-news/golang-proto v0.3.1
	google.golang.org/protobuf v1.28.0
)

require (
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/myesui/uuid v1.0.0 // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.0.0-20220325170049-de3da57026de // indirect
	golang.org/x/sys v0.0.0-20220325203850-36772127a21f // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220324131243-acbaeb5b85eb // indirect
	gopkg.in/stretchr/testify.v1 v1.2.2 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)

require (
	github.com/BurntSushi/toml v1.0.0
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/lib/pq v1.10.4
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/testify v1.7.1
	golang.org/x/crypto v0.0.0-20220321153916-2c7772ba3064
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
)

replace google.golang.org/genproto => ./libs/genproto
