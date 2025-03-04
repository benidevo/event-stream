package error_codes

// Error codes as defined in the Kafka protocol specification.
// These represent possible error conditions that can be returned
// in Kafka protocol responses.
const (
	// Error codes
	UNKNOWN_SERVER_ERROR                  int16 = -1
	NONE                                  int16 = 0
	OFFSET_OUT_OF_RANGE                   int16 = 1
	CORRUPT_MESSAGE                       int16 = 2
	UNKNOWN_TOPIC_OR_PARTITION            int16 = 3
	INVALID_FETCH_SIZE                    int16 = 4
	LEADER_NOT_AVAILABLE                  int16 = 5
	NOT_LEADER_OR_FOLLOWER                int16 = 6
	REQUEST_TIMED_OUT                     int16 = 7
	BROKER_NOT_AVAILABLE                  int16 = 8
	REPLICA_NOT_AVAILABLE                 int16 = 9
	MESSAGE_TOO_LARGE                     int16 = 10
	STALE_CONTROLLER_EPOCH                int16 = 11
	OFFSET_METADATA_TOO_LARGE             int16 = 12
	NETWORK_EXCEPTION                     int16 = 13
	COORDINATOR_LOAD_IN_PROGRESS          int16 = 14
	COORDINATOR_NOT_AVAILABLE             int16 = 15
	NOT_COORDINATOR                       int16 = 16
	INVALID_TOPIC_EXCEPTION               int16 = 17
	RECORD_LIST_TOO_LARGE                 int16 = 18
	NOT_ENOUGH_REPLICAS                   int16 = 19
	NOT_ENOUGH_REPLICAS_AFTER_APPEND      int16 = 20
	INVALID_REQUIRED_ACKS                 int16 = 21
	ILLEGAL_GENERATION                    int16 = 22
	INCONSISTENT_GROUP_PROTOCOL           int16 = 23
	INVALID_GROUP_ID                      int16 = 24
	UNKNOWN_MEMBER_ID                     int16 = 25
	INVALID_SESSION_TIMEOUT               int16 = 26
	REBALANCE_IN_PROGRESS                 int16 = 27
	INVALID_COMMIT_OFFSET_SIZE            int16 = 28
	TOPIC_AUTHORIZATION_FAILED            int16 = 29
	GROUP_AUTHORIZATION_FAILED            int16 = 30
	CLUSTER_AUTHORIZATION_FAILED          int16 = 31
	INVALID_TIMESTAMP                     int16 = 32
	UNSUPPORTED_SASL_MECHANISM            int16 = 33
	ILLEGAL_SASL_STATE                    int16 = 34
	UNSUPPORTED_VERSION                   int16 = 35
	TOPIC_ALREADY_EXISTS                  int16 = 36
	INVALID_PARTITIONS                    int16 = 37
	INVALID_REPLICATION_FACTOR            int16 = 38
	INVALID_REPLICA_ASSIGNMENT            int16 = 39
	INVALID_CONFIG                        int16 = 40
	NOT_CONTROLLER                        int16 = 41
	INVALID_REQUEST                       int16 = 42
	UNSUPPORTED_FOR_MESSAGE_FORMAT        int16 = 43
	POLICY_VIOLATION                      int16 = 44
	OUT_OF_ORDER_SEQUENCE_NUMBER          int16 = 45
	DUPLICATE_SEQUENCE_NUMBER             int16 = 46
	INVALID_PRODUCER_EPOCH                int16 = 47
	INVALID_TXN_STATE                     int16 = 48
	INVALID_PRODUCER_ID_MAPPING           int16 = 49
	INVALID_TRANSACTION_TIMEOUT           int16 = 50
	CONCURRENT_TRANSACTIONS               int16 = 51
	TRANSACTION_COORDINATOR_FENCED        int16 = 52
	TRANSACTIONAL_ID_AUTHORIZATION_FAILED int16 = 53
	SECURITY_DISABLED                     int16 = 54
	OPERATION_NOT_ATTEMPTED               int16 = 55
	KAFKA_STORAGE_ERROR                   int16 = 56
	LOG_DIR_NOT_FOUND                     int16 = 57
	SASL_AUTHENTICATION_FAILED            int16 = 58
	UNKNOWN_PRODUCER_ID                   int16 = 59
	REASSIGNMENT_IN_PROGRESS              int16 = 60
	DELEGATION_TOKEN_AUTH_DISABLED        int16 = 61
	DELEGATION_TOKEN_NOT_FOUND            int16 = 62
	DELEGATION_TOKEN_OWNER_MISMATCH       int16 = 63
	DELEGATION_TOKEN_REQUEST_NOT_ALLOWED  int16 = 64
	DELEGATION_TOKEN_AUTHORIZATION_FAILED int16 = 65
	DELEGATION_TOKEN_EXPIRED              int16 = 66
	INVALID_PRINCIPAL_TYPE                int16 = 67
	NON_EMPTY_GROUP                       int16 = 68
	GROUP_ID_NOT_FOUND                    int16 = 69
	FETCH_SESSION_ID_NOT_FOUND            int16 = 70
	INVALID_FETCH_SESSION_EPOCH           int16 = 71
	LISTENER_NOT_FOUND                    int16 = 72
	TOPIC_DELETION_DISABLED               int16 = 73
	FENCED_LEADER_EPOCH                   int16 = 74
	UNKNOWN_LEADER_EPOCH                  int16 = 75
	UNSUPPORTED_COMPRESSION_TYPE          int16 = 76
	STALE_BROKER_EPOCH                    int16 = 77
	OFFSET_NOT_AVAILABLE                  int16 = 78
	MEMBER_ID_REQUIRED                    int16 = 79
	PREFERRED_LEADER_NOT_AVAILABLE        int16 = 80
	GROUP_MAX_SIZE_REACHED                int16 = 81
	FENCED_INSTANCE_ID                    int16 = 82
	ELIGIBLE_LEADERS_NOT_AVAILABLE        int16 = 83
	ELECTION_NOT_NEEDED                   int16 = 84
	NO_REASSIGNMENT_IN_PROGRESS           int16 = 85
	GROUP_SUBSCRIBED_TO_TOPIC             int16 = 86
	INVALID_RECORD                        int16 = 87
	UNSTABLE_OFFSET_COMMIT                int16 = 88
	THROTTLING_QUOTA_EXCEEDED             int16 = 89
	PRODUCER_FENCED                       int16 = 90
	RESOURCE_NOT_FOUND                    int16 = 91
	DUPLICATE_RESOURCE                    int16 = 92
	UNACCEPTABLE_CREDENTIAL               int16 = 93
	INCONSISTENT_VOTER_SET                int16 = 94
	INVALID_UPDATE_VERSION                int16 = 95
	FEATURE_UPDATE_FAILED                 int16 = 96
	PRINCIPAL_DESERIALIZATION_FAILURE     int16 = 97
	SNAPSHOT_NOT_FOUND                    int16 = 98
	POSITION_OUT_OF_RANGE                 int16 = 99
	UNKNOWN_TOPIC_ID                      int16 = 100
	DUPLICATE_BROKER_REGISTRATION         int16 = 101
	BROKER_ID_NOT_REGISTERED              int16 = 102
	INCONSISTENT_TOPIC_ID                 int16 = 103
	INCONSISTENT_CLUSTER_ID               int16 = 104
	TRANSACTIONAL_ID_NOT_FOUND            int16 = 105
	FETCH_SESSION_TOPIC_ID_ERROR          int16 = 106
	INELIGIBLE_REPLICA                    int16 = 107
	NEW_LEADER_ELECTED                    int16 = 108
	OFFSET_MOVED_TO_TIERED_STORAGE        int16 = 109
	FENCED_MEMBER_EPOCH                   int16 = 110
	UNRELEASED_INSTANCE_ID                int16 = 111
	UNSUPPORTED_ASSIGNOR                  int16 = 112
	STALE_MEMBER_EPOCH                    int16 = 113
	MISMATCHED_ENDPOINT_TYPE              int16 = 114
	UNSUPPORTED_ENDPOINT_TYPE             int16 = 115
	UNKNOWN_CONTROLLER_ID                 int16 = 116
	UNKNOWN_SUBSCRIPTION_ID               int16 = 117
	TELEMETRY_TOO_LARGE                   int16 = 118
	INVALID_REGISTRATION                  int16 = 119
	TRANSACTION_ABORTABLE                 int16 = 120
	INVALID_RECORD_STATE                  int16 = 121
	SHARE_SESSION_NOT_FOUND               int16 = 122
	INVALID_SHARE_SESSION_EPOCH           int16 = 123
	FENCED_STATE_EPOCH                    int16 = 124
	INVALID_VOTER_KEY                     int16 = 125
	DUPLICATE_VOTER                       int16 = 126
	VOTER_NOT_FOUND                       int16 = 127
)

// ErrorCodeNames maps numeric error codes to their string names for debugging and logging.
var ErrorCodeNames = map[int16]string{
	UNKNOWN_SERVER_ERROR:                  "UNKNOWN_SERVER_ERROR",
	NONE:                                  "NONE",
	OFFSET_OUT_OF_RANGE:                   "OFFSET_OUT_OF_RANGE",
	CORRUPT_MESSAGE:                       "CORRUPT_MESSAGE",
	UNKNOWN_TOPIC_OR_PARTITION:            "UNKNOWN_TOPIC_OR_PARTITION",
	INVALID_FETCH_SIZE:                    "INVALID_FETCH_SIZE",
	LEADER_NOT_AVAILABLE:                  "LEADER_NOT_AVAILABLE",
	NOT_LEADER_OR_FOLLOWER:                "NOT_LEADER_OR_FOLLOWER",
	REQUEST_TIMED_OUT:                     "REQUEST_TIMED_OUT",
	BROKER_NOT_AVAILABLE:                  "BROKER_NOT_AVAILABLE",
	REPLICA_NOT_AVAILABLE:                 "REPLICA_NOT_AVAILABLE",
	MESSAGE_TOO_LARGE:                     "MESSAGE_TOO_LARGE",
	STALE_CONTROLLER_EPOCH:                "STALE_CONTROLLER_EPOCH",
	OFFSET_METADATA_TOO_LARGE:             "OFFSET_METADATA_TOO_LARGE",
	NETWORK_EXCEPTION:                     "NETWORK_EXCEPTION",
	COORDINATOR_LOAD_IN_PROGRESS:          "COORDINATOR_LOAD_IN_PROGRESS",
	COORDINATOR_NOT_AVAILABLE:             "COORDINATOR_NOT_AVAILABLE",
	NOT_COORDINATOR:                       "NOT_COORDINATOR",
	INVALID_TOPIC_EXCEPTION:               "INVALID_TOPIC_EXCEPTION",
	RECORD_LIST_TOO_LARGE:                 "RECORD_LIST_TOO_LARGE",
	NOT_ENOUGH_REPLICAS:                   "NOT_ENOUGH_REPLICAS",
	NOT_ENOUGH_REPLICAS_AFTER_APPEND:      "NOT_ENOUGH_REPLICAS_AFTER_APPEND",
	INVALID_REQUIRED_ACKS:                 "INVALID_REQUIRED_ACKS",
	ILLEGAL_GENERATION:                    "ILLEGAL_GENERATION",
	INCONSISTENT_GROUP_PROTOCOL:           "INCONSISTENT_GROUP_PROTOCOL",
	INVALID_GROUP_ID:                      "INVALID_GROUP_ID",
	UNKNOWN_MEMBER_ID:                     "UNKNOWN_MEMBER_ID",
	INVALID_SESSION_TIMEOUT:               "INVALID_SESSION_TIMEOUT",
	REBALANCE_IN_PROGRESS:                 "REBALANCE_IN_PROGRESS",
	INVALID_COMMIT_OFFSET_SIZE:            "INVALID_COMMIT_OFFSET_SIZE",
	TOPIC_AUTHORIZATION_FAILED:            "TOPIC_AUTHORIZATION_FAILED",
	GROUP_AUTHORIZATION_FAILED:            "GROUP_AUTHORIZATION_FAILED",
	CLUSTER_AUTHORIZATION_FAILED:          "CLUSTER_AUTHORIZATION_FAILED",
	INVALID_TIMESTAMP:                     "INVALID_TIMESTAMP",
	UNSUPPORTED_SASL_MECHANISM:            "UNSUPPORTED_SASL_MECHANISM",
	ILLEGAL_SASL_STATE:                    "ILLEGAL_SASL_STATE",
	UNSUPPORTED_VERSION:                   "UNSUPPORTED_VERSION",
	TOPIC_ALREADY_EXISTS:                  "TOPIC_ALREADY_EXISTS",
	INVALID_PARTITIONS:                    "INVALID_PARTITIONS",
	INVALID_REPLICATION_FACTOR:            "INVALID_REPLICATION_FACTOR",
	INVALID_REPLICA_ASSIGNMENT:            "INVALID_REPLICA_ASSIGNMENT",
	INVALID_CONFIG:                        "INVALID_CONFIG",
	NOT_CONTROLLER:                        "NOT_CONTROLLER",
	INVALID_REQUEST:                       "INVALID_REQUEST",
	UNSUPPORTED_FOR_MESSAGE_FORMAT:        "UNSUPPORTED_FOR_MESSAGE_FORMAT",
	POLICY_VIOLATION:                      "POLICY_VIOLATION",
	OUT_OF_ORDER_SEQUENCE_NUMBER:          "OUT_OF_ORDER_SEQUENCE_NUMBER",
	DUPLICATE_SEQUENCE_NUMBER:             "DUPLICATE_SEQUENCE_NUMBER",
	INVALID_PRODUCER_EPOCH:                "INVALID_PRODUCER_EPOCH",
	INVALID_TXN_STATE:                     "INVALID_TXN_STATE",
	INVALID_PRODUCER_ID_MAPPING:           "INVALID_PRODUCER_ID_MAPPING",
	INVALID_TRANSACTION_TIMEOUT:           "INVALID_TRANSACTION_TIMEOUT",
	CONCURRENT_TRANSACTIONS:               "CONCURRENT_TRANSACTIONS",
	TRANSACTION_COORDINATOR_FENCED:        "TRANSACTION_COORDINATOR_FENCED",
	TRANSACTIONAL_ID_AUTHORIZATION_FAILED: "TRANSACTIONAL_ID_AUTHORIZATION_FAILED",
	SECURITY_DISABLED:                     "SECURITY_DISABLED",
	OPERATION_NOT_ATTEMPTED:               "OPERATION_NOT_ATTEMPTED",
	KAFKA_STORAGE_ERROR:                   "KAFKA_STORAGE_ERROR",
	LOG_DIR_NOT_FOUND:                     "LOG_DIR_NOT_FOUND",
	SASL_AUTHENTICATION_FAILED:            "SASL_AUTHENTICATION_FAILED",
	UNKNOWN_PRODUCER_ID:                   "UNKNOWN_PRODUCER_ID",
	REASSIGNMENT_IN_PROGRESS:              "REASSIGNMENT_IN_PROGRESS",
	DELEGATION_TOKEN_AUTH_DISABLED:        "DELEGATION_TOKEN_AUTH_DISABLED",
	DELEGATION_TOKEN_NOT_FOUND:            "DELEGATION_TOKEN_NOT_FOUND",
	DELEGATION_TOKEN_OWNER_MISMATCH:       "DELEGATION_TOKEN_OWNER_MISMATCH",
	DELEGATION_TOKEN_REQUEST_NOT_ALLOWED:  "DELEGATION_TOKEN_REQUEST_NOT_ALLOWED",
	DELEGATION_TOKEN_AUTHORIZATION_FAILED: "DELEGATION_TOKEN_AUTHORIZATION_FAILED",
	DELEGATION_TOKEN_EXPIRED:              "DELEGATION_TOKEN_EXPIRED",
	INVALID_PRINCIPAL_TYPE:                "INVALID_PRINCIPAL_TYPE",
	NON_EMPTY_GROUP:                       "NON_EMPTY_GROUP",
	GROUP_ID_NOT_FOUND:                    "GROUP_ID_NOT_FOUND",
	FETCH_SESSION_ID_NOT_FOUND:            "FETCH_SESSION_ID_NOT_FOUND",
	INVALID_FETCH_SESSION_EPOCH:           "INVALID_FETCH_SESSION_EPOCH",
	LISTENER_NOT_FOUND:                    "LISTENER_NOT_FOUND",
	TOPIC_DELETION_DISABLED:               "TOPIC_DELETION_DISABLED",
	FENCED_LEADER_EPOCH:                   "FENCED_LEADER_EPOCH",
	UNKNOWN_LEADER_EPOCH:                  "UNKNOWN_LEADER_EPOCH",
	UNSUPPORTED_COMPRESSION_TYPE:          "UNSUPPORTED_COMPRESSION_TYPE",
	STALE_BROKER_EPOCH:                    "STALE_BROKER_EPOCH",
	OFFSET_NOT_AVAILABLE:                  "OFFSET_NOT_AVAILABLE",
	MEMBER_ID_REQUIRED:                    "MEMBER_ID_REQUIRED",
	PREFERRED_LEADER_NOT_AVAILABLE:        "PREFERRED_LEADER_NOT_AVAILABLE",
	GROUP_MAX_SIZE_REACHED:                "GROUP_MAX_SIZE_REACHED",
	FENCED_INSTANCE_ID:                    "FENCED_INSTANCE_ID",
	ELIGIBLE_LEADERS_NOT_AVAILABLE:        "ELIGIBLE_LEADERS_NOT_AVAILABLE",
	ELECTION_NOT_NEEDED:                   "ELECTION_NOT_NEEDED",
	NO_REASSIGNMENT_IN_PROGRESS:           "NO_REASSIGNMENT_IN_PROGRESS",
	GROUP_SUBSCRIBED_TO_TOPIC:             "GROUP_SUBSCRIBED_TO_TOPIC",
	INVALID_RECORD:                        "INVALID_RECORD",
	UNSTABLE_OFFSET_COMMIT:                "UNSTABLE_OFFSET_COMMIT",
	THROTTLING_QUOTA_EXCEEDED:             "THROTTLING_QUOTA_EXCEEDED",
	PRODUCER_FENCED:                       "PRODUCER_FENCED",
	RESOURCE_NOT_FOUND:                    "RESOURCE_NOT_FOUND",
	DUPLICATE_RESOURCE:                    "DUPLICATE_RESOURCE",
	UNACCEPTABLE_CREDENTIAL:               "UNACCEPTABLE_CREDENTIAL",
	INCONSISTENT_VOTER_SET:                "INCONSISTENT_VOTER_SET",
	INVALID_UPDATE_VERSION:                "INVALID_UPDATE_VERSION",
	FEATURE_UPDATE_FAILED:                 "FEATURE_UPDATE_FAILED",
	PRINCIPAL_DESERIALIZATION_FAILURE:     "PRINCIPAL_DESERIALIZATION_FAILURE",
	SNAPSHOT_NOT_FOUND:                    "SNAPSHOT_NOT_FOUND",
	POSITION_OUT_OF_RANGE:                 "POSITION_OUT_OF_RANGE",
	UNKNOWN_TOPIC_ID:                      "UNKNOWN_TOPIC_ID",
	DUPLICATE_BROKER_REGISTRATION:         "DUPLICATE_BROKER_REGISTRATION",
	BROKER_ID_NOT_REGISTERED:              "BROKER_ID_NOT_REGISTERED",
	INCONSISTENT_TOPIC_ID:                 "INCONSISTENT_TOPIC_ID",
	INCONSISTENT_CLUSTER_ID:               "INCONSISTENT_CLUSTER_ID",
	TRANSACTIONAL_ID_NOT_FOUND:            "TRANSACTIONAL_ID_NOT_FOUND",
	FETCH_SESSION_TOPIC_ID_ERROR:          "FETCH_SESSION_TOPIC_ID_ERROR",
	INELIGIBLE_REPLICA:                    "INELIGIBLE_REPLICA",
	NEW_LEADER_ELECTED:                    "NEW_LEADER_ELECTED",
	OFFSET_MOVED_TO_TIERED_STORAGE:        "OFFSET_MOVED_TO_TIERED_STORAGE",
	FENCED_MEMBER_EPOCH:                   "FENCED_MEMBER_EPOCH",
	UNRELEASED_INSTANCE_ID:                "UNRELEASED_INSTANCE_ID",
	UNSUPPORTED_ASSIGNOR:                  "UNSUPPORTED_ASSIGNOR",
	STALE_MEMBER_EPOCH:                    "STALE_MEMBER_EPOCH",
	MISMATCHED_ENDPOINT_TYPE:              "MISMATCHED_ENDPOINT_TYPE",
	UNSUPPORTED_ENDPOINT_TYPE:             "UNSUPPORTED_ENDPOINT_TYPE",
	UNKNOWN_CONTROLLER_ID:                 "UNKNOWN_CONTROLLER_ID",
	UNKNOWN_SUBSCRIPTION_ID:               "UNKNOWN_SUBSCRIPTION_ID",
	TELEMETRY_TOO_LARGE:                   "TELEMETRY_TOO_LARGE",
	INVALID_REGISTRATION:                  "INVALID_REGISTRATION",
	TRANSACTION_ABORTABLE:                 "TRANSACTION_ABORTABLE",
	INVALID_RECORD_STATE:                  "INVALID_RECORD_STATE",
	SHARE_SESSION_NOT_FOUND:               "SHARE_SESSION_NOT_FOUND",
	INVALID_SHARE_SESSION_EPOCH:           "INVALID_SHARE_SESSION_EPOCH",
	FENCED_STATE_EPOCH:                    "FENCED_STATE_EPOCH",
	INVALID_VOTER_KEY:                     "INVALID_VOTER_KEY",
	DUPLICATE_VOTER:                       "DUPLICATE_VOTER",
	VOTER_NOT_FOUND:                       "VOTER_NOT_FOUND",
}
