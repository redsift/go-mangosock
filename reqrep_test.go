package mangosock

import (
	"crypto/rand"
	"os"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/redsift/go-mangosock/nano"
)

const count = 1

func testRep(t *testing.T, ready, wg *sync.WaitGroup, path string) {
	defer wg.Done()
	readyNow := sync.OnceFunc(ready.Done)
	defer readyNow()

	t.Logf("testRep(%q)", path)

	var err error
	var s nano.Rep
	if s, err = NewRepSocket(); err != nil {
		t.Errorf("NewRepSocket error: %v", err)
		return
	}
	if err = s.Bind("ipc://" + path + "/nano1.sock"); err != nil {
		t.Errorf("Bind error: %v", err)
		return
	}

	t.Logf("testRep(%q) ready", path)
	readyNow()

	i := 0
	for {
		t.Logf("testRep.for(i=%d)", i)
		rsp, err := s.Recv()
		if err != nil {
			t.Errorf("Recv error: %v", err)
			return
		}

		//t.Log("Received: ", string(rsp))
		t.Logf("Received: %d", len(rsp))

		_, err = s.Send([]byte("bye " + strconv.Itoa(i)))
		if err != nil {
			t.Errorf("Error sending request: %v", err)
			return
		}

		t.Log("Sent response to req socket")

		i += 1
		if i >= count {
			break
		}
	}
}

func TestReqRep(t *testing.T) {
	st := time.Now()
	var wg, ready sync.WaitGroup
	wg.Add(1)
	ready.Add(1)

	path := t.TempDir()

	go testRep(t, &ready, &wg, path)

	newSock := func() nano.Req {
		var err error
		var s nano.Req

		t.Logf("newSock(%q)", path)

		if s, err = NewReqSocket(); err != nil {
			t.Log("NewReqSocket error: ", err)
			os.Exit(1)
		}

		_ = s.SetSendTimeout(2 * time.Second)
		_ = s.SetRecvTimeout(2 * time.Second)

		t.Logf("newSock(%q) -> connect", path)

		err = s.Connect("ipc://" + path + "/nano1.sock")
		if err != nil {
			t.Log("Error connecting socket:", err)
		}

		t.Log("Connected to req socket")

		return s
	}

	ready.Wait()

	s := newSock()
	//t1, _ := s.RecvTimeout()
	//t.Log("RecvTimeout=", t1)

	token := make([]byte, 4*1024*1024)
	rand.Read(token)

	for i := 0; i < count; i++ {
		//_, err := s.Send([]byte("hello " + strconv.Itoa(i)))
		_, err := s.Send(token)

		//t.Log("Send err:", err, err == syscall.ETIMEDOUT)
		if err != nil {
			t.Log("Error sending request:", err)
		}

		t.Log("Sent request")

		rsp, err := s.Recv()
		//t.Log("Received err:", err, err == syscall.ETIMEDOUT)
		if err != nil {
			t.Log("Error receiving response:", err)
			/*err = s.Close()
			s = newSock()
			if err != nil {
				t.Log("Error closing socket:", err)
			}*/
		} else {
			t.Log("Received: ", string(rsp))
		}
	}

	wg.Wait()

	t.Log("Time:", time.Since(st))
	t.Log("dummy log")
}
