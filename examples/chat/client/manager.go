package client

import (
	"fmt"
	"lark/pkg/common/xtimer"
	"math"
	"sync"
	"time"
)

var msgTimer = xtimer.NewTimer(2 * time.Second)

type Manager struct {
	rwLock     sync.RWMutex
	unregister chan *Client
	clients    map[int64]*Client
}

func NewManager() (mgr *Manager) {
	mgr = &Manager{clients: make(map[int64]*Client), unregister: make(chan *Client, 100)}
	return
}

func (m *Manager) Run() {
	//go m.listener()
	go m.batchCreate(1000)
}

func (m *Manager) unregisterClient(client *Client) {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	var (
		ok bool
	)
	if _, ok = m.clients[client.uid]; ok {
		delete(m.clients, client.uid)
	}
}

func (m *Manager) listener() {
	ticker := time.NewTicker(5 * time.Minute)
	var (
		client *Client
	)
	for {
		select {
		case client = <-m.unregister:
			m.unregisterClient(client)
		case <-ticker.C:
			m.batchCreate(1000)
		}
	}
}

func (m *Manager) batchCreate(count int64) {
	var (
		i             int64
		step          int64 = 20 //需要偶数
		consumerCount       = int64(math.Ceil(float64(count) / float64(step)))
		wg                  = &sync.WaitGroup{}
	)

	for i = 0; i < consumerCount; i++ {
		wg.Add(1)
		go func(index int64, wg *sync.WaitGroup) {
			defer wg.Done()
			var (
				j int64
			)
			for j = index * step; j < (index+1)*step; j = j + 2 {
				m.newConnection(j, j+1)
			}
		}(i, wg)
	}
	wg.Wait()
	fmt.Println("准备发送消息:", time.Now())
	time.Sleep(5 * time.Second)
	fmt.Println("开始发送消息:", time.Now())
	msgTimer.Run()
	m.loopSend()
}

func (m *Manager) loopSend() {
	go func() {
		var (
			i      int64
			count  int
			ticker = time.NewTicker(time.Second * 1)
		)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				for i = 1; i < 10; i++ {
					if client, ok := m.clients[i]; ok {
						go client.SendUser(i + 1)
					}
				}
				count++
				if count > 0 {
					return
				}
			}
		}
	}()
}

func (m *Manager) newConnection(uid1 int64, uid2 int64) {
	var (
		client1 *Client
		client2 *Client
	)
	client1 = NewClient(uid1, m)
	client2 = NewClient(uid2, m)

	if client1.conn != nil && client2.conn != nil {
		m.rwLock.Lock()
		m.clients[uid1] = client1
		m.clients[uid2] = client2
		m.rwLock.Unlock()
	}
}
