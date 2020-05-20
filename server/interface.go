package main

import (
	"context"
	"os"
	"unsafe"

	lg "github.com/DavidSantia/grpc-agent/log"
	pb "github.com/DavidSantia/grpc-agent/protos"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// gRPC wrappers
type server struct {
	pb.UnimplementedGoAgentServer
}

// Datatypes that hold newrelic agent data
// Indexed with uint64 values (instead of memory pointers)

type TxnWithSeg struct {
	Txn  *newrelic.Transaction
	Seg  []*newrelic.Segment
	Host string
}

type App struct {
	Idx  uint64
	App  *newrelic.Application
	Host string
}

type AppMap map[uint64]App
type AppLookup map[string]uint64

type TxnSegMap map[uint64]TxnWithSeg

var appMap AppMap
var appLookup AppLookup
var txnSegMap TxnSegMap

func initData() {
	appMap = make(AppMap)
	appLookup = make(AppLookup)
	txnSegMap = make(TxnSegMap)
}

// implement CreateApp
func (s *server) CreateApp(ctx context.Context, in *pb.Config) (*pb.Index, error) {
	//log.Printf("CreateApp(%q)", in.GetName())

	idx, err := newApp(in.GetName(), in.GetLicense(), in.GetHost())
	return &pb.Index{Idx: idx}, err
}

func newApp(name, key, host string) (appPtr uint64, err error) {
	var app App

	appPtr = appLookup[name]
	if appPtr != 0 {
		return
	}

	// Create App
	app.App, err = newrelic.NewApplication(
		newrelic.ConfigAppName(name),
		newrelic.ConfigLicense(key),
		newrelic.ConfigInfoLogger(os.Stdout),
		newrelic.ConfigDistributedTracerEnabled(true),
	)
	if err != nil {
		return
	}

	// Convert pointer to uint64 and save in map
	appPtr = uint64(uintptr(unsafe.Pointer(app.App)))
	appLookup[name] = appPtr
	app.Idx = appPtr
	app.Host = host
	appMap[appPtr] = app
	return
}

// implement NewTxn
func (s *server) NewTxn(ctx context.Context, in *pb.NameIndex) (*pb.Index, error) {
	//log.Printf("NewTxn(%q)", in.GetName())

	idx := newTxn(in.GetIdx(), in.GetName())
	return &pb.Index{Idx: idx}, nil
}

func newTxn(appPtr uint64, name string) (txnSegPtr uint64) {
	var app App
	var txn *newrelic.Transaction

	app = appMap[appPtr]
	txn = app.App.StartTransaction(name)

	// Convert pointer to uint64 and save in map
	txnSegPtr = uint64(uintptr(unsafe.Pointer(txn)))
	txnSegMap[txnSegPtr] = TxnWithSeg{Txn: txn, Host: app.Host}
	return
}

// implement NewSeg
func (s *server) NewSeg(ctx context.Context, in *pb.NameIndex) (*pb.Index, error) {
	//log.Printf("NewSeg(%q)", in.GetName())

	newSeg(in.GetIdx(), in.GetName())
	return &pb.Index{Idx: in.GetIdx()}, nil
}

func newSeg(txnSegPtr uint64, name string) {
	var txnWithSeg TxnWithSeg
	var seg *newrelic.Segment

	txnWithSeg = txnSegMap[txnSegPtr]

	// Create new segment, append slice and update map
	seg = newrelic.StartSegment(txnWithSeg.Txn, name)
	txnWithSeg.Seg = append(txnWithSeg.Seg, seg)
	txnSegMap[txnSegPtr] = txnWithSeg
	return
}

// implement EndSeg
func (s *server) EndSeg(ctx context.Context, in *pb.Index) (*pb.Index, error) {
	//log.Printf("EndSeg()")

	endSeg(in.GetIdx())
	return &pb.Index{Idx: in.GetIdx()}, nil
}

func endSeg(txnSegPtr uint64) {
	var txnWithSeg TxnWithSeg
	var seg *newrelic.Segment

	txnWithSeg = txnSegMap[txnSegPtr]

	// End last segment in list, and update map
	i := len(txnWithSeg.Seg) - 1
	seg = txnWithSeg.Seg[i]
	seg.End()
	txnWithSeg.Seg = txnWithSeg.Seg[:i]
	txnSegMap[txnSegPtr] = txnWithSeg
}

// implement EndTxn
func (s *server) EndTxn(ctx context.Context, in *pb.Index) (*pb.Index, error) {
	//log.Printf("EndTxn()")

	endAll(in.GetIdx())
	return &pb.Index{Idx: in.GetIdx()}, nil
}

func endAll(txnSegPtr uint64) {
	var txnWithSeg TxnWithSeg
	var seg *newrelic.Segment

	txnWithSeg = txnSegMap[txnSegPtr]

	// End all segments in list
	i := len(txnWithSeg.Seg)
	for i > 0 {
		i--
		seg = txnWithSeg.Seg[i]
		seg.End()
		txnWithSeg.Seg = txnWithSeg.Seg[:i]
	}

	// End txn
	txnWithSeg.Txn.End()
	delete(txnSegMap, txnSegPtr)
}

// implement LogTxn
func (s *server) LogTxn(ctx context.Context, in *pb.LogIndex) (*pb.Index, error) {
	var txnWithSeg TxnWithSeg

	//log.Printf("LogTxn(%d): %q", in.GetIdx(), in.GetMessage())

	// Log message with txn context
	txnWithSeg = txnSegMap[in.GetIdx()]
	lg.DoLog(txnWithSeg.Txn, in.GetLevel(), in.GetMessage(), txnWithSeg.Host)
	return &pb.Index{Idx: in.GetIdx()}, nil
}
