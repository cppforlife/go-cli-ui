package errors_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/cppforlife/go-cli-ui/errors"
)

func TestSemiStructuredError(t *testing.T) {
	tests := []semiStructuredErrorTest{
		{
			Description: "multiple items",
			Actual:      `Job.batch "pi" is invalid: [spec.selector: Required value, spec.template.metadata.labels: Invalid value: map[string]string{"kapp.k14s.io/app":"1586905796363557000", "kapp.k14s.io/association":"v1.a4db8f96450049336d37eb62d798d883"}: selector does not match template labels, spec.selector: Invalid value: "null": field is immutable, spec.template: Invalid value: core.PodTemplateSpec{ObjectMeta:v1.ObjectMeta{Name:"", GenerateName:"", Namespace:"", SelfLink:"", UID:"", ResourceVersion:"", Generation:0, CreationTimestamp:v1.Time{Time:time.Time{wall:0x0, ext:0, loc:(*time.Location)(nil)}}, DeletionTimestamp:(*v1.Time)(nil), DeletionGracePeriodSeconds:(*int64)(nil), Labels:map[string]string{"kapp.k14s.io/app":"1586905796363557000", "kapp.k14s.io/association":"v1.a4db8f96450049336d37eb62d798d883"}, Annotations:map[string]string(nil), OwnerReferences:[]v1.OwnerReference(nil), Initializers:(*v1.Initializers)(nil), Finalizers:[]string(nil), ClusterName:"", ManagedFields:[]v1.ManagedFieldsEntry(nil)}, Spec:core.PodSpec{Volumes:[]core.Volume(nil), InitContainers:[]core.Container(nil), Containers:[]core.Container{core.Container{Name:"pi", Image:"perl", Command:[]string{"perl", "-Mbignum=bpi", "-wle", "print bpi(2000)"}, Args:[]string(nil), WorkingDir:"", Ports:[]core.ContainerPort(nil), EnvFrom:[]core.EnvFromSource(nil), Env:[]core.EnvVar(nil), Resources:core.ResourceRequirements{Limits:core.ResourceList(nil), Requests:core.ResourceList(nil)}, VolumeMounts:[]core.VolumeMount(nil), VolumeDevices:[]core.VolumeDevice(nil), LivenessProbe:(*core.Probe)(nil), ReadinessProbe:(*core.Probe)(nil), Lifecycle:(*core.Lifecycle)(nil), TerminationMessagePath:"/dev/termination-log", TerminationMessagePolicy:"File", ImagePullPolicy:"Always", SecurityContext:(*core.SecurityContext)(nil), Stdin:false, StdinOnce:false, TTY:false}}, RestartPolicy:"Never", TerminationGracePeriodSeconds:(*int64)(0xc00e08a8d0), ActiveDeadlineSeconds:(*int64)(nil), DNSPolicy:"ClusterFirst", NodeSelector:map[string]string(nil), ServiceAccountName:"", AutomountServiceAccountToken:(*bool)(nil), NodeName:"", SecurityContext:(*core.PodSecurityContext)(0xc012c22bd0), ImagePullSecrets:[]core.LocalObjectReference(nil), Hostname:"", Subdomain:"", Affinity:(*core.Affinity)(nil), SchedulerName:"default-scheduler", Tolerations:[]core.Toleration(nil), HostAliases:[]core.HostAlias(nil), PriorityClassName:"", Priority:(*int32)(nil), PreemptionPolicy:(*core.PreemptionPolicy)(nil), DNSConfig:(*core.PodDNSConfig)(nil), ReadinessGates:[]core.PodReadinessGate(nil), RuntimeClassName:(*string)(nil), EnableServiceLinks:(*bool)(nil)}}: field is immutable] (reason: Invalid)`,
			Expected: `
Job.batch "pi" is invalid: 

  - spec.selector: Required value

  - spec.template.metadata.labels: Invalid value: map[string]string{"kapp.k14s.io/app":"1586905796363557000", "kapp.k14s.io/association":"v1.a4db8f96450049336d37eb62d798d883"}: selector does not match template labels

  - spec.selector: Invalid value: "null": field is immutable

  - spec.template: Invalid value: core.PodTemplateSpec{ObjectMeta:v1.ObjectMeta{Name:"", GenerateName:"", Namespace:"", SelfLink:"", UID:"", ResourceVersion:"", Generation:0, CreationTimestamp:v1.Time{Time:time.Time{wall:0x0, ext:0, loc:(*time.Location)(nil)}}, DeletionTimestamp:(*v1.Time)(nil), DeletionGracePeriodSeconds:(*int64)(nil), Labels:map[string]string{"kapp.k14s.io/app":"1586905796363557000", "kapp.k14s.io/association":"v1.a4db8f96450049336d37eb62d798d883"}, Annotations:map[string]string(nil), OwnerReferences:[]v1.OwnerReference(nil), Initializers:(*v1.Initializers)(nil), Finalizers:[]string(nil), ClusterName:"", ManagedFields:[]v1.ManagedFieldsEntry(nil)}, Spec:core.PodSpec{Volumes:[]core.Volume(nil), InitContainers:[]core.Container(nil), Containers:[]core.Container{core.Container{Name:"pi", Image:"perl", Command:[]string{"perl", "-Mbignum=bpi", "-wle", "print bpi(2000)"}, Args:[]string(nil), WorkingDir:"", Ports:[]core.ContainerPort(nil), EnvFrom:[]core.EnvFromSource(nil), Env:[]core.EnvVar(nil), Resources:core.ResourceRequirements{Limits:core.ResourceList(nil), Requests:core.ResourceList(nil)}, VolumeMounts:[]core.VolumeMount(nil), VolumeDevices:[]core.VolumeDevice(nil), LivenessProbe:(*core.Probe)(nil), ReadinessProbe:(*core.Probe)(nil), Lifecycle:(*core.Lifecycle)(nil), TerminationMessagePath:"/dev/termination-log", TerminationMessagePolicy:"File", ImagePullPolicy:"Always", SecurityContext:(*core.SecurityContext)(nil), Stdin:false, StdinOnce:false, TTY:false}}, RestartPolicy:"Never", TerminationGracePeriodSeconds:(*int64)(0xc00e08a8d0), ActiveDeadlineSeconds:(*int64)(nil), DNSPolicy:"ClusterFirst", NodeSelector:map[string]string(nil), ServiceAccountName:"", AutomountServiceAccountToken:(*bool)(nil), NodeName:"", SecurityContext:(*core.PodSecurityContext)(0xc012c22bd0), ImagePullSecrets:[]core.LocalObjectReference(nil), Hostname:"", Subdomain:"", Affinity:(*core.Affinity)(nil), SchedulerName:"default-scheduler", Tolerations:[]core.Toleration(nil), HostAliases:[]core.HostAlias(nil), PriorityClassName:"", Priority:(*int32)(nil), PreemptionPolicy:(*core.PreemptionPolicy)(nil), DNSConfig:(*core.PodDNSConfig)(nil), ReadinessGates:[]core.PodReadinessGate(nil), RuntimeClassName:(*string)(nil), EnableServiceLinks:(*bool)(nil)}}: field is immutable

 (reason: Invalid)
`,
		},
		{
			Description: "item contains ',' for inside-item formatting",
			Actual:      `Job.batch "successful-job" is invalid: [spec.selector: Invalid value: v1.LabelSelector{MatchLabels:map[string]string{"blah":"balh", "controller-uid":"374ab0c4-8a21-4a9b-b814-fe85cf99a69a"}, MatchExpressions:[]v1.LabelSelectorRequirement(nil)}: selector not auto-generated, spec.template.spec.restartPolicy: Unsupported value: "Always": supported values: "OnFailure", "Never", spec.template.metadata.labels: Invalid value: map[string]string{"controller-uid":"374ab0c4-8a21-4a9b-b814-fe85cf99a69a", "foo":"foo", "job-name":"successful-job", "kapp.k14s.io/app":"1588294182746647000", "kapp.k14s.io/association":"v1.627b27b70a5aaeaa8dbd44b5cee9b165"}: selector does not match template labels] (reason: Invalid)`,
			Expected: `
Job.batch "successful-job" is invalid: 

  - spec.selector: Invalid value: v1.LabelSelector{MatchLabels:map[string]string{"blah":"balh", "controller-uid":"374ab0c4-8a21-4a9b-b814-fe85cf99a69a"}, MatchExpressions:[]v1.LabelSelectorRequirement(nil)}: selector not auto-generated

  - spec.template.spec.restartPolicy: Unsupported value: "Always": supported values: "OnFailure", "Never"

  - spec.template.metadata.labels: Invalid value: map[string]string{"controller-uid":"374ab0c4-8a21-4a9b-b814-fe85cf99a69a", "foo":"foo", "job-name":"successful-job", "kapp.k14s.io/app":"1588294182746647000", "kapp.k14s.io/association":"v1.627b27b70a5aaeaa8dbd44b5cee9b165"}: selector does not match template labels

 (reason: Invalid)
`,
		},
		{
			Description: "single item, not in brackets",
			Actual:      `Deployment.apps "frontend" is invalid: spec.template.metadata.labels: Invalid value: map[string]string{"app":"guestbook", "kapp.k14s.io/app":"1588343775866234000", "kapp.k14s.io/association":"v1.95c1511bde234f3b1296c5e2be3c6864", "tier":"frontend"}: selector does not match template labels (reason: Invalid)`,
			Expected:    `Deployment.apps "frontend" is invalid: spec.template.metadata.labels: Invalid value: map[string]string{"app":"guestbook", "kapp.k14s.io/app":"1588343775866234000", "kapp.k14s.io/association":"v1.95c1511bde234f3b1296c5e2be3c6864", "tier":"frontend"}: selector does not match template labels (reason: Invalid)`,
		},
		{
			Description: "structure is not symmetric",
			Actual:      `Job.batch "pi" is invalid: [spec.selector: Required value, spec.template.metadata.labels: Invalid value: map[string]string{"kapp.k14s.io/app":"1586905796363557000", "kapp.k14s.io/association":"v1.a4db8f96450049336d37eb62d798d883"}: selector does not match template labels, spec.selector: Invalid value: "null": field is immutable, spec.template: Invalid value: core.PodTemplateSpec{ObjectMeta:v1.ObjectMeta{Name:"", GenerateName:"",...: field is immutable] (reason: Invalid)`,
			Expected:    `Job.batch "pi" is invalid: [spec.selector: Required value, spec.template.metadata.labels: Invalid value: map[string]string{"kapp.k14s.io/app":"1586905796363557000", "kapp.k14s.io/association":"v1.a4db8f96450049336d37eb62d798d883"}: selector does not match template labels, spec.selector: Invalid value: "null": field is immutable, spec.template: Invalid value: core.PodTemplateSpec{ObjectMeta:v1.ObjectMeta{Name:"", GenerateName:"",...: field is immutable] (reason: Invalid)`,
		},
		{
			Description: "truncated content",
			Actual:      `Applying create service/redis-master (v1) namespace: default: Creating resource service/redis-master (v1) namespace: default: Service in version "v1" cannot be handled as a Service: v1.Service.Spec: v1.ServiceSpec.Ports: []v1.ServicePort: v1.ServicePort.Port: readUint32: unexpected character: �, error found in #10 byte of ...|[{"port":"6380s","ta|..., bigger context ...|,"namespace":"default"},"spec":{"ports":[{"port":"6380s","targetPort":6380}],"selector":{"app":"redi|... (reason: BadRequest)`,
			Expected:    `Applying create service/redis-master (v1) namespace: default: Creating resource service/redis-master (v1) namespace: default: Service in version "v1" cannot be handled as a Service: v1.Service.Spec: v1.ServiceSpec.Ports: []v1.ServicePort: v1.ServicePort.Port: readUint32: unexpected character: �, error found in #10 byte of ...|[{"port":"6380s","ta|..., bigger context ...|,"namespace":"default"},"spec":{"ports":[{"port":"6380s","targetPort":6380}],"selector":{"app":"redi|... (reason: BadRequest)`,
		},
	}

	for _, test := range tests {
		test.Check(t)
	}
}

type semiStructuredErrorTest struct {
	Description string
	Actual      string
	Expected    string
}

func (e semiStructuredErrorTest) Check(t *testing.T) {
	apiErr := errors.NewSemiStructuredError(fmt.Errorf("%s", e.Actual))
	e.Expected = strings.TrimSpace(e.Expected)

	if apiErr.Error() != e.Expected {
		t.Fatalf("(%s) expected error to match:\n%d chars >>>%s<<< vs \n%d chars >>>%s<<<",
			e.Description, len(apiErr.Error()), apiErr, len(e.Expected), e.Expected)
	}
}
