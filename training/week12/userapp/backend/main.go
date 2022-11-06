package main

import (
	"github.com/go-redis/redis/v9"
	"gitlab.xchch.top/zhangsai/go-101/training/cache"
	"gitlab.xchch.top/zhangsai/go-101/training/web/session"
	"gitlab.xchch.top/zhangsai/go-101/training/web/session/cookie"
	rstore "gitlab.xchch.top/zhangsai/go-101/training/web/session/redis"
	"gitlab.xchch.top/zhangsai/go-101/training/week12/userapp/backend/internal/repository"
	"gitlab.xchch.top/zhangsai/go-101/training/week12/userapp/backend/internal/repository/dao"
	"gitlab.xchch.top/zhangsai/go-101/training/week12/userapp/backend/internal/service"
	"gitlab.xchch.top/zhangsai/go-101/training/week12/userapp/backend/internal/web/handler"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"go.uber.org/zap"
	"log"
	_ "net/http/pprof"
	"os"
)

// 这里各种地址都是直接写死的，在真实的环境替换为从配置文件里面读取就可以
// 随便用一个配置框架，大体上都差不多的
func main() {
	// 在 main 函数的入口里面完成所有的依赖组装。
	// 这个部分你可以考虑替换为 google 的 wire 框架，达成依赖注入的效果
	l, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(l)

	db := initDB()
	rc := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "abc",
	})
	c := cache.NewRedisCache(rc)
	repo := repository.NewUserRepository(dao.NewUserDAO(db), c)
	userSvr := service.NewUserService(repo)
	sessMgr := session.Manager{
		Store:      rstore.NewStore(rc),
		Propagator: cookie.NewPropagator("sess_id"),
		SessCtxKey: "_sess",
	}

	initZipkin()

	// 路由注册和 middleware 注册，可以抽取出来作为一个单独的方法，也可以将路由注册部分下沉到 handler 包
	// 例如为 Handler 定义一个新的方法，该方法会注册所有的路由
	// 我一般喜欢在一个集中的地方注册
	userHdl := handler.NewUserHandler(userSvr, sessMgr)
	server := initSever()
	server.Post("/signup", userHdl.SignUp)
	server.Post("/login", userHdl.Login)
	server.Get("/profile", userHdl.Profile)
	server.Post("/update", userHdl.Update)

	if err = server.Start(":8081"); err != nil {
		panic(err)
	}
}

func initZipkin() {
	exporter, err := zipkin.New(
		"http://localhost:19411/api/v2/spans",
		zipkin.WithLogger(log.New(os.Stderr, "userapp", log.Ldate|log.Ltime|log.Llongfile)),
	)
	if err != nil {
		panic(err)
	}
	batcher := sdktrace.NewBatchSpanProcessor(exporter)
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSpanProcessor(batcher),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("userapp"),
		)),
	)
	otel.SetTracerProvider(tp)
}
