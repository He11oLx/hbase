package hbase

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// TCell - Used to transport a cell value (byte[]) and the timestamp it was
// stored with together as a result for get and getRow methods. This promotes
// the timestamp of a cell to a first-class value, making it easy to take
// note of temporal data. Cell is used all the way from HStore up to HTable.
//
// Attributes:
//  - Value
//  - Timestamp
type TCell struct {
	Value     []byte `thrift:"value,1" db:"value" json:"value"`
	Timestamp int64  `thrift:"timestamp,2" db:"timestamp" json:"timestamp"`
}

func NewTCell() *TCell {
	return &TCell{}
}

func (p *TCell) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.I64 {
				if err := p.ReadField2(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TCell) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Value = v
	}
	return nil
}

func (p *TCell) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Timestamp = v
	}
	return nil
}

func (p *TCell) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TCell"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField2(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TCell) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("value", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:value: ", p), err)
	}
	if err := oprot.WriteBinary(p.Value); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.value (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:value: ", p), err)
	}
	return err
}

func (p *TCell) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("timestamp", thrift.I64, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:timestamp: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.Timestamp)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.timestamp (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:timestamp: ", p), err)
	}
	return err
}

func (p *TCell) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TCell(%+v)", *p)
}

// An HColumnDescriptor contains information about a column family
// such as the number of versions, compression settings, etc. It is
// used as input when creating a table or adding a column.
//
// Attributes:
//  - Name
//  - MaxVersions
//  - Compression
//  - InMemory
//  - BloomFilterType
//  - BloomFilterVectorSize
//  - BloomFilterNbHashes
//  - BlockCacheEnabled
//  - TimeToLive
type ColumnDescriptor struct {
	Name                  []byte `thrift:"name,1" db:"name" json:"name"`
	MaxVersions           int32  `thrift:"maxVersions,2" db:"maxVersions" json:"maxVersions"`
	Compression           string `thrift:"compression,3" db:"compression" json:"compression"`
	InMemory              bool   `thrift:"inMemory,4" db:"inMemory" json:"inMemory"`
	BloomFilterType       string `thrift:"bloomFilterType,5" db:"bloomFilterType" json:"bloomFilterType"`
	BloomFilterVectorSize int32  `thrift:"bloomFilterVectorSize,6" db:"bloomFilterVectorSize" json:"bloomFilterVectorSize"`
	BloomFilterNbHashes   int32  `thrift:"bloomFilterNbHashes,7" db:"bloomFilterNbHashes" json:"bloomFilterNbHashes"`
	BlockCacheEnabled     bool   `thrift:"blockCacheEnabled,8" db:"blockCacheEnabled" json:"blockCacheEnabled"`
	TimeToLive            int32  `thrift:"timeToLive,9" db:"timeToLive" json:"timeToLive"`
}

func NewColumnDescriptor() *ColumnDescriptor {
	return &ColumnDescriptor{
		MaxVersions: 3,

		Compression: "NONE",

		BloomFilterType: "NONE",

		TimeToLive: 2147483647,
	}
}

func (p *ColumnDescriptor) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField2(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField3(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 4:
			if fieldTypeId == thrift.BOOL {
				if err := p.ReadField4(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 5:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField5(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 6:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField6(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 7:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField7(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 8:
			if fieldTypeId == thrift.BOOL {
				if err := p.ReadField8(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 9:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField9(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *ColumnDescriptor) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Name = v
	}
	return nil
}

func (p *ColumnDescriptor) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.MaxVersions = v
	}
	return nil
}

func (p *ColumnDescriptor) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Compression = v
	}
	return nil
}

func (p *ColumnDescriptor) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.InMemory = v
	}
	return nil
}

func (p *ColumnDescriptor) ReadField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.BloomFilterType = v
	}
	return nil
}

func (p *ColumnDescriptor) ReadField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 6: ", err)
	} else {
		p.BloomFilterVectorSize = v
	}
	return nil
}

func (p *ColumnDescriptor) ReadField7(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 7: ", err)
	} else {
		p.BloomFilterNbHashes = v
	}
	return nil
}

func (p *ColumnDescriptor) ReadField8(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return thrift.PrependError("error reading field 8: ", err)
	} else {
		p.BlockCacheEnabled = v
	}
	return nil
}

func (p *ColumnDescriptor) ReadField9(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 9: ", err)
	} else {
		p.TimeToLive = v
	}
	return nil
}

func (p *ColumnDescriptor) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ColumnDescriptor"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField2(oprot); err != nil {
			return err
		}
		if err := p.writeField3(oprot); err != nil {
			return err
		}
		if err := p.writeField4(oprot); err != nil {
			return err
		}
		if err := p.writeField5(oprot); err != nil {
			return err
		}
		if err := p.writeField6(oprot); err != nil {
			return err
		}
		if err := p.writeField7(oprot); err != nil {
			return err
		}
		if err := p.writeField8(oprot); err != nil {
			return err
		}
		if err := p.writeField9(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *ColumnDescriptor) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("name", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:name: ", p), err)
	}
	if err := oprot.WriteBinary(p.Name); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.name (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:name: ", p), err)
	}
	return err
}

func (p *ColumnDescriptor) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("maxVersions", thrift.I32, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:maxVersions: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.MaxVersions)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.maxVersions (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:maxVersions: ", p), err)
	}
	return err
}

func (p *ColumnDescriptor) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("compression", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:compression: ", p), err)
	}
	if err := oprot.WriteString(string(p.Compression)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.compression (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:compression: ", p), err)
	}
	return err
}

func (p *ColumnDescriptor) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("inMemory", thrift.BOOL, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:inMemory: ", p), err)
	}
	if err := oprot.WriteBool(bool(p.InMemory)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.inMemory (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:inMemory: ", p), err)
	}
	return err
}

func (p *ColumnDescriptor) writeField5(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("bloomFilterType", thrift.STRING, 5); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:bloomFilterType: ", p), err)
	}
	if err := oprot.WriteString(string(p.BloomFilterType)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.bloomFilterType (5) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 5:bloomFilterType: ", p), err)
	}
	return err
}

func (p *ColumnDescriptor) writeField6(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("bloomFilterVectorSize", thrift.I32, 6); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:bloomFilterVectorSize: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.BloomFilterVectorSize)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.bloomFilterVectorSize (6) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 6:bloomFilterVectorSize: ", p), err)
	}
	return err
}

func (p *ColumnDescriptor) writeField7(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("bloomFilterNbHashes", thrift.I32, 7); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:bloomFilterNbHashes: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.BloomFilterNbHashes)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.bloomFilterNbHashes (7) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 7:bloomFilterNbHashes: ", p), err)
	}
	return err
}

func (p *ColumnDescriptor) writeField8(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("blockCacheEnabled", thrift.BOOL, 8); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 8:blockCacheEnabled: ", p), err)
	}
	if err := oprot.WriteBool(bool(p.BlockCacheEnabled)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.blockCacheEnabled (8) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 8:blockCacheEnabled: ", p), err)
	}
	return err
}

func (p *ColumnDescriptor) writeField9(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("timeToLive", thrift.I32, 9); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 9:timeToLive: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.TimeToLive)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.timeToLive (9) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 9:timeToLive: ", p), err)
	}
	return err
}

func (p *ColumnDescriptor) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ColumnDescriptor(%+v)", *p)
}

// A TRegionInfo contains information about an HTable region.
//
// Attributes:
//  - StartKey
//  - EndKey
//  - ID
//  - Name
//  - Version
//  - ServerName
//  - Port
type TRegionInfo struct {
	StartKey   []byte `thrift:"startKey,1" db:"startKey" json:"startKey"`
	EndKey     []byte `thrift:"endKey,2" db:"endKey" json:"endKey"`
	ID         int64  `thrift:"id,3" db:"id" json:"id"`
	Name       []byte `thrift:"name,4" db:"name" json:"name"`
	Version    int8   `thrift:"version,5" db:"version" json:"version"`
	ServerName []byte `thrift:"serverName,6" db:"serverName" json:"serverName"`
	Port       int32  `thrift:"port,7" db:"port" json:"port"`
}

func NewTRegionInfo() *TRegionInfo {
	return &TRegionInfo{}
}

func (p *TRegionInfo) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.I64 {
				if err := p.ReadField3(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 4:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField4(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 5:
			if fieldTypeId == thrift.BYTE {
				if err := p.ReadField5(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 6:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField6(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 7:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField7(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TRegionInfo) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.StartKey = v
	}
	return nil
}

func (p *TRegionInfo) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.EndKey = v
	}
	return nil
}

func (p *TRegionInfo) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.ID = v
	}
	return nil
}

func (p *TRegionInfo) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.Name = v
	}
	return nil
}

func (p *TRegionInfo) ReadField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadByte(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.Version = int8(v)
	}
	return nil
}

func (p *TRegionInfo) ReadField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 6: ", err)
	} else {
		p.ServerName = v
	}
	return nil
}

func (p *TRegionInfo) ReadField7(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 7: ", err)
	} else {
		p.Port = v
	}
	return nil
}

func (p *TRegionInfo) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TRegionInfo"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField2(oprot); err != nil {
			return err
		}
		if err := p.writeField3(oprot); err != nil {
			return err
		}
		if err := p.writeField4(oprot); err != nil {
			return err
		}
		if err := p.writeField5(oprot); err != nil {
			return err
		}
		if err := p.writeField6(oprot); err != nil {
			return err
		}
		if err := p.writeField7(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TRegionInfo) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("startKey", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:startKey: ", p), err)
	}
	if err := oprot.WriteBinary(p.StartKey); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.startKey (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:startKey: ", p), err)
	}
	return err
}

func (p *TRegionInfo) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("endKey", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:endKey: ", p), err)
	}
	if err := oprot.WriteBinary(p.EndKey); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.endKey (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:endKey: ", p), err)
	}
	return err
}

func (p *TRegionInfo) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("id", thrift.I64, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:id: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.ID)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.id (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:id: ", p), err)
	}
	return err
}

func (p *TRegionInfo) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("name", thrift.STRING, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:name: ", p), err)
	}
	if err := oprot.WriteBinary(p.Name); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.name (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:name: ", p), err)
	}
	return err
}

func (p *TRegionInfo) writeField5(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("version", thrift.BYTE, 5); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:version: ", p), err)
	}
	if err := oprot.WriteByte(int8(p.Version)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.version (5) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 5:version: ", p), err)
	}
	return err
}

func (p *TRegionInfo) writeField6(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("serverName", thrift.STRING, 6); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:serverName: ", p), err)
	}
	if err := oprot.WriteBinary(p.ServerName); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.serverName (6) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 6:serverName: ", p), err)
	}
	return err
}

func (p *TRegionInfo) writeField7(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("port", thrift.I32, 7); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:port: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Port)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.port (7) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 7:port: ", p), err)
	}
	return err
}

func (p *TRegionInfo) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TRegionInfo(%+v)", *p)
}

// A Mutation object is used to either update or delete a column-value.
//
// Attributes:
//  - IsDelete
//  - Column
//  - Value
//  - WriteToWAL
type Mutation struct {
	IsDelete   bool   `thrift:"isDelete,1" db:"isDelete" json:"isDelete"`
	Column     []byte `thrift:"column,2" db:"column" json:"column"`
	Value      []byte `thrift:"value,3" db:"value" json:"value"`
	WriteToWAL bool   `thrift:"writeToWAL,4" db:"writeToWAL" json:"writeToWAL"`
}

func NewMutation() *Mutation {
	return &Mutation{
		WriteToWAL: true,
	}
}

func (p *Mutation) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.BOOL {
				if err := p.ReadField1(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField3(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 4:
			if fieldTypeId == thrift.BOOL {
				if err := p.ReadField4(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *Mutation) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.IsDelete = v
	}
	return nil
}

func (p *Mutation) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Column = v
	}
	return nil
}

func (p *Mutation) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Value = v
	}
	return nil
}

func (p *Mutation) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.WriteToWAL = v
	}
	return nil
}

func (p *Mutation) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Mutation"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField2(oprot); err != nil {
			return err
		}
		if err := p.writeField3(oprot); err != nil {
			return err
		}
		if err := p.writeField4(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *Mutation) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("isDelete", thrift.BOOL, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:isDelete: ", p), err)
	}
	if err := oprot.WriteBool(bool(p.IsDelete)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.isDelete (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:isDelete: ", p), err)
	}
	return err
}

func (p *Mutation) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("column", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:column: ", p), err)
	}
	if err := oprot.WriteBinary(p.Column); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.column (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:column: ", p), err)
	}
	return err
}

func (p *Mutation) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("value", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:value: ", p), err)
	}
	if err := oprot.WriteBinary(p.Value); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.value (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:value: ", p), err)
	}
	return err
}

func (p *Mutation) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("writeToWAL", thrift.BOOL, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:writeToWAL: ", p), err)
	}
	if err := oprot.WriteBool(bool(p.WriteToWAL)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.writeToWAL (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:writeToWAL: ", p), err)
	}
	return err
}

func (p *Mutation) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Mutation(%+v)", *p)
}

// A BatchMutation object is used to apply a number of Mutations to a single row.
//
// Attributes:
//  - Row
//  - Mutations
type BatchMutation struct {
	Row       []byte      `thrift:"row,1" db:"row" json:"row"`
	Mutations []*Mutation `thrift:"mutations,2" db:"mutations" json:"mutations"`
}

func NewBatchMutation() *BatchMutation {
	return &BatchMutation{}
}

func (p *BatchMutation) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.LIST {
				if err := p.ReadField2(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *BatchMutation) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Row = v
	}
	return nil
}

func (p *BatchMutation) ReadField2(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*Mutation, 0, size)
	p.Mutations = tSlice
	for i := 0; i < size; i++ {
		_elem0 := &Mutation{
			WriteToWAL: true,
		}
		if err := _elem0.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem0), err)
		}
		p.Mutations = append(p.Mutations, _elem0)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *BatchMutation) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("BatchMutation"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField2(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *BatchMutation) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("row", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:row: ", p), err)
	}
	if err := oprot.WriteBinary(p.Row); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.row (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:row: ", p), err)
	}
	return err
}

func (p *BatchMutation) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("mutations", thrift.LIST, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:mutations: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Mutations)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Mutations {
		if err := v.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:mutations: ", p), err)
	}
	return err
}

func (p *BatchMutation) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("BatchMutation(%+v)", *p)
}

// For increments that are not incrementColumnValue
// equivalents.
//
// Attributes:
//  - Table
//  - Row
//  - Column
//  - Ammount
type TIncrement struct {
	Table   []byte `thrift:"table,1" db:"table" json:"table"`
	Row     []byte `thrift:"row,2" db:"row" json:"row"`
	Column  []byte `thrift:"column,3" db:"column" json:"column"`
	Ammount int64  `thrift:"ammount,4" db:"ammount" json:"ammount"`
}

func NewTIncrement() *TIncrement {
	return &TIncrement{}
}

func (p *TIncrement) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField3(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 4:
			if fieldTypeId == thrift.I64 {
				if err := p.ReadField4(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TIncrement) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Table = v
	}
	return nil
}

func (p *TIncrement) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Row = v
	}
	return nil
}

func (p *TIncrement) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Column = v
	}
	return nil
}

func (p *TIncrement) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.Ammount = v
	}
	return nil
}

func (p *TIncrement) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TIncrement"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField2(oprot); err != nil {
			return err
		}
		if err := p.writeField3(oprot); err != nil {
			return err
		}
		if err := p.writeField4(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TIncrement) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("table", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:table: ", p), err)
	}
	if err := oprot.WriteBinary(p.Table); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.table (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:table: ", p), err)
	}
	return err
}

func (p *TIncrement) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("row", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:row: ", p), err)
	}
	if err := oprot.WriteBinary(p.Row); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.row (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:row: ", p), err)
	}
	return err
}

func (p *TIncrement) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("column", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:column: ", p), err)
	}
	if err := oprot.WriteBinary(p.Column); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.column (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:column: ", p), err)
	}
	return err
}

func (p *TIncrement) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("ammount", thrift.I64, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:ammount: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.Ammount)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.ammount (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:ammount: ", p), err)
	}
	return err
}

func (p *TIncrement) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TIncrement(%+v)", *p)
}

// Holds column name and the cell.
//
// Attributes:
//  - ColumnName
//  - Cell
type TColumn struct {
	ColumnName []byte `thrift:"columnName,1" db:"columnName" json:"columnName"`
	Cell       *TCell `thrift:"cell,2" db:"cell" json:"cell"`
}

func NewTColumn() *TColumn {
	return &TColumn{}
}

func (p *TColumn) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField2(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TColumn) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.ColumnName = v
	}
	return nil
}

func (p *TColumn) ReadField2(iprot thrift.TProtocol) error {
	p.Cell = &TCell{}
	if err := p.Cell.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Cell), err)
	}
	return nil
}

func (p *TColumn) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TColumn"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField2(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TColumn) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("columnName", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:columnName: ", p), err)
	}
	if err := oprot.WriteBinary(p.ColumnName); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.columnName (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:columnName: ", p), err)
	}
	return err
}

func (p *TColumn) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("cell", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:cell: ", p), err)
	}
	if err := p.Cell.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Cell), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:cell: ", p), err)
	}
	return err
}

func (p *TColumn) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TColumn(%+v)", *p)
}

// Holds row name and then a map of columns to cells.
//
// Attributes:
//  - Row
//  - Columns
//  - SortedColumns
type TRowResult_ struct {
	Row           []byte            `thrift:"row,1" db:"row" json:"row"`
	Columns       map[string]*TCell `thrift:"columns,2" db:"columns" json:"columns,omitempty"`
	SortedColumns []*TColumn        `thrift:"sortedColumns,3" db:"sortedColumns" json:"sortedColumns,omitempty"`
}

func NewTRowResult_() *TRowResult_ {
	return &TRowResult_{}
}

func (p *TRowResult_) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.MAP {
				if err := p.ReadField2(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.LIST {
				if err := p.ReadField3(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TRowResult_) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Row = v
	}
	return nil
}

func (p *TRowResult_) ReadField2(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return thrift.PrependError("error reading map begin: ", err)
	}
	tMap := make(map[string]*TCell, size)
	p.Columns = tMap
	for i := 0; i < size; i++ {
		var _key1 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_key1 = v
		}
		_val2 := &TCell{}
		if err := _val2.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _val2), err)
		}
		p.Columns[_key1] = _val2
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return thrift.PrependError("error reading map end: ", err)
	}
	return nil
}

func (p *TRowResult_) ReadField3(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*TColumn, 0, size)
	p.SortedColumns = tSlice
	for i := 0; i < size; i++ {
		_elem3 := &TColumn{}
		if err := _elem3.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem3), err)
		}
		p.SortedColumns = append(p.SortedColumns, _elem3)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *TRowResult_) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TRowResult"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField2(oprot); err != nil {
			return err
		}
		if err := p.writeField3(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TRowResult_) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("row", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:row: ", p), err)
	}
	if err := oprot.WriteBinary(p.Row); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.row (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:row: ", p), err)
	}
	return err
}

func (p *TRowResult_) writeField2(oprot thrift.TProtocol) (err error) {
	if p.Columns != nil {
		if err := oprot.WriteFieldBegin("columns", thrift.MAP, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:columns: ", p), err)
		}
		if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRUCT, len(p.Columns)); err != nil {
			return thrift.PrependError("error writing map begin: ", err)
		}
		for k, v := range p.Columns {
			if err := oprot.WriteString(string(k)); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
			}
			if err := v.Write(oprot); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
			}
		}
		if err := oprot.WriteMapEnd(); err != nil {
			return thrift.PrependError("error writing map end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:columns: ", p), err)
		}
	}
	return err
}

func (p *TRowResult_) writeField3(oprot thrift.TProtocol) (err error) {
	if p.SortedColumns != nil {
		if err := oprot.WriteFieldBegin("sortedColumns", thrift.LIST, 3); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:sortedColumns: ", p), err)
		}
		if err := oprot.WriteListBegin(thrift.STRUCT, len(p.SortedColumns)); err != nil {
			return thrift.PrependError("error writing list begin: ", err)
		}
		for _, v := range p.SortedColumns {
			if err := v.Write(oprot); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return thrift.PrependError("error writing list end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 3:sortedColumns: ", p), err)
		}
	}
	return err
}

func (p *TRowResult_) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TRowResult_(%+v)", *p)
}

// A Scan object is used to specify scanner parameters when opening a scanner.
//
// Attributes:
//  - StartRow
//  - StopRow
//  - Timestamp
//  - Columns
//  - Caching
//  - FilterString
//  - BatchSize
//  - SortColumns
//  - Reversed
//  - CacheBlocks
type TScan struct {
	StartRow     []byte   `thrift:"startRow,1" db:"startRow" json:"startRow,omitempty"`
	StopRow      []byte   `thrift:"stopRow,2" db:"stopRow" json:"stopRow,omitempty"`
	Timestamp    *int64   `thrift:"timestamp,3" db:"timestamp" json:"timestamp,omitempty"`
	Columns      [][]byte `thrift:"columns,4" db:"columns" json:"columns,omitempty"`
	Caching      *int32   `thrift:"caching,5" db:"caching" json:"caching,omitempty"`
	FilterString []byte   `thrift:"filterString,6" db:"filterString" json:"filterString,omitempty"`
	BatchSize    *int32   `thrift:"batchSize,7" db:"batchSize" json:"batchSize,omitempty"`
	SortColumns  *bool    `thrift:"sortColumns,8" db:"sortColumns" json:"sortColumns,omitempty"`
	Reversed     *bool    `thrift:"reversed,9" db:"reversed" json:"reversed,omitempty"`
	CacheBlocks  *bool    `thrift:"cacheBlocks,10" db:"cacheBlocks" json:"cacheBlocks,omitempty"`
}

func NewTScan() *TScan {
	return &TScan{}
}

func (p *TScan) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.I64 {
				if err := p.ReadField3(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 4:
			if fieldTypeId == thrift.LIST {
				if err := p.ReadField4(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 5:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField5(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 6:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField6(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 7:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField7(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 8:
			if fieldTypeId == thrift.BOOL {
				if err := p.ReadField8(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 9:
			if fieldTypeId == thrift.BOOL {
				if err := p.ReadField9(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 10:
			if fieldTypeId == thrift.BOOL {
				if err := p.ReadField10(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TScan) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.StartRow = v
	}
	return nil
}

func (p *TScan) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.StopRow = v
	}
	return nil
}

func (p *TScan) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Timestamp = &v
	}
	return nil
}

func (p *TScan) ReadField4(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([][]byte, 0, size)
	p.Columns = tSlice
	for i := 0; i < size; i++ {
		var _elem4 []byte
		if v, err := iprot.ReadBinary(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_elem4 = v
		}
		p.Columns = append(p.Columns, _elem4)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *TScan) ReadField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.Caching = &v
	}
	return nil
}

func (p *TScan) ReadField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 6: ", err)
	} else {
		p.FilterString = v
	}
	return nil
}

func (p *TScan) ReadField7(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 7: ", err)
	} else {
		p.BatchSize = &v
	}
	return nil
}

func (p *TScan) ReadField8(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return thrift.PrependError("error reading field 8: ", err)
	} else {
		p.SortColumns = &v
	}
	return nil
}

func (p *TScan) ReadField9(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return thrift.PrependError("error reading field 9: ", err)
	} else {
		p.Reversed = &v
	}
	return nil
}

func (p *TScan) ReadField10(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return thrift.PrependError("error reading field 10: ", err)
	} else {
		p.CacheBlocks = &v
	}
	return nil
}

func (p *TScan) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TScan"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField2(oprot); err != nil {
			return err
		}
		if err := p.writeField3(oprot); err != nil {
			return err
		}
		if err := p.writeField4(oprot); err != nil {
			return err
		}
		if err := p.writeField5(oprot); err != nil {
			return err
		}
		if err := p.writeField6(oprot); err != nil {
			return err
		}
		if err := p.writeField7(oprot); err != nil {
			return err
		}
		if err := p.writeField8(oprot); err != nil {
			return err
		}
		if err := p.writeField9(oprot); err != nil {
			return err
		}
		if err := p.writeField10(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TScan) writeField1(oprot thrift.TProtocol) (err error) {
	if p.StartRow != nil {
		if err := oprot.WriteFieldBegin("startRow", thrift.STRING, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:startRow: ", p), err)
		}
		if err := oprot.WriteBinary(p.StartRow); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.startRow (1) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:startRow: ", p), err)
		}
	}
	return err
}

func (p *TScan) writeField2(oprot thrift.TProtocol) (err error) {
	if p.StopRow != nil {
		if err := oprot.WriteFieldBegin("stopRow", thrift.STRING, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:stopRow: ", p), err)
		}
		if err := oprot.WriteBinary(p.StopRow); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.stopRow (2) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:stopRow: ", p), err)
		}
	}
	return err
}

func (p *TScan) writeField3(oprot thrift.TProtocol) (err error) {
	if p.Timestamp != nil {
		if err := oprot.WriteFieldBegin("timestamp", thrift.I64, 3); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:timestamp: ", p), err)
		}
		if err := oprot.WriteI64(int64(*p.Timestamp)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.timestamp (3) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 3:timestamp: ", p), err)
		}
	}
	return err
}

func (p *TScan) writeField4(oprot thrift.TProtocol) (err error) {
	if p.Columns != nil {
		if err := oprot.WriteFieldBegin("columns", thrift.LIST, 4); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:columns: ", p), err)
		}
		if err := oprot.WriteListBegin(thrift.STRING, len(p.Columns)); err != nil {
			return thrift.PrependError("error writing list begin: ", err)
		}
		for _, v := range p.Columns {
			if err := oprot.WriteBinary(v); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return thrift.PrependError("error writing list end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 4:columns: ", p), err)
		}
	}
	return err
}

func (p *TScan) writeField5(oprot thrift.TProtocol) (err error) {
	if p.Caching != nil {
		if err := oprot.WriteFieldBegin("caching", thrift.I32, 5); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:caching: ", p), err)
		}
		if err := oprot.WriteI32(int32(*p.Caching)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.caching (5) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 5:caching: ", p), err)
		}
	}
	return err
}

func (p *TScan) writeField6(oprot thrift.TProtocol) (err error) {
	if p.FilterString != nil {
		if err := oprot.WriteFieldBegin("filterString", thrift.STRING, 6); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:filterString: ", p), err)
		}
		if err := oprot.WriteBinary(p.FilterString); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.filterString (6) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 6:filterString: ", p), err)
		}
	}
	return err
}

func (p *TScan) writeField7(oprot thrift.TProtocol) (err error) {
	if p.BatchSize != nil {
		if err := oprot.WriteFieldBegin("batchSize", thrift.I32, 7); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:batchSize: ", p), err)
		}
		if err := oprot.WriteI32(int32(*p.BatchSize)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.batchSize (7) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 7:batchSize: ", p), err)
		}
	}
	return err
}

func (p *TScan) writeField8(oprot thrift.TProtocol) (err error) {
	if p.SortColumns != nil {
		if err := oprot.WriteFieldBegin("sortColumns", thrift.BOOL, 8); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 8:sortColumns: ", p), err)
		}
		if err := oprot.WriteBool(bool(*p.SortColumns)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.sortColumns (8) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 8:sortColumns: ", p), err)
		}
	}
	return err
}

func (p *TScan) writeField9(oprot thrift.TProtocol) (err error) {
	if p.Reversed != nil {
		if err := oprot.WriteFieldBegin("reversed", thrift.BOOL, 9); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 9:reversed: ", p), err)
		}
		if err := oprot.WriteBool(bool(*p.Reversed)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.reversed (9) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 9:reversed: ", p), err)
		}
	}
	return err
}

func (p *TScan) writeField10(oprot thrift.TProtocol) (err error) {
	if p.CacheBlocks != nil {
		if err := oprot.WriteFieldBegin("cacheBlocks", thrift.BOOL, 10); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 10:cacheBlocks: ", p), err)
		}
		if err := oprot.WriteBool(bool(*p.CacheBlocks)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.cacheBlocks (10) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 10:cacheBlocks: ", p), err)
		}
	}
	return err
}

func (p *TScan) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TScan(%+v)", *p)
}

// An Append object is used to specify the parameters for performing the append operation.
//
// Attributes:
//  - Table
//  - Row
//  - Columns
//  - Values
type TAppend struct {
	Table   []byte   `thrift:"table,1" db:"table" json:"table"`
	Row     []byte   `thrift:"row,2" db:"row" json:"row"`
	Columns [][]byte `thrift:"columns,3" db:"columns" json:"columns"`
	Values  [][]byte `thrift:"values,4" db:"values" json:"values"`
}

func NewTAppend() *TAppend {
	return &TAppend{}
}

func (p *TAppend) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.LIST {
				if err := p.ReadField3(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 4:
			if fieldTypeId == thrift.LIST {
				if err := p.ReadField4(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TAppend) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Table = v
	}
	return nil
}

func (p *TAppend) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Row = v
	}
	return nil
}

func (p *TAppend) ReadField3(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([][]byte, 0, size)
	p.Columns = tSlice
	for i := 0; i < size; i++ {
		var _elem5 []byte
		if v, err := iprot.ReadBinary(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_elem5 = v
		}
		p.Columns = append(p.Columns, _elem5)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *TAppend) ReadField4(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([][]byte, 0, size)
	p.Values = tSlice
	for i := 0; i < size; i++ {
		var _elem6 []byte
		if v, err := iprot.ReadBinary(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_elem6 = v
		}
		p.Values = append(p.Values, _elem6)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *TAppend) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TAppend"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField2(oprot); err != nil {
			return err
		}
		if err := p.writeField3(oprot); err != nil {
			return err
		}
		if err := p.writeField4(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TAppend) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("table", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:table: ", p), err)
	}
	if err := oprot.WriteBinary(p.Table); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.table (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:table: ", p), err)
	}
	return err
}

func (p *TAppend) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("row", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:row: ", p), err)
	}
	if err := oprot.WriteBinary(p.Row); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.row (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:row: ", p), err)
	}
	return err
}

func (p *TAppend) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("columns", thrift.LIST, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:columns: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRING, len(p.Columns)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Columns {
		if err := oprot.WriteBinary(v); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:columns: ", p), err)
	}
	return err
}

func (p *TAppend) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("values", thrift.LIST, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:values: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRING, len(p.Values)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Values {
		if err := oprot.WriteBinary(v); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:values: ", p), err)
	}
	return err
}

func (p *TAppend) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TAppend(%+v)", *p)
}
