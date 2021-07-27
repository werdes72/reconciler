package compreconciler

import (
	"fmt"
	"github.com/kyma-incubator/reconciler/pkg/test"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRemoteCallbackHandler(t *testing.T) {
	if !test.RunExpensiveTests() {
		return
	}
	t.Run("Test successful remote status update", func(t *testing.T) {
		rcb, err := newRemoteCallbackHandler("https://httpbin.org/status/200", true)
		require.NoError(t, err)
		require.NoError(t, rcb.Callback(Running))
	})

	t.Run("Test failed remote status update", func(t *testing.T) {
		rcb, err := newRemoteCallbackHandler("https://httpbin.org/status/400", true)
		require.NoError(t, err)
		require.Error(t, rcb.Callback(Running))
	})
}

func TestLocalCallbackHandler(t *testing.T) {
	t.Run("Test successful local status update", func(t *testing.T) {
		var localFctCalled bool
		rcb, err := newLocalCallbackHandler(func(status Status) error {
			localFctCalled = true
			return nil
		}, true)
		require.NoError(t, err)
		require.NoError(t, rcb.Callback(Running))
		require.True(t, localFctCalled)
	})

	t.Run("Test failed local status update", func(t *testing.T) {
		rcb, err := newLocalCallbackHandler(func(status Status) error {
			return fmt.Errorf("I failed")
		}, true)
		require.NoError(t, err)
		require.Error(t, rcb.Callback(Running))
	})
}