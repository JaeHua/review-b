package data

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/hashicorp/consul/api"
	v1 "review-b/api/review/v1"
	"review-b/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewDicovery, NewData, NewBusinessRepo, NewReviewServiceClient)

// Data .
type Data struct {
	// TODO wrapped database client

	// 嵌入 RPC 客户端，通过这个客户端去调用 review-service 的服务
	rc  v1.ReviewClient
	log *log.Helper
}

// NewData .
func NewData(c *conf.Data, rc v1.ReviewClient, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		rc:  rc,
		log: log.NewHelper(logger),
	}, cleanup, nil
}

func NewDicovery(conf *conf.Registry) registry.Discovery {
	// TODO 创建服务发现实例

	//new consul client
	c := api.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	client, err := api.NewClient(c)
	if err != nil {
		panic(err)
	}
	// new disc with consul client
	dis := consul.New(client, consul.WithHealthCheck(true))
	return dis
}

func NewReviewServiceClient(d registry.Discovery) v1.ReviewClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		//grpc.WithEndpoint("127.0.0.1:9002"),
		grpc.WithEndpoint("discovery:///review.service"),
		grpc.WithDiscovery(d),
		grpc.WithMiddleware(
			recovery.Recovery(),
			validate.Validator()),
	)
	if err != nil {
		panic(err)
	}
	return v1.NewReviewClient(conn)
}
