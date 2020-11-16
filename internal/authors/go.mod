module github.com/elsagg/graphapis/internal/authors

go 1.15

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/elsagg/graphapis/pkg v0.0.0-20201116221635-d91710a526bf
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-contrib/logger v0.0.2
	github.com/gin-gonic/gin v1.6.3
	github.com/joho/godotenv v1.3.0
	github.com/rs/zerolog v1.20.0
	github.com/vektah/gqlparser/v2 v2.1.0
	go.mongodb.org/mongo-driver v1.4.1
)

// replace github.com/elsagg/graphapis/pkg => ../../pkg
