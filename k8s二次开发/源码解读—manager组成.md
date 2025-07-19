# controllerManager组成
实现了 ctrl.Manager 接口。它是 Operator 或控制器程序的“管理中心”，负责 统一启动、停止、调度 controller、webhook、metrics、探针、leader election 等子系统
```go
type controllerManager struct {
	sync.Mutex
	started bool  // 控制 manager 的启动状态，防止重复启动

	stopProcedureEngaged *int64  // 	标记 shutdown 是否已开始
	errChan              chan error  // 所有子组件的错误会通过这个通道上报
	runnables            *runnables

	// cluster holds a variety of methods to interact with a cluster. Required.
	cluster cluster.Cluster  /*与 Kubernetes 集群交互的多种方法，内部包含：Client: 与 API Server 通信（读写资源）
Cache: informer 缓存
Scheme: GVK 到 Go 类型的映射
RESTMapper: GVK 到资源路径的映射 */

	// recorderProvider is used to generate event recorders that will be injected into Controllers
	// (and EventHandlers, Sources and Predicates).
	recorderProvider *intrec.Provider

	// resourceLock forms the basis for leader election
	resourceLock resourcelock.Interface

	// leaderElectionReleaseOnCancel defines if the manager should step back from the leader lease
	// on shutdown
	leaderElectionReleaseOnCancel bool

	// metricsServer is used to serve prometheus metrics
	metricsServer metricsserver.Server

	// healthProbeListener is used to serve liveness probe
	healthProbeListener net.Listener

	// Readiness probe endpoint name
	readinessEndpointName string

	// Liveness probe endpoint name
	livenessEndpointName string

	// Readyz probe handler
	readyzHandler *healthz.Handler

	// Healthz probe handler
	healthzHandler *healthz.Handler

	// pprofListener is used to serve pprof
	pprofListener net.Listener

	// controllerConfig are the global controller options.
	controllerConfig config.Controller

	// Logger is the logger that should be used by this manager.
	// If none is set, it defaults to log.Log global logger.
	logger logr.Logger

	// leaderElectionStopped is an internal channel used to signal the stopping procedure that the
	// LeaderElection.Run(...) function has returned and the shutdown can proceed.
	leaderElectionStopped chan struct{}

	// leaderElectionCancel is used to cancel the leader election. It is distinct from internalStopper,
	// because for safety reasons we need to os.Exit() when we lose the leader election, meaning that
	// it must be deferred until after gracefulShutdown is done.
	leaderElectionCancel context.CancelFunc

	// elected is closed when this manager becomes the leader of a group of
	// managers, either because it won a leader election or because no leader
	// election was configured.
	elected chan struct{}

	webhookServer webhook.Server
	// webhookServerOnce will be called in GetWebhookServer() to optionally initialize
	// webhookServer if unset, and Add() it to controllerManager.
	webhookServerOnce sync.Once

	// leaderElectionID is the name of the resource that leader election
	// will use for holding the leader lock.
	leaderElectionID string
	// leaseDuration is the duration that non-leader candidates will
	// wait to force acquire leadership.
	leaseDuration time.Duration
	// renewDeadline is the duration that the acting controlplane will retry
	// refreshing leadership before giving up.
	renewDeadline time.Duration
	// retryPeriod is the duration the LeaderElector clients should wait
	// between tries of actions.
	retryPeriod time.Duration

	// gracefulShutdownTimeout is the duration given to runnable to stop
	// before the manager actually returns on stop.
	gracefulShutdownTimeout time.Duration

	// onStoppedLeading is callled when the leader election lease is lost.
	// It can be overridden for tests.
	onStoppedLeading func()

	// shutdownCtx is the context that can be used during shutdown. It will be cancelled
	// after the gracefulShutdownTimeout ended. It must not be accessed before internalStop
	// is closed because it will be nil.
	shutdownCtx context.Context

	internalCtx    context.Context
	internalCancel context.CancelFunc

	// internalProceduresStop channel is used internally to the manager when coordinating
	// the proper shutdown of servers. This channel is also used for dependency injection.
	internalProceduresStop chan struct{}
}
```
