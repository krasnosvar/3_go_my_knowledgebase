# gRPC services

gRPC — RPC-фреймворк на основе HTTP/2 и Protocol Buffers. Стандарт для межсервисного взаимодействия в Go-микросервисах.

## Когда выбирать

- Service-to-service внутри организации.
- Строго типизированный контракт.
- Высокая нагрузка, важна скорость.
- Streaming (server-streaming, client-streaming, bi-directional).

## Когда НЕ gRPC

- Публичный API для веб-фронта или сторонних разработчиков → REST/JSON.
- Простые случаи, где REST хватает с запасом.
- Окружение, где нет HTTP/2 (старые прокси, некоторые CDN).

## Стек

- `google.golang.org/grpc` — основной runtime.
- `google.golang.org/protobuf` — proto runtime.
- `protoc` + `protoc-gen-go` + `protoc-gen-go-grpc` — генерация кода.
- `buf` — современный `protoc`-обёртка с linter'ом и breaking-change-detector'ом.

## Proto contract

```proto
syntax = "proto3";
package user.v1;

service UserService {
  rpc Get (GetRequest) returns (GetResponse);
  rpc List (ListRequest) returns (stream UserMessage);
}

message GetRequest { int64 id = 1; }
message GetResponse { UserMessage user = 1; }
message UserMessage {
  int64 id = 1;
  string email = 2;
}
```

Поля имеют номер — это часть бинарного протокола. Менять номера = breaking change.

## Server skeleton

```go
type userServer struct {
    user.UnimplementedUserServiceServer
    svc *service.UserService
}

func (s *userServer) Get(ctx context.Context, req *user.GetRequest) (*user.GetResponse, error) {
    u, err := s.svc.Get(ctx, req.Id)
    if err != nil {
        if errors.Is(err, ErrNotFound) {
            return nil, status.Errorf(codes.NotFound, "user %d not found", req.Id)
        }
        return nil, status.Errorf(codes.Internal, "get user: %v", err)
    }
    return &user.GetResponse{User: toProto(u)}, nil
}

func main() {
    lis, _ := net.Listen("tcp", ":50051")
    s := grpc.NewServer(
        grpc.UnaryInterceptor(loggingInterceptor),
    )
    user.RegisterUserServiceServer(s, &userServer{svc: ...})
    s.Serve(lis)
}
```

## Status codes

`codes.OK`, `NotFound`, `InvalidArgument`, `PermissionDenied`, `Internal`, `DeadlineExceeded`, `Unavailable`. Возвращай `status.Error(codes.X, msg)`, не голый Go-error.

## Interceptors

Аналог middleware:

- Unary: `grpc.UnaryInterceptor`.
- Stream: `grpc.StreamInterceptor`.

Используются для логов, метрик, трассировки, auth, recovery.

## Контекст и deadlines

Каждый RPC имеет дедлайн от клиента (если он его поставил). На сервере — `ctx.Deadline()` показывает, сколько осталось.

Клиент:

```go
ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
defer cancel()
resp, err := client.Get(ctx, &req)
```

## Что отдельно

- `proto3` имеет особенности с zero-value полями (нельзя отличить «не задано» от 0 без `optional`).
- gRPC error → нужен `status.FromError` для извлечения деталей.
- Versioning: `package user.v1; user.v2;` — рядом, со временем deprecation v1.

## Правило

- Чёткий proto-contract → код-ген на обеих сторонах.
- Status codes из `codes.*`, не свой.
- Дедлайны на стороне клиента обязательны.
- Interceptors для cross-cutting.
