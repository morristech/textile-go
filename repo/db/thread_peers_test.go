package db

import (
	"database/sql"
	"github.com/segmentio/ksuid"
	"github.com/textileio/textile-go/repo"
	"sync"
	"testing"
)

var threadpeerdb repo.ThreadPeerStore

func init() {
	setupThreadPeerDB()
}

func setupThreadPeerDB() {
	conn, _ := sql.Open("sqlite3", ":memory:")
	initDatabaseTables(conn, "")
	threadpeerdb = NewThreadPeerStore(conn, new(sync.Mutex))
}

func TestThreadPeerDB_Add(t *testing.T) {
	err := threadpeerdb.Add(&repo.ThreadPeer{
		Id:       "abc",
		ThreadId: ksuid.New().String(),
		Welcomed: false,
	})
	if err != nil {
		t.Error(err)
	}
	stmt, err := threadpeerdb.PrepareQuery("select id from thread_peers where id=?")
	defer stmt.Close()
	var id string
	err = stmt.QueryRow("abc").Scan(&id)
	if err != nil {
		t.Error(err)
	}
	if id != "abc" {
		t.Errorf(`expected id "abc" got %s`, id)
	}
}

func TestThreadPeerDB_ListById(t *testing.T) {
	setupThreadPeerDB()
	err := threadpeerdb.Add(&repo.ThreadPeer{
		Id:       ksuid.New().String(),
		ThreadId: ksuid.New().String(),
		Welcomed: false,
	})
	if err != nil {
		t.Error(err)
	}
	err = threadpeerdb.Add(&repo.ThreadPeer{
		Id:       "boo",
		ThreadId: ksuid.New().String(),
		Welcomed: false,
	})
	if err != nil {
		t.Error(err)
	}
	filtered := threadpeerdb.ListById("boo")
	if len(filtered) != 1 {
		t.Error("returned incorrect number of peers")
		return
	}
}

func TestThreadPeerDB_ListByThread(t *testing.T) {
	setupThreadPeerDB()
	err := threadpeerdb.Add(&repo.ThreadPeer{
		Id:       ksuid.New().String(),
		ThreadId: "foo",
		Welcomed: false,
	})
	if err != nil {
		t.Error(err)
	}
	err = threadpeerdb.Add(&repo.ThreadPeer{
		Id:       ksuid.New().String(),
		ThreadId: "boo",
		Welcomed: false,
	})
	if err != nil {
		t.Error(err)
	}
	filtered := threadpeerdb.ListByThread("boo")
	if len(filtered) != 1 {
		t.Error("returned incorrect number of peers")
		return
	}
}

func TestThreadPeerDB_Count(t *testing.T) {
	setupThreadPeerDB()
	err := threadpeerdb.Add(&repo.ThreadPeer{
		Id:       "bar",
		ThreadId: "1",
		Welcomed: false,
	})
	if err != nil {
		t.Error(err)
	}
	err = threadpeerdb.Add(&repo.ThreadPeer{
		Id:       "bar",
		ThreadId: "2",
		Welcomed: false,
	})
	if err != nil {
		t.Error(err)
	}
	err = threadpeerdb.Add(&repo.ThreadPeer{
		Id:       "bar2",
		ThreadId: "2",
		Welcomed: false,
	})
	if err != nil {
		t.Error(err)
	}
	cnt := threadpeerdb.Count(false)
	if cnt != 3 {
		t.Error("returned incorrect count of peers")
		return
	}
	distinct := threadpeerdb.Count(true)
	if distinct != 2 {
		t.Error("returned incorrect count of peers")
		return
	}
}

func TestThreadPeerDB_Delete(t *testing.T) {
	err := threadpeerdb.Add(&repo.ThreadPeer{
		Id:       "car",
		ThreadId: "3",
		Welcomed: false,
	})
	if err != nil {
		t.Error(err)
	}
	err = threadpeerdb.Delete("car", "3")
	if err != nil {
		t.Error(err)
	}
	stmt, err := threadpeerdb.PrepareQuery("select id from thread_peers where id=?")
	defer stmt.Close()
	var id string
	err = stmt.QueryRow("car").Scan(&id)
	if err == nil {
		t.Error("delete failed")
	}
}

func TestThreadPeerDB_DeleteById(t *testing.T) {
	err := threadpeerdb.DeleteById("bar")
	if err != nil {
		t.Error(err)
	}
	stmt, err := threadpeerdb.PrepareQuery("select id from thread_peers where id=?")
	defer stmt.Close()
	var id string
	err = stmt.QueryRow("bar").Scan(&id)
	if err == nil {
		t.Error("delete failed")
	}
}

func TestThreadPeerDB_DeleteByThread(t *testing.T) {
	err := threadpeerdb.DeleteByThread("2")
	if err != nil {
		t.Error(err)
	}
	stmt, err := threadpeerdb.PrepareQuery("select id from thread_peers where id=?")
	defer stmt.Close()
	var id string
	err = stmt.QueryRow("bar2").Scan(&id)
	if err == nil {
		t.Error("delete failed")
	}
}
