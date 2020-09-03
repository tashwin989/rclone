// Test Crypt filesystem interface
package press

import (
	"os"
	"path/filepath"
	"testing"

	_ "github.com/rclone/rclone/backend/dropbox"
	_ "github.com/rclone/rclone/backend/local"
	"github.com/rclone/rclone/fstest"
	"github.com/rclone/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	if *fstest.RemoteName == "" {
		t.Skip("Skipping as -remote not set")
	}
	fstests.Run(t, &fstests.Opt{
		RemoteName: *fstest.RemoteName,
		NilObject:  (*Object)(nil),
		UnimplementableFsMethods: []string{
			"OpenWriterAt",
			"MergeDirs",
			"DirCacheFlush",
			"PutUnchecked",
			"PutStream",
			"UserInfo",
			"Disconnect",
		},
		UnimplementableObjectMethods: []string{
			"GetTier",
			"SetTier",
		},
	})
}

// TestRemoteGzip tests GZIP compression
func TestRemoteGzip(t *testing.T) {
	if *fstest.RemoteName != "" {
		t.Skip("Skipping as -remote set")
	}
	tempdir := filepath.Join(os.TempDir(), "rclone-press-test-gzip")
	name := "TestPressGzip"
	fstests.Run(t, &fstests.Opt{
		RemoteName: name + ":",
		NilObject:  (*Object)(nil),
		UnimplementableFsMethods: []string{
			"OpenWriterAt",
			"MergeDirs",
			"DirCacheFlush",
			"PutUnchecked",
			"PutStream",
			"UserInfo",
			"Disconnect",
		},
		UnimplementableObjectMethods: []string{
			"GetTier",
			"SetTier",
		},
		ExtraConfig: []fstests.ExtraConfigItem{
			{Name: name, Key: "type", Value: "press"},
			{Name: name, Key: "remote", Value: tempdir},
			{Name: name, Key: "compression_mode", Value: "gzip"},
		},
	})
}
