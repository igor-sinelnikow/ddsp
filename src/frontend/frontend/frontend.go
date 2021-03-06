package frontend

import (
	"sync"
	"time"

	rclient "router/client"
	"router/router"
	"storage"
)

// InitTimeout is a timeout to wait after unsuccessful List() request to Router.
//
// InitTimeout -- количество времени, которое нужно подождать до следующей попытки
// отправки запроса List() в Router.
const InitTimeout = 100 * time.Millisecond

// Config stores configuration for a Frontend service.
//
// Config -- содержит конфигурацию Frontend.
type Config struct {
	// Addr is an address to listen at.
	// Addr -- слушающий адрес Frontend.
	Addr storage.ServiceAddr
	// Router is an address of Router service.
	// Router -- адрес Router service.
	Router storage.ServiceAddr

	// NC specifies client for Node.
	// NC -- клиент для node.
	NC storage.Client `yaml:"-"`
	// RC specifies client for Router.
	// RC -- клиент для router.
	RC rclient.Client `yaml:"-"`
	// NodesFinder specifies a NodeFinder to use.
	// NodesFinder -- NodesFinder, который нужно использовать в Frontend.
	NF router.NodesFinder `yaml:"-"`
}

// Frontend is a frontend service.
type Frontend struct {
	cfg   Config
	nodes []storage.ServiceAddr
	once  sync.Once
}

// New creates a new Frontend with a given cfg.
//
// New создает новый Frontend с данным cfg.
func New(cfg Config) *Frontend {
	return &Frontend{cfg: cfg}
}

func (fe *Frontend) handleErrors(k storage.RecordID,
	action func(node storage.ServiceAddr) error) error {
	nodes, err := fe.cfg.RC.NodesFind(fe.cfg.Router, k)
	if err != nil {
		return err
	}
	if len(nodes) < storage.MinRedundancy {
		return storage.ErrNotEnoughDaemons
	}

	errors := make(chan error, len(nodes))
	for _, node := range nodes {
		go func(node storage.ServiceAddr) {
			errors <- action(node)
		}(node)
	}

	count := make(map[error]int)
	for range nodes {
		count[<-errors]++
	}

	for err, num := range count {
		if num >= storage.MinRedundancy {
			return err
		}
	}

	return storage.ErrQuorumNotReached
}

// Put an item to the storage if an item for the given key doesn't exist.
// Returns error otherwise.
//
// Put -- добавить запись в хранилище, если запись для данного ключа
// не существует. Иначе вернуть ошибку.
func (fe *Frontend) Put(k storage.RecordID, d []byte) error {
	return fe.handleErrors(k, func(node storage.ServiceAddr) error {
		return fe.cfg.NC.Put(node, k, d)
	})
}

// Del an item from the storage if an item exists for the given key.
// Returns error otherwise.
//
// Del -- удалить запись из хранилища, если запись для данного ключа
// существует. Иначе вернуть ошибку.
func (fe *Frontend) Del(k storage.RecordID) error {
	return fe.handleErrors(k, func(node storage.ServiceAddr) error {
		return fe.cfg.NC.Del(node, k)
	})
}

type getReply struct {
	d   []byte
	err error
}

// Get an item from the storage if an item exists for the given key.
// Returns error otherwise.
//
// Get -- получить запись из хранилища, если запись для данного ключа
// существует. Иначе вернуть ошибку.
func (fe *Frontend) Get(k storage.RecordID) ([]byte, error) {
	fe.once.Do(func() {
		var err error
		for {
			fe.nodes, err = fe.cfg.RC.List(fe.cfg.Router)
			if err == nil {
				break
			}
			time.Sleep(InitTimeout)
		}
	})

	nodes := fe.cfg.NF.NodesFind(k, fe.nodes)
	if len(nodes) < storage.MinRedundancy {
		return nil, storage.ErrNotEnoughDaemons
	}

	replies := make(chan getReply, len(nodes))
	for _, node := range nodes {
		go func(node storage.ServiceAddr) {
			d, err := fe.cfg.NC.Get(node, k)
			replies <- getReply{d, err}
		}(node)
	}

	countData := make(map[string]int)
	countErr := make(map[error]int)
	for range nodes {
		reply := <-replies
		if reply.err == nil {
			countData[string(reply.d)]++
			if countData[string(reply.d)] >= storage.MinRedundancy {
				return reply.d, nil
			}
		} else {
			countErr[reply.err]++
			if countErr[reply.err] >= storage.MinRedundancy {
				return nil, reply.err
			}
		}
	}

	return nil, storage.ErrQuorumNotReached
}
