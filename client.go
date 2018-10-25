package hbase

import (
	"context"
	"git.apache.org/thrift.git/lib/go/thrift"
)

type Client struct {
	c thrift.TClient
}

// Deprecated: Use NewHbase instead
func NewClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *Client {
	return &Client{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

// Deprecated: Use NewHbase instead
func NewClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *Client {
	return &Client{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewClient(c thrift.TClient) *Client {
	return &Client{
		c: c,
	}
}

// Brings a table on-line (enables it)
//
// Parameters:
//  - TableName: name of the table
func (p *Client) EnableTable(ctx context.Context, tableName []byte) (err error) {
	var _args7 EnableTableArgs
	_args7.TableName = tableName
	var _result8 EnableTableResult
	if err = p.c.Call(ctx, "enableTable", &_args7, &_result8); err != nil {
		return
	}
	switch {
	case _result8.Io != nil:
		return _result8.Io
	}
	return nil
}

// Disables a table (takes it off-line) If it is being served, the master
// will tell the servers to stop serving it.
//
// Parameters:
//  - TableName: name of the table
func (p *Client) DisableTable(ctx context.Context, tableName []byte) (err error) {
	var _args9 DisableTableArgs
	_args9.TableName = tableName
	var _result10 DisableTableResult
	if err = p.c.Call(ctx, "disableTable", &_args9, &_result10); err != nil {
		return
	}
	switch {
	case _result10.Io != nil:
		return _result10.Io
	}

	return nil
}

// @return true if table is on-line
//
// Parameters:
//  - TableName: name of the table to check
func (p *Client) IsTableEnabled(ctx context.Context, tableName []byte) (r bool, err error) {
	var _args11 IsTableEnabledArgs
	_args11.TableName = tableName
	var _result12 IsTableEnabledResult
	if err = p.c.Call(ctx, "isTableEnabled", &_args11, &_result12); err != nil {
		return
	}
	switch {
	case _result12.Io != nil:
		return r, _result12.Io
	}

	return _result12.GetSuccess(), nil
}

// Parameters:
//  - TableNameOrRegionName
func (p *Client) Compact(ctx context.Context, tableNameOrRegionName []byte) (err error) {
	var _args13 CompactArgs
	_args13.TableNameOrRegionName = tableNameOrRegionName
	var _result14 CompactResult
	if err = p.c.Call(ctx, "compact", &_args13, &_result14); err != nil {
		return
	}
	switch {
	case _result14.Io != nil:
		return _result14.Io
	}

	return nil
}

// Parameters:
//  - TableNameOrRegionName
func (p *Client) MajorCompact(ctx context.Context, tableNameOrRegionName []byte) (err error) {
	var _args15 MajorCompactArgs
	_args15.TableNameOrRegionName = tableNameOrRegionName
	var _result16 MajorCompactResult
	if err = p.c.Call(ctx, "majorCompact", &_args15, &_result16); err != nil {
		return
	}
	switch {
	case _result16.Io != nil:
		return _result16.Io
	}

	return nil
}

// List all the userspace tables.
//
// @return returns a list of names
func (p *Client) GetTableNames(ctx context.Context) (r [][]byte, err error) {
	var _args17 GetTableNamesArgs
	var _result18 GetTableNamesResult
	if err = p.c.Call(ctx, "getTableNames", &_args17, &_result18); err != nil {
		return
	}
	switch {
	case _result18.Io != nil:
		return r, _result18.Io
	}

	return _result18.GetSuccess(), nil
}

// List all the column families assoicated with a table.
//
// @return list of column family descriptors
//
// Parameters:
//  - TableName: table name
func (p *Client) GetColumnDescriptors(ctx context.Context, tableName []byte) (r map[string]*ColumnDescriptor, err error) {
	var _args19 GetColumnDescriptorsArgs
	_args19.TableName = tableName
	var _result20 GetColumnDescriptorsResult
	if err = p.c.Call(ctx, "getColumnDescriptors", &_args19, &_result20); err != nil {
		return
	}
	switch {
	case _result20.Io != nil:
		return r, _result20.Io
	}

	return _result20.GetSuccess(), nil
}

// List the regions associated with a table.
//
// @return list of region descriptors
//
// Parameters:
//  - TableName: table name
func (p *Client) GetTableRegions(ctx context.Context, tableName []byte) (r []*TRegionInfo, err error) {
	var _args21 GetTableRegionsArgs
	_args21.TableName = tableName
	var _result22 GetTableRegionsResult
	if err = p.c.Call(ctx, "getTableRegions", &_args21, &_result22); err != nil {
		return
	}
	switch {
	case _result22.Io != nil:
		return r, _result22.Io
	}

	return _result22.GetSuccess(), nil
}

// Create a table with the specified column families.  The name
// field for each ColumnDescriptor must be set and must end in a
// colon (:). All other fields are optional and will get default
// values if not explicitly specified.
//
// @throws IllegalArgument if an input parameter is invalid
//
// @throws AlreadyExists if the table name already exists
//
// Parameters:
//  - TableName: name of table to create
//  - ColumnFamilies: list of column family descriptors
func (p *Client) CreateTable(ctx context.Context, tableName []byte, columnFamilies []*ColumnDescriptor) (err error) {
	var _args23 CreateTableArgs
	_args23.TableName = tableName
	_args23.ColumnFamilies = columnFamilies
	var _result24 CreateTableResult
	if err = p.c.Call(ctx, "createTable", &_args23, &_result24); err != nil {
		return
	}
	switch {
	case _result24.Io != nil:
		return _result24.Io
	case _result24.Ia != nil:
		return _result24.Ia
	case _result24.Exist != nil:
		return _result24.Exist
	}

	return nil
}

// Deletes a table
//
// @throws IOError if table doesn't exist on server or there was some other
// problem
//
// Parameters:
//  - TableName: name of table to delete
func (p *Client) DeleteTable(ctx context.Context, tableName []byte) (err error) {
	var _args25 DeleteTableArgs
	_args25.TableName = tableName
	var _result26 DeleteTableResult
	if err = p.c.Call(ctx, "deleteTable", &_args25, &_result26); err != nil {
		return
	}
	switch {
	case _result26.Io != nil:
		return _result26.Io
	}

	return nil
}

// Get a single TCell for the specified table, row, and column at the
// latest timestamp. Returns an empty list if no such value exists.
//
// @return value for specified row/column
//
// Parameters:
//  - TableName: name of table
//  - Row: row key
//  - Column: column name
//  - Attributes: Get attributes
func (p *Client) Get(ctx context.Context, tableName []byte, row []byte, column []byte, attributes map[string][]byte) (r []*TCell, err error) {
	var _args27 GetArgs
	_args27.TableName = tableName
	_args27.Row = row
	_args27.Column = column
	_args27.Attributes = attributes
	var _result28 GetResult
	if err = p.c.Call(ctx, "get", &_args27, &_result28); err != nil {
		return
	}
	switch {
	case _result28.Io != nil:
		return r, _result28.Io
	}

	return _result28.GetSuccess(), nil
}

// Get the specified number of versions for the specified table,
// row, and column.
//
// @return list of cells for specified row/column
//
// Parameters:
//  - TableName: name of table
//  - Row: row key
//  - Column: column name
//  - NumVersions: number of versions to retrieve
//  - Attributes: Get attributes
func (p *Client) GetVer(ctx context.Context, tableName []byte, row []byte, column []byte, numVersions int32, attributes map[string][]byte) (r []*TCell, err error) {
	var _args29 GetVerArgs
	_args29.TableName = tableName
	_args29.Row = row
	_args29.Column = column
	_args29.NumVersions = numVersions
	_args29.Attributes = attributes
	var _result30 GetVerResult
	if err = p.c.Call(ctx, "getVer", &_args29, &_result30); err != nil {
		return
	}
	switch {
	case _result30.Io != nil:
		return r, _result30.Io
	}

	return _result30.GetSuccess(), nil
}

// Get the specified number of versions for the specified table,
// row, and column.  Only versions less than or equal to the specified
// timestamp will be returned.
//
// @return list of cells for specified row/column
//
// Parameters:
//  - TableName: name of table
//  - Row: row key
//  - Column: column name
//  - Timestamp: timestamp
//  - NumVersions: number of versions to retrieve
//  - Attributes: Get attributes
func (p *Client) GetVerTs(ctx context.Context, tableName []byte, row []byte, column []byte, timestamp int64, numVersions int32, attributes map[string][]byte) (r []*TCell, err error) {
	var _args31 GetVerTsArgs
	_args31.TableName = tableName
	_args31.Row = row
	_args31.Column = column
	_args31.Timestamp = timestamp
	_args31.NumVersions = numVersions
	_args31.Attributes = attributes
	var _result32 GetVerTsResult
	if err = p.c.Call(ctx, "getVerTs", &_args31, &_result32); err != nil {
		return
	}
	switch {
	case _result32.Io != nil:
		return r, _result32.Io
	}

	return _result32.GetSuccess(), nil
}

// Get all the data for the specified table and row at the latest
// timestamp. Returns an empty list if the row does not exist.
//
// @return TRowResult containing the row and map of columns to TCells
//
// Parameters:
//  - TableName: name of table
//  - Row: row key
//  - Attributes: Get attributes
func (p *Client) GetRow(ctx context.Context, tableName []byte, row []byte, attributes map[string][]byte) (r []*TRowResult_, err error) {
	var _args33 GetRowArgs
	_args33.TableName = tableName
	_args33.Row = row
	_args33.Attributes = attributes
	var _result34 GetRowResult
	if err = p.c.Call(ctx, "getRow", &_args33, &_result34); err != nil {
		return
	}
	switch {
	case _result34.Io != nil:
		return r, _result34.Io
	}

	return _result34.GetSuccess(), nil
}

// Get the specified columns for the specified table and row at the latest
// timestamp. Returns an empty list if the row does not exist.
//
// @return TRowResult containing the row and map of columns to TCells
//
// Parameters:
//  - TableName: name of table
//  - Row: row key
//  - Columns: List of columns to return, null for all columns
//  - Attributes: Get attributes
func (p *Client) GetRowWithColumns(ctx context.Context, tableName []byte, row []byte, columns [][]byte, attributes map[string][]byte) (r []*TRowResult_, err error) {
	var _args35 GetRowWithColumnsArgs
	_args35.TableName = tableName
	_args35.Row = row
	_args35.Columns = columns
	_args35.Attributes = attributes
	var _result36 GetRowWithColumnsResult
	if err = p.c.Call(ctx, "getRowWithColumns", &_args35, &_result36); err != nil {
		return
	}
	switch {
	case _result36.Io != nil:
		return r, _result36.Io
	}
	return _result36.GetSuccess(), nil
}

// Get all the data for the specified table and row at the specified
// timestamp. Returns an empty list if the row does not exist.
//
// @return TRowResult containing the row and map of columns to TCells
//
// Parameters:
//  - TableName: name of the table
//  - Row: row key
//  - Timestamp: timestamp
//  - Attributes: Get attributes
func (p *Client) GetRowTs(ctx context.Context, tableName []byte, row []byte, timestamp int64, attributes map[string][]byte) (r []*TRowResult_, err error) {
	var _args37 GetRowTsArgs
	_args37.TableName = tableName
	_args37.Row = row
	_args37.Timestamp = timestamp
	_args37.Attributes = attributes
	var _result38 GetRowTsResult
	if err = p.c.Call(ctx, "getRowTs", &_args37, &_result38); err != nil {
		return
	}
	switch {
	case _result38.Io != nil:
		return r, _result38.Io
	}

	return _result38.GetSuccess(), nil
}

// Get the specified columns for the specified table and row at the specified
// timestamp. Returns an empty list if the row does not exist.
//
// @return TRowResult containing the row and map of columns to TCells
//
// Parameters:
//  - TableName: name of table
//  - Row: row key
//  - Columns: List of columns to return, null for all columns
//  - Timestamp
//  - Attributes: Get attributes
func (p *Client) GetRowWithColumnsTs(ctx context.Context, tableName []byte, row []byte, columns [][]byte, timestamp int64, attributes map[string][]byte) (r []*TRowResult_, err error) {
	var _args39 GetRowWithColumnsTsArgs
	_args39.TableName = tableName
	_args39.Row = row
	_args39.Columns = columns
	_args39.Timestamp = timestamp
	_args39.Attributes = attributes
	var _result40 GetRowWithColumnsTsResult
	if err = p.c.Call(ctx, "getRowWithColumnsTs", &_args39, &_result40); err != nil {
		return
	}
	switch {
	case _result40.Io != nil:
		return r, _result40.Io
	}

	return _result40.GetSuccess(), nil
}

// Get all the data for the specified table and rows at the latest
// timestamp. Returns an empty list if no rows exist.
//
// @return TRowResult containing the rows and map of columns to TCells
//
// Parameters:
//  - TableName: name of table
//  - Rows: row keys
//  - Attributes: Get attributes
func (p *Client) GetRows(ctx context.Context, tableName []byte, rows [][]byte, attributes map[string][]byte) (r []*TRowResult_, err error) {
	var _args41 GetRowsArgs
	_args41.TableName = tableName
	_args41.Rows = rows
	_args41.Attributes = attributes
	var _result42 GetRowsResult
	if err = p.c.Call(ctx, "getRows", &_args41, &_result42); err != nil {
		return
	}
	switch {
	case _result42.Io != nil:
		return r, _result42.Io
	}

	return _result42.GetSuccess(), nil
}

// Get the specified columns for the specified table and rows at the latest
// timestamp. Returns an empty list if no rows exist.
//
// @return TRowResult containing the rows and map of columns to TCells
//
// Parameters:
//  - TableName: name of table
//  - Rows: row keys
//  - Columns: List of columns to return, null for all columns
//  - Attributes: Get attributes
func (p *Client) GetRowsWithColumns(ctx context.Context, tableName []byte, rows [][]byte, columns [][]byte, attributes map[string][]byte) (r []*TRowResult_, err error) {
	var _args43 GetRowsWithColumnsArgs
	_args43.TableName = tableName
	_args43.Rows = rows
	_args43.Columns = columns
	_args43.Attributes = attributes
	var _result44 GetRowsWithColumnsResult
	if err = p.c.Call(ctx, "getRowsWithColumns", &_args43, &_result44); err != nil {
		return
	}
	switch {
	case _result44.Io != nil:
		return r, _result44.Io
	}

	return _result44.GetSuccess(), nil
}

// Get all the data for the specified table and rows at the specified
// timestamp. Returns an empty list if no rows exist.
//
// @return TRowResult containing the rows and map of columns to TCells
//
// Parameters:
//  - TableName: name of the table
//  - Rows: row keys
//  - Timestamp: timestamp
//  - Attributes: Get attributes
func (p *Client) GetRowsTs(ctx context.Context, tableName []byte, rows [][]byte, timestamp int64, attributes map[string][]byte) (r []*TRowResult_, err error) {
	var _args45 GetRowsTsArgs
	_args45.TableName = tableName
	_args45.Rows = rows
	_args45.Timestamp = timestamp
	_args45.Attributes = attributes
	var _result46 GetRowsTsResult
	if err = p.c.Call(ctx, "getRowsTs", &_args45, &_result46); err != nil {
		return
	}
	switch {
	case _result46.Io != nil:
		return r, _result46.Io
	}

	return _result46.GetSuccess(), nil
}

// Get the specified columns for the specified table and rows at the specified
// timestamp. Returns an empty list if no rows exist.
//
// @return TRowResult containing the rows and map of columns to TCells
//
// Parameters:
//  - TableName: name of table
//  - Rows: row keys
//  - Columns: List of columns to return, null for all columns
//  - Timestamp
//  - Attributes: Get attributes
func (p *Client) GetRowsWithColumnsTs(ctx context.Context, tableName []byte, rows [][]byte, columns [][]byte, timestamp int64, attributes map[string][]byte) (r []*TRowResult_, err error) {
	var _args47 GetRowsWithColumnsTsArgs
	_args47.TableName = tableName
	_args47.Rows = rows
	_args47.Columns = columns
	_args47.Timestamp = timestamp
	_args47.Attributes = attributes
	var _result48 GetRowsWithColumnsTsResult
	if err = p.c.Call(ctx, "getRowsWithColumnsTs", &_args47, &_result48); err != nil {
		return
	}
	switch {
	case _result48.Io != nil:
		return r, _result48.Io
	}

	return _result48.GetSuccess(), nil
}

// Apply a series of mutations (updates/deletes) to a row in a
// single transaction.  If an exception is thrown, then the
// transaction is aborted.  Default current timestamp is used, and
// all entries will have an identical timestamp.
//
// Parameters:
//  - TableName: name of table
//  - Row: row key
//  - Mutations: list of mutation commands
//  - Attributes: Mutation attributes
func (p *Client) MutateRow(ctx context.Context, tableName []byte, row []byte, mutations []*Mutation, attributes map[string][]byte) (err error) {
	var _args49 MutateRowArgs
	_args49.TableName = tableName
	_args49.Row = row
	_args49.Mutations = mutations
	_args49.Attributes = attributes
	var _result50 MutateRowResult
	if err = p.c.Call(ctx, "mutateRow", &_args49, &_result50); err != nil {
		return
	}
	switch {
	case _result50.Io != nil:
		return _result50.Io
	case _result50.Ia != nil:
		return _result50.Ia
	}

	return nil
}

// Apply a series of mutations (updates/deletes) to a row in a
// single transaction.  If an exception is thrown, then the
// transaction is aborted.  The specified timestamp is used, and
// all entries will have an identical timestamp.
//
// Parameters:
//  - TableName: name of table
//  - Row: row key
//  - Mutations: list of mutation commands
//  - Timestamp: timestamp
//  - Attributes: Mutation attributes
func (p *Client) MutateRowTs(ctx context.Context, tableName []byte, row []byte, mutations []*Mutation, timestamp int64, attributes map[string][]byte) (err error) {
	var _args51 MutateRowTsArgs
	_args51.TableName = tableName
	_args51.Row = row
	_args51.Mutations = mutations
	_args51.Timestamp = timestamp
	_args51.Attributes = attributes
	var _result52 MutateRowTsResult
	if err = p.c.Call(ctx, "mutateRowTs", &_args51, &_result52); err != nil {
		return
	}
	switch {
	case _result52.Io != nil:
		return _result52.Io
	case _result52.Ia != nil:
		return _result52.Ia
	}

	return nil
}

// Apply a series of batches (each a series of mutations on a single row)
// in a single transaction.  If an exception is thrown, then the
// transaction is aborted.  Default current timestamp is used, and
// all entries will have an identical timestamp.
//
// Parameters:
//  - TableName: name of table
//  - RowBatches: list of row batches
//  - Attributes: Mutation attributes
func (p *Client) MutateRows(ctx context.Context, tableName []byte, rowBatches []*BatchMutation, attributes map[string][]byte) (err error) {
	var _args53 MutateRowsArgs
	_args53.TableName = tableName
	_args53.RowBatches = rowBatches
	_args53.Attributes = attributes
	var _result54 MutateRowsResult
	if err = p.c.Call(ctx, "mutateRows", &_args53, &_result54); err != nil {
		return
	}
	switch {
	case _result54.Io != nil:
		return _result54.Io
	case _result54.Ia != nil:
		return _result54.Ia
	}

	return nil
}

// Apply a series of batches (each a series of mutations on a single row)
// in a single transaction.  If an exception is thrown, then the
// transaction is aborted.  The specified timestamp is used, and
// all entries will have an identical timestamp.
//
// Parameters:
//  - TableName: name of table
//  - RowBatches: list of row batches
//  - Timestamp: timestamp
//  - Attributes: Mutation attributes
func (p *Client) MutateRowsTs(ctx context.Context, tableName []byte, rowBatches []*BatchMutation, timestamp int64, attributes map[string][]byte) (err error) {
	var _args55 MutateRowsTsArgs
	_args55.TableName = tableName
	_args55.RowBatches = rowBatches
	_args55.Timestamp = timestamp
	_args55.Attributes = attributes
	var _result56 MutateRowsTsResult
	if err = p.c.Call(ctx, "mutateRowsTs", &_args55, &_result56); err != nil {
		return
	}
	switch {
	case _result56.Io != nil:
		return _result56.Io
	case _result56.Ia != nil:
		return _result56.Ia
	}

	return nil
}

// Atomically increment the column value specified.  Returns the next value post increment.
//
// Parameters:
//  - TableName: name of table
//  - Row: row to increment
//  - Column: name of column
//  - Value: amount to increment by
func (p *Client) AtomicIncrement(ctx context.Context, tableName []byte, row []byte, column []byte, value int64) (r int64, err error) {
	var _args57 AtomicIncrementArgs
	_args57.TableName = tableName
	_args57.Row = row
	_args57.Column = column
	_args57.Value = value
	var _result58 AtomicIncrementResult
	if err = p.c.Call(ctx, "atomicIncrement", &_args57, &_result58); err != nil {
		return
	}
	switch {
	case _result58.Io != nil:
		return r, _result58.Io
	case _result58.Ia != nil:
		return r, _result58.Ia
	}

	return _result58.GetSuccess(), nil
}

// Delete all cells that match the passed row and column.
//
// Parameters:
//  - TableName: name of table
//  - Row: Row to update
//  - Column: name of column whose value is to be deleted
//  - Attributes: Delete attributes
func (p *Client) DeleteAll(ctx context.Context, tableName []byte, row []byte, column []byte, attributes map[string][]byte) (err error) {
	var _args59 DeleteAllArgs
	_args59.TableName = tableName
	_args59.Row = row
	_args59.Column = column
	_args59.Attributes = attributes
	var _result60 DeleteAllResult
	if err = p.c.Call(ctx, "deleteAll", &_args59, &_result60); err != nil {
		return
	}
	switch {
	case _result60.Io != nil:
		return _result60.Io
	}

	return nil
}

// Delete all cells that match the passed row and column and whose
// timestamp is equal-to or older than the passed timestamp.
//
// Parameters:
//  - TableName: name of table
//  - Row: Row to update
//  - Column: name of column whose value is to be deleted
//  - Timestamp: timestamp
//  - Attributes: Delete attributes
func (p *Client) DeleteAllTs(ctx context.Context, tableName []byte, row []byte, column []byte, timestamp int64, attributes map[string][]byte) (err error) {
	var _args61 DeleteAllTsArgs
	_args61.TableName = tableName
	_args61.Row = row
	_args61.Column = column
	_args61.Timestamp = timestamp
	_args61.Attributes = attributes
	var _result62 DeleteAllTsResult
	if err = p.c.Call(ctx, "deleteAllTs", &_args61, &_result62); err != nil {
		return
	}
	switch {
	case _result62.Io != nil:
		return _result62.Io
	}

	return nil
}

// Completely delete the row's cells.
//
// Parameters:
//  - TableName: name of table
//  - Row: key of the row to be completely deleted.
//  - Attributes: Delete attributes
func (p *Client) DeleteAllRow(ctx context.Context, tableName []byte, row []byte, attributes map[string][]byte) (err error) {
	var _args63 DeleteAllRowArgs
	_args63.TableName = tableName
	_args63.Row = row
	_args63.Attributes = attributes
	var _result64 DeleteAllRowResult
	if err = p.c.Call(ctx, "deleteAllRow", &_args63, &_result64); err != nil {
		return
	}
	switch {
	case _result64.Io != nil:
		return _result64.Io
	}

	return nil
}

// Increment a cell by the ammount.
// Increments can be applied async if hbase.regionserver.thrift.coalesceIncrement is set to true.
// False is the default.  Turn to true if you need the extra performance and can accept some
// data loss if a thrift server dies with increments still in the queue.
//
// Parameters:
//  - Increment: The single increment to apply
func (p *Client) Increment(ctx context.Context, increment *TIncrement) (err error) {
	var _args65 IncrementArgs
	_args65.Increment = increment
	var _result66 IncrementResult
	if err = p.c.Call(ctx, "increment", &_args65, &_result66); err != nil {
		return
	}
	switch {
	case _result66.Io != nil:
		return _result66.Io
	}

	return nil
}

// Parameters:
//  - Increments: The list of increments
func (p *Client) IncrementRows(ctx context.Context, increments []*TIncrement) (err error) {
	var _args67 IncrementRowsArgs
	_args67.Increments = increments
	var _result68 IncrementRowsResult
	if err = p.c.Call(ctx, "incrementRows", &_args67, &_result68); err != nil {
		return
	}
	switch {
	case _result68.Io != nil:
		return _result68.Io
	}

	return nil
}

// Completely delete the row's cells marked with a timestamp
// equal-to or older than the passed timestamp.
//
// Parameters:
//  - TableName: name of table
//  - Row: key of the row to be completely deleted.
//  - Timestamp: timestamp
//  - Attributes: Delete attributes
func (p *Client) DeleteAllRowTs(ctx context.Context, tableName []byte, row []byte, timestamp int64, attributes map[string][]byte) (err error) {
	var _args69 DeleteAllRowTsArgs
	_args69.TableName = tableName
	_args69.Row = row
	_args69.Timestamp = timestamp
	_args69.Attributes = attributes
	var _result70 DeleteAllRowTsResult
	if err = p.c.Call(ctx, "deleteAllRowTs", &_args69, &_result70); err != nil {
		return
	}
	switch {
	case _result70.Io != nil:
		return _result70.Io
	}

	return nil
}

// Get a scanner on the current table, using the Scan instance
// for the scan parameters.
//
// Parameters:
//  - TableName: name of table
//  - Scan: Scan instance
//  - Attributes: Scan attributes
func (p *Client) ScannerOpenWithScan(ctx context.Context, tableName []byte, scan *TScan, attributes map[string][]byte) (r ScannerID, err error) {
	var _args71 ScannerOpenWithScanArgs
	_args71.TableName = tableName
	_args71.Scan = scan
	_args71.Attributes = attributes
	var _result72 ScannerOpenWithScanResult
	if err = p.c.Call(ctx, "scannerOpenWithScan", &_args71, &_result72); err != nil {
		return
	}
	switch {
	case _result72.Io != nil:
		return r, _result72.Io
	}

	return _result72.GetSuccess(), nil
}

// Get a scanner on the current table starting at the specified row and
// ending at the last row in the table.  Return the specified columns.
//
// @return scanner id to be used with other scanner procedures
//
// Parameters:
//  - TableName: name of table
//  - StartRow: Starting row in table to scan.
// Send "" (empty string) to start at the first row.
//  - Columns: columns to scan. If column name is a column family, all
// columns of the specified column family are returned. It's also possible
// to pass a regex in the column qualifier.
//  - Attributes: Scan attributes
func (p *Client) ScannerOpen(ctx context.Context, tableName []byte, startRow []byte, columns [][]byte, attributes map[string][]byte) (r ScannerID, err error) {
	var _args73 ScannerOpenArgs
	_args73.TableName = tableName
	_args73.StartRow = startRow
	_args73.Columns = columns
	_args73.Attributes = attributes
	var _result74 ScannerOpenResult
	if err = p.c.Call(ctx, "scannerOpen", &_args73, &_result74); err != nil {
		return
	}
	switch {
	case _result74.Io != nil:
		return r, _result74.Io
	}

	return _result74.GetSuccess(), nil
}

// Get a scanner on the current table starting and stopping at the
// specified rows.  ending at the last row in the table.  Return the
// specified columns.
//
// @return scanner id to be used with other scanner procedures
//
// Parameters:
//  - TableName: name of table
//  - StartRow: Starting row in table to scan.
// Send "" (empty string) to start at the first row.
//  - StopRow: row to stop scanning on. This row is *not* included in the
// scanner's results
//  - Columns: columns to scan. If column name is a column family, all
// columns of the specified column family are returned. It's also possible
// to pass a regex in the column qualifier.
//  - Attributes: Scan attributes
func (p *Client) ScannerOpenWithStop(ctx context.Context, tableName []byte, startRow []byte, stopRow []byte, columns [][]byte, attributes map[string][]byte) (r ScannerID, err error) {
	var _args75 ScannerOpenWithStopArgs
	_args75.TableName = tableName
	_args75.StartRow = startRow
	_args75.StopRow = stopRow
	_args75.Columns = columns
	_args75.Attributes = attributes
	var _result76 ScannerOpenWithStopResult
	if err = p.c.Call(ctx, "scannerOpenWithStop", &_args75, &_result76); err != nil {
		return
	}
	switch {
	case _result76.Io != nil:
		return r, _result76.Io
	}

	return _result76.GetSuccess(), nil
}

// Open a scanner for a given prefix.  That is all rows will have the specified
// prefix. No other rows will be returned.
//
// @return scanner id to use with other scanner calls
//
// Parameters:
//  - TableName: name of table
//  - StartAndPrefix: the prefix (and thus start row) of the keys you want
//  - Columns: the columns you want returned
//  - Attributes: Scan attributes
func (p *Client) ScannerOpenWithPrefix(ctx context.Context, tableName []byte, startAndPrefix []byte, columns [][]byte, attributes map[string][]byte) (r ScannerID, err error) {
	var _args77 ScannerOpenWithPrefixArgs
	_args77.TableName = tableName
	_args77.StartAndPrefix = startAndPrefix
	_args77.Columns = columns
	_args77.Attributes = attributes
	var _result78 ScannerOpenWithPrefixResult
	if err = p.c.Call(ctx, "scannerOpenWithPrefix", &_args77, &_result78); err != nil {
		return
	}
	switch {
	case _result78.Io != nil:
		return r, _result78.Io
	}

	return _result78.GetSuccess(), nil
}

// Get a scanner on the current table starting at the specified row and
// ending at the last row in the table.  Return the specified columns.
// Only values with the specified timestamp are returned.
//
// @return scanner id to be used with other scanner procedures
//
// Parameters:
//  - TableName: name of table
//  - StartRow: Starting row in table to scan.
// Send "" (empty string) to start at the first row.
//  - Columns: columns to scan. If column name is a column family, all
// columns of the specified column family are returned. It's also possible
// to pass a regex in the column qualifier.
//  - Timestamp: timestamp
//  - Attributes: Scan attributes
func (p *Client) ScannerOpenTs(ctx context.Context, tableName []byte, startRow []byte, columns [][]byte, timestamp int64, attributes map[string][]byte) (r ScannerID, err error) {
	var _args79 ScannerOpenTsArgs
	_args79.TableName = tableName
	_args79.StartRow = startRow
	_args79.Columns = columns
	_args79.Timestamp = timestamp
	_args79.Attributes = attributes
	var _result80 ScannerOpenTsResult
	if err = p.c.Call(ctx, "scannerOpenTs", &_args79, &_result80); err != nil {
		return
	}
	switch {
	case _result80.Io != nil:
		return r, _result80.Io
	}

	return _result80.GetSuccess(), nil
}

// Get a scanner on the current table starting and stopping at the
// specified rows.  ending at the last row in the table.  Return the
// specified columns.  Only values with the specified timestamp are
// returned.
//
// @return scanner id to be used with other scanner procedures
//
// Parameters:
//  - TableName: name of table
//  - StartRow: Starting row in table to scan.
// Send "" (empty string) to start at the first row.
//  - StopRow: row to stop scanning on. This row is *not* included in the
// scanner's results
//  - Columns: columns to scan. If column name is a column family, all
// columns of the specified column family are returned. It's also possible
// to pass a regex in the column qualifier.
//  - Timestamp: timestamp
//  - Attributes: Scan attributes
func (p *Client) ScannerOpenWithStopTs(ctx context.Context, tableName []byte, startRow []byte, stopRow []byte, columns [][]byte, timestamp int64, attributes map[string][]byte) (r ScannerID, err error) {
	var _args81 ScannerOpenWithStopTsArgs
	_args81.TableName = tableName
	_args81.StartRow = startRow
	_args81.StopRow = stopRow
	_args81.Columns = columns
	_args81.Timestamp = timestamp
	_args81.Attributes = attributes
	var _result82 ScannerOpenWithStopTsResult
	if err = p.c.Call(ctx, "scannerOpenWithStopTs", &_args81, &_result82); err != nil {
		return
	}
	switch {
	case _result82.Io != nil:
		return r, _result82.Io
	}

	return _result82.GetSuccess(), nil
}

// Returns the scanner's current row value and advances to the next
// row in the table.  When there are no more rows in the table, or a key
// greater-than-or-equal-to the scanner's specified stopRow is reached,
// an empty list is returned.
//
// @return a TRowResult containing the current row and a map of the columns to TCells.
//
// @throws IllegalArgument if ScannerID is invalid
//
// @throws NotFound when the scanner reaches the end
//
// Parameters:
//  - ID: id of a scanner returned by scannerOpen
func (p *Client) ScannerGet(ctx context.Context, id ScannerID) (r []*TRowResult_, err error) {
	var _args83 ScannerGetArgs
	_args83.ID = id
	var _result84 ScannerGetResult
	if err = p.c.Call(ctx, "scannerGet", &_args83, &_result84); err != nil {
		return
	}
	switch {
	case _result84.Io != nil:
		return r, _result84.Io
	case _result84.Ia != nil:
		return r, _result84.Ia
	}

	return _result84.GetSuccess(), nil
}

// Returns, starting at the scanner's current row value nbRows worth of
// rows and advances to the next row in the table.  When there are no more
// rows in the table, or a key greater-than-or-equal-to the scanner's
// specified stopRow is reached,  an empty list is returned.
//
// @return a TRowResult containing the current row and a map of the columns to TCells.
//
// @throws IllegalArgument if ScannerID is invalid
//
// @throws NotFound when the scanner reaches the end
//
// Parameters:
//  - ID: id of a scanner returned by scannerOpen
//  - NbRows: number of results to return
func (p *Client) ScannerGetList(ctx context.Context, id ScannerID, nbRows int32) (r []*TRowResult_, err error) {
	var _args85 ScannerGetListArgs
	_args85.ID = id
	_args85.NbRows = nbRows
	var _result86 ScannerGetListResult
	if err = p.c.Call(ctx, "scannerGetList", &_args85, &_result86); err != nil {
		return
	}
	switch {
	case _result86.Io != nil:
		return r, _result86.Io
	case _result86.Ia != nil:
		return r, _result86.Ia
	}

	return _result86.GetSuccess(), nil
}

// Closes the server-state associated with an open scanner.
//
// @throws IllegalArgument if ScannerID is invalid
//
// Parameters:
//  - ID: id of a scanner returned by scannerOpen
func (p *Client) ScannerClose(ctx context.Context, id ScannerID) (err error) {
	var _args87 ScannerCloseArgs
	_args87.ID = id
	var _result88 ScannerCloseResult
	if err = p.c.Call(ctx, "scannerClose", &_args87, &_result88); err != nil {
		return
	}
	switch {
	case _result88.Io != nil:
		return _result88.Io
	case _result88.Ia != nil:
		return _result88.Ia
	}

	return nil
}

// Get the regininfo for the specified row. It scans
// the metatable to find region's start and end keys.
//
// @return value for specified row/column
//
// Parameters:
//  - Row: row key
func (p *Client) GetRegionInfo(ctx context.Context, row []byte) (r *TRegionInfo, err error) {
	var _args89 GetRegionInfoArgs
	_args89.Row = row
	var _result90 GetRegionInfoResult
	if err = p.c.Call(ctx, "getRegionInfo", &_args89, &_result90); err != nil {
		return
	}
	switch {
	case _result90.Io != nil:
		return r, _result90.Io
	}

	return _result90.GetSuccess(), nil
}

// Appends values to one or more columns within a single row.
//
// @return values of columns after the append operation.
//
// Parameters:
//  - Append: The single append operation to apply
func (p *Client) Append(ctx context.Context, append *TAppend) (r []*TCell, err error) {
	var _args91 AppendArgs
	_args91.Append = append
	var _result92 AppendResult
	if err = p.c.Call(ctx, "append", &_args91, &_result92); err != nil {
		return
	}
	switch {
	case _result92.Io != nil:
		return r, _result92.Io
	}

	return _result92.GetSuccess(), nil
}

// Atomically checks if a row/family/qualifier value matches the expected
// value. If it does, it adds the corresponding mutation operation for put.
//
// @return true if the new put was executed, false otherwise
//
// Parameters:
//  - TableName: name of table
//  - Row: row key
//  - Column: column name
//  - Value: the expected value for the column parameter, if not
// provided the check is for the non-existence of the
// column in question
//  - Mput: mutation for the put
//  - Attributes: Mutation attributes
func (p *Client) CheckAndPut(ctx context.Context, tableName []byte, row []byte, column []byte, value []byte, mput *Mutation, attributes map[string][]byte) (r bool, err error) {
	var _args93 CheckAndPutArgs
	_args93.TableName = tableName
	_args93.Row = row
	_args93.Column = column
	_args93.Value = value
	_args93.Mput = mput
	_args93.Attributes = attributes
	var _result94 CheckAndPutResult
	if err = p.c.Call(ctx, "checkAndPut", &_args93, &_result94); err != nil {
		return
	}
	switch {
	case _result94.Io != nil:
		return r, _result94.Io
	case _result94.Ia != nil:
		return r, _result94.Ia
	}

	return _result94.GetSuccess(), nil
}
