package service

import (
	"content-system/internal/api/content"
	"context"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type CmsApp struct {
	db         *gorm.DB
	rdb        *redis.Client
	contentRpc content.AppClient
}

func NewCmsApp() *CmsApp {
	app := &CmsApp{}
	connDB(app)
	connRdb(app)
	return app
}

func connDB(app *CmsApp) {
	mysqlDB, err := gorm.Open(mysql.Open("root:123456789@tcp(ro)/?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	db, err := mysqlDB.DB()
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(4)
	db.SetMaxIdleConns(2)
	app.db = mysqlDB
}

func connRpcClient(app *CmsApp) {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithMiddleware(
			recovery.Recovery()))

	if err != nil {
		panic(err)
	}
	rpcClient := content.NewAppClient(conn)
	app.contentRpc = rpcClient
}

func connRdb(app *CmsApp) {
	// redis-cli
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	app.rdb = rdb
}
