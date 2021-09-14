package kong

type RouterGroup struct {
	ServiceRouter
	RouteRouter
	UpstreamRouter
	TargetRouter
	ConsumerRouter
	PluginRouter
	CertificateRouter
	CACertificateRouter
	SniRouter
}
