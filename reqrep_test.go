package mangosock

import (
	"crypto/rand"
	"fmt"
	"os"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/redsift/go-mangosock/nano"
)

const count = 1

func testRep(wg *sync.WaitGroup) {
	defer wg.Done()

	var err error
	var s nano.Rep
	if s, err = NewRepSocket(); err != nil {
		fmt.Println("NewRepSocket error: ", err)
		os.Exit(1)
	}
	if err = s.Bind("ipc:///tmp/nano1.sock"); err != nil {
		fmt.Println("Bind error: ", err)
		os.Exit(1)
	}

	i := 0
	for {
		rsp, err := s.Recv()
		if err != nil {
			fmt.Println("Recv error: ", err)
			os.Exit(1)
		}

		//fmt.Println("Received: ", string(rsp))
		fmt.Println("Received: ", len(rsp))

		_, err = s.Send([]byte("bye " + strconv.Itoa(i)))
		if err != nil {
			fmt.Println("Error sending request:", err)
			os.Exit(1)
		}

		fmt.Println("Sent response to req socket")

		i += 1
		if i >= count {
			break
		}
	}
}

func TestReqRep(t *testing.T) {
	st := time.Now()
	var wg sync.WaitGroup
	wg.Add(1)
	go testRep(&wg)

	newSock := func() nano.Req {
		var err error
		var s nano.Req
		if s, err = NewReqSocket(); err != nil {
			fmt.Println("NewReqSocket error: ", err)
			os.Exit(1)
		}

		_ = s.SetSendTimeout(2 * time.Second)
		_ = s.SetRecvTimeout(2 * time.Second)

		err = s.Connect("ipc:///tmp/nano1.sock")
		if err != nil {
			fmt.Println("Error connecting socket:", err)
		}

		fmt.Println("Connected to req socket")

		return s
	}

	s := newSock()
	//t1, _ := s.RecvTimeout()
	//fmt.Println("RecvTimeout=", t1)

	token := make([]byte, 4*1024*1024)
	rand.Read(token)

	for i := 0; i < count; i++ {
		//_, err := s.Send([]byte("hello " + strconv.Itoa(i)))
		_, err := s.Send(token)

		//fmt.Println("Send err:", err, err == syscall.ETIMEDOUT)
		if err != nil {
			fmt.Println("Error sending request:", err)
		}

		fmt.Println("Sent request")

		rsp, err := s.Recv()
		//fmt.Println("Received err:", err, err == syscall.ETIMEDOUT)
		if err != nil {
			fmt.Println("Error receiving response:", err)
			/*err = s.Close()
			s = newSock()
			if err != nil {
				fmt.Println("Error closing socket:", err)
			}*/
		} else {
			fmt.Println("Received: ", string(rsp))
		}
	}

	wg.Wait()

	fmt.Println("Time:", time.Since(st))
	fmt.Println("dummy log")
}
