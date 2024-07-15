package grpcapp

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"sync"

	grpccontent "github.com/PolyAbit/content/internal/grpc/content"
	"github.com/PolyAbit/content/internal/lib/logger/sl"
	middleware "github.com/PolyAbit/content/internal/lib/middlewares"
	contentv1 "github.com/PolyAbit/protos/gen/go/content"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	grpcPort   int
	httpPort   int
}

func New(log *slog.Logger, contentService grpccontent.Content, permProvider middleware.PermissionProvider, gRPCPort int, httpPort int, jwtSecret string) *App {
	authMiddleware := &middleware.AuthMiddleware{AuthFunc: middleware.New(jwtSecret, permProvider)}

	unaryInterceptors := []grpc.UnaryServerInterceptor{
		recovery.UnaryServerInterceptor(),
		authMiddleware.UnaryInterceptor,
	}

	gRPCServer := grpc.NewServer(grpc.ChainUnaryInterceptor(unaryInterceptors...))

	grpccontent.Register(gRPCServer, contentService)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		grpcPort:   gRPCPort,
		httpPort:   httpPort,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "grpcapp.Run"

	log := a.log.With(slog.String("op", op))

	ctx := context.Background()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		if err := a.startGrpcServer(); err != nil {
			log.Error("failed start grpc", sl.Err(err))
		}
	}()

	go func() {
		defer wg.Done()

		if err := a.startHttpServer(ctx); err != nil {
			log.Error("failed start http", sl.Err(err))
		}
	}()

	wg.Wait()

	return nil
}

func (a *App) startGrpcServer() error {
	const op = "app.grpc.startGrpcServer"

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.grpcPort))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	a.log.Info("grpc server started", slog.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func allowedOrigin(origin string) bool {
	return true
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if allowedOrigin(r.Header.Get("Origin")) {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
		}
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}

func (a *App) startHttpServer(ctx context.Context) error {
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	err := contentv1.RegisterContentHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%d", a.grpcPort), opts)
	if err != nil {
		return err
	}

	a.log.Info("gateway server started", slog.String("addr", fmt.Sprintf(":%d", a.httpPort)))

	return http.ListenAndServe(fmt.Sprintf("localhost:%d", a.httpPort), cors(mux))
}

func (a *App) Stop() {
	const op = "grpcapp.Stop"

	a.log.Info("stopping gRPC server", slog.Int("port", a.grpcPort), slog.String("op", op))

	a.gRPCServer.GracefulStop()
}
