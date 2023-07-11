package byzantine

import (
	"context"
	"github.com/autonity/autonity/autonity"
	"github.com/autonity/autonity/consensus"
	"github.com/autonity/autonity/consensus/tendermint/core"
	"github.com/autonity/autonity/consensus/tendermint/core/message"
	et "github.com/autonity/autonity/e2e_test"
	"github.com/autonity/autonity/node"
	"testing"
)

type accountableGarbageProposalBroadcaster struct {
	*core.Core
}

func simulateAccountableGarbageMsg(ctx context.Context, c *core.Core, msg *message.Message, code uint8) {
	logger := c.Logger().New("step", c.Step())
	logger.Info("Broadcasting accountable garbage msg bytes")

	randomBytes, err := et.GenerateRandomBytes(2048)
	if err != nil {
		logger.Error("Failed to generate random bytes ", "err", err)
		return
	}

	maliciousMsg := &message.Message{
		Code:          code,
		Payload:       randomBytes,
		Address:       c.Address(),
		CommittedSeal: []byte{},
	}
	maliciousBytes, err := c.SignMessage(maliciousMsg)
	if err != nil {
		logger.Error("Failed to finalize accountable garbage maliciousMsg msg", "err", err)
		return
	}
	logger.Info("Misbehaviour of GarbageMessage rule is simulated by garbage maliciousMsg msg.")
	et.DefaultSignAndBroadcast(ctx, c, msg)
	_ = c.Backend().Broadcast(ctx, c.CommitteeSet().Committee(), maliciousBytes)

}

func (s *accountableGarbageProposalBroadcaster) SignAndBroadcast(ctx context.Context, msg *message.Message) {
	simulateAccountableGarbageMsg(ctx, s.Core, msg, consensus.MsgProposal)
}

type accountableGarbagePrevoteBroadcaster struct {
	*core.Core
}

func (s *accountableGarbagePrevoteBroadcaster) SignAndBroadcast(ctx context.Context, msg *message.Message) {
	simulateAccountableGarbageMsg(ctx, s.Core, msg, consensus.MsgPrevote)
}

type accountableGarbagePrecommitBroadcaster struct {
	*core.Core
}

func (s *accountableGarbagePrecommitBroadcaster) SignAndBroadcast(ctx context.Context, msg *message.Message) {
	simulateAccountableGarbageMsg(ctx, s.Core, msg, consensus.MsgPrecommit)
}

type invalidMsgCodeBroadcaster struct {
	*core.Core
}

func (s *invalidMsgCodeBroadcaster) SignAndBroadcast(ctx context.Context, msg *message.Message) {
	simulateAccountableGarbageMsg(ctx, s.Core, msg, consensus.MsgPrecommit+100)
}

func AccountableGarbageMsgTests(t *testing.T) {
	t.Run("MisbehaviourRuleAccountableGarbageProposal", func(t *testing.T) {
		handler := &node.TendermintServices{Broadcaster: &accountableGarbageProposalBroadcaster{}}
		tp := autonity.Misbehaviour
		rule := autonity.GarbageMessage
		runTest(t, handler, tp, rule, 45)
	})

	t.Run("MisbehaviourRuleAccountableGarbagePrevote", func(t *testing.T) {
		handler := &node.TendermintServices{Broadcaster: &accountableGarbagePrevoteBroadcaster{}}
		tp := autonity.Misbehaviour
		rule := autonity.GarbageMessage
		runTest(t, handler, tp, rule, 45)
	})

	t.Run("MisbehaviourRuleAccountableGarbagePrecommit", func(t *testing.T) {
		handler := &node.TendermintServices{Broadcaster: &accountableGarbagePrecommitBroadcaster{}}
		tp := autonity.Misbehaviour
		rule := autonity.GarbageMessage
		runTest(t, handler, tp, rule, 45)
	})

	t.Run("MisbehaviourRuleAccountableGarbagePrecommit with invalid consensus msg code", func(t *testing.T) {
		handler := &node.TendermintServices{Broadcaster: &invalidMsgCodeBroadcaster{}}
		tp := autonity.Misbehaviour
		rule := autonity.GarbageMessage
		runTest(t, handler, tp, rule, 45)
	})
}
