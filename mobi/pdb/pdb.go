package pdb

import (
	"bytes"
	"encoding/binary"
	"io"
	"io/ioutil"
)

// Endian of values encoded in Palm databases.
var Endian = binary.BigEndian

// Database represents an in-memory Palm database.
type Database struct {
	Name    string
	Records []Record
}

// NewDatabase creates an empty Palm database with name.
func NewDatabase(name string) Database {
	return Database{
		Name:    trimZeroes(name),
		Records: []Record{},
	}
}

// AddRecord adds a generic record to the Palm database.
func (d *Database) AddRecord(r Record) {
	d.Records = append(d.Records, r)
}

// Write writes out the binary representation of the Palm database to w.
func (d Database) Write(w io.Writer) error {
	rlen := len(d.Records)
	palmDBHeader := NewPalmDBHeader(d.Name, uint16(rlen), uint32(rlen))
	err := binary.Write(w, Endian, palmDBHeader)
	if err != nil {
		return err
	}

	// Write record headers
	offset := PalmDBHeaderLength + RecordHeaderLength*len(d.Records) + 2
	for i, rec := range d.Records {
		h := RecordHeader{
			Offset:    uint32(offset),
			Attribute: 1,
			Skip:      0,
			UniqueID:  uint16(i),
		}
		offset += rec.Length()
		err := binary.Write(w, Endian, h)
		if err != nil {
			return err
		}
	}

	// Write 2-byte padding
	pad := make([]byte, 2)
	_, err = w.Write(pad)
	if err != nil {
		return err
	}

	// Write records
	for _, rec := range d.Records {
		err := rec.Write(w)
		if err != nil {
			return err
		}
	}

	// Success
	return nil
}

// ReadDatabase reads an uninterpreted Palm database from r.
func ReadDatabase(r io.Reader) (*Database, error) {
	data, err := ioutil.ReadAll(r)
	b := bytes.NewReader(data)
	if err != nil {
		return nil, err
	}

	palmDBHeader := PalmDBHeader{}
	binary.Read(b, Endian, &palmDBHeader)

	offsets := make([]RecordHeader, palmDBHeader.NumRecords)
	binary.Read(b, Endian, &offsets)

	records := make([]Record, 0)
	for i := 1; i < len(offsets); i++ {
		curr := offsets[i].Offset
		prev := offsets[i-1].Offset
		records = append(records, RawRecord(data[prev:curr]))
	}
	last := offsets[len(offsets)-1].Offset
	records = append(records, RawRecord(data[last:]))

	name := trimZeroes(string(palmDBHeader.Name[:]))

	return &Database{
		Name:    name,
		Records: records,
	}, nil
}
