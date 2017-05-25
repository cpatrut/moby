// Code generated by protoc-gen-go.
// source: google/logging/type/log_severity.proto
// DO NOT EDIT!

package ltype

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The severity of the event described in a log entry, expressed as one of the
// standard severity levels listed below.  For your reference, the levels are
// assigned the listed numeric values. The effect of using numeric values other
// than those listed is undefined.
//
// You can filter for log entries by severity.  For example, the following
// filter expression will match log entries with severities `INFO`, `NOTICE`,
// and `WARNING`:
//
//     severity > DEBUG AND severity <= WARNING
//
// If you are writing log entries, you should map other severity encodings to
// one of these standard levels. For example, you might map all of Java's FINE,
// FINER, and FINEST levels to `LogSeverity.DEBUG`. You can preserve the
// original severity level in the log entry payload if you wish.
type LogSeverity int32

const (
	// (0) The log entry has no assigned severity level.
	LogSeverity_DEFAULT LogSeverity = 0
	// (100) Debug or trace information.
	LogSeverity_DEBUG LogSeverity = 100
	// (200) Routine information, such as ongoing status or performance.
	LogSeverity_INFO LogSeverity = 200
	// (300) Normal but significant events, such as start up, shut down, or
	// a configuration change.
	LogSeverity_NOTICE LogSeverity = 300
	// (400) Warning events might cause problems.
	LogSeverity_WARNING LogSeverity = 400
	// (500) Error events are likely to cause problems.
	LogSeverity_ERROR LogSeverity = 500
	// (600) Critical events cause more severe problems or outages.
	LogSeverity_CRITICAL LogSeverity = 600
	// (700) A person must take an action immediately.
	LogSeverity_ALERT LogSeverity = 700
	// (800) One or more systems are unusable.
	LogSeverity_EMERGENCY LogSeverity = 800
)

var LogSeverity_name = map[int32]string{
	0:   "DEFAULT",
	100: "DEBUG",
	200: "INFO",
	300: "NOTICE",
	400: "WARNING",
	500: "ERROR",
	600: "CRITICAL",
	700: "ALERT",
	800: "EMERGENCY",
}
var LogSeverity_value = map[string]int32{
	"DEFAULT":   0,
	"DEBUG":     100,
	"INFO":      200,
	"NOTICE":    300,
	"WARNING":   400,
	"ERROR":     500,
	"CRITICAL":  600,
	"ALERT":     700,
	"EMERGENCY": 800,
}

func (x LogSeverity) String() string {
	return proto.EnumName(LogSeverity_name, int32(x))
}
func (LogSeverity) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func init() {
	proto.RegisterEnum("google.logging.type.LogSeverity", LogSeverity_name, LogSeverity_value)
}

func init() { proto.RegisterFile("google/logging/type/log_severity.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4b, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0xcf, 0xc9, 0x4f, 0x4f, 0xcf, 0xcc, 0x4b, 0xd7, 0x2f, 0xa9, 0x2c, 0x00, 0x73,
	0xe2, 0x8b, 0x53, 0xcb, 0x52, 0x8b, 0x32, 0x4b, 0x2a, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85,
	0x84, 0x21, 0xea, 0xf4, 0xa0, 0xea, 0xf4, 0x40, 0xea, 0xa4, 0x64, 0xa0, 0x9a, 0x13, 0x0b, 0x32,
	0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0x8a, 0x21, 0x5a, 0xb4, 0x9a,
	0x18, 0xb9, 0xb8, 0x7d, 0xf2, 0xd3, 0x83, 0xa1, 0x06, 0x09, 0x71, 0x73, 0xb1, 0xbb, 0xb8, 0xba,
	0x39, 0x86, 0xfa, 0x84, 0x08, 0x30, 0x08, 0x71, 0x72, 0xb1, 0xba, 0xb8, 0x3a, 0x85, 0xba, 0x0b,
	0xa4, 0x08, 0x71, 0x72, 0xb1, 0x78, 0xfa, 0xb9, 0xf9, 0x0b, 0x9c, 0x60, 0x14, 0xe2, 0xe6, 0x62,
	0xf3, 0xf3, 0x0f, 0xf1, 0x74, 0x76, 0x15, 0x58, 0xc3, 0x24, 0xc4, 0xc3, 0xc5, 0x1e, 0xee, 0x18,
	0xe4, 0xe7, 0xe9, 0xe7, 0x2e, 0x30, 0x81, 0x59, 0x88, 0x8b, 0x8b, 0xd5, 0x35, 0x28, 0xc8, 0x3f,
	0x48, 0xe0, 0x0b, 0xb3, 0x10, 0x2f, 0x17, 0x87, 0x73, 0x90, 0x67, 0x88, 0xa7, 0xb3, 0xa3, 0x8f,
	0xc0, 0x0d, 0x16, 0x90, 0x94, 0xa3, 0x8f, 0x6b, 0x50, 0x88, 0xc0, 0x1e, 0x56, 0x21, 0x3e, 0x2e,
	0x4e, 0x57, 0x5f, 0xd7, 0x20, 0x77, 0x57, 0x3f, 0xe7, 0x48, 0x81, 0x05, 0x6c, 0x4e, 0xcd, 0x8c,
	0x5c, 0xe2, 0xc9, 0xf9, 0xb9, 0x7a, 0x58, 0x9c, 0xef, 0x24, 0x80, 0xe4, 0xba, 0x00, 0x90, 0x93,
	0x03, 0x18, 0xa3, 0x2c, 0xa0, 0x0a, 0xd3, 0xf3, 0x73, 0x12, 0xf3, 0xd2, 0xf5, 0xf2, 0x8b, 0xd2,
	0xf5, 0xd3, 0x53, 0xf3, 0xc0, 0x1e, 0xd2, 0x87, 0x48, 0x25, 0x16, 0x64, 0x16, 0xa3, 0x04, 0x97,
	0x75, 0x0e, 0x88, 0x5c, 0xc5, 0x24, 0xe9, 0x0e, 0xd1, 0xea, 0x9c, 0x93, 0x5f, 0x9a, 0xa2, 0xe7,
	0x03, 0xb5, 0x29, 0xa4, 0xb2, 0x20, 0x35, 0x89, 0x0d, 0x6c, 0x80, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0x1b, 0x91, 0x99, 0x37, 0x6e, 0x01, 0x00, 0x00,
}
