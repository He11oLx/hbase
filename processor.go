package hbase

import (
	"context"
	"git.apache.org/thrift.git/lib/go/thrift"
)

type Processor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      Hbase
}

func NewProcessor(handler Hbase) *Processor {
	self95 := &Processor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self95.processorMap["enableTable"] = &ProcessorEnableTable{handler: handler}
	self95.processorMap["disableTable"] = &ProcessorDisableTable{handler: handler}
	self95.processorMap["isTableEnabled"] = &ProcessorIsTableEnabled{handler: handler}
	self95.processorMap["compact"] = &ProcessorCompact{handler: handler}
	self95.processorMap["majorCompact"] = &ProcessorMajorCompact{handler: handler}
	self95.processorMap["getTableNames"] = &ProcessorGetTableNames{handler: handler}
	self95.processorMap["getColumnDescriptors"] = &ProcessorGetColumnDescriptors{handler: handler}
	self95.processorMap["getTableRegions"] = &ProcessorGetTableRegions{handler: handler}
	self95.processorMap["createTable"] = &ProcessorCreateTable{handler: handler}
	self95.processorMap["deleteTable"] = &ProcessorDeleteTable{handler: handler}
	self95.processorMap["get"] = &ProcessorGet{handler: handler}
	self95.processorMap["getVer"] = &ProcessorGetVer{handler: handler}
	self95.processorMap["getVerTs"] = &ProcessorGetVerTs{handler: handler}
	self95.processorMap["getRow"] = &ProcessorGetRow{handler: handler}
	self95.processorMap["getRowWithColumns"] = &ProcessorGetRowWithColumns{handler: handler}
	self95.processorMap["getRowTs"] = &ProcessorGetRowTs{handler: handler}
	self95.processorMap["getRowWithColumnsTs"] = &ProcessorGetRowWithColumnsTs{handler: handler}
	self95.processorMap["getRows"] = &ProcessorGetRows{handler: handler}
	self95.processorMap["getRowsWithColumns"] = &ProcessorGetRowsWithColumns{handler: handler}
	self95.processorMap["getRowsTs"] = &ProcessorGetRowsTs{handler: handler}
	self95.processorMap["getRowsWithColumnsTs"] = &ProcessorGetRowsWithColumnsTs{handler: handler}
	self95.processorMap["mutateRow"] = &ProcessorMutateRow{handler: handler}
	self95.processorMap["mutateRowTs"] = &ProcessorMutateRowTs{handler: handler}
	self95.processorMap["mutateRows"] = &ProcessorMutateRows{handler: handler}
	self95.processorMap["mutateRowsTs"] = &ProcessorMutateRowsTs{handler: handler}
	self95.processorMap["atomicIncrement"] = &ProcessorAtomicIncrement{handler: handler}
	self95.processorMap["deleteAll"] = &ProcessorDeleteAll{handler: handler}
	self95.processorMap["deleteAllTs"] = &ProcessorDeleteAllTs{handler: handler}
	self95.processorMap["deleteAllRow"] = &ProcessorDeleteAllRow{handler: handler}
	self95.processorMap["increment"] = &ProcessorIncrement{handler: handler}
	self95.processorMap["incrementRows"] = &ProcessorIncrementRows{handler: handler}
	self95.processorMap["deleteAllRowTs"] = &ProcessorDeleteAllRowTs{handler: handler}
	self95.processorMap["scannerOpenWithScan"] = &ProcessorScannerOpenWithScan{handler: handler}
	self95.processorMap["scannerOpen"] = &ProcessorScannerOpen{handler: handler}
	self95.processorMap["scannerOpenWithStop"] = &ProcessorScannerOpenWithStop{handler: handler}
	self95.processorMap["scannerOpenWithPrefix"] = &ProcessorScannerOpenWithPrefix{handler: handler}
	self95.processorMap["scannerOpenTs"] = &ProcessorScannerOpenTs{handler: handler}
	self95.processorMap["scannerOpenWithStopTs"] = &ProcessorScannerOpenWithStopTs{handler: handler}
	self95.processorMap["scannerGet"] = &ProcessorScannerGet{handler: handler}
	self95.processorMap["scannerGetList"] = &ProcessorScannerGetList{handler: handler}
	self95.processorMap["scannerClose"] = &ProcessorScannerClose{handler: handler}
	self95.processorMap["getRegionInfo"] = &ProcessorGetRegionInfo{handler: handler}
	self95.processorMap["append"] = &ProcessorAppend{handler: handler}
	self95.processorMap["checkAndPut"] = &ProcessorCheckAndPut{handler: handler}
	return self95
}

func (p *Processor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *Processor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *Processor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func (p *Processor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x96 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x96.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x96

}

type ProcessorEnableTable struct {
	handler Hbase
}

func (p *ProcessorEnableTable) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := EnableTableArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("enableTable", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := EnableTableResult{}
	var err2 error
	if err2 = p.handler.EnableTable(ctx, args.TableName); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing enableTable: "+err2.Error())
			oprot.WriteMessageBegin("enableTable", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	}
	if err2 = oprot.WriteMessageBegin("enableTable", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorDisableTable struct {
	handler Hbase
}

func (p *ProcessorDisableTable) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := DisableTableArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("disableTable", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := DisableTableResult{}
	var err2 error
	if err2 = p.handler.DisableTable(ctx, args.TableName); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing disableTable: "+err2.Error())
			oprot.WriteMessageBegin("disableTable", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	}
	if err2 = oprot.WriteMessageBegin("disableTable", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorIsTableEnabled struct {
	handler Hbase
}

func (p *ProcessorIsTableEnabled) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := IsTableEnabledArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("isTableEnabled", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := IsTableEnabledResult{}
	var retval bool
	var err2 error
	if retval, err2 = p.handler.IsTableEnabled(ctx, args.TableName); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing isTableEnabled: "+err2.Error())
			oprot.WriteMessageBegin("isTableEnabled", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("isTableEnabled", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorCompact struct {
	handler Hbase
}

func (p *ProcessorCompact) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := CompactArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("compact", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := CompactResult{}
	var err2 error
	if err2 = p.handler.Compact(ctx, args.TableNameOrRegionName); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing compact: "+err2.Error())
			oprot.WriteMessageBegin("compact", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	}
	if err2 = oprot.WriteMessageBegin("compact", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorMajorCompact struct {
	handler Hbase
}

func (p *ProcessorMajorCompact) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := MajorCompactArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("majorCompact", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := MajorCompactResult{}
	var err2 error
	if err2 = p.handler.MajorCompact(ctx, args.TableNameOrRegionName); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing majorCompact: "+err2.Error())
			oprot.WriteMessageBegin("majorCompact", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	}
	if err2 = oprot.WriteMessageBegin("majorCompact", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorGetTableNames struct {
	handler Hbase
}

func (p *ProcessorGetTableNames) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := GetTableNamesArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getTableNames", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := GetTableNamesResult{}
	var retval [][]byte
	var err2 error
	if retval, err2 = p.handler.GetTableNames(ctx); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getTableNames: "+err2.Error())
			oprot.WriteMessageBegin("getTableNames", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getTableNames", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorGetColumnDescriptors struct {
	handler Hbase
}

func (p *ProcessorGetColumnDescriptors) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := GetColumnDescriptorsArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getColumnDescriptors", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := GetColumnDescriptorsResult{}
	var retval map[string]*ColumnDescriptor
	var err2 error
	if retval, err2 = p.handler.GetColumnDescriptors(ctx, args.TableName); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getColumnDescriptors: "+err2.Error())
			oprot.WriteMessageBegin("getColumnDescriptors", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getColumnDescriptors", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorGetTableRegions struct {
	handler Hbase
}

func (p *ProcessorGetTableRegions) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := GetTableRegionsArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getTableRegions", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := GetTableRegionsResult{}
	var retval []*TRegionInfo
	var err2 error
	if retval, err2 = p.handler.GetTableRegions(ctx, args.TableName); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getTableRegions: "+err2.Error())
			oprot.WriteMessageBegin("getTableRegions", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getTableRegions", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorCreateTable struct {
	handler Hbase
}

func (p *ProcessorCreateTable) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := CreateTableArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("createTable", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := CreateTableResult{}
	var err2 error
	if err2 = p.handler.CreateTable(ctx, args.TableName, args.ColumnFamilies); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		case *IllegalArgument:
			result.Ia = v
		case *AlreadyExists:
			result.Exist = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing createTable: "+err2.Error())
			oprot.WriteMessageBegin("createTable", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	}
	if err2 = oprot.WriteMessageBegin("createTable", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorDeleteTable struct {
	handler Hbase
}

func (p *ProcessorDeleteTable) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := DeleteTableArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("deleteTable", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := DeleteTableResult{}
	var err2 error
	if err2 = p.handler.DeleteTable(ctx, args.TableName); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing deleteTable: "+err2.Error())
			oprot.WriteMessageBegin("deleteTable", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	}
	if err2 = oprot.WriteMessageBegin("deleteTable", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorGet struct {
	handler Hbase
}

func (p *ProcessorGet) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := GetArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("get", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := GetResult{}
	var retval []*TCell
	var err2 error
	if retval, err2 = p.handler.Get(ctx, args.TableName, args.Row, args.Column, args.Attributes); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing get: "+err2.Error())
			oprot.WriteMessageBegin("get", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("get", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorGetVer struct {
	handler Hbase
}

func (p *ProcessorGetVer) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := GetVerArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getVer", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := GetVerResult{}
	var retval []*TCell
	var err2 error
	if retval, err2 = p.handler.GetVer(ctx, args.TableName, args.Row, args.Column, args.NumVersions, args.Attributes); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getVer: "+err2.Error())
			oprot.WriteMessageBegin("getVer", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getVer", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorGetVerTs struct {
	handler Hbase
}

func (p *ProcessorGetVerTs) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := GetVerTsArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getVerTs", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := GetVerTsResult{}
	var retval []*TCell
	var err2 error
	if retval, err2 = p.handler.GetVerTs(ctx, args.TableName, args.Row, args.Column, args.Timestamp, args.NumVersions, args.Attributes); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getVerTs: "+err2.Error())
			oprot.WriteMessageBegin("getVerTs", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getVerTs", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorGetRow struct {
	handler Hbase
}

func (p *ProcessorGetRow) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := GetRowArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getRow", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := GetRowResult{}
	var retval []*TRowResult_
	var err2 error
	if retval, err2 = p.handler.GetRow(ctx, args.TableName, args.Row, args.Attributes); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getRow: "+err2.Error())
			oprot.WriteMessageBegin("getRow", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getRow", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorGetRowWithColumns struct {
	handler Hbase
}

func (p *ProcessorGetRowWithColumns) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := GetRowWithColumnsArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getRowWithColumns", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := GetRowWithColumnsResult{}
	var retval []*TRowResult_
	var err2 error
	if retval, err2 = p.handler.GetRowWithColumns(ctx, args.TableName, args.Row, args.Columns, args.Attributes); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getRowWithColumns: "+err2.Error())
			oprot.WriteMessageBegin("getRowWithColumns", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getRowWithColumns", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorGetRowTs struct {
	handler Hbase
}

func (p *ProcessorGetRowTs) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := GetRowTsArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getRowTs", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := GetRowTsResult{}
	var retval []*TRowResult_
	var err2 error
	if retval, err2 = p.handler.GetRowTs(ctx, args.TableName, args.Row, args.Timestamp, args.Attributes); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getRowTs: "+err2.Error())
			oprot.WriteMessageBegin("getRowTs", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getRowTs", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorGetRowWithColumnsTs struct {
	handler Hbase
}

func (p *ProcessorGetRowWithColumnsTs) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := GetRowWithColumnsTsArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getRowWithColumnsTs", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := GetRowWithColumnsTsResult{}
	var retval []*TRowResult_
	var err2 error
	if retval, err2 = p.handler.GetRowWithColumnsTs(ctx, args.TableName, args.Row, args.Columns, args.Timestamp, args.Attributes); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getRowWithColumnsTs: "+err2.Error())
			oprot.WriteMessageBegin("getRowWithColumnsTs", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getRowWithColumnsTs", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorGetRows struct {
	handler Hbase
}

func (p *ProcessorGetRows) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := GetRowsArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getRows", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := GetRowsResult{}
	var retval []*TRowResult_
	var err2 error
	if retval, err2 = p.handler.GetRows(ctx, args.TableName, args.Rows, args.Attributes); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getRows: "+err2.Error())
			oprot.WriteMessageBegin("getRows", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getRows", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorGetRowsWithColumns struct {
	handler Hbase
}

func (p *ProcessorGetRowsWithColumns) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := GetRowsWithColumnsArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getRowsWithColumns", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := GetRowsWithColumnsResult{}
	var retval []*TRowResult_
	var err2 error
	if retval, err2 = p.handler.GetRowsWithColumns(ctx, args.TableName, args.Rows, args.Columns, args.Attributes); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getRowsWithColumns: "+err2.Error())
			oprot.WriteMessageBegin("getRowsWithColumns", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getRowsWithColumns", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorGetRowsTs struct {
	handler Hbase
}

func (p *ProcessorGetRowsTs) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := GetRowsTsArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getRowsTs", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := GetRowsTsResult{}
	var retval []*TRowResult_
	var err2 error
	if retval, err2 = p.handler.GetRowsTs(ctx, args.TableName, args.Rows, args.Timestamp, args.Attributes); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getRowsTs: "+err2.Error())
			oprot.WriteMessageBegin("getRowsTs", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getRowsTs", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorGetRowsWithColumnsTs struct {
	handler Hbase
}

func (p *ProcessorGetRowsWithColumnsTs) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := GetRowsWithColumnsTsArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getRowsWithColumnsTs", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := GetRowsWithColumnsTsResult{}
	var retval []*TRowResult_
	var err2 error
	if retval, err2 = p.handler.GetRowsWithColumnsTs(ctx, args.TableName, args.Rows, args.Columns, args.Timestamp, args.Attributes); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getRowsWithColumnsTs: "+err2.Error())
			oprot.WriteMessageBegin("getRowsWithColumnsTs", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getRowsWithColumnsTs", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorMutateRow struct {
	handler Hbase
}

func (p *ProcessorMutateRow) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := MutateRowArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("mutateRow", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := MutateRowResult{}
	var err2 error
	if err2 = p.handler.MutateRow(ctx, args.TableName, args.Row, args.Mutations, args.Attributes); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		case *IllegalArgument:
			result.Ia = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing mutateRow: "+err2.Error())
			oprot.WriteMessageBegin("mutateRow", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	}
	if err2 = oprot.WriteMessageBegin("mutateRow", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorMutateRowTs struct {
	handler Hbase
}

func (p *ProcessorMutateRowTs) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := MutateRowTsArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("mutateRowTs", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := MutateRowTsResult{}
	var err2 error
	if err2 = p.handler.MutateRowTs(ctx, args.TableName, args.Row, args.Mutations, args.Timestamp, args.Attributes); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		case *IllegalArgument:
			result.Ia = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing mutateRowTs: "+err2.Error())
			oprot.WriteMessageBegin("mutateRowTs", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	}
	if err2 = oprot.WriteMessageBegin("mutateRowTs", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorMutateRows struct {
	handler Hbase
}

func (p *ProcessorMutateRows) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := MutateRowsArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("mutateRows", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := MutateRowsResult{}
	var err2 error
	if err2 = p.handler.MutateRows(ctx, args.TableName, args.RowBatches, args.Attributes); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		case *IllegalArgument:
			result.Ia = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing mutateRows: "+err2.Error())
			oprot.WriteMessageBegin("mutateRows", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	}
	if err2 = oprot.WriteMessageBegin("mutateRows", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorMutateRowsTs struct {
	handler Hbase
}

func (p *ProcessorMutateRowsTs) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := MutateRowsTsArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("mutateRowsTs", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := MutateRowsTsResult{}
	var err2 error
	if err2 = p.handler.MutateRowsTs(ctx, args.TableName, args.RowBatches, args.Timestamp, args.Attributes); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		case *IllegalArgument:
			result.Ia = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing mutateRowsTs: "+err2.Error())
			oprot.WriteMessageBegin("mutateRowsTs", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	}
	if err2 = oprot.WriteMessageBegin("mutateRowsTs", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorAtomicIncrement struct {
	handler Hbase
}

func (p *ProcessorAtomicIncrement) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := AtomicIncrementArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("atomicIncrement", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := AtomicIncrementResult{}
	var retval int64
	var err2 error
	if retval, err2 = p.handler.AtomicIncrement(ctx, args.TableName, args.Row, args.Column, args.Value); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		case *IllegalArgument:
			result.Ia = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing atomicIncrement: "+err2.Error())
			oprot.WriteMessageBegin("atomicIncrement", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("atomicIncrement", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorDeleteAll struct {
	handler Hbase
}

func (p *ProcessorDeleteAll) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := DeleteAllArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("deleteAll", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := DeleteAllResult{}
	var err2 error
	if err2 = p.handler.DeleteAll(ctx, args.TableName, args.Row, args.Column, args.Attributes); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing deleteAll: "+err2.Error())
			oprot.WriteMessageBegin("deleteAll", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	}
	if err2 = oprot.WriteMessageBegin("deleteAll", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorDeleteAllTs struct {
	handler Hbase
}

func (p *ProcessorDeleteAllTs) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := DeleteAllTsArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("deleteAllTs", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := DeleteAllTsResult{}
	var err2 error
	if err2 = p.handler.DeleteAllTs(ctx, args.TableName, args.Row, args.Column, args.Timestamp, args.Attributes); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing deleteAllTs: "+err2.Error())
			oprot.WriteMessageBegin("deleteAllTs", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	}
	if err2 = oprot.WriteMessageBegin("deleteAllTs", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorDeleteAllRow struct {
	handler Hbase
}

func (p *ProcessorDeleteAllRow) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := DeleteAllRowArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("deleteAllRow", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := DeleteAllRowResult{}
	var err2 error
	if err2 = p.handler.DeleteAllRow(ctx, args.TableName, args.Row, args.Attributes); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing deleteAllRow: "+err2.Error())
			oprot.WriteMessageBegin("deleteAllRow", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	}
	if err2 = oprot.WriteMessageBegin("deleteAllRow", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorIncrement struct {
	handler Hbase
}

func (p *ProcessorIncrement) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := IncrementArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("increment", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := IncrementResult{}
	var err2 error
	if err2 = p.handler.Increment(ctx, args.Increment); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing increment: "+err2.Error())
			oprot.WriteMessageBegin("increment", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	}
	if err2 = oprot.WriteMessageBegin("increment", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorIncrementRows struct {
	handler Hbase
}

func (p *ProcessorIncrementRows) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := IncrementRowsArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("incrementRows", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := IncrementRowsResult{}
	var err2 error
	if err2 = p.handler.IncrementRows(ctx, args.Increments); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing incrementRows: "+err2.Error())
			oprot.WriteMessageBegin("incrementRows", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	}
	if err2 = oprot.WriteMessageBegin("incrementRows", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorDeleteAllRowTs struct {
	handler Hbase
}

func (p *ProcessorDeleteAllRowTs) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := DeleteAllRowTsArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("deleteAllRowTs", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := DeleteAllRowTsResult{}
	var err2 error
	if err2 = p.handler.DeleteAllRowTs(ctx, args.TableName, args.Row, args.Timestamp, args.Attributes); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing deleteAllRowTs: "+err2.Error())
			oprot.WriteMessageBegin("deleteAllRowTs", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	}
	if err2 = oprot.WriteMessageBegin("deleteAllRowTs", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorScannerOpenWithScan struct {
	handler Hbase
}

func (p *ProcessorScannerOpenWithScan) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := ScannerOpenWithScanArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("scannerOpenWithScan", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := ScannerOpenWithScanResult{}
	var retval ScannerID
	var err2 error
	if retval, err2 = p.handler.ScannerOpenWithScan(ctx, args.TableName, args.Scan, args.Attributes); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing scannerOpenWithScan: "+err2.Error())
			oprot.WriteMessageBegin("scannerOpenWithScan", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("scannerOpenWithScan", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorScannerOpen struct {
	handler Hbase
}

func (p *ProcessorScannerOpen) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := ScannerOpenArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("scannerOpen", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := ScannerOpenResult{}
	var retval ScannerID
	var err2 error
	if retval, err2 = p.handler.ScannerOpen(ctx, args.TableName, args.StartRow, args.Columns, args.Attributes); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing scannerOpen: "+err2.Error())
			oprot.WriteMessageBegin("scannerOpen", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("scannerOpen", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorScannerOpenWithStop struct {
	handler Hbase
}

func (p *ProcessorScannerOpenWithStop) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := ScannerOpenWithStopArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("scannerOpenWithStop", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := ScannerOpenWithStopResult{}
	var retval ScannerID
	var err2 error
	if retval, err2 = p.handler.ScannerOpenWithStop(ctx, args.TableName, args.StartRow, args.StopRow, args.Columns, args.Attributes); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing scannerOpenWithStop: "+err2.Error())
			oprot.WriteMessageBegin("scannerOpenWithStop", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("scannerOpenWithStop", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorScannerOpenWithPrefix struct {
	handler Hbase
}

func (p *ProcessorScannerOpenWithPrefix) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := ScannerOpenWithPrefixArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("scannerOpenWithPrefix", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := ScannerOpenWithPrefixResult{}
	var retval ScannerID
	var err2 error
	if retval, err2 = p.handler.ScannerOpenWithPrefix(ctx, args.TableName, args.StartAndPrefix, args.Columns, args.Attributes); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing scannerOpenWithPrefix: "+err2.Error())
			oprot.WriteMessageBegin("scannerOpenWithPrefix", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("scannerOpenWithPrefix", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorScannerOpenTs struct {
	handler Hbase
}

func (p *ProcessorScannerOpenTs) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := ScannerOpenTsArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("scannerOpenTs", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := ScannerOpenTsResult{}
	var retval ScannerID
	var err2 error
	if retval, err2 = p.handler.ScannerOpenTs(ctx, args.TableName, args.StartRow, args.Columns, args.Timestamp, args.Attributes); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing scannerOpenTs: "+err2.Error())
			oprot.WriteMessageBegin("scannerOpenTs", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("scannerOpenTs", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorScannerOpenWithStopTs struct {
	handler Hbase
}

func (p *ProcessorScannerOpenWithStopTs) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := ScannerOpenWithStopTsArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("scannerOpenWithStopTs", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := ScannerOpenWithStopTsResult{}
	var retval ScannerID
	var err2 error
	if retval, err2 = p.handler.ScannerOpenWithStopTs(ctx, args.TableName, args.StartRow, args.StopRow, args.Columns, args.Timestamp, args.Attributes); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing scannerOpenWithStopTs: "+err2.Error())
			oprot.WriteMessageBegin("scannerOpenWithStopTs", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("scannerOpenWithStopTs", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorScannerGet struct {
	handler Hbase
}

func (p *ProcessorScannerGet) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := ScannerGetArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("scannerGet", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := ScannerGetResult{}
	var retval []*TRowResult_
	var err2 error
	if retval, err2 = p.handler.ScannerGet(ctx, args.ID); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		case *IllegalArgument:
			result.Ia = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing scannerGet: "+err2.Error())
			oprot.WriteMessageBegin("scannerGet", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("scannerGet", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorScannerGetList struct {
	handler Hbase
}

func (p *ProcessorScannerGetList) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := ScannerGetListArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("scannerGetList", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := ScannerGetListResult{}
	var retval []*TRowResult_
	var err2 error
	if retval, err2 = p.handler.ScannerGetList(ctx, args.ID, args.NbRows); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		case *IllegalArgument:
			result.Ia = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing scannerGetList: "+err2.Error())
			oprot.WriteMessageBegin("scannerGetList", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("scannerGetList", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorScannerClose struct {
	handler Hbase
}

func (p *ProcessorScannerClose) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := ScannerCloseArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("scannerClose", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := ScannerCloseResult{}
	var err2 error
	if err2 = p.handler.ScannerClose(ctx, args.ID); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		case *IllegalArgument:
			result.Ia = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing scannerClose: "+err2.Error())
			oprot.WriteMessageBegin("scannerClose", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	}
	if err2 = oprot.WriteMessageBegin("scannerClose", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorGetRegionInfo struct {
	handler Hbase
}

func (p *ProcessorGetRegionInfo) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := GetRegionInfoArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getRegionInfo", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := GetRegionInfoResult{}
	var retval *TRegionInfo
	var err2 error
	if retval, err2 = p.handler.GetRegionInfo(ctx, args.Row); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getRegionInfo: "+err2.Error())
			oprot.WriteMessageBegin("getRegionInfo", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getRegionInfo", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorAppend struct {
	handler Hbase
}

func (p *ProcessorAppend) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := AppendArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("append", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := AppendResult{}
	var retval []*TCell
	var err2 error
	if retval, err2 = p.handler.Append(ctx, args.Append); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing append: "+err2.Error())
			oprot.WriteMessageBegin("append", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("append", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ProcessorCheckAndPut struct {
	handler Hbase
}

func (p *ProcessorCheckAndPut) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := CheckAndPutArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("checkAndPut", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := CheckAndPutResult{}
	var retval bool
	var err2 error
	if retval, err2 = p.handler.CheckAndPut(ctx, args.TableName, args.Row, args.Column, args.Value, args.Mput, args.Attributes); err2 != nil {
		switch v := err2.(type) {
		case *IOError:
			result.Io = v
		case *IllegalArgument:
			result.Ia = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing checkAndPut: "+err2.Error())
			oprot.WriteMessageBegin("checkAndPut", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("checkAndPut", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}
