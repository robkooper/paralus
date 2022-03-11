package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net"
	"net/http"
	goruntime "runtime"
	"strings"
	"sync"
	"time"

	"github.com/RafaySystems/rcloud-base/internal/fixtures"
	providers "github.com/RafaySystems/rcloud-base/internal/persistence/provider/kratos"
	authv3 "github.com/RafaySystems/rcloud-base/pkg/auth/v3"
	"github.com/RafaySystems/rcloud-base/pkg/common"
	"github.com/RafaySystems/rcloud-base/pkg/enforcer"
	"github.com/RafaySystems/rcloud-base/pkg/gateway"
	"github.com/RafaySystems/rcloud-base/pkg/grpc"
	"github.com/RafaySystems/rcloud-base/pkg/leaderelection"
	"github.com/RafaySystems/rcloud-base/pkg/log"
	"github.com/RafaySystems/rcloud-base/pkg/notify"
	"github.com/RafaySystems/rcloud-base/pkg/reconcile"
	"github.com/RafaySystems/rcloud-base/pkg/sentry/peering"
	"github.com/RafaySystems/rcloud-base/pkg/service"
	auditrpc "github.com/RafaySystems/rcloud-base/proto/rpc/audit"
	rolerpc "github.com/RafaySystems/rcloud-base/proto/rpc/role"
	schedulerrpc "github.com/RafaySystems/rcloud-base/proto/rpc/scheduler"
	sentryrpc "github.com/RafaySystems/rcloud-base/proto/rpc/sentry"
	systemrpc "github.com/RafaySystems/rcloud-base/proto/rpc/system"
	userrpc "github.com/RafaySystems/rcloud-base/proto/rpc/user"
	"github.com/RafaySystems/rcloud-base/server"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	kclient "github.com/ory/kratos-client-go"
	"github.com/rs/xid"
	"github.com/spf13/viper"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	_grpc "google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
)

const (
	rpcPortEnv                = "RPC_PORT"
	apiPortEnv                = "API_PORT"
	debugPortEnv              = "DEBUG_PORT"
	dbAddrEnv                 = "DB_ADDR"
	dbNameEnv                 = "DB_NAME"
	dbUserEnv                 = "DB_USER"
	dbPasswordEnv             = "DB_PASSWORD"
	devEnv                    = "DEV"
	apiAddrEnv                = "API_ADDR"
	sentryPeeringHostEnv      = "SENTRY_PEERING_HOST"
	coreRelayConnectorHostEnv = "CORE_RELAY_CONNECTOR_HOST"
	coreRelayUserHostEnv      = "CORE_RELAY_USER_HOST"
	bootstrapKEKEnv           = "BOOTSTRAP_KEK"
	authAddrEnv               = "AUTH_ADDR"
	sentryBootstrapEnv        = "SENTRY_BOOTSTRAP_ADDR"
	relayImageEnv             = "RELAY_IMAGE"
	controlAddrEnv            = "CONTROL_ADDR"

	// audit
	esEndPointEnv              = "ES_END_POINT"
	esIndexPrefixEnv           = "ES_INDEX_PREFIX"
	relayAuditESIndexPrefixEnv = "RELAY_AUDITS_ES_INDEX_PREFIX"
	relayCommandESIndexPrefix  = "RELAY_COMMANDS_ES_INDEX_PREFIX"
	RpcPort                    = "AUDIT_RPC_PORT"
	ApiPort                    = "AUDIT_API_PORT"

	// cd relay
	coreCDRelayUserHostEnv      = "CORE_CD_RELAY_USER_HOST"
	coreCDRelayConnectorHostEnv = "CORE_CD_RELAY_CONNECTOR_HOST"
	schedulerNamespaceEnv       = "SCHEDULER_NAMESPACE"

	kratosSchemeEnv = "KRATOS_SCHEME"
	kratosAddrEnv   = "KRATOS_ADDR"
)

var (
	rpcPort                int
	apiPort                int
	debugPort              int
	rpcRelayPeeringPort    int
	dbAddr                 string
	dbName                 string
	dbUser                 string
	dbPassword             string
	apiAddr                string
	dev                    bool
	db                     *bun.DB
	gormDb                 *gorm.DB
	kratosScheme           string
	kratosAddr             string
	kc                     *kclient.APIClient
	ps                     service.PartnerService
	os                     service.OrganizationService
	pps                    service.ProjectService
	bs                     service.BootstrapService
	aps                    service.AccountPermissionService
	gps                    service.GroupPermissionService
	krs                    service.KubeconfigRevocationService
	kss                    service.KubeconfigSettingService
	kcs                    service.KubectlClusterSettingsService
	as                     service.AuthzService
	cs                     service.ClusterService
	ms                     service.MetroService
	us                     service.UserService
	gs                     service.GroupService
	rs                     service.RoleService
	rrs                    service.RolepermissionService
	is                     service.IdpService
	oidcs                  service.OIDCProviderService
	aus                    *service.AuditLogService
	ras                    *service.RelayAuditService
	rcs                    *service.AuditLogService
	_log                   = log.GetLogger()
	schedulerPool          schedulerrpc.SchedulerPool
	schedulerAddr          string
	bootstrapKEK           string
	sentryPeeringHost      string
	coreRelayConnectorHost string
	coreRelayUserHost      string
	downloadData           *common.DownloadData
	controlAddr            string

	// audit
	elasticSearchUrl           string
	esIndexPrefix              string
	relayAuditsESIndexPrefix   string
	relayCommandsESIndexPrefix string

	// cd relay
	coreCDRelayUserHost      string
	coreCDRelayConnectorHost string
	schedulerNamespace       string

	kekFunc = func() ([]byte, error) {
		if len(bootstrapKEK) == 0 {
			return nil, errors.New("empty KEK")
		}
		return []byte(bootstrapKEK), nil
	}
)

func setup() {
	viper.SetDefault(rpcPortEnv, 10000)
	viper.SetDefault(apiPortEnv, 11000)
	viper.SetDefault(debugPortEnv, 12000)
	viper.SetDefault(dbAddrEnv, "localhost:5432")
	viper.SetDefault(dbNameEnv, "admindb")
	viper.SetDefault(dbUserEnv, "admindbuser")
	viper.SetDefault(dbPasswordEnv, "admindbpassword")
	viper.SetDefault(devEnv, true)
	viper.SetDefault(apiAddrEnv, "localhost:11000")
	viper.SetDefault(kratosSchemeEnv, "http")
	viper.SetDefault(kratosAddrEnv, "localhost:4433")
	viper.SetDefault(sentryPeeringHostEnv, "peering.sentry.rafay.local:10001")
	viper.SetDefault(coreRelayConnectorHostEnv, "*.core-connector.relay.rafay.local:10002")
	viper.SetDefault(coreRelayUserHostEnv, "*.user.relay.rafay.local:10002")
	viper.SetDefault(bootstrapKEKEnv, "rafay")
	viper.SetDefault(authAddrEnv, "authsrv.rcloud-admin.svc.cluster.local:50011")
	viper.SetDefault(coreCDRelayUserHostEnv, "*.user.cdrelay.rafay.local:10012")
	viper.SetDefault(coreCDRelayConnectorHostEnv, "*.core-connector.cdrelay.rafay.local:10012")
	viper.SetDefault(sentryBootstrapEnv, "console.rafay.dev:443")
	viper.SetDefault(relayImageEnv, "registry.rafay-edge.net/rafay/rafay-relay-agent:r1.10.0-24")
	viper.SetDefault(controlAddrEnv, "localhost:5002")
	viper.SetDefault(schedulerNamespaceEnv, "rafay-system")
	viper.SetDefault(esEndPointEnv, "http://127.0.0.1:9200")
	viper.SetDefault(esIndexPrefixEnv, "auditlog-system")
	viper.SetDefault(relayAuditESIndexPrefixEnv, "auditlog-relay")
	viper.SetDefault(relayCommandESIndexPrefix, "auditlog-commands")

	viper.BindEnv(rpcPortEnv)
	viper.BindEnv(apiPortEnv)
	viper.BindEnv(debugPortEnv)
	viper.BindEnv(dbAddrEnv)
	viper.BindEnv(dbNameEnv)
	viper.BindEnv(dbUserEnv)
	viper.BindEnv(dbPasswordEnv)
	viper.BindEnv(devEnv)
	viper.BindEnv(apiAddrEnv)
	viper.BindEnv(kratosSchemeEnv)
	viper.BindEnv(kratosAddrEnv)
	viper.BindEnv(bootstrapKEKEnv)
	viper.BindEnv(authAddrEnv)
	viper.BindEnv(sentryPeeringHostEnv)
	viper.BindEnv(coreRelayConnectorHostEnv)
	viper.BindEnv(coreRelayUserHostEnv)
	viper.BindEnv(coreCDRelayConnectorHostEnv)
	viper.BindEnv(coreCDRelayUserHostEnv)
	viper.BindEnv(sentryBootstrapEnv)
	viper.BindEnv(relayImageEnv)
	viper.BindEnv(controlAddrEnv)
	viper.BindEnv(schedulerNamespaceEnv)
	viper.BindEnv(esEndPointEnv)
	viper.BindEnv(esIndexPrefixEnv)
	viper.BindEnv(relayAuditESIndexPrefixEnv)
	viper.BindEnv(relayCommandESIndexPrefix)

	rpcPort = viper.GetInt(rpcPortEnv)
	apiPort = viper.GetInt(apiPortEnv)
	debugPort = viper.GetInt(debugPortEnv)
	dbAddr = viper.GetString(dbAddrEnv)
	dbName = viper.GetString(dbNameEnv)
	dbUser = viper.GetString(dbUserEnv)
	dbPassword = viper.GetString(dbPasswordEnv)
	apiAddr = viper.GetString(apiAddrEnv)
	dev = viper.GetBool(devEnv)
	kratosScheme = viper.GetString(kratosSchemeEnv)
	kratosAddr = viper.GetString(kratosAddrEnv)
	bootstrapKEK = viper.GetString(bootstrapKEKEnv)
	sentryPeeringHost = viper.GetString(sentryPeeringHostEnv)
	coreRelayConnectorHost = viper.GetString(coreRelayConnectorHostEnv)
	coreRelayUserHost = viper.GetString(coreRelayUserHostEnv)
	coreCDRelayConnectorHost = viper.GetString(coreCDRelayConnectorHostEnv)
	coreCDRelayUserHost = viper.GetString(coreCDRelayUserHostEnv)
	controlAddr = viper.GetString(controlAddrEnv)
	schedulerNamespace = viper.GetString(schedulerNamespaceEnv)
	elasticSearchUrl = viper.GetString(esEndPointEnv)
	esIndexPrefix = viper.GetString(esIndexPrefixEnv)
	relayAuditsESIndexPrefix = viper.GetString(relayAuditESIndexPrefixEnv)
	relayCommandsESIndexPrefix = viper.GetString(relayCommandESIndexPrefix)

	rpcRelayPeeringPort = rpcPort + 1

	// Kratos client setup
	kratosConfig := kclient.NewConfiguration()
	kratosUrl := kratosScheme + "://" + kratosAddr
	kratosConfig.Servers[0].URL = kratosUrl
	kc = kclient.NewAPIClient(kratosConfig)

	// db setup
	dsn := "postgres://" + dbUser + ":" + dbPassword + "@" + dbAddr + "/" + dbName + "?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db = bun.NewDB(sqldb, pgdialect.New())

	if dev {
		db.AddQueryHook(bundebug.NewQueryHook(
			bundebug.WithVerbose(true),
			bundebug.FromEnv("BUNDEBUG"),
		))
		lc := make(chan string)
		go _log.ChangeLevel(lc)
		lc <- "debug"
		_log.Debugw("Debug mode set in log because this is a dev environment")
	}

	_log.Infow("printing db", "db", db)

	schedulerPool = schedulerrpc.NewSchedulerPool(schedulerAddr, 5*goruntime.NumCPU())

	ps = service.NewPartnerService(db)
	os = service.NewOrganizationService(db)
	pps = service.NewProjectService(db)

	// authz services
	gormDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		_log.Fatalw("unable to create db connection", "error", err)
	}
	enforcer, err := enforcer.NewCasbinEnforcer(gormDb).Init()
	if err != nil {
		_log.Fatalw("unable to init enforcer", "error", err)
	}
	as = service.NewAuthzService(gormDb, enforcer)

	// users and role management services
	us = service.NewUserService(providers.NewKratosAuthProvider(kc), db, as)
	gs = service.NewGroupService(db, as)
	rs = service.NewRoleService(db, as)
	rrs = service.NewRolepermissionService(db)
	is = service.NewIdpService(db, apiAddr)
	oidcs = service.NewOIDCProviderService(db, kratosUrl)

	//sentry related services
	bs = service.NewBootstrapService(db)
	krs = service.NewKubeconfigRevocationService(db)
	kss = service.NewKubeconfigSettingService(db)
	kcs = service.NewkubectlClusterSettingsService(db)
	aps = service.NewAccountPermissionService(db)
	gps = service.NewGroupPermissionService(db)

	// audit services
	aus, err = service.NewAuditLogService(elasticSearchUrl, esIndexPrefix+"-*", "AuditLog API: ")
	if err != nil {
		if dev && strings.Contains(err.Error(), "connect: connection refused") {
			// This is primarily from ES not being available. ES being
			// pretty heavy, you might not always wanna have it
			// running in the background. This way, you can continue
			// working on rcloud-base with ES eating up all the cpu.
			_log.Warn("unable to create auditLog service: ", err)
		} else {
			_log.Fatalw("unable to create auditLog service", "error", err)
		}
	}
	ras, err = service.NewRelayAuditService(elasticSearchUrl, relayAuditsESIndexPrefix+"-*", "RelayAudit API: ")
	if err != nil {
		if dev && strings.Contains(err.Error(), "connect: connection refused") {
			_log.Warn("unable to create relayAudit service: ", err)
		} else {
			_log.Fatalw("unable to create relayAudit service", "error", err)
		}
	}
	rcs, err = service.NewAuditLogService(elasticSearchUrl, relayCommandsESIndexPrefix+"-*", "RelayCommand API: ")
	if err != nil {
		if dev && strings.Contains(err.Error(), "connect: connection refused") {
			_log.Warn("unable to create auditLog service:", err)
		} else {
			_log.Fatalw("unable to create auditLog service", "error", err)
		}
	}

	// cluster bootstrap
	downloadData = &common.DownloadData{
		ControlAddr:     controlAddr,
		APIAddr:         apiAddr,
		RelayAgentImage: relayImageEnv,
	}

	cs = service.NewClusterService(db, downloadData, bs)
	ms = service.NewMetroService(db)

	notify.Init(cs)

	_log.Infow("queried number of cpus", "numCPUs", goruntime.NumCPU())
}

func run() {

	ctx := signals.SetupSignalHandler()

	notify.Start(ctx.Done())

	replace := map[string]interface{}{
		"sentryPeeringHost":   sentryPeeringHost,
		"coreRelayServerHost": coreRelayConnectorHost,
		"coreRelayUserHost":   coreRelayUserHost,

		// cd relay
		"coreCDRelayUserHost":      coreCDRelayUserHost,
		"coreCDRelayConnectorHost": coreCDRelayConnectorHost,
	}

	_log.Infow("loading fixtures", "data", replace)

	fixtures.Load(ctx, bs, replace, kekFunc)

	healthServer, err := grpc.NewServer()
	if err != nil {
		_log.Infow("failed to initialize grpc for health server")
	}
	// health server
	_log.Infow("registering grpc health server")
	hs := health.NewServer()
	hs.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)
	grpc_health_v1.RegisterHealthServer(healthServer, hs)
	_log.Infow("registered grpc health server")

	var wg sync.WaitGroup
	wg.Add(5)

	go runAPI(&wg, ctx)
	go runRPC(&wg, ctx)
	go runRelayPeerRPC(&wg, ctx)
	go runDebug(&wg, ctx)
	go runEventHandlers(&wg, ctx)

	<-ctx.Done()
	_log.Infow("shutting down, waiting for children to die")
	wg.Wait()

}

func runAPI(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()

	gwHandler, err := gateway.NewGateway(
		ctx,
		fmt.Sprintf(":%d", rpcPort),
		make([]runtime.ServeMuxOption, 0),
		systemrpc.RegisterPartnerHandlerFromEndpoint,
		systemrpc.RegisterOrganizationHandlerFromEndpoint,
		systemrpc.RegisterProjectHandlerFromEndpoint,
		sentryrpc.RegisterBootstrapHandlerFromEndpoint,
		sentryrpc.RegisterKubeConfigHandlerFromEndpoint,
		sentryrpc.RegisterKubectlClusterSettingsHandlerFromEndpoint,
		sentryrpc.RegisterClusterAuthorizationHandlerFromEndpoint,
		schedulerrpc.RegisterClusterHandlerFromEndpoint,
		systemrpc.RegisterLocationHandlerFromEndpoint,
		userrpc.RegisterUserHandlerFromEndpoint,
		userrpc.RegisterGroupHandlerFromEndpoint,
		rolerpc.RegisterRoleHandlerFromEndpoint,
		rolerpc.RegisterRolepermissionHandlerFromEndpoint,
		systemrpc.RegisterIdpHandlerFromEndpoint,
		systemrpc.RegisterOIDCProviderHandlerFromEndpoint,
		auditrpc.RegisterAuditLogHandlerFromEndpoint,
		auditrpc.RegisterRelayAuditHandlerFromEndpoint,
	)
	if err != nil {
		_log.Fatalw("unable to create gateway", "error", err)
	}
	mux.Handle("/", gwHandler)

	s := http.Server{
		Addr:    fmt.Sprintf(":%d", apiPort),
		Handler: mux,
	}
	go func() {
		defer s.Shutdown(context.TODO())
		<-ctx.Done()
	}()

	_log.Infow("starting gateway server", "port", apiPort)
	err = s.ListenAndServe()
	if err != nil {
		_log.Fatalw("unable to start gateway", "error", err)
	}

}

func runRelayPeerRPC(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()

	_log.Infow("waiting to fetch peering server creds")
	time.Sleep(time.Second * 25)
	cert, key, ca, err := peering.GetPeeringServerCreds(context.Background(), bs, rpcPort, sentryPeeringHost)
	if err != nil {
		_log.Fatalw("unable to get peering server cerds", "error", err)
	}

	relayPeerService, err := server.NewRelayPeerService()
	if err != nil {
		_log.Fatalw("unable to get create relay peer service")
	}
	clusterAuthzServer := server.NewClusterAuthzServer(bs, aps, gps, krs, kcs, kss)

	/*
		auditInfoServer := server.NewAuditInfoServer(bs, aps)
	*/

	s, err := grpc.NewSecureServerWithPEM(cert, key, ca)
	if err != nil {
		_log.Fatalw("cannot grpc secure server failed", "error", err)

	}

	go func() {
		defer s.GracefulStop()

		<-ctx.Done()
		_log.Infow("peer service stopped due to context done")
	}()

	sentryrpc.RegisterRelayPeerServiceServer(s, relayPeerService)
	sentryrpc.RegisterClusterAuthorizationServer(s, clusterAuthzServer)
	/*sentryrpc.RegisterAuditInformationServer(s, auditInfoServer)*/

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", rpcRelayPeeringPort))
	if err != nil {
		_log.Fatalw("failed to listen relay peer service port", "port", rpcRelayPeeringPort, "error", err)
		return
	}

	go server.RunRelaySurveyHandler(ctx.Done(), relayPeerService)

	_log.Infow("started relay rpc service ", "port", rpcRelayPeeringPort)
	if err = s.Serve(l); err != nil {
		_log.Fatalw("failed to serve relay peer service", "error", err)
	}

}

func runRPC(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	defer ps.Close()
	defer schedulerPool.Close()
	defer gs.Close()
	defer rs.Close()
	defer rrs.Close()

	partnerServer := server.NewPartnerServer(ps)
	organizationServer := server.NewOrganizationServer(os)
	projectServer := server.NewProjectServer(pps)

	bootstrapServer := server.NewBootstrapServer(bs, kekFunc, cs)
	kubeConfigServer := server.NewKubeConfigServer(bs, aps, gps, kss, krs, kekFunc)
	/*auditInfoServer := rpcv2.NewAuditInfoServer(bs, aps)*/
	clusterAuthzServer := server.NewClusterAuthzServer(bs, aps, gps, krs, kcs, kss)
	kubectlClusterSettingsServer := server.NewKubectlClusterSettingsServer(bs, kcs)
	crpc := server.NewClusterServer(cs, downloadData)
	mserver := server.NewLocationServer(ms)

	userServer := server.NewUserServer(us)
	groupServer := server.NewGroupServer(gs)
	roleServer := server.NewRoleServer(rs)
	rolepermissionServer := server.NewRolePermissionServer(rrs)
	idpServer := server.NewIdpServer(is)
	oidcProviderServer := server.NewOIDCServer(oidcs)

	// audit
	auditLogServer, err := server.NewAuditLogServer(aus)
	if err != nil {
		_log.Fatalw("unable to create auditLog server", "error", err)
	}
	relayAuditServer, err := server.NewRelayAuditServer(ras, rcs)
	if err != nil {
		_log.Fatalw("unable to create relayAudit server", "error", err)
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", rpcPort))
	if err != nil {
		_log.Fatalw("unable to start rpc listener", "error", err)
	}

	var opts []_grpc.ServerOption
	if !dev {
		_log.Infow("adding auth interceptor")
		ac := authv3.NewAuthContext()
		o := authv3.Option{}
		opts = append(opts, _grpc.UnaryInterceptor(
			ac.NewAuthUnaryInterceptor(o),
		))
	}
	s, err := grpc.NewServer(opts...)
	if err != nil {
		_log.Fatalw("unable to create grpc server", "error", err)
	}

	go func() {
		defer s.GracefulStop()

		<-ctx.Done()
		_log.Infow("context done")
	}()

	systemrpc.RegisterPartnerServer(s, partnerServer)
	systemrpc.RegisterOrganizationServer(s, organizationServer)
	systemrpc.RegisterProjectServer(s, projectServer)
	sentryrpc.RegisterBootstrapServer(s, bootstrapServer)
	sentryrpc.RegisterKubeConfigServer(s, kubeConfigServer)
	sentryrpc.RegisterClusterAuthorizationServer(s, clusterAuthzServer)
	/*pbrpcv2.RegisterAuditInformationServer(s, auditInfoServer)*/
	sentryrpc.RegisterKubectlClusterSettingsServer(s, kubectlClusterSettingsServer)
	schedulerrpc.RegisterClusterServer(s, crpc)
	systemrpc.RegisterLocationServer(s, mserver)
	userrpc.RegisterUserServer(s, userServer)
	userrpc.RegisterGroupServer(s, groupServer)
	rolerpc.RegisterRoleServer(s, roleServer)
	rolerpc.RegisterRolepermissionServer(s, rolepermissionServer)
	systemrpc.RegisterIdpServer(s, idpServer)
	systemrpc.RegisterOIDCProviderServer(s, oidcProviderServer)
	auditrpc.RegisterAuditLogServer(s, auditLogServer)
	auditrpc.RegisterRelayAuditServer(s, relayAuditServer)

	_log.Infow("starting rpc server", "port", rpcPort)
	err = s.Serve(l)
	if err != nil {
		_log.Fatalw("unable to start rpc server", "error", err)
	}

}

func runEventHandlers(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()

	//TODO: need to add a bunch of other handlers with gitops
	ceh := reconcile.NewClusterEventHandler(cs)
	_log.Infow("starting cluster event handler")
	go ceh.Handle(ctx.Done())

	// listen to cluster events
	cs.AddEventHandler(ceh.ClusterHook())

	if !dev {
		rl, err := leaderelection.NewConfigMapLock("cluster-scheduler", schedulerNamespace, xid.New().String())
		if err != nil {
			_log.Fatalw("unable to create configmap lock", "error", err)
		}
		go func() {
			err := leaderelection.Run(rl, func(stop <-chan struct{}) {
			}, ctx.Done())

			if err != nil {
				_log.Fatalw("unable to run leader election", "error", err)
			}
		}()
	}

	<-ctx.Done()
}

func runDebug(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	s := http.Server{
		Addr: fmt.Sprintf(":%d", debugPort),
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			_log.Fatalw("unable to start debug server", "error", err)
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	s.Shutdown(ctx)
}

func main() {
	setup()
	run()
}