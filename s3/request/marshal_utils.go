package request

// StringValue provides encoding of string for AWS protocols.
type StringValue string

// MarshalValue formats the value into a string for encoding.
func (v StringValue) MarshalValue() (string, error) {
	return string(v), nil
}
