module github.com/kuaukutsu/auth/sso

go 1.25.6

require (
	github.com/brianvoe/gofakeit/v6 v6.28.0
	github.com/golang-jwt/jwt/v5 v5.3.1
	github.com/google/uuid v1.6.0
	github.com/ilyakaznacheev/cleanenv v1.5.0
	github.com/kuaukutsu/auth/protos v0.0.0
	github.com/mattn/go-sqlite3 v1.14.33
	github.com/stretchr/testify v1.11.1
	golang.org/x/crypto v0.47.0
	google.golang.org/grpc v1.78.0
)

replace github.com/kuaukutsu/auth/protos => ../protos

require (
	github.com/BurntSushi/toml v1.2.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.48.0 // indirect
	golang.org/x/sys v0.40.0 // indirect
	golang.org/x/text v0.33.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251029180050-ab9386a59fda // indirect
	google.golang.org/protobuf v1.36.11 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)
