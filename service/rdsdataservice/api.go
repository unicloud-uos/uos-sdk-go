// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsdataservice

import (
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/awsutil"
	"gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/request"
)

const opExecuteSql = "ExecuteSql"

// ExecuteSqlRequest generates a "aws/request.Request" representing the
// client's request for the ExecuteSql operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See ExecuteSql for more information on using the ExecuteSql
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the ExecuteSqlRequest method.
//    req, resp := client.ExecuteSqlRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/rds-data-2018-08-01/ExecuteSql
func (c *RDSDataService) ExecuteSqlRequest(input *ExecuteSqlInput) (req *request.Request, output *ExecuteSqlOutput) {
	op := &request.Operation{
		Name:       opExecuteSql,
		HTTPMethod: "POST",
		HTTPPath:   "/ExecuteSql",
	}

	if input == nil {
		input = &ExecuteSqlInput{}
	}

	output = &ExecuteSqlOutput{}
	req = c.newRequest(op, input, output)
	return
}

// ExecuteSql API operation for AWS RDS DataService.
//
// Executes any SQL statement on the target database synchronously
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS RDS DataService's
// API operation ExecuteSql for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeBadRequestException "BadRequestException"
//   Invalid Request exception
//
//   * ErrCodeForbiddenException "ForbiddenException"
//   Access denied exception
//
//   * ErrCodeInternalServerErrorException "InternalServerErrorException"
//   Internal service error
//
//   * ErrCodeServiceUnavailableError "ServiceUnavailableError"
//   Internal service unavailable error
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/rds-data-2018-08-01/ExecuteSql
func (c *RDSDataService) ExecuteSql(input *ExecuteSqlInput) (*ExecuteSqlOutput, error) {
	req, out := c.ExecuteSqlRequest(input)
	return out, req.Send()
}

// ExecuteSqlWithContext is the same as ExecuteSql with the addition of
// the ability to pass a context and additional request options.
//
// See ExecuteSql for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSDataService) ExecuteSqlWithContext(ctx aws.Context, input *ExecuteSqlInput, opts ...request.Option) (*ExecuteSqlOutput, error) {
	req, out := c.ExecuteSqlRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

// Column Metadata
type ColumnMetadata struct {
	_ struct{} `type:"structure"`

	// Homogenous array base SQL type from java.sql.Types.
	ArrayBaseColumnType *int64 `locationName:"arrayBaseColumnType" type:"integer"`

	// Whether the designated column is automatically numbered
	IsAutoIncrement *bool `locationName:"isAutoIncrement" type:"boolean"`

	// Whether values in the designated column's case matters
	IsCaseSensitive *bool `locationName:"isCaseSensitive" type:"boolean"`

	// Whether values in the designated column is a cash value
	IsCurrency *bool `locationName:"isCurrency" type:"boolean"`

	// Whether values in the designated column are signed numbers
	IsSigned *bool `locationName:"isSigned" type:"boolean"`

	// Usually specified by the SQL AS. If not specified, return column name.
	Label *string `locationName:"label" type:"string"`

	// Name of the column.
	Name *string `locationName:"name" type:"string"`

	// Indicates the nullability of values in the designated column. One of columnNoNulls
	// (0), columnNullable (1), columnNullableUnknown (2)
	Nullable *int64 `locationName:"nullable" type:"integer"`

	// Get the designated column's specified column size.For numeric data, this
	// is the maximum precision. For character data, this is the length in characters.
	// For datetime datatypes, this is the length in characters of the String representation
	// (assuming the maximum allowed precision of the fractional seconds component).
	// For binary data, this is the length in bytes. For the ROWID datatype, this
	// is the length in bytes. 0 is returned for data types where the column size
	// is not applicable.
	Precision *int64 `locationName:"precision" type:"integer"`

	// Designated column's number of digits to right of the decimal point. 0 is
	// returned for data types where the scale is not applicable.
	Scale *int64 `locationName:"scale" type:"integer"`

	// Designated column's table's schema
	SchemaName *string `locationName:"schemaName" type:"string"`

	// Designated column's table name
	TableName *string `locationName:"tableName" type:"string"`

	// SQL type from java.sql.Types.
	Type *int64 `locationName:"type" type:"integer"`

	// Database-specific type name.
	TypeName *string `locationName:"typeName" type:"string"`
}

// String returns the string representation
func (s ColumnMetadata) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ColumnMetadata) GoString() string {
	return s.String()
}

// SetArrayBaseColumnType sets the ArrayBaseColumnType field's value.
func (s *ColumnMetadata) SetArrayBaseColumnType(v int64) *ColumnMetadata {
	s.ArrayBaseColumnType = &v
	return s
}

// SetIsAutoIncrement sets the IsAutoIncrement field's value.
func (s *ColumnMetadata) SetIsAutoIncrement(v bool) *ColumnMetadata {
	s.IsAutoIncrement = &v
	return s
}

// SetIsCaseSensitive sets the IsCaseSensitive field's value.
func (s *ColumnMetadata) SetIsCaseSensitive(v bool) *ColumnMetadata {
	s.IsCaseSensitive = &v
	return s
}

// SetIsCurrency sets the IsCurrency field's value.
func (s *ColumnMetadata) SetIsCurrency(v bool) *ColumnMetadata {
	s.IsCurrency = &v
	return s
}

// SetIsSigned sets the IsSigned field's value.
func (s *ColumnMetadata) SetIsSigned(v bool) *ColumnMetadata {
	s.IsSigned = &v
	return s
}

// SetLabel sets the Label field's value.
func (s *ColumnMetadata) SetLabel(v string) *ColumnMetadata {
	s.Label = &v
	return s
}

// SetName sets the Name field's value.
func (s *ColumnMetadata) SetName(v string) *ColumnMetadata {
	s.Name = &v
	return s
}

// SetNullable sets the Nullable field's value.
func (s *ColumnMetadata) SetNullable(v int64) *ColumnMetadata {
	s.Nullable = &v
	return s
}

// SetPrecision sets the Precision field's value.
func (s *ColumnMetadata) SetPrecision(v int64) *ColumnMetadata {
	s.Precision = &v
	return s
}

// SetScale sets the Scale field's value.
func (s *ColumnMetadata) SetScale(v int64) *ColumnMetadata {
	s.Scale = &v
	return s
}

// SetSchemaName sets the SchemaName field's value.
func (s *ColumnMetadata) SetSchemaName(v string) *ColumnMetadata {
	s.SchemaName = &v
	return s
}

// SetTableName sets the TableName field's value.
func (s *ColumnMetadata) SetTableName(v string) *ColumnMetadata {
	s.TableName = &v
	return s
}

// SetType sets the Type field's value.
func (s *ColumnMetadata) SetType(v int64) *ColumnMetadata {
	s.Type = &v
	return s
}

// SetTypeName sets the TypeName field's value.
func (s *ColumnMetadata) SetTypeName(v string) *ColumnMetadata {
	s.TypeName = &v
	return s
}

// Execute SQL Request
type ExecuteSqlInput struct {
	_ struct{} `type:"structure"`

	// ARN of the db credentials in AWS Secret Store or the friendly secret name
	//
	// AwsSecretStoreArn is a required field
	AwsSecretStoreArn *string `locationName:"awsSecretStoreArn" type:"string" required:"true"`

	// Target DB name
	Database *string `locationName:"database" type:"string"`

	// ARN of the target db cluster or instance
	//
	// DbClusterOrInstanceArn is a required field
	DbClusterOrInstanceArn *string `locationName:"dbClusterOrInstanceArn" type:"string" required:"true"`

	// Target Schema name
	Schema *string `locationName:"schema" type:"string"`

	// SQL statement(s) to be executed. Statements can be chained by using semicolons
	//
	// SqlStatements is a required field
	SqlStatements *string `locationName:"sqlStatements" type:"string" required:"true"`
}

// String returns the string representation
func (s ExecuteSqlInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ExecuteSqlInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ExecuteSqlInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ExecuteSqlInput"}
	if s.AwsSecretStoreArn == nil {
		invalidParams.Add(request.NewErrParamRequired("AwsSecretStoreArn"))
	}
	if s.DbClusterOrInstanceArn == nil {
		invalidParams.Add(request.NewErrParamRequired("DbClusterOrInstanceArn"))
	}
	if s.SqlStatements == nil {
		invalidParams.Add(request.NewErrParamRequired("SqlStatements"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAwsSecretStoreArn sets the AwsSecretStoreArn field's value.
func (s *ExecuteSqlInput) SetAwsSecretStoreArn(v string) *ExecuteSqlInput {
	s.AwsSecretStoreArn = &v
	return s
}

// SetDatabase sets the Database field's value.
func (s *ExecuteSqlInput) SetDatabase(v string) *ExecuteSqlInput {
	s.Database = &v
	return s
}

// SetDbClusterOrInstanceArn sets the DbClusterOrInstanceArn field's value.
func (s *ExecuteSqlInput) SetDbClusterOrInstanceArn(v string) *ExecuteSqlInput {
	s.DbClusterOrInstanceArn = &v
	return s
}

// SetSchema sets the Schema field's value.
func (s *ExecuteSqlInput) SetSchema(v string) *ExecuteSqlInput {
	s.Schema = &v
	return s
}

// SetSqlStatements sets the SqlStatements field's value.
func (s *ExecuteSqlInput) SetSqlStatements(v string) *ExecuteSqlInput {
	s.SqlStatements = &v
	return s
}

// Execute SQL response
type ExecuteSqlOutput struct {
	_ struct{} `type:"structure"`

	// Results returned by executing the sql statement(s)
	//
	// SqlStatementResults is a required field
	SqlStatementResults []*SqlStatementResult `locationName:"sqlStatementResults" type:"list" required:"true"`
}

// String returns the string representation
func (s ExecuteSqlOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ExecuteSqlOutput) GoString() string {
	return s.String()
}

// SetSqlStatementResults sets the SqlStatementResults field's value.
func (s *ExecuteSqlOutput) SetSqlStatementResults(v []*SqlStatementResult) *ExecuteSqlOutput {
	s.SqlStatementResults = v
	return s
}

// Row or Record
type Record struct {
	_ struct{} `type:"structure"`

	// Record
	Values []*Value `locationName:"values" type:"list"`
}

// String returns the string representation
func (s Record) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Record) GoString() string {
	return s.String()
}

// SetValues sets the Values field's value.
func (s *Record) SetValues(v []*Value) *Record {
	s.Values = v
	return s
}

// Result Frame
type ResultFrame struct {
	_ struct{} `type:"structure"`

	// ResultSet Metadata.
	Records []*Record `locationName:"records" type:"list"`

	// ResultSet Metadata.
	ResultSetMetadata *ResultSetMetadata `locationName:"resultSetMetadata" type:"structure"`
}

// String returns the string representation
func (s ResultFrame) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ResultFrame) GoString() string {
	return s.String()
}

// SetRecords sets the Records field's value.
func (s *ResultFrame) SetRecords(v []*Record) *ResultFrame {
	s.Records = v
	return s
}

// SetResultSetMetadata sets the ResultSetMetadata field's value.
func (s *ResultFrame) SetResultSetMetadata(v *ResultSetMetadata) *ResultFrame {
	s.ResultSetMetadata = v
	return s
}

// List of columns and their types.
type ResultSetMetadata struct {
	_ struct{} `type:"structure"`

	// Number of columns
	ColumnCount *int64 `locationName:"columnCount" type:"long"`

	// List of columns and their types
	ColumnMetadata []*ColumnMetadata `locationName:"columnMetadata" type:"list"`
}

// String returns the string representation
func (s ResultSetMetadata) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ResultSetMetadata) GoString() string {
	return s.String()
}

// SetColumnCount sets the ColumnCount field's value.
func (s *ResultSetMetadata) SetColumnCount(v int64) *ResultSetMetadata {
	s.ColumnCount = &v
	return s
}

// SetColumnMetadata sets the ColumnMetadata field's value.
func (s *ResultSetMetadata) SetColumnMetadata(v []*ColumnMetadata) *ResultSetMetadata {
	s.ColumnMetadata = v
	return s
}

// SQL statement execution result
type SqlStatementResult struct {
	_ struct{} `type:"structure"`

	// Number of rows updated.
	NumberOfRecordsUpdated *int64 `locationName:"numberOfRecordsUpdated" type:"long"`

	// ResultFrame returned by executing the sql statement
	ResultFrame *ResultFrame `locationName:"resultFrame" type:"structure"`
}

// String returns the string representation
func (s SqlStatementResult) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s SqlStatementResult) GoString() string {
	return s.String()
}

// SetNumberOfRecordsUpdated sets the NumberOfRecordsUpdated field's value.
func (s *SqlStatementResult) SetNumberOfRecordsUpdated(v int64) *SqlStatementResult {
	s.NumberOfRecordsUpdated = &v
	return s
}

// SetResultFrame sets the ResultFrame field's value.
func (s *SqlStatementResult) SetResultFrame(v *ResultFrame) *SqlStatementResult {
	s.ResultFrame = v
	return s
}

// User Defined Type
type StructValue struct {
	_ struct{} `type:"structure"`

	// Struct or UDT
	Attributes []*Value `locationName:"attributes" type:"list"`
}

// String returns the string representation
func (s StructValue) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s StructValue) GoString() string {
	return s.String()
}

// SetAttributes sets the Attributes field's value.
func (s *StructValue) SetAttributes(v []*Value) *StructValue {
	s.Attributes = v
	return s
}

// Column value
type Value struct {
	_ struct{} `type:"structure"`

	// Arbitrarily nested arrays
	ArrayValues []*Value `locationName:"arrayValues" type:"list"`

	// Long value
	BigIntValue *int64 `locationName:"bigIntValue" type:"long"`

	// Bit value
	BitValue *bool `locationName:"bitValue" type:"boolean"`

	// Blob value
	//
	// BlobValue is automatically base64 encoded/decoded by the SDK.
	BlobValue []byte `locationName:"blobValue" type:"blob"`

	// Double value
	DoubleValue *float64 `locationName:"doubleValue" type:"double"`

	// Integer value
	IntValue *int64 `locationName:"intValue" type:"integer"`

	// Is column null
	IsNull *bool `locationName:"isNull" type:"boolean"`

	// Float value
	RealValue *float64 `locationName:"realValue" type:"float"`

	// String value
	StringValue *string `locationName:"stringValue" type:"string"`

	// Struct or UDT
	StructValue *StructValue `locationName:"structValue" type:"structure"`
}

// String returns the string representation
func (s Value) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Value) GoString() string {
	return s.String()
}

// SetArrayValues sets the ArrayValues field's value.
func (s *Value) SetArrayValues(v []*Value) *Value {
	s.ArrayValues = v
	return s
}

// SetBigIntValue sets the BigIntValue field's value.
func (s *Value) SetBigIntValue(v int64) *Value {
	s.BigIntValue = &v
	return s
}

// SetBitValue sets the BitValue field's value.
func (s *Value) SetBitValue(v bool) *Value {
	s.BitValue = &v
	return s
}

// SetBlobValue sets the BlobValue field's value.
func (s *Value) SetBlobValue(v []byte) *Value {
	s.BlobValue = v
	return s
}

// SetDoubleValue sets the DoubleValue field's value.
func (s *Value) SetDoubleValue(v float64) *Value {
	s.DoubleValue = &v
	return s
}

// SetIntValue sets the IntValue field's value.
func (s *Value) SetIntValue(v int64) *Value {
	s.IntValue = &v
	return s
}

// SetIsNull sets the IsNull field's value.
func (s *Value) SetIsNull(v bool) *Value {
	s.IsNull = &v
	return s
}

// SetRealValue sets the RealValue field's value.
func (s *Value) SetRealValue(v float64) *Value {
	s.RealValue = &v
	return s
}

// SetStringValue sets the StringValue field's value.
func (s *Value) SetStringValue(v string) *Value {
	s.StringValue = &v
	return s
}

// SetStructValue sets the StructValue field's value.
func (s *Value) SetStructValue(v *StructValue) *Value {
	s.StructValue = v
	return s
}
