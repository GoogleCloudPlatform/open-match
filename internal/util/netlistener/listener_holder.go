package netlistener

import (
	"fmt"
	"net"

	"sync"

	"github.com/pkg/errors"
)

// Port is a preallocated port that is guaranteed to only hand off go routine.
type ListenerHolder struct {
	number   int
	listener net.Listener
	sync.RWMutex
}

// Obtain returns the TCP listener. This method can only be called once and is thread-safe.
func (lh *ListenerHolder) Obtain() (net.Listener, error) {
	lh.Lock()
	defer lh.Unlock()
	listener := lh.listener
	lh.listener = nil
	if listener == nil {
		return nil, errors.WithStack(fmt.Errorf("cannot Obtain() listener for %d because already handed off", lh.number))
	}
	return listener, nil
}

// Number returns the port number.
func (lh *ListenerHolder) Number() int {
	return lh.number
}

// Close shutsdown the TCP listener.
func (lh *ListenerHolder) Close() error {
	lh.Lock()
	defer lh.Unlock()
	if lh.listener != nil {
		err := lh.listener.Close()
		lh.listener = nil
		return err
	}
	return nil
}

// NewFromPortNumber opens a TCP based on the port number provided.
func NewFromPortNumber(portNumber int) (*ListenerHolder, error) {
	addr := fmt.Sprintf(":%d", portNumber)
	conn, err := net.Listen("tcp", addr)

	tcpConn, ok := conn.Addr().(*net.TCPAddr)
	if !ok || tcpConn == nil {
		return nil, fmt.Errorf("net.Listen(\"tcp\", %s) did not return a *net.TCPAddr", addr)
	}

	if err == nil {
		return &ListenerHolder{
			number:   tcpConn.Port,
			listener: conn,
		}, nil
	}
	return nil, err
}

/*
// CreatePort creates a Port object from TCP listener.
func Create(portNumber int, listener net.Listener) *ListenerHolder {
	return &ListenerHolder{
		number:   portNumber,
		listener: listener,
	}
}
*/
