package tests

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/project-flogo/rules/common"
	"github.com/project-flogo/rules/common/model"
	"github.com/project-flogo/rules/ruleapi"
)

func createRuleSession() (model.RuleSession, error) {
	rs, _ := ruleapi.GetOrCreateRuleSession("test")

	tupleDescFileAbsPath := common.GetAbsPathForResource("src/github.com/project-flogo/rules/ruleapi/tests/tests.json")

	dat, err := ioutil.ReadFile(tupleDescFileAbsPath)
	if err != nil {
		log.Fatal(err)
	}
	err = model.RegisterTupleDescriptors(string(dat))
	if err != nil {
		return nil, err
	}
	return rs, nil
}

//conditions and actions
func trueCondition(ruleName string, condName string, tuples map[model.TupleType]model.Tuple, ctx model.RuleContext) bool {
	return true
}
func falseCondition(ruleName string, condName string, tuples map[model.TupleType]model.Tuple, ctx model.RuleContext) bool {
	return false
}
func emptyAction(ctx context.Context, rs model.RuleSession, ruleName string, tuples map[model.TupleType]model.Tuple, ruleCtx model.RuleContext) {
}

func printTuples(t *testing.T, oprn string, tupleMap map[string]map[string]model.Tuple) {

	for k, v := range tupleMap {
		t.Logf("%s tuples for type [%s]\n", oprn, k)
		for k1, _ := range v {
			t.Logf("    tuples key [%s]\n", k1)
		}
	}
}
func printModified(t *testing.T, modified map[string]map[string]model.RtcModified) {

	for k, v := range modified {
		t.Logf("%s tuples for type [%s]\n", "Modified", k)
		for k1, _ := range v {
			t.Logf("    tuples key [%s]\n", k1)
		}
	}
}

type txnCtx struct {
	Testing *testing.T
	TxnCnt  int
}

// func createBActionServiceFromFunction(b *testing.B, actionFunction model.ActionFunction) model.ActionService {
// 	fname := runtime.FuncForPC(reflect.ValueOf(actionFunction).Pointer()).Name()
// 	cfg := &config.ServiceDescriptor{
// 		Name:        fname,
// 		Description: fname,
// 		Type:        config.TypeServiceFunction,
// 		Function:    actionFunction,
// 	}
// 	aService, err := ruleapi.NewActionService(cfg)
// 	assert.Nil(b, err)
// 	assert.NotNil(b, aService)
// 	return aService
// }

func Command(name string, arg ...string) {
	fmt.Printf("%s %v\n", name, arg)
	output, err := exec.Command(name, arg...).CombinedOutput()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	if len(output) > 0 {
		fmt.Printf("%s", string(output))
	}
}

func Drain(port string) {
	for {
		conn, err := net.DialTimeout("tcp", net.JoinHostPort("", port), time.Second)
		if conn != nil {
			conn.Close()
		}
		if err != nil && strings.Contains(err.Error(), "connect: connection refused") {
			break
		}
	}
}

func Pour(port string) {
	for {
		conn, _ := net.Dial("tcp", net.JoinHostPort("", port))
		if conn != nil {
			conn.Close()
			break
		}
	}
}

type TestKey struct{}
