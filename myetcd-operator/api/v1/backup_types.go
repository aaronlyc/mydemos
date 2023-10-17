package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type BackupStorageType string

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BackupSpec defines the desired state of Backup
type BackupSpec struct {
	// EtcdEndpoints specifies the endpoints of an etcd cluster.
	// When multiple endpoints are given, the backup operator retrieves
	// the backup from the endpoint that has the most up-to-date state.
	// The given endpoints must belong to the same etcd cluster.
	EtcdEndpoints []string `json:"etcdEndpoints,omitempty"`
	// StorageType is the etcd backup storage type.
	// We need this field because CRD doesn't support validation against invalid fields
	// and we cannot verify invalid backup storage source.
	StorageType BackupStorageType `json:"storageType"`
	// BackupPolicy configures the backup process.
	BackupPolicy *BackupPolicy `json:"backupPolicy,omitempty"`
	// BackupSource is the backup storage source.
	BackupSource `json:",inline"`
	// ClientTLSSecret is the secret containing the etcd TLS client certs and
	// must contain the following data items:
	// data:
	//    "etcd-client.crt": <pem-encoded-cert>
	//    "etcd-client.key": <pem-encoded-key>
	//    "etcd-client-ca.crt": <pem-encoded-ca-cert>
	ClientTLSSecret string `json:"clientTLSSecret,omitempty"`
}

type BackupSource struct {
	// S3 defines the S3 backup source spec.
	S3 *S3BackupSource `json:"s3,omitempty"`
	// ABS defines the ABS backup source spec.
	ABS *ABSBackupSource `json:"abs,omitempty"`
	// GCS defines the GCS backup source spec.
	GCS *GCSBackupSource `json:"gcs,omitempty"`
	// OSS defines the OSS backup source spec.
	OSS *OSSBackupSource `json:"oss,omitempty"`
}

type BackupPolicy struct {
	// TimeoutInSecond is the maximal allowed time in second of the entire backup process.
	TimeoutInSecond int64 `json:"timeoutInSecond,omitempty"`
	// BackupIntervalInSecond is to specify how often operator take snapshot
	// 0 is magic number to indicate one-shot backup
	BackupIntervalInSecond int64 `json:"backupIntervalInSecond,omitempty"`
	// MaxBackups is to specify how many backups we want to keep
	// 0 is magic number to indicate un-limited backups
	MaxBackups int `json:"maxBackups,omitempty"`
}

// BackupStatus defines the observed state of Backup
type BackupStatus struct {
	// Succeeded indicates if the backup has Succeeded.
	Succeeded bool `json:"succeeded"`
	// Reason indicates the reason for any backup related failures.
	Reason string `json:"Reason,omitempty"`
	// EtcdVersion is the version of the backup etcd server.
	EtcdVersion string `json:"etcdVersion,omitempty"`
	// EtcdRevision is the revision of etcd's KV store where the backup is performed on.
	EtcdRevision int64 `json:"etcdRevision,omitempty"`
	// LastSuccessDate indicate the time to get snapshot last time
	LastSuccessDate metav1.Time `json:"lastSuccessDate,omitempty"`
}

// S3BackupSource provides the spec how to store backups on S3.
type S3BackupSource struct {
	// Path is the full s3 path where the backup is saved.
	// The format of the path must be: "<s3-bucket-name>/<path-to-backup-file>"
	// e.g: "mybucket/etcd.backup"
	Path string `json:"path"`

	// The name of the secret object that stores the AWS credential and config files.
	// The file name of the credential MUST be 'credentials'.
	// The file name of the config MUST be 'config'.
	// The profile to use in both files will be 'default'.
	//
	// AWSSecret overwrites the default etcd operator wide AWS credential and config.
	AWSSecret string `json:"awsSecret"`

	// Endpoint if blank points to aws. If specified, can point to s3 compatible object
	// stores.
	Endpoint string `json:"endpoint,omitempty"`

	// ForcePathStyle forces to use path style over the default subdomain style.
	// This is useful when you have an s3 compatible endpoint that doesn't support
	// subdomain buckets.
	ForcePathStyle bool `json:"forcePathStyle"`
}

// ABSBackupSource provides the spec how to store backups on ABS.
type ABSBackupSource struct {
	// Path is the full abs path where the backup is saved.
	// The format of the path must be: "<abs-container-name>/<path-to-backup-file>"
	// e.g: "myabscontainer/etcd.backup"
	Path string `json:"path"`

	// The name of the secret object that stores the Azure storage credential
	ABSSecret string `json:"absSecret"`
}

// GCSBackupSource provides the spec how to store backups on GCS.
type GCSBackupSource struct {
	// Path is the full GCS path where the backup is saved.
	// The format of the path must be: "<gcs-bucket-name>/<path-to-backup-file>"
	// e.g: "mygcsbucket/etcd.backup"
	Path string `json:"path"`

	// The name of the secret object that stores the Google storage credential
	// containing at most ONE of the following:
	// An access token with file name of 'access-token'.
	// JSON credentials with file name of 'credentials.json'.
	//
	// If omitted, client will use the default application credentials.
	GCPSecret string `json:"gcpSecret,omitempty"`
}

// OSSBackupSource provides the spec how to store backups on OSS.
type OSSBackupSource struct {
	// Path is the full abs path where the backup is saved.
	// The format of the path must be: "<oss-bucket-name>/<path-to-backup-file>"
	// e.g: "mybucket/etcd.backup"
	Path string `json:"path"`

	// The name of the secret object that stores the credential which will be used
	// to access Alibaba Cloud OSS.
	//
	// The secret must contain the following keys/fields:
	//     accessKeyID
	//     accessKeySecret
	//
	// The format of secret:
	//
	//   apiVersion: v1
	//   kind: Secret
	//   metadata:
	//     name: <my-credential-name>
	//   type: Opaque
	//   data:
	//     accessKeyID: <base64 of my-access-key-id>
	//     accessKeySecret: <base64 of my-access-key-secret>
	//
	OSSSecret string `json:"ossSecret"`

	// Endpoint is the OSS service endpoint on alibaba cloud, defaults to
	// "http://oss-cn-hangzhou.aliyuncs.com".
	//
	// Details about regions and endpoints, see:
	//  https://www.alibabacloud.com/help/doc-detail/31837.htm
	Endpoint string `json:"endpoint,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Backup is the Schema for the backups API
type Backup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BackupSpec   `json:"spec,omitempty"`
	Status BackupStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// BackupList contains a list of Backup
type BackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Backup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Backup{}, &BackupList{})
}
