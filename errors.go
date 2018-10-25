package hbase

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// An IOError exception signals that an error occurred communicating
// to the Hbase master or an Hbase region server.  Also used to return
// more general Hbase error conditions.
//
// Attributes:
//  - Message
type IOError struct {
	Message string `thrift:"message,1" db:"message" json:"message"`
}

func NewIOError() *IOError {
	return &IOError{}
}

func (p *IOError) Read(iprot thrift.TProtocol) error {
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

func (p *IOError) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Message = v
	}
	return nil
}

func (p *IOError) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("IOError"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
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

func (p *IOError) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("message", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:message: ", p), err)
	}
	if err := oprot.WriteString(string(p.Message)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.message (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:message: ", p), err)
	}
	return err
}

func (p *IOError) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("IOError(%+v)", *p)
}

func (p *IOError) Error() string {
	return p.String()
}

// An IllegalArgument exception indicates an illegal or invalid
// argument was passed into a procedure.
//
// Attributes:
//  - Message
type IllegalArgument struct {
	Message string `thrift:"message,1" db:"message" json:"message"`
}

func NewIllegalArgument() *IllegalArgument {
	return &IllegalArgument{}
}

func (p *IllegalArgument) Read(iprot thrift.TProtocol) error {
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

func (p *IllegalArgument) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Message = v
	}
	return nil
}

func (p *IllegalArgument) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("IllegalArgument"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
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

func (p *IllegalArgument) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("message", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:message: ", p), err)
	}
	if err := oprot.WriteString(string(p.Message)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.message (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:message: ", p), err)
	}
	return err
}

func (p *IllegalArgument) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("IllegalArgument(%+v)", *p)
}

func (p *IllegalArgument) Error() string {
	return p.String()
}

// An AlreadyExists exceptions signals that a table with the specified
// name already exists
//
// Attributes:
//  - Message
type AlreadyExists struct {
	Message string `thrift:"message,1" db:"message" json:"message"`
}

func NewAlreadyExists() *AlreadyExists {
	return &AlreadyExists{}
}

func (p *AlreadyExists) Read(iprot thrift.TProtocol) error {
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

func (p *AlreadyExists) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Message = v
	}
	return nil
}

func (p *AlreadyExists) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("AlreadyExists"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
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

func (p *AlreadyExists) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("message", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:message: ", p), err)
	}
	if err := oprot.WriteString(string(p.Message)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.message (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:message: ", p), err)
	}
	return err
}

func (p *AlreadyExists) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AlreadyExists(%+v)", *p)
}

func (p *AlreadyExists) Error() string {
	return p.String()
}
