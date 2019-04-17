package kafka

import (
	"log"
	"testing"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&KafkaActivity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestPlain(t *testing.T) {

	act := &KafkaActivity{}
	tc := test.NewActivityContext(act.Metadata())

	//setup attrs
	tc.SetInput("brokerurls", "localhost:9092")
	tc.SetInput("topic", "syslog")
	tc.SetInput("message", "######### PLAIN ###########  Mary had a little lamb, its fleece was white as snow.")
	act.Eval(tc)
	log.Printf("TestEval successfull.  partition [%d]  offset [%d]", tc.GetOutput("partition"), tc.GetOutput("offset"))
}
