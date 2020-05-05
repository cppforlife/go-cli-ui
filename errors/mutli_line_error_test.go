package errors_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/cppforlife/go-cli-ui/errors"
)

func TestMultiLineError(t *testing.T) {
	tests := []multiLineErrorTest{
		{
			Description: "single line error",
			Actual:      `Applying create deployment/frontend (apps/v1) namespace: default: Creating resource deployment/frontend (apps/v1) namespace: default: Deployment.apps "frontend" is invalid: spec.template.metadata.labels: Invalid value: map[string]string{"app":"guestbook", "kapp.k14s.io/app":"1588343775866234000", "kapp.k14s.io/association":"v1.95c1511bde234f3b1296c5e2be3c6864", "tier":"frontend"}: selector does not match template labels (reason: Invalid)`,
			Expected: `
Applying create deployment/frontend (apps/v1) namespace: default:
  Creating resource deployment/frontend (apps/v1) namespace: default:
    Deployment.apps "frontend" is invalid: spec.template.metadata.labels:
      Invalid value: map[string]string{"app":"guestbook", "kapp.k14s.io/app":"1588343775866234000", "kapp.k14s.io/association":"v1.95c1511bde234f3b1296c5e2be3c6864", "tier":"frontend"}: selector does not match template labels (reason: Invalid)
`,
		},
		{
			Description: "multi line error",
			Actual: `
Applying create deployment/frontend (apps/v1) namespace: default: Creating resource deployment/frontend (apps/v1) namespace: default: Job.batch "successful-job" is invalid: 

  - spec.selector: Invalid value: v1.LabelSelector{MatchLabels:map[string]string{"blah":"balh", "controller-uid":"374ab0c4-8a21-4a9b-b814-fe85cf99a69a"}, MatchExpressions:[]v1.LabelSelectorRequirement(nil)}: selector not auto-generated

  - spec.template.spec.restartPolicy: Unsupported value: "Always": supported values: "OnFailure", "Never"

 (reason: Invalid)
`,
			Expected: `
Applying create deployment/frontend (apps/v1) namespace: default:
  Creating resource deployment/frontend (apps/v1) namespace: default:
    Job.batch "successful-job" is invalid: 

  - spec.selector: Invalid value: v1.LabelSelector{MatchLabels:map[string]string{"blah":"balh", "controller-uid":"374ab0c4-8a21-4a9b-b814-fe85cf99a69a"}, MatchExpressions:[]v1.LabelSelectorRequirement(nil)}: selector not auto-generated

  - spec.template.spec.restartPolicy: Unsupported value: "Always": supported values: "OnFailure", "Never"

 (reason: Invalid)
`,
		},
	}

	for _, test := range tests {
		test.Check(t)
	}
}

type multiLineErrorTest struct {
	Description string
	Actual      string
	Expected    string
}

func (e multiLineErrorTest) Check(t *testing.T) {
	apiErr := errors.NewMultiLineError(fmt.Errorf("%s", strings.TrimSpace(e.Actual)))
	e.Expected = strings.TrimSpace(e.Expected)

	if apiErr.Error() != e.Expected {
		t.Fatalf("(%s) expected error to match:\n%d chars >>>%s<<< vs \n%d chars >>>%s<<<",
			e.Description, len(apiErr.Error()), apiErr, len(e.Expected), e.Expected)
	}
}
