package Rpc

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/rpc"
	"sort"
	"sync"
)
var(
	server *rpc.Server
)

func init()  {
	server = rpc.NewServer()
}


var errMissingParams = errors.New("jsonrpc: request body missing params")
type serverRequest struct {
	Method string           `json:"method"`
	Params *json.RawMessage `json:"params"`
	Id     *json.RawMessage `json:"id"`
}

func (r *serverRequest) reset() {
	r.Method = ""
	r.Params = nil
	r.Id = nil
}

type serverResponse struct {
	Id     *json.RawMessage `json:"id"`
	Result interface{}      `json:"result"`
	Error  interface{}      `json:"error"`
}

type jsonServerCodec struct {
	dec *json.Decoder // for reading JSON values
	enc *json.Encoder // for writing JSON values
	c   io.Closer

	// temporary work space
	req serverRequest

	// JSON-RPC clients can use arbitrary json values as request IDs.
	// Package rpc expects uint64 request IDs.
	// We assign uint64 sequence numbers to incoming requests
	// but save the original request ID in the pending map.
	// When rpc responds, we use the sequence number in
	// the response to find the original request ID.
	mutex   sync.Mutex // protects seq, pending
	seq     uint64
	pending map[uint64]*json.RawMessage
}

// NewServerCodec returns a new rpc.ServerCodec using JSON-RPC on conn.
func NewServerCodec(conn io.ReadWriteCloser) rpc.ServerCodec {
	return &jsonServerCodec{
		dec:     json.NewDecoder(conn),
		enc:     json.NewEncoder(conn),
		c:       conn,
		pending: make(map[uint64]*json.RawMessage),
	}
}

func (c *jsonServerCodec) ReadRequestHeader(r *rpc.Request) error {
	c.req.reset()
	if err := c.dec.Decode(&c.req); err != nil {
		return err
	}
	r.ServiceMethod = c.req.Method

	// JSON request id can be any JSON value;
	// RPC package expects uint64.  Translate to
	// internal uint64 and save JSON on the side.
	c.mutex.Lock()
	c.seq++
	c.pending[c.seq] = c.req.Id
	c.req.Id = nil
	r.Seq = c.seq
	c.mutex.Unlock()

	return nil
}

func (c *jsonServerCodec) ReadRequestBody(x interface{}) error {
	if x == nil {
		return nil
	}
	if c.req.Params == nil {
		return errMissingParams
	}
	// JSON params is array value.
	// RPC params is struct.
	// Unmarshal into array containing struct for now.
	// Should think about making RPC more general.
	var params [1]interface{}
	params[0] = x
	return json.Unmarshal(*c.req.Params, &params)
}

var null = json.RawMessage([]byte("null"))

func (c *jsonServerCodec) WriteResponse(r *rpc.Response, x interface{}) error {
	c.mutex.Lock()
	b, ok := c.pending[r.Seq]
	if !ok {
		c.mutex.Unlock()
		return errors.New("invalid sequence number in response")
	}
	delete(c.pending, r.Seq)
	c.mutex.Unlock()

	if b == nil {
		// Invalid request so no id. Use JSON null.
		b = &null
	}
	resp := serverResponse{Id: b}
	if r.Error == "" {
		resp.Result = x
	} else {
		resp.Error = r.Error
	}
	return c.enc.Encode(resp)
}

func (c *jsonServerCodec) Close() error {
	return c.c.Close()
}
type Server struct {
	codec jsonServerCodec
}

func NewRpcServer() *Server {
	return &Server{
	}
}

var connected = "200 Connected to Go RPC"

func (rs *Server) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	if req.Method != "CONNECT" {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = io.WriteString(w, "405 must CONNECT\n")
		return
	}
	conn, _, err := w.(http.Hijacker).Hijack()
	if err != nil {
		log.Print("rpc hijacking ", req.RemoteAddr, ": ", err.Error())
		return
	}
	_, _ = io.WriteString(conn, "HTTP/1.0 "+connected+"\n\n")

	rs.ServeConn(conn)
}

func (rs *Server) ServeConn(conn io.ReadWriteCloser) {
	server.ServeCodec(NewServerCodec(conn))
}

// HandleHTTP registers an HTTP handler for RPC messages on rpcPath,
// and a debugging handler on debugPath.
// It is still necessary to invoke http.Serve(), typically in a go statement.
func (rs *Server) HandleHTTP(rpcPath, debugPath string) {
	http.Handle(rpcPath, rs)
	http.Handle(debugPath, debugHTTP{rs})
}

func ListenAndServe(addr string, rs *Server) {
	for{
		err := http.ListenAndServe(addr, rs)
		if  err != nil{
			log.Println("")
		}
	}
}

func Register(v interface{}) error {
	return rpc.Register(v)
}

