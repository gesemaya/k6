package execution

import (
	"errors"
	"testing"

	"github.com/gesemaya/k6/lib"
	"github.com/gesemaya/k6/lib/testutils"
	"github.com/gesemaya/k6/lib/testutils/minirunner"
	"github.com/gesemaya/k6/metrics"
	"github.com/stretchr/testify/require"
)

func getBogusTestRunState(tb testing.TB) *lib.TestRunState {
	reg := metrics.NewRegistry()
	piState := &lib.TestPreInitState{
		Logger:         testutils.NewLogger(tb),
		RuntimeOptions: lib.RuntimeOptions{},
		Registry:       reg,
		BuiltinMetrics: metrics.RegisterBuiltinMetrics(reg),
	}

	return &lib.TestRunState{
		TestPreInitState: piState,
		Options:          lib.Options{},
		Runner:           &minirunner.MiniRunner{},
		RunTags:          piState.Registry.RootTagSet(),
	}
}

// Just a lib.PausableExecutor implementation that can return an error
type pausableExecutor struct {
	lib.Executor
	err error
}

func (p pausableExecutor) SetPaused(bool) error {
	return p.err
}

func TestSetPaused(t *testing.T) {
	t.Parallel()
	t.Run("second pause is an error", func(t *testing.T) {
		t.Parallel()
		testRunState := getBogusTestRunState(t)
		sched, err := NewScheduler(testRunState)
		require.NoError(t, err)
		sched.executors = []lib.Executor{pausableExecutor{err: nil}}

		require.NoError(t, sched.SetPaused(true))
		err = sched.SetPaused(true)
		require.Error(t, err)
		require.Contains(t, err.Error(), "execution is already paused")
	})

	t.Run("unpause at the start is an error", func(t *testing.T) {
		t.Parallel()
		testRunState := getBogusTestRunState(t)
		sched, err := NewScheduler(testRunState)
		require.NoError(t, err)
		sched.executors = []lib.Executor{pausableExecutor{err: nil}}
		err = sched.SetPaused(false)
		require.Error(t, err)
		require.Contains(t, err.Error(), "execution wasn't paused")
	})

	t.Run("second unpause is an error", func(t *testing.T) {
		t.Parallel()
		testRunState := getBogusTestRunState(t)
		sched, err := NewScheduler(testRunState)
		require.NoError(t, err)
		sched.executors = []lib.Executor{pausableExecutor{err: nil}}
		require.NoError(t, sched.SetPaused(true))
		require.NoError(t, sched.SetPaused(false))
		err = sched.SetPaused(false)
		require.Error(t, err)
		require.Contains(t, err.Error(), "execution wasn't paused")
	})

	t.Run("an error on pausing is propagated", func(t *testing.T) {
		t.Parallel()
		testRunState := getBogusTestRunState(t)
		sched, err := NewScheduler(testRunState)
		require.NoError(t, err)
		expectedErr := errors.New("testing pausable executor error")
		sched.executors = []lib.Executor{pausableExecutor{err: expectedErr}}
		err = sched.SetPaused(true)
		require.Error(t, err)
		require.Equal(t, err, expectedErr)
	})
}
