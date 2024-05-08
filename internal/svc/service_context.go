package svc

import (
	"github.com/kebin6/simple-file-api/internal/config"
	i18n2 "github.com/kebin6/simple-file-api/internal/i18n"
	"github.com/kebin6/simple-file-api/internal/middleware"
	"github.com/kebin6/simple-file-api/internal/utils/cloud"
	"github.com/suyuan32/simple-admin-core/rpc/coreclient"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/kebin6/simple-file-api/ent"
	_ "github.com/kebin6/simple-file-api/ent/runtime"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config       config.Config
	Casbin       *casbin.Enforcer
	Authority    rest.Middleware
	DB           *ent.Client
	Trans        *i18n.Translator
	CoreRpc      coreclient.Core
	CloudStorage *cloud.CloudServiceGroup
}

func NewServiceContext(c config.Config) *ServiceContext {

	rds := c.RedisConf.MustNewUniversalRedis()

	cbn := c.CasbinConf.MustNewCasbinWithOriginalRedisWatcher(c.CasbinDatabaseConf.Type, c.CasbinDatabaseConf.GetDSN(), c.RedisConf)

	trans := i18n.NewTranslator(i18n2.LocaleFS)

	db := ent.NewClient(
		ent.Log(logx.Info), // logger
		ent.Driver(c.DatabaseConf.NewNoCacheDriver()),
		ent.Debug(), // debug mode
	)

	return &ServiceContext{
		Config:       c,
		Authority:    middleware.NewAuthorityMiddleware(cbn, rds, trans).Handle,
		Trans:        trans,
		DB:           db,
		CoreRpc:      coreclient.NewCore(zrpc.MustNewClient(c.CoreRpc)),
		CloudStorage: cloud.NewCloudServiceGroup(db),
	}
}
