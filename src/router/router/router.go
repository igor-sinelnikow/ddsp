package router

import (
	"sync"
	"time"

	"storage"
)

// Config stores configuration for a Router service.
//
// Config -- содержит конфигурацию Router.
type Config struct {
	// Addr is an address to listen at.
	// Addr -- слушающий адрес.
	Addr storage.ServiceAddr

	// Nodes is a list of nodes served by the Router.
	// Nodes -- список node обслуживаемых Router.
	Nodes []storage.ServiceAddr

	// ForgetTimeout is a timeout after node is considered to be unavailable
	// in absence of hearbeats.
	// ForgetTimeout -- если в течении ForgetTimeout node не посылала heartbeats, то
	// node считается недоступной.
	ForgetTimeout time.Duration `yaml:"forget_timeout"`

	// NodesFinder specifies a NodesFinder to use.
	// NodesFinder -- NodesFinder, который нужно использовать в Router.
	NodesFinder NodesFinder `yaml:"-"`
}

// Router is a router service.
type Router struct {
	cfg  Config
	hb   map[storage.ServiceAddr]time.Time
	lock sync.RWMutex
}

// New creates a new Router with a given cfg.
// Returns storage.ErrNotEnoughDaemons error if less then storage.ReplicationFactor
// nodes was provided in cfg.Nodes.
//
// New создает новый Router с данным cfg.
// Возвращает ошибку storage.ErrNotEnoughDaemons если в cfg.Nodes
// меньше чем storage.ReplicationFactor nodes.
func New(cfg Config) (*Router, error) {
	if len(cfg.Nodes) < storage.ReplicationFactor {
		return nil, storage.ErrNotEnoughDaemons
	}

	r := Router{
		cfg: cfg,
		hb:  make(map[storage.ServiceAddr]time.Time, len(cfg.Nodes))}

	now := time.Now()
	for _, node := range cfg.Nodes {
		r.hb[node] = now
	}

	return &r, nil
}

// Heartbeat registers node in the router.
// Returns storage.ErrUnknownDaemon error if node is not served by the Router.
//
// Heartbeat регистритрует node в router.
// Возвращает ошибку storage.ErrUnknownDaemon если node не
// обслуживается Router.
func (r *Router) Heartbeat(node storage.ServiceAddr) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	_, ok := r.hb[node]
	if !ok {
		return storage.ErrUnknownDaemon
	}
	r.hb[node] = time.Now()
	return nil
}

// NodesFind returns a list of available nodes, where record with associated key k
// should be stored. Returns storage.ErrNotEnoughDaemons error
// if less then storage.MinRedundancy can be returned.
//
// NodesFind возвращает cписок достпуных node, на которых должна храниться
// запись с ключом k. Возвращает ошибку storage.ErrNotEnoughDaemons
// если меньше, чем storage.MinRedundancy найдено.
func (r *Router) NodesFind(k storage.RecordID) ([]storage.ServiceAddr, error) {
	found := r.cfg.NodesFinder.NodesFind(k, r.cfg.Nodes)
	available := make([]storage.ServiceAddr, 0, len(found))
	for _, node := range found {
		r.lock.Lock()
		if time.Now().Sub(r.hb[node]) < r.cfg.ForgetTimeout {
			available = append(available, node)
		}
		r.lock.Unlock()
	}
	if len(available) < storage.MinRedundancy {
		return available, storage.ErrNotEnoughDaemons
	}
	return available, nil
}

// List returns a list of all nodes served by Router.
//
// List возвращает cписок всех node, обслуживаемых Router.
func (r *Router) List() []storage.ServiceAddr {
	return r.cfg.Nodes
}
