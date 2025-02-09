package api

import (
	"github.com/opentracing-contrib/go-gin/ginhttp"
	"github.com/opentracing/opentracing-go"
	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	reporter "github.com/openzipkin/zipkin-go/reporter/http"
	"net/http"

	//"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin"
)

// opentracingMiddleware 上报中间件 上报地址：http://localhost:9411/api/v2/spans
func opentracingMiddleware() gin.HandlerFunc {
	// 创建Reporter
	report := reporter.NewReporter("http://localhost:9411/api/v2/spans")
	// 创建本地节点
	endpoint, err := zipkin.NewEndpoint("content-system", "localhost:8080")
	if err != nil {
		panic(err)
	}
	// 创建Zipkin Tracer
	tracer, err := zipkin.NewTracer(report,
		zipkin.WithLocalEndpoint(endpoint),
		zipkin.WithTraceID128Bit(true))
	if err != nil {
		panic(err)
	}
	zipTracer := zipkinot.Wrap(tracer)
	opentracing.SetGlobalTracer(zipTracer)
	// 创建中间件
	return ginhttp.Middleware(zipTracer, ginhttp.OperationNameFunc(func(r *http.Request) string {
		return r.URL.Path
	}))
}
