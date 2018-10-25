package hbase

import "context"

type Hbase interface {
	// Brings a table on-line (enables it)
	//
	// Parameters:
	//  - TableName: name of the table
	EnableTable(ctx context.Context, tableName []byte) (err error)
	// Disables a table (takes it off-line) If it is being served, the master
	// will tell the servers to stop serving it.
	//
	// Parameters:
	//  - TableName: name of the table
	DisableTable(ctx context.Context, tableName []byte) (err error)
	// @return true if table is on-line
	//
	// Parameters:
	//  - TableName: name of the table to check
	IsTableEnabled(ctx context.Context, tableName []byte) (r bool, err error)
	// Parameters:
	//  - TableNameOrRegionName
	Compact(ctx context.Context, tableNameOrRegionName []byte) (err error)
	// Parameters:
	//  - TableNameOrRegionName
	MajorCompact(ctx context.Context, tableNameOrRegionName []byte) (err error)
	// List all the userspace tables.
	//
	// @return returns a list of names
	GetTableNames(ctx context.Context) (r [][]byte, err error)
	// List all the column families assoicated with a table.
	//
	// @return list of column family descriptors
	//
	// Parameters:
	//  - TableName: table name
	GetColumnDescriptors(ctx context.Context, tableName []byte) (r map[string]*ColumnDescriptor, err error)
	// List the regions associated with a table.
	//
	// @return list of region descriptors
	//
	// Parameters:
	//  - TableName: table name
	GetTableRegions(ctx context.Context, tableName []byte) (r []*TRegionInfo, err error)
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
	CreateTable(ctx context.Context, tableName []byte, columnFamilies []*ColumnDescriptor) (err error)
	// Deletes a table
	//
	// @throws IOError if table doesn't exist on server or there was some other
	// problem
	//
	// Parameters:
	//  - TableName: name of table to delete
	DeleteTable(ctx context.Context, tableName []byte) (err error)
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
	Get(ctx context.Context, tableName []byte, row []byte, column []byte, attributes map[string][]byte) (r []*TCell, err error)
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
	GetVer(ctx context.Context, tableName []byte, row []byte, column []byte, numVersions int32, attributes map[string][]byte) (r []*TCell, err error)
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
	GetVerTs(ctx context.Context, tableName []byte, row []byte, column []byte, timestamp int64, numVersions int32, attributes map[string][]byte) (r []*TCell, err error)
	// Get all the data for the specified table and row at the latest
	// timestamp. Returns an empty list if the row does not exist.
	//
	// @return TRowResult containing the row and map of columns to TCells
	//
	// Parameters:
	//  - TableName: name of table
	//  - Row: row key
	//  - Attributes: Get attributes
	GetRow(ctx context.Context, tableName []byte, row []byte, attributes map[string][]byte) (r []*TRowResult_, err error)
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
	GetRowWithColumns(ctx context.Context, tableName []byte, row []byte, columns [][]byte, attributes map[string][]byte) (r []*TRowResult_, err error)
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
	GetRowTs(ctx context.Context, tableName []byte, row []byte, timestamp int64, attributes map[string][]byte) (r []*TRowResult_, err error)
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
	GetRowWithColumnsTs(ctx context.Context, tableName []byte, row []byte, columns [][]byte, timestamp int64, attributes map[string][]byte) (r []*TRowResult_, err error)
	// Get all the data for the specified table and rows at the latest
	// timestamp. Returns an empty list if no rows exist.
	//
	// @return TRowResult containing the rows and map of columns to TCells
	//
	// Parameters:
	//  - TableName: name of table
	//  - Rows: row keys
	//  - Attributes: Get attributes
	GetRows(ctx context.Context, tableName []byte, rows [][]byte, attributes map[string][]byte) (r []*TRowResult_, err error)
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
	GetRowsWithColumns(ctx context.Context, tableName []byte, rows [][]byte, columns [][]byte, attributes map[string][]byte) (r []*TRowResult_, err error)
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
	GetRowsTs(ctx context.Context, tableName []byte, rows [][]byte, timestamp int64, attributes map[string][]byte) (r []*TRowResult_, err error)
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
	GetRowsWithColumnsTs(ctx context.Context, tableName []byte, rows [][]byte, columns [][]byte, timestamp int64, attributes map[string][]byte) (r []*TRowResult_, err error)
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
	MutateRow(ctx context.Context, tableName []byte, row []byte, mutations []*Mutation, attributes map[string][]byte) (err error)
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
	MutateRowTs(ctx context.Context, tableName []byte, row []byte, mutations []*Mutation, timestamp int64, attributes map[string][]byte) (err error)
	// Apply a series of batches (each a series of mutations on a single row)
	// in a single transaction.  If an exception is thrown, then the
	// transaction is aborted.  Default current timestamp is used, and
	// all entries will have an identical timestamp.
	//
	// Parameters:
	//  - TableName: name of table
	//  - RowBatches: list of row batches
	//  - Attributes: Mutation attributes
	MutateRows(ctx context.Context, tableName []byte, rowBatches []*BatchMutation, attributes map[string][]byte) (err error)
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
	MutateRowsTs(ctx context.Context, tableName []byte, rowBatches []*BatchMutation, timestamp int64, attributes map[string][]byte) (err error)
	// Atomically increment the column value specified.  Returns the next value post increment.
	//
	// Parameters:
	//  - TableName: name of table
	//  - Row: row to increment
	//  - Column: name of column
	//  - Value: amount to increment by
	AtomicIncrement(ctx context.Context, tableName []byte, row []byte, column []byte, value int64) (r int64, err error)
	// Delete all cells that match the passed row and column.
	//
	// Parameters:
	//  - TableName: name of table
	//  - Row: Row to update
	//  - Column: name of column whose value is to be deleted
	//  - Attributes: Delete attributes
	DeleteAll(ctx context.Context, tableName []byte, row []byte, column []byte, attributes map[string][]byte) (err error)
	// Delete all cells that match the passed row and column and whose
	// timestamp is equal-to or older than the passed timestamp.
	//
	// Parameters:
	//  - TableName: name of table
	//  - Row: Row to update
	//  - Column: name of column whose value is to be deleted
	//  - Timestamp: timestamp
	//  - Attributes: Delete attributes
	DeleteAllTs(ctx context.Context, tableName []byte, row []byte, column []byte, timestamp int64, attributes map[string][]byte) (err error)
	// Completely delete the row's cells.
	//
	// Parameters:
	//  - TableName: name of table
	//  - Row: key of the row to be completely deleted.
	//  - Attributes: Delete attributes
	DeleteAllRow(ctx context.Context, tableName []byte, row []byte, attributes map[string][]byte) (err error)
	// Increment a cell by the ammount.
	// Increments can be applied async if hbase.regionserver.thrift.coalesceIncrement is set to true.
	// False is the default.  Turn to true if you need the extra performance and can accept some
	// data loss if a thrift server dies with increments still in the queue.
	//
	// Parameters:
	//  - Increment: The single increment to apply
	Increment(ctx context.Context, increment *TIncrement) (err error)
	// Parameters:
	//  - Increments: The list of increments
	IncrementRows(ctx context.Context, increments []*TIncrement) (err error)
	// Completely delete the row's cells marked with a timestamp
	// equal-to or older than the passed timestamp.
	//
	// Parameters:
	//  - TableName: name of table
	//  - Row: key of the row to be completely deleted.
	//  - Timestamp: timestamp
	//  - Attributes: Delete attributes
	DeleteAllRowTs(ctx context.Context, tableName []byte, row []byte, timestamp int64, attributes map[string][]byte) (err error)
	// Get a scanner on the current table, using the Scan instance
	// for the scan parameters.
	//
	// Parameters:
	//  - TableName: name of table
	//  - Scan: Scan instance
	//  - Attributes: Scan attributes
	ScannerOpenWithScan(ctx context.Context, tableName []byte, scan *TScan, attributes map[string][]byte) (r ScannerID, err error)
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
	ScannerOpen(ctx context.Context, tableName []byte, startRow []byte, columns [][]byte, attributes map[string][]byte) (r ScannerID, err error)
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
	ScannerOpenWithStop(ctx context.Context, tableName []byte, startRow []byte, stopRow []byte, columns [][]byte, attributes map[string][]byte) (r ScannerID, err error)
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
	ScannerOpenWithPrefix(ctx context.Context, tableName []byte, startAndPrefix []byte, columns [][]byte, attributes map[string][]byte) (r ScannerID, err error)
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
	ScannerOpenTs(ctx context.Context, tableName []byte, startRow []byte, columns [][]byte, timestamp int64, attributes map[string][]byte) (r ScannerID, err error)
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
	ScannerOpenWithStopTs(ctx context.Context, tableName []byte, startRow []byte, stopRow []byte, columns [][]byte, timestamp int64, attributes map[string][]byte) (r ScannerID, err error)
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
	ScannerGet(ctx context.Context, id ScannerID) (r []*TRowResult_, err error)
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
	ScannerGetList(ctx context.Context, id ScannerID, nbRows int32) (r []*TRowResult_, err error)
	// Closes the server-state associated with an open scanner.
	//
	// @throws IllegalArgument if ScannerID is invalid
	//
	// Parameters:
	//  - ID: id of a scanner returned by scannerOpen
	ScannerClose(ctx context.Context, id ScannerID) (err error)
	// Get the regininfo for the specified row. It scans
	// the metatable to find region's start and end keys.
	//
	// @return value for specified row/column
	//
	// Parameters:
	//  - Row: row key
	GetRegionInfo(ctx context.Context, row []byte) (r *TRegionInfo, err error)
	// Appends values to one or more columns within a single row.
	//
	// @return values of columns after the append operation.
	//
	// Parameters:
	//  - Append: The single append operation to apply
	Append(ctx context.Context, append *TAppend) (r []*TCell, err error)
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
	CheckAndPut(ctx context.Context, tableName []byte, row []byte, column []byte, value []byte, mput *Mutation, attributes map[string][]byte) (r bool, err error)
}
