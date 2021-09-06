package kong

type ApiGroup struct {
	ServiceApi
	RouteApi
	UpstreamApi
	TargetApi
	ConsumerApi
	PluginApi
	CertificateApi
	CACertificateApi
	SniApi
}
