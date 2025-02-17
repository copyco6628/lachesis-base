package abft

import (
	"github.com/copyco6628/lachesis-base/abft/dagidx"
	"github.com/copyco6628/lachesis-base/abft/election"
	"github.com/copyco6628/lachesis-base/hash"
	"github.com/copyco6628/lachesis-base/inter/idx"
	"github.com/copyco6628/lachesis-base/inter/pos"
)

type OrdererCallbacks struct {
	ApplyAtropos func(decidedFrame idx.Frame, atropos hash.Event) (sealEpoch *pos.Validators)

	EpochDBLoaded func(idx.Epoch)
}

type OrdererDagIndex interface {
	dagidx.ForklessCause
}

// Unlike processes events to reach finality on their order.
// Unlike abft.Lachesis, this raw level of abstraction doesn't track cheaters detection
type Orderer struct {
	config Config
	crit   func(error)
	store  *Store
	input  EventSource

	election *election.Election
	dagIndex OrdererDagIndex

	callback OrdererCallbacks
}

// New creates Orderer instance.
// Unlike Lachesis, Orderer doesn't updates DAG indexes for events, and doesn't detect cheaters
// It has only one purpose - reaching consensus on events order.
func NewOrderer(store *Store, input EventSource, dagIndex OrdererDagIndex, crit func(error), config Config) *Orderer {
	p := &Orderer{
		config:   config,
		store:    store,
		input:    input,
		crit:     crit,
		dagIndex: dagIndex,
	}

	return p
}
