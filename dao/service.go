package dao

type ServiceDetail struct {
	Info              *ServiceInfo   `json:"info" description:"基本信息"`
	HttpRule          *HttpRule      `json:"http" description:"http_rule"`
	TCPRule           *TcpRule       `json:"tcp" description:"tcp_rule"`
	GRPCRule          *GrpcRule      `json:"grpc" description:"grpc_rule"`
	LoadBalanceRule   *LoadBalance   `json:"grpc" description:"load_balance"`
	AccessControlRule *AccessControl `json:"grpc" description:"access_control"`
}
