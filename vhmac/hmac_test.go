package vhmac_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vhmac"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestSha256(t *testing.T) {
	vtest.Equal(t, "88cd2108b5347d973cf39cdf9053d7dd42704876d8c9a9bd8e2d168259d3ddf7", vhmac.Sha256("test", "test"))
	vtest.Equal(t, "43b0cef99265f9e34c10ea9d3501926d27b39f57c6d674561d8ba236e7a819fb", vhmac.Sha256("test", ""))
	vtest.Equal(t, "b613679a0814d9ec772f95d778c35fc5ff1697c493715653c6c712144292c5ad", vhmac.Sha256("", ""))
}
