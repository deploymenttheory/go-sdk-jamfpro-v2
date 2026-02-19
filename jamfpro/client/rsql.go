package client

import (
	"fmt"
	"strings"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
)

// rsqlFilterBuilder implements interfaces.RSQLFilterBuilder.
// It accumulates RSQL tokens into a strings.Builder and is not thread-safe;
// create a new instance per expression.
//
// See: https://developer.jamf.com/jamf-pro/docs/filtering-with-rsql
type rsqlFilterBuilder struct {
	buf strings.Builder
}

// NewRSQLFilterBuilder returns a new, empty RSQL filter expression builder.
func NewRSQLFilterBuilder() interfaces.RSQLFilterBuilder {
	return &rsqlFilterBuilder{}
}

// EqualTo produces: field=="value"
// Wildcards (*) already present in value are preserved as RSQL wildcards.
func (b *rsqlFilterBuilder) EqualTo(field, value string) interfaces.RSQLFilterBuilder {
	fmt.Fprintf(&b.buf, `%s==%s`, field, rsqlQuote(value))
	return b
}

// NotEqualTo produces: field!="value"
func (b *rsqlFilterBuilder) NotEqualTo(field, value string) interfaces.RSQLFilterBuilder {
	fmt.Fprintf(&b.buf, `%s!=%s`, field, rsqlQuote(value))
	return b
}

// LessThan produces: field<"value"
func (b *rsqlFilterBuilder) LessThan(field, value string) interfaces.RSQLFilterBuilder {
	fmt.Fprintf(&b.buf, `%s<%s`, field, rsqlQuote(value))
	return b
}

// LessOrEqual produces: field<="value"
func (b *rsqlFilterBuilder) LessOrEqual(field, value string) interfaces.RSQLFilterBuilder {
	fmt.Fprintf(&b.buf, `%s<=%s`, field, rsqlQuote(value))
	return b
}

// GreaterThan produces: field>"value"
func (b *rsqlFilterBuilder) GreaterThan(field, value string) interfaces.RSQLFilterBuilder {
	fmt.Fprintf(&b.buf, `%s>%s`, field, rsqlQuote(value))
	return b
}

// GreaterOrEqual produces: field>="value"
func (b *rsqlFilterBuilder) GreaterOrEqual(field, value string) interfaces.RSQLFilterBuilder {
	fmt.Fprintf(&b.buf, `%s>=%s`, field, rsqlQuote(value))
	return b
}

// In produces: field=in=(v1,v2,...)
func (b *rsqlFilterBuilder) In(field string, values ...string) interfaces.RSQLFilterBuilder {
	quoted := make([]string, len(values))
	for i, v := range values {
		quoted[i] = rsqlQuote(v)
	}
	fmt.Fprintf(&b.buf, `%s=in=(%s)`, field, strings.Join(quoted, ","))
	return b
}

// NotIn produces: field=out=(v1,v2,...)
func (b *rsqlFilterBuilder) NotIn(field string, values ...string) interfaces.RSQLFilterBuilder {
	quoted := make([]string, len(values))
	for i, v := range values {
		quoted[i] = rsqlQuote(v)
	}
	fmt.Fprintf(&b.buf, `%s=out=(%s)`, field, strings.Join(quoted, ","))
	return b
}

// Contains produces: field=="*value*"
// Literal asterisks in value are escaped so they are not treated as wildcards.
func (b *rsqlFilterBuilder) Contains(field, value string) interfaces.RSQLFilterBuilder {
	fmt.Fprintf(&b.buf, `%s=="*%s*"`, field, rsqlEscapeLiteralWildcard(value))
	return b
}

// StartsWith produces: field=="value*"
func (b *rsqlFilterBuilder) StartsWith(field, value string) interfaces.RSQLFilterBuilder {
	fmt.Fprintf(&b.buf, `%s=="%s*"`, field, rsqlEscapeLiteralWildcard(value))
	return b
}

// EndsWith produces: field=="*value"
func (b *rsqlFilterBuilder) EndsWith(field, value string) interfaces.RSQLFilterBuilder {
	fmt.Fprintf(&b.buf, `%s=="*%s"`, field, rsqlEscapeLiteralWildcard(value))
	return b
}

// And appends a semicolon — logical AND in RSQL.
func (b *rsqlFilterBuilder) And() interfaces.RSQLFilterBuilder {
	b.buf.WriteByte(';')
	return b
}

// Or appends a comma — logical OR in RSQL.
func (b *rsqlFilterBuilder) Or() interfaces.RSQLFilterBuilder {
	b.buf.WriteByte(',')
	return b
}

// OpenGroup appends a left parenthesis.
func (b *rsqlFilterBuilder) OpenGroup() interfaces.RSQLFilterBuilder {
	b.buf.WriteByte('(')
	return b
}

// CloseGroup appends a right parenthesis.
func (b *rsqlFilterBuilder) CloseGroup() interfaces.RSQLFilterBuilder {
	b.buf.WriteByte(')')
	return b
}

// Build returns the completed RSQL expression string.
func (b *rsqlFilterBuilder) Build() string {
	return b.buf.String()
}

// IsEmpty reports whether no tokens have been written yet.
func (b *rsqlFilterBuilder) IsEmpty() bool {
	return b.buf.Len() == 0
}

// rsqlQuote wraps value in double quotes and escapes embedded double quotes.
// All values are quoted for safety; Jamf Pro accepts quoted numeric values.
// Wildcards (*) already present in the value are preserved (not escaped),
// allowing callers to embed wildcards in EqualTo/NotEqualTo expressions.
func rsqlQuote(value string) string {
	escaped := strings.ReplaceAll(value, `"`, `\"`)
	return `"` + escaped + `"`
}

// rsqlEscapeLiteralWildcard escapes literal asterisks and backslashes in a
// value used with Contains/StartsWith/EndsWith, so the user's value is treated
// as a plain string rather than a wildcard pattern. The wildcard characters
// added by those methods themselves are NOT passed through this function.
func rsqlEscapeLiteralWildcard(value string) string {
	value = strings.ReplaceAll(value, `\`, `\\`)
	value = strings.ReplaceAll(value, `*`, `\*`)
	value = strings.ReplaceAll(value, `"`, `\"`)
	return value
}
