module git.merostm.com/scm/goq/go-sse-quote-service.git

go 1.14

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis/v8 v8.3.2
	github.com/googollee/go-socket.io v1.4.4
	github.com/joho/godotenv v1.3.0
	github.com/r3labs/sse v0.0.0-20201007160420-c638e5516aa7
)

replace git.merostm.com/scm/goq/go-sse-quote-service.git/sse-quote => ./sse-quote
