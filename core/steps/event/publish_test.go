package event_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	g "github.com/onsi/gomega"
	"github.com/weaveworks/flintlock/api/events"
	"github.com/weaveworks/flintlock/core/steps/event"
	"github.com/weaveworks/flintlock/infrastructure/mock"
)

const (
	testTopic     = "test-topic"
	testVMID      = "testvm1"
	testNamespace = "testns"
)

func TestNewPublish(t *testing.T) {
	g.RegisterTestingT(t)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	eventService := mock.NewMockEventService(mockCtrl)
	ctx := context.Background()
	evt := &events.MicroVMSpecDeleted{
		ID:        testVMID,
		Namespace: testNamespace,
	}

	eventService.
		EXPECT().
		Publish(gomock.Any(), gomock.Eq(testTopic), gomock.Eq(evt)).
		Return(nil)

	step := event.NewPublish(testTopic, evt, eventService)

	// Lame test, it can be only true now, but better to document what we expect
	// now than crying later if the system does not do what we want.
	shouldDo, _ := step.ShouldDo(ctx)
	subSteps, err := step.Do(ctx)
	verifyErr := step.Verify(ctx)

	g.Expect(shouldDo).To(g.BeTrue())
	g.Expect(subSteps).To(g.BeEmpty())
	g.Expect(err).To(g.BeNil())
	g.Expect(verifyErr).To(g.BeNil())
}

func TestNewPublish_eventServiceFailure(t *testing.T) {
	g.RegisterTestingT(t)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	eventService := mock.NewMockEventService(mockCtrl)
	ctx := context.Background()
	evt := &events.MicroVMSpecDeleted{
		ID:        testVMID,
		Namespace: testNamespace,
	}

	eventService.
		EXPECT().
		Publish(gomock.Any(), gomock.Eq(testTopic), gomock.Eq(evt)).
		Return(errors.New("something went terribly wrong, that's sad"))

	step := event.NewPublish(testTopic, evt, eventService)

	subSteps, err := step.Do(ctx)
	verifyErr := step.Verify(ctx)

	g.Expect(subSteps).To(g.BeEmpty())
	g.Expect(err).ToNot(g.BeNil())
	g.Expect(verifyErr).To(g.BeNil())
}
