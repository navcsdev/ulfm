# Install Dependencis

go get github.com/gin-gonic/gin
go get github.com/joho/godotenv
go get go.mongodb.org/mongo-driver/mongo

# Start mongo
 
`cd db`
`docker-compose up -d`

# Start app

`go run main`