package Rpc

import (
	"Juliet/Utils"
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/rpc"
	"sync"
)

type clientCodec struct {
	rwc    io.ReadWriteCloser
	dec    *json.Decoder
	enc    *json.Encoder
	req    clientRequest
	resp   clientResponse
	mutex   sync.Mutex
	pending map[uint64]string
	closer io.Closer
}

type clientRequest struct {
	Method string
	Params [1]interface{}
	Id uint64
}

type clientResponse struct {
	Id uint64
	Result *json.RawMessage
	Error interface{}
	mutex   sync.Mutex
}
func (r *clientResponse) reset() {
	r.Id = 0
	r.Result = nil
	r.Error = nil
}

func (c *clientCodec) WriteRequest(r *rpc.Request, param interface{}) error {
	c.mutex.Lock()
	c.pending[r.Seq] = r.ServiceMethod
	c.mutex.Unlock()
	c.req.Method = r.ServiceMethod
	c.req.Params[0] = param
	c.req.Id = r.Seq
	return c.enc.Encode(&c.req)
}

func (c *clientCodec) ReadResponseHeader(r *rpc.Response) error  {
	c.resp.reset()
	if err := c.dec.Decode(&c.resp); err != nil{
		return err
	}

	c.mutex.Lock()
	r.ServiceMethod = c.pending[c.resp.Id]
	delete(c.pending, c.resp.Id)
	c.mutex.Unlock()
	r.Error = ""
	r.Seq = c.resp.Id
	if c.resp.Error != nil || c.resp.Result == nil {
		x, ok := c.resp.Error.(string)
		if !ok {
			return fmt.Errorf("invalid error %v", c.resp.Error)
		}
		if x == "" {
			x = "unspecified error"
		}
		r.Error = x
	}
	return nil
}

func (c *clientCodec) ReadResponseBody(v interface{}) error {
	if v == nil {
		return nil
	}
	return json.Unmarshal(*c.resp.Result, v)
}

func (c *clientCodec) Close() error {
	return c.closer.Close()
}

type RpcClient struct {
	client *rpc.Client
}

func NewRpcClient(conn io.ReadWriteCloser) *RpcClient {
	codec := &clientCodec{
		rwc:    conn,
		dec:    json.NewDecoder(conn),
		enc:    json.NewEncoder(conn),
	}
	return &RpcClient{
		client: rpc.NewClientWithCodec(codec),
	}
}
func DialHTTPPath(network, address, path string) (*RpcClient, error) {
	var err error
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	_, _ = io.WriteString(conn, "CONNECT "+path+" HTTP/1.0\n\n")

	// Require successful HTTP response
	// before switching to RPC protocol.
	resp, err := http.ReadResponse(bufio.NewReader(conn), &http.Request{Method: "CONNECT"})
	if err == nil && resp.Status == connected {
		return NewRpcClient(conn), nil
	}
	if err == nil {
		err = errors.New("unexpected HTTP response: " + resp.Status)
	}
	Utils.Log(conn.Close())
	return nil, &net.OpError{
		Op:   "dial-http",
		Net:  network + " " + address,
		Addr: nil,
		Err:  err,
	}
}

func Dial(network, address string) (*RpcClient, error)  {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return NewRpcClient(conn), nil
}

func (client *RpcClient) Call(serviceMethod string, args interface{}, reply interface{}) error {
	call := <- client.client.Go(serviceMethod, args, reply, make(chan *rpc.Call, 1)).Done
	return call.Error
}

