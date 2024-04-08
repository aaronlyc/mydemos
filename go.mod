module mydemos

go 1.17

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.6
	go.etcd.io/etcd/api/v3 => go.etcd.io/etcd/api/v3 v3.5.1
	go.etcd.io/etcd/client/pkg/v3 => go.etcd.io/etcd/client/pkg/v3 v3.5.1
	go.etcd.io/etcd/client/v3 => go.etcd.io/etcd/client/v3 v3.5.1
	go.etcd.io/etcd/pkg/v3 => go.etcd.io/etcd/pkg/v3 v3.5.1
	go.etcd.io/etcd/raft/v3 => go.etcd.io/etcd/raft/v3 v3.5.1
	go.etcd.io/etcd/server/v3 => go.etcd.io/etcd/server/v3 v3.5.1
	google.golang.org/grpc => google.golang.org/grpc v1.40.0
)

require (
	gopkg.in/yaml.v3 v3.0.1
	k8s.io/apimachinery v0.29.3
)
