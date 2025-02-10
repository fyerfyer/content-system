package service

import (
	"content-system/internal/api/content"
	"context"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/opentracing/opentracing-go"
	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	reporter "github.com/openzipkin/zipkin-go/reporter/http"
	"github.com/redis/go-redis/v9"
	clientv3 "go.etcd.io/etcd/client/v3"
	gormopentracing "gorm.io/plugin/opentracing"

	//goflow "github.com/s8sg/goflow/v1"
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
	mysqlDB, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	db, err := mysqlDB.DB()
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(4)
	db.SetMaxIdleConns(2)
	report := reporter.NewReporter("http://localhost:9411/api/v2/spans")
	endpoint, err := zipkin.NewEndpoint("content-system", "localhost:80")
	if err != nil {
		panic(err)
	}
	tracer, err := zipkin.NewTracer(report,
		zipkin.WithTraceID128Bit(true),
		zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		panic(err)
	}

	zipkinTracer := zipkinot.Wrap(tracer)
	opentracing.SetGlobalTracer(zipkinTracer)
	err = mysqlDB.Use(gormopentracing.New(gormopentracing.WithTracer(zipkinTracer)))
	if err != nil {
		panic(err)
	}
	app.db = mysqlDB
}

func connRpcClient(app *CmsApp) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	})

	if err != nil {
		panic(err)
	}

	discovery := etcd.New(client)
	endpoint := "discovery:///content-system"
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
		grpc.WithEndpoint(endpoint),
		grpc.WithDiscovery(discovery),
	)

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
