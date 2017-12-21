package schedulerd

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/sensu/sensu-go/backend/messaging"
	"github.com/sensu/sensu-go/testing/mockstore"
	"github.com/sensu/sensu-go/types"
	"github.com/stretchr/testify/suite"
)

type CheckSchedulerSuite struct {
	suite.Suite
	check     *types.CheckConfig
	scheduler *CheckScheduler
	msgBus    *messaging.WizardBus
}

func (suite *CheckSchedulerSuite) SetupTest() {
	suite.check = types.FixtureCheckConfig("check1")
	suite.check.Interval = 1
	suite.msgBus = &messaging.WizardBus{}

	manager := NewStateManager(&mockstore.MockStore{})
	manager.Update(func(state *SchedulerState) {
		state.SetChecks([]*types.CheckConfig{suite.check})
	})

	suite.scheduler = &CheckScheduler{
		CheckName:     suite.check.Name,
		CheckEnv:      suite.check.Environment,
		CheckOrg:      suite.check.Organization,
		CheckInterval: suite.check.Interval,
		StateManager:  manager,
		MessageBus:    suite.msgBus,
		WaitGroup:     &sync.WaitGroup{},
	}

	suite.NoError(suite.msgBus.Start())
}

func (suite *CheckSchedulerSuite) TestStart() {
	// Set interval to smallest valid value
	check := suite.check
	check.Subscriptions = []string{"subscription1"}

	c1 := make(chan interface{}, 10)
	topic := fmt.Sprintf(
		"%s:%s:%s:subscription1",
		messaging.TopicSubscriptions,
		check.Organization,
		check.Environment,
	)
	suite.NoError(suite.msgBus.Subscribe(topic, "channel1", c1))

	suite.NoError(suite.scheduler.Start())
	time.Sleep(1 * time.Second)
	suite.NoError(suite.scheduler.Stop())
	suite.NoError(suite.msgBus.Stop())
	close(c1)

	messages := []*types.CheckRequest{}
	for msg := range c1 {
		res, ok := msg.(*types.CheckRequest)
		suite.True(ok)
		messages = append(messages, res)
	}
	res := messages[0]
	suite.Equal(1, len(messages))
	suite.Equal("check1", res.Config.Name)
}

type CheckSubdueSuite struct {
	suite.Suite
	check     *types.CheckConfig
	scheduler *CheckScheduler
	msgBus    *messaging.WizardBus
}

func (suite *CheckSubdueSuite) SetupTest() {
	suite.check = types.FixtureCheckConfig("check1")
	suite.check.Interval = 1
	suite.msgBus = &messaging.WizardBus{}

	manager := NewStateManager(&mockstore.MockStore{})
	manager.Update(func(state *SchedulerState) {
		state.SetChecks([]*types.CheckConfig{suite.check})
	})

	suite.scheduler = &CheckScheduler{
		CheckName:     suite.check.Name,
		CheckEnv:      suite.check.Environment,
		CheckOrg:      suite.check.Organization,
		CheckInterval: suite.check.Interval,
		StateManager:  manager,
		MessageBus:    suite.msgBus,
		WaitGroup:     &sync.WaitGroup{},
	}

	suite.NoError(suite.msgBus.Start())
}

func (suite *CheckSubdueSuite) TestStart() {
	// Set interval to smallest valid value
	check := suite.check
	check.Subscriptions = []string{"subscription1"}
	check.Subdue = &types.TimeWindowWhen{
		Days: types.TimeWindowDays{
			All: []*types.TimeWindowTimeRange{
				{
					Begin: "1:00 AM",
					End:   "11:00 PM",
				},
				{
					Begin: "10:00 PM",
					End:   "12:30 AM",
				},
			},
		},
	}

	c1 := make(chan interface{}, 10)
	topic := fmt.Sprintf(
		"%s:%s:%s:subscription1",
		messaging.TopicSubscriptions,
		check.Organization,
		check.Environment,
	)
	suite.NoError(suite.msgBus.Subscribe(topic, "channel1", c1))

	suite.NoError(suite.scheduler.Start())
	time.Sleep(1 * time.Second)
	suite.NoError(suite.scheduler.Stop())
	suite.NoError(suite.msgBus.Stop())
	close(c1)

	messages := []*types.CheckRequest{}
	for msg := range c1 {
		res, ok := msg.(*types.CheckRequest)
		suite.True(ok)
		messages = append(messages, res)
	}
	// Check should have been subdued at this time, so expect no messages
	suite.Equal(0, len(messages))
}

type TimerSuite struct {
	suite.Suite
}

func (suite *TimerSuite) TestSplay() {
	timer := NewCheckTimer("check1", 10)

	suite.Condition(func() bool { return timer.splay > 0 })

	timer2 := NewCheckTimer("check1", 10)
	suite.Equal(timer.splay, timer2.splay)
}

func (suite *TimerSuite) TestInitialOffset() {
	inputs := []uint{1, 10, 60}
	for _, intervalSeconds := range inputs {
		now := time.Now()
		timer := NewCheckTimer("check1", intervalSeconds)
		nextExecution := timer.calcInitialOffset()
		executionTime := now.Add(nextExecution)

		// We've scheduled it in the future.
		suite.Condition(func() bool { return executionTime.Sub(now) > 0 })
		// The offset is less than the check interval.
		suite.Condition(func() bool { return nextExecution < (time.Duration(intervalSeconds) * time.Second) })
		// The next execution occurs _before_ now + interval.
		suite.Condition(func() bool { return executionTime.Before(now.Add(time.Duration(intervalSeconds) * time.Second)) })
	}
}

func (suite *TimerSuite) TestStop() {
	timer := NewCheckTimer("check1", 10)
	timer.Start()

	result := timer.Stop()
	suite.True(result)
}

type CheckExecSuite struct {
	suite.Suite
	check  *types.CheckConfig
	exec   *CheckExecutor
	msgBus messaging.MessageBus
}

func (suite *CheckExecSuite) SetupTest() {
	suite.msgBus = &messaging.WizardBus{}
	suite.NoError(suite.msgBus.Start())

	request := types.FixtureCheckRequest("check1")
	asset := request.Assets[0]
	hook := request.Hooks[0]
	suite.check = request.Config

	state := &SchedulerState{}
	state.SetChecks([]*types.CheckConfig{request.Config})
	state.SetAssets([]*types.Asset{&asset})
	state.SetHooks([]*types.HookConfig{&hook})

	suite.exec = &CheckExecutor{
		State: state,
		Bus:   suite.msgBus,
	}
}

func (suite *CheckExecSuite) AfterTest() {
	suite.NoError(suite.msgBus.Stop())
}

func (suite *CheckExecSuite) TestBuild() {
	check := suite.check
	request := suite.exec.BuildRequest(check)
	suite.NotNil(request)
	suite.NotNil(request.Config)
	suite.NotNil(request.Assets)
	suite.NotEmpty(request.Assets)
	suite.Len(request.Assets, 1)
	suite.NotNil(request.Hooks)
	suite.NotEmpty(request.Hooks)
	suite.Len(request.Hooks, 1)

	check.RuntimeAssets = []string{}
	check.CheckHooks = []types.HookList{}
	request = suite.exec.BuildRequest(check)
	suite.NotNil(request)
	suite.NotNil(request.Config)
	suite.Empty(request.Assets)
	suite.Empty(request.Hooks)
}

func TestRunExecSuite(t *testing.T) {
	suite.Run(t, new(TimerSuite))
	suite.Run(t, new(CheckSchedulerSuite))
	suite.Run(t, new(CheckExecSuite))
	suite.Run(t, new(CheckSubdueSuite))
}
