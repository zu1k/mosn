package types

type ContextKey string

const (
	ContextKeyStreamId                   ContextKey = "StreamId"
	ContextKeyConnectionId               ContextKey = "ConnectionId"
	ContextKeyLogger                     ContextKey = "Logger"
	ContextKeyListenerPort               ContextKey = "ListenerPort"
	ContextKeyListenerName               ContextKey = "ListenerName"
	ContextKeyListenerStatsNameSpace     ContextKey = "ListenerStatsNameSpace"
	ContextKeyNetworkFilterChainFactory  ContextKey = "NetworkFilterChainFactory"
	ContextKeyStreamFilterChainFactories ContextKey = "StreamFilterChainFactory"
	ContextKeyConnectionCodecMapPool     ContextKey = "ContextKeyConnectionCodecMapPool"
)

const (
	GlobalStatsNamespace = ""
)
