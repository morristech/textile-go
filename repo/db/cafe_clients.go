package db

import (
	"database/sql"
	"github.com/textileio/textile-go/repo"
	"sync"
	"time"
)

type CafeClientDB struct {
	modelStore
}

func NewCafeClientStore(db *sql.DB, lock *sync.Mutex) repo.CafeClientStore {
	return &CafeClientDB{modelStore{db, lock}}
}

func (c *CafeClientDB) Add(client *repo.CafeClient) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}
	stm := `insert into cafe_clients(id, address, created, lastSeen) values(?,?,?,?)`
	stmt, err := tx.Prepare(stm)
	if err != nil {
		log.Errorf("error in tx prepare: %s", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		client.Id,
		client.Address,
		int(client.Created.Unix()),
		int(client.LastSeen.Unix()),
	)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (c *CafeClientDB) Get(id string) *repo.CafeClient {
	c.lock.Lock()
	defer c.lock.Unlock()
	ret := c.handleQuery("select * from cafe_clients where id='" + id + "';")
	if len(ret) == 0 {
		return nil
	}
	return &ret[0]
}

func (c *CafeClientDB) Count() int {
	c.lock.Lock()
	defer c.lock.Unlock()
	row := c.db.QueryRow("select Count(*) from cafe_clients;")
	var count int
	row.Scan(&count)
	return count
}

func (c *CafeClientDB) List() []repo.CafeClient {
	c.lock.Lock()
	defer c.lock.Unlock()
	stm := "select * from cafe_clients order by lastSeen desc;"
	return c.handleQuery(stm)
}

func (c *CafeClientDB) ListByAddress(address string) []repo.CafeClient {
	c.lock.Lock()
	defer c.lock.Unlock()
	stm := "select * from cafe_clients where address='" + address + "' order by lastSeen desc;"
	return c.handleQuery(stm)
}

func (c *CafeClientDB) UpdateLastSeen(id string, date time.Time) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	_, err := c.db.Exec("update cafe_clients set lastSeen=? where id=?", int(date.Unix()), id)
	return err
}

func (c *CafeClientDB) Delete(id string) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	_, err := c.db.Exec("delete from cafe_clients where id=?", id)
	return err
}

func (c *CafeClientDB) handleQuery(stm string) []repo.CafeClient {
	var ret []repo.CafeClient
	rows, err := c.db.Query(stm)
	if err != nil {
		log.Errorf("error in db query: %s", err)
		return nil
	}
	for rows.Next() {
		var id, address string
		var createdInt, lastSeenInt int
		if err := rows.Scan(&id, &address, &createdInt, &lastSeenInt); err != nil {
			log.Errorf("error in db scan: %s", err)
			continue
		}
		accnt := repo.CafeClient{
			Id:       id,
			Address:  address,
			Created:  time.Unix(int64(createdInt), 0),
			LastSeen: time.Unix(int64(lastSeenInt), 0),
		}
		ret = append(ret, accnt)
	}
	return ret
}
