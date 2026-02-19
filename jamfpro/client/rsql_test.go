package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRSQLFilterBuilder(t *testing.T) {
	b := NewRSQLFilterBuilder()
	assert.True(t, b.IsEmpty())

	s := b.EqualTo("name", "x").Build()
	assert.Equal(t, `name=="x"`, s)
}

func TestRSQL_NotEqualTo_LessThan_GreaterThan(t *testing.T) {
	s := NewRSQLFilterBuilder().NotEqualTo("a", "b").Build()
	assert.Equal(t, `a!="b"`, s)
	s = NewRSQLFilterBuilder().LessThan("n", "10").Build()
	assert.Equal(t, `n<"10"`, s)
	s = NewRSQLFilterBuilder().LessOrEqual("n", "10").Build()
	assert.Equal(t, `n<="10"`, s)
	s = NewRSQLFilterBuilder().GreaterThan("n", "0").Build()
	assert.Equal(t, `n>"0"`, s)
	s = NewRSQLFilterBuilder().GreaterOrEqual("n", "0").Build()
	assert.Equal(t, `n>="0"`, s)
}

func TestRSQL_In_NotIn(t *testing.T) {
	s := NewRSQLFilterBuilder().In("id", "1", "2").Build()
	assert.Equal(t, `id=in=("1","2")`, s)
	s = NewRSQLFilterBuilder().NotIn("id", "a").Build()
	assert.Equal(t, `id=out=("a")`, s)
}

func TestRSQL_Contains_StartsWith_EndsWith(t *testing.T) {
	s := NewRSQLFilterBuilder().Contains("name", "Mac").Build()
	assert.Equal(t, `name=="*Mac*"`, s)
	s = NewRSQLFilterBuilder().StartsWith("name", "Test").Build()
	assert.Equal(t, `name=="Test*"`, s)
	s = NewRSQLFilterBuilder().EndsWith("name", "x").Build()
	assert.Equal(t, `name=="*x"`, s)
	s = NewRSQLFilterBuilder().Contains("field", "a*b").Build()
	assert.Equal(t, `field=="*a\*b*"`, s)
}

func TestRSQL_And_Or_Groups(t *testing.T) {
	s := NewRSQLFilterBuilder().
		EqualTo("a", "1").
		And().
		EqualTo("b", "2").
		OpenGroup().
		EqualTo("c", "3").
		Or().
		EqualTo("d", "4").
		CloseGroup().
		Build()
	assert.NotEmpty(t, s)
	assert.Equal(t, byte(')'), s[len(s)-1])
}

func TestRSQL_Quote_Escape(t *testing.T) {
	s := NewRSQLFilterBuilder().EqualTo("msg", `say "hello"`).Build()
	assert.Equal(t, `msg=="say \"hello\""`, s)
}
