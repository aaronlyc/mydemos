package v1

import "errors"

// TLSPolicy defines the TLS policy of an etcd cluster
type TLSPolicy struct {
	// StaticTLS enables user to generate static x509 certificates and keys,
	// put them into Kubernetes secrets, and specify them into here.
	Static *StaticTLS `json:"static,omitempty"`
}

type StaticTLS struct {
	// Member contains secrets containing TLS certs used by each etcd member pod.
	Member *MemberSecret `json:"member,omitempty"`
	// OperatorSecret is the secret containing TLS certs used by operator to
	// talk securely to this cluster.
	OperatorSecret string `json:"operatorSecret,omitempty"`
}

type MemberSecret struct {
	// PeerSecret is the secret containing TLS certs used by each etcd member pod
	// for the communication between etcd peers.
	PeerSecret string `json:"peerSecret,omitempty"`
	// ServerSecret is the secret containing TLS certs used by each etcd member pod
	// for the communication between etcd server and its clients.
	ServerSecret string `json:"serverSecret,omitempty"`
}

func (tp *TLSPolicy) Validate() error {
	if tp.Static == nil {
		return nil
	}
	st := tp.Static

	if len(st.OperatorSecret) != 0 {
		if len(st.Member.ServerSecret) == 0 {
			return errors.New("operator secret set but member serverSecret not set")
		}
	} else if st.Member != nil && len(st.Member.ServerSecret) != 0 {
		return errors.New("member serverSecret set but operator secret not set")
	}
	return nil
}

func (tp *TLSPolicy) IsSecureClient() bool {
	if tp == nil || tp.Static == nil {
		return false
	}
	return len(tp.Static.OperatorSecret) != 0
}

func (tp *TLSPolicy) IsSecurePeer() bool {
	if tp == nil || tp.Static == nil || tp.Static.Member == nil {
		return false
	}
	return len(tp.Static.Member.PeerSecret) != 0
}
