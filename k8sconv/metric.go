// Code generated from semantic convention specification. DO NOT EDIT.

package k8sconv

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/noop"
)

// NamespacePhaseAttr is an attribute conforming to the k8s.namespace.phase
// semantic conventions. It represents the phase of the K8s namespace.
type NamespacePhaseAttr string

var (
	// NamespacePhaseActive is the active namespace phase as described by [K8s API]
	// .
	//
	// [K8s API]: https://pkg.go.dev/k8s.io/api@v0.31.3/core/v1#NamespacePhase
	NamespacePhaseActive NamespacePhaseAttr = "active"
	// NamespacePhaseTerminating is the terminating namespace phase as described by
	// [K8s API].
	//
	// [K8s API]: https://pkg.go.dev/k8s.io/api@v0.31.3/core/v1#NamespacePhase
	NamespacePhaseTerminating NamespacePhaseAttr = "terminating"
)

// NetworkIODirectionAttr is an attribute conforming to the network.io.direction
// semantic conventions. It represents the network IO operation direction.
type NetworkIODirectionAttr string

var (
	// NetworkIODirectionTransmit is the none.
	NetworkIODirectionTransmit NetworkIODirectionAttr = "transmit"
	// NetworkIODirectionReceive is the none.
	NetworkIODirectionReceive NetworkIODirectionAttr = "receive"
)

// CronJobActiveJobs is an instrument used to record metric values conforming to
// the "k8s.cronjob.active_jobs" semantic conventions. It represents the number
// of actively running jobs for a cronjob.
type CronJobActiveJobs struct {
	inst metric.Int64UpDownCounter
}

// NewCronJobActiveJobs returns a new CronJobActiveJobs instrument.
func NewCronJobActiveJobs(m metric.Meter) (CronJobActiveJobs, error) {
	i, err := m.Int64UpDownCounter(
	    "k8s.cronjob.active_jobs",
	    metric.WithDescription("The number of actively running jobs for a cronjob"),
	    metric.WithUnit("{job}"),
	)
	if err != nil {
	    return CronJobActiveJobs{inst: noop.Int64UpDownCounter{}}, err
	}
	return CronJobActiveJobs{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (CronJobActiveJobs) Name() string {
	return "k8s.cronjob.active_jobs"
}

// Unit returns the semantic convention unit of the instrument
func (CronJobActiveJobs) Unit() string {
	return "{job}"
}

// Description returns the semantic convention description of the instrument
func (CronJobActiveJobs) Description() string {
	return "The number of actively running jobs for a cronjob"
}

// Add adds incr to the existing count.
//
// This metric aligns with the `active` field of the
// [K8s CronJobStatus].
//
// This metric SHOULD, at a minimum, be reported against a
// [`k8s.cronjob`] resource.
//
// [K8s CronJobStatus]: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#cronjobstatus-v1-batch
// [`k8s.cronjob`]: ../resource/k8s.md#cronjob
func (m CronJobActiveJobs) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Add(ctx, incr)
	} else {
		m.inst.Add(ctx, incr, metric.WithAttributes(attrs...))
	}
}

// DaemonSetCurrentScheduledNodes is an instrument used to record metric values
// conforming to the "k8s.daemonset.current_scheduled_nodes" semantic
// conventions. It represents the number of nodes that are running at least 1
// daemon pod and are supposed to run the daemon pod.
type DaemonSetCurrentScheduledNodes struct {
	inst metric.Int64UpDownCounter
}

// NewDaemonSetCurrentScheduledNodes returns a new DaemonSetCurrentScheduledNodes
// instrument.
func NewDaemonSetCurrentScheduledNodes(m metric.Meter) (DaemonSetCurrentScheduledNodes, error) {
	i, err := m.Int64UpDownCounter(
	    "k8s.daemonset.current_scheduled_nodes",
	    metric.WithDescription("Number of nodes that are running at least 1 daemon pod and are supposed to run the daemon pod"),
	    metric.WithUnit("{node}"),
	)
	if err != nil {
	    return DaemonSetCurrentScheduledNodes{inst: noop.Int64UpDownCounter{}}, err
	}
	return DaemonSetCurrentScheduledNodes{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (DaemonSetCurrentScheduledNodes) Name() string {
	return "k8s.daemonset.current_scheduled_nodes"
}

// Unit returns the semantic convention unit of the instrument
func (DaemonSetCurrentScheduledNodes) Unit() string {
	return "{node}"
}

// Description returns the semantic convention description of the instrument
func (DaemonSetCurrentScheduledNodes) Description() string {
	return "Number of nodes that are running at least 1 daemon pod and are supposed to run the daemon pod"
}

// Add adds incr to the existing count.
//
// This metric aligns with the `currentNumberScheduled` field of the
// [K8s DaemonSetStatus].
//
// This metric SHOULD, at a minimum, be reported against a
// [`k8s.daemonset`] resource.
//
// [K8s DaemonSetStatus]: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#daemonsetstatus-v1-apps
// [`k8s.daemonset`]: ../resource/k8s.md#daemonset
func (m DaemonSetCurrentScheduledNodes) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Add(ctx, incr)
	} else {
		m.inst.Add(ctx, incr, metric.WithAttributes(attrs...))
	}
}

// DaemonSetDesiredScheduledNodes is an instrument used to record metric values
// conforming to the "k8s.daemonset.desired_scheduled_nodes" semantic
// conventions. It represents the number of nodes that should be running the
// daemon pod (including nodes currently running the daemon pod).
type DaemonSetDesiredScheduledNodes struct {
	inst metric.Int64UpDownCounter
}

// NewDaemonSetDesiredScheduledNodes returns a new DaemonSetDesiredScheduledNodes
// instrument.
func NewDaemonSetDesiredScheduledNodes(m metric.Meter) (DaemonSetDesiredScheduledNodes, error) {
	i, err := m.Int64UpDownCounter(
	    "k8s.daemonset.desired_scheduled_nodes",
	    metric.WithDescription("Number of nodes that should be running the daemon pod (including nodes currently running the daemon pod)"),
	    metric.WithUnit("{node}"),
	)
	if err != nil {
	    return DaemonSetDesiredScheduledNodes{inst: noop.Int64UpDownCounter{}}, err
	}
	return DaemonSetDesiredScheduledNodes{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (DaemonSetDesiredScheduledNodes) Name() string {
	return "k8s.daemonset.desired_scheduled_nodes"
}

// Unit returns the semantic convention unit of the instrument
func (DaemonSetDesiredScheduledNodes) Unit() string {
	return "{node}"
}

// Description returns the semantic convention description of the instrument
func (DaemonSetDesiredScheduledNodes) Description() string {
	return "Number of nodes that should be running the daemon pod (including nodes currently running the daemon pod)"
}

// Add adds incr to the existing count.
//
// This metric aligns with the `desiredNumberScheduled` field of the
// [K8s DaemonSetStatus].
//
// This metric SHOULD, at a minimum, be reported against a
// [`k8s.daemonset`] resource.
//
// [K8s DaemonSetStatus]: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#daemonsetstatus-v1-apps
// [`k8s.daemonset`]: ../resource/k8s.md#daemonset
func (m DaemonSetDesiredScheduledNodes) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Add(ctx, incr)
	} else {
		m.inst.Add(ctx, incr, metric.WithAttributes(attrs...))
	}
}

// DaemonSetMisscheduledNodes is an instrument used to record metric values
// conforming to the "k8s.daemonset.misscheduled_nodes" semantic conventions. It
// represents the number of nodes that are running the daemon pod, but are not
// supposed to run the daemon pod.
type DaemonSetMisscheduledNodes struct {
	inst metric.Int64UpDownCounter
}

// NewDaemonSetMisscheduledNodes returns a new DaemonSetMisscheduledNodes
// instrument.
func NewDaemonSetMisscheduledNodes(m metric.Meter) (DaemonSetMisscheduledNodes, error) {
	i, err := m.Int64UpDownCounter(
	    "k8s.daemonset.misscheduled_nodes",
	    metric.WithDescription("Number of nodes that are running the daemon pod, but are not supposed to run the daemon pod"),
	    metric.WithUnit("{node}"),
	)
	if err != nil {
	    return DaemonSetMisscheduledNodes{inst: noop.Int64UpDownCounter{}}, err
	}
	return DaemonSetMisscheduledNodes{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (DaemonSetMisscheduledNodes) Name() string {
	return "k8s.daemonset.misscheduled_nodes"
}

// Unit returns the semantic convention unit of the instrument
func (DaemonSetMisscheduledNodes) Unit() string {
	return "{node}"
}

// Description returns the semantic convention description of the instrument
func (DaemonSetMisscheduledNodes) Description() string {
	return "Number of nodes that are running the daemon pod, but are not supposed to run the daemon pod"
}

// Add adds incr to the existing count.
//
// This metric aligns with the `numberMisscheduled` field of the
// [K8s DaemonSetStatus].
//
// This metric SHOULD, at a minimum, be reported against a
// [`k8s.daemonset`] resource.
//
// [K8s DaemonSetStatus]: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#daemonsetstatus-v1-apps
// [`k8s.daemonset`]: ../resource/k8s.md#daemonset
func (m DaemonSetMisscheduledNodes) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Add(ctx, incr)
	} else {
		m.inst.Add(ctx, incr, metric.WithAttributes(attrs...))
	}
}

// DaemonSetReadyNodes is an instrument used to record metric values conforming
// to the "k8s.daemonset.ready_nodes" semantic conventions. It represents the
// number of nodes that should be running the daemon pod and have one or more of
// the daemon pod running and ready.
type DaemonSetReadyNodes struct {
	inst metric.Int64UpDownCounter
}

// NewDaemonSetReadyNodes returns a new DaemonSetReadyNodes instrument.
func NewDaemonSetReadyNodes(m metric.Meter) (DaemonSetReadyNodes, error) {
	i, err := m.Int64UpDownCounter(
	    "k8s.daemonset.ready_nodes",
	    metric.WithDescription("Number of nodes that should be running the daemon pod and have one or more of the daemon pod running and ready"),
	    metric.WithUnit("{node}"),
	)
	if err != nil {
	    return DaemonSetReadyNodes{inst: noop.Int64UpDownCounter{}}, err
	}
	return DaemonSetReadyNodes{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (DaemonSetReadyNodes) Name() string {
	return "k8s.daemonset.ready_nodes"
}

// Unit returns the semantic convention unit of the instrument
func (DaemonSetReadyNodes) Unit() string {
	return "{node}"
}

// Description returns the semantic convention description of the instrument
func (DaemonSetReadyNodes) Description() string {
	return "Number of nodes that should be running the daemon pod and have one or more of the daemon pod running and ready"
}

// Add adds incr to the existing count.
//
// This metric aligns with the `numberReady` field of the
// [K8s DaemonSetStatus].
//
// This metric SHOULD, at a minimum, be reported against a
// [`k8s.daemonset`] resource.
//
// [K8s DaemonSetStatus]: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#daemonsetstatus-v1-apps
// [`k8s.daemonset`]: ../resource/k8s.md#daemonset
func (m DaemonSetReadyNodes) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Add(ctx, incr)
	} else {
		m.inst.Add(ctx, incr, metric.WithAttributes(attrs...))
	}
}

// DeploymentAvailablePods is an instrument used to record metric values
// conforming to the "k8s.deployment.available_pods" semantic conventions. It
// represents the total number of available replica pods (ready for at least
// minReadySeconds) targeted by this deployment.
type DeploymentAvailablePods struct {
	inst metric.Int64UpDownCounter
}

// NewDeploymentAvailablePods returns a new DeploymentAvailablePods instrument.
func NewDeploymentAvailablePods(m metric.Meter) (DeploymentAvailablePods, error) {
	i, err := m.Int64UpDownCounter(
	    "k8s.deployment.available_pods",
	    metric.WithDescription("Total number of available replica pods (ready for at least minReadySeconds) targeted by this deployment"),
	    metric.WithUnit("{pod}"),
	)
	if err != nil {
	    return DeploymentAvailablePods{inst: noop.Int64UpDownCounter{}}, err
	}
	return DeploymentAvailablePods{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (DeploymentAvailablePods) Name() string {
	return "k8s.deployment.available_pods"
}

// Unit returns the semantic convention unit of the instrument
func (DeploymentAvailablePods) Unit() string {
	return "{pod}"
}

// Description returns the semantic convention description of the instrument
func (DeploymentAvailablePods) Description() string {
	return "Total number of available replica pods (ready for at least minReadySeconds) targeted by this deployment"
}

// Add adds incr to the existing count.
//
// This metric aligns with the `availableReplicas` field of the
// [K8s DeploymentStatus].
//
// This metric SHOULD, at a minimum, be reported against a
// [`k8s.deployment`] resource.
//
// [K8s DeploymentStatus]: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#deploymentstatus-v1-apps
// [`k8s.deployment`]: ../resource/k8s.md#deployment
func (m DeploymentAvailablePods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Add(ctx, incr)
	} else {
		m.inst.Add(ctx, incr, metric.WithAttributes(attrs...))
	}
}

// DeploymentDesiredPods is an instrument used to record metric values conforming
// to the "k8s.deployment.desired_pods" semantic conventions. It represents the
// number of desired replica pods in this deployment.
type DeploymentDesiredPods struct {
	inst metric.Int64UpDownCounter
}

// NewDeploymentDesiredPods returns a new DeploymentDesiredPods instrument.
func NewDeploymentDesiredPods(m metric.Meter) (DeploymentDesiredPods, error) {
	i, err := m.Int64UpDownCounter(
	    "k8s.deployment.desired_pods",
	    metric.WithDescription("Number of desired replica pods in this deployment"),
	    metric.WithUnit("{pod}"),
	)
	if err != nil {
	    return DeploymentDesiredPods{inst: noop.Int64UpDownCounter{}}, err
	}
	return DeploymentDesiredPods{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (DeploymentDesiredPods) Name() string {
	return "k8s.deployment.desired_pods"
}

// Unit returns the semantic convention unit of the instrument
func (DeploymentDesiredPods) Unit() string {
	return "{pod}"
}

// Description returns the semantic convention description of the instrument
func (DeploymentDesiredPods) Description() string {
	return "Number of desired replica pods in this deployment"
}

// Add adds incr to the existing count.
//
// This metric aligns with the `replicas` field of the
// [K8s DeploymentSpec].
//
// This metric SHOULD, at a minimum, be reported against a
// [`k8s.deployment`] resource.
//
// [K8s DeploymentSpec]: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#deploymentspec-v1-apps
// [`k8s.deployment`]: ../resource/k8s.md#deployment
func (m DeploymentDesiredPods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Add(ctx, incr)
	} else {
		m.inst.Add(ctx, incr, metric.WithAttributes(attrs...))
	}
}

// HpaCurrentPods is an instrument used to record metric values conforming to the
// "k8s.hpa.current_pods" semantic conventions. It represents the current number
// of replica pods managed by this horizontal pod autoscaler, as last seen by the
// autoscaler.
type HpaCurrentPods struct {
	inst metric.Int64UpDownCounter
}

// NewHpaCurrentPods returns a new HpaCurrentPods instrument.
func NewHpaCurrentPods(m metric.Meter) (HpaCurrentPods, error) {
	i, err := m.Int64UpDownCounter(
	    "k8s.hpa.current_pods",
	    metric.WithDescription("Current number of replica pods managed by this horizontal pod autoscaler, as last seen by the autoscaler"),
	    metric.WithUnit("{pod}"),
	)
	if err != nil {
	    return HpaCurrentPods{inst: noop.Int64UpDownCounter{}}, err
	}
	return HpaCurrentPods{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (HpaCurrentPods) Name() string {
	return "k8s.hpa.current_pods"
}

// Unit returns the semantic convention unit of the instrument
func (HpaCurrentPods) Unit() string {
	return "{pod}"
}

// Description returns the semantic convention description of the instrument
func (HpaCurrentPods) Description() string {
	return "Current number of replica pods managed by this horizontal pod autoscaler, as last seen by the autoscaler"
}

// Add adds incr to the existing count.
//
// This metric aligns with the `currentReplicas` field of the
// [K8s HorizontalPodAutoscalerStatus]
//
// This metric SHOULD, at a minimum, be reported against a
// [`k8s.hpa`] resource.
//
// [K8s HorizontalPodAutoscalerStatus]: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#horizontalpodautoscalerstatus-v2-autoscaling
// [`k8s.hpa`]: ../resource/k8s.md#horizontalpodautoscaler
func (m HpaCurrentPods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Add(ctx, incr)
	} else {
		m.inst.Add(ctx, incr, metric.WithAttributes(attrs...))
	}
}

// HpaDesiredPods is an instrument used to record metric values conforming to the
// "k8s.hpa.desired_pods" semantic conventions. It represents the desired number
// of replica pods managed by this horizontal pod autoscaler, as last calculated
// by the autoscaler.
type HpaDesiredPods struct {
	inst metric.Int64UpDownCounter
}

// NewHpaDesiredPods returns a new HpaDesiredPods instrument.
func NewHpaDesiredPods(m metric.Meter) (HpaDesiredPods, error) {
	i, err := m.Int64UpDownCounter(
	    "k8s.hpa.desired_pods",
	    metric.WithDescription("Desired number of replica pods managed by this horizontal pod autoscaler, as last calculated by the autoscaler"),
	    metric.WithUnit("{pod}"),
	)
	if err != nil {
	    return HpaDesiredPods{inst: noop.Int64UpDownCounter{}}, err
	}
	return HpaDesiredPods{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (HpaDesiredPods) Name() string {
	return "k8s.hpa.desired_pods"
}

// Unit returns the semantic convention unit of the instrument
func (HpaDesiredPods) Unit() string {
	return "{pod}"
}

// Description returns the semantic convention description of the instrument
func (HpaDesiredPods) Description() string {
	return "Desired number of replica pods managed by this horizontal pod autoscaler, as last calculated by the autoscaler"
}

// Add adds incr to the existing count.
//
// This metric aligns with the `desiredReplicas` field of the
// [K8s HorizontalPodAutoscalerStatus]
//
// This metric SHOULD, at a minimum, be reported against a
// [`k8s.hpa`] resource.
//
// [K8s HorizontalPodAutoscalerStatus]: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#horizontalpodautoscalerstatus-v2-autoscaling
// [`k8s.hpa`]: ../resource/k8s.md#horizontalpodautoscaler
func (m HpaDesiredPods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Add(ctx, incr)
	} else {
		m.inst.Add(ctx, incr, metric.WithAttributes(attrs...))
	}
}

// HpaMaxPods is an instrument used to record metric values conforming to the
// "k8s.hpa.max_pods" semantic conventions. It represents the upper limit for the
// number of replica pods to which the autoscaler can scale up.
type HpaMaxPods struct {
	inst metric.Int64UpDownCounter
}

// NewHpaMaxPods returns a new HpaMaxPods instrument.
func NewHpaMaxPods(m metric.Meter) (HpaMaxPods, error) {
	i, err := m.Int64UpDownCounter(
	    "k8s.hpa.max_pods",
	    metric.WithDescription("The upper limit for the number of replica pods to which the autoscaler can scale up"),
	    metric.WithUnit("{pod}"),
	)
	if err != nil {
	    return HpaMaxPods{inst: noop.Int64UpDownCounter{}}, err
	}
	return HpaMaxPods{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (HpaMaxPods) Name() string {
	return "k8s.hpa.max_pods"
}

// Unit returns the semantic convention unit of the instrument
func (HpaMaxPods) Unit() string {
	return "{pod}"
}

// Description returns the semantic convention description of the instrument
func (HpaMaxPods) Description() string {
	return "The upper limit for the number of replica pods to which the autoscaler can scale up"
}

// Add adds incr to the existing count.
//
// This metric aligns with the `maxReplicas` field of the
// [K8s HorizontalPodAutoscalerSpec]
//
// This metric SHOULD, at a minimum, be reported against a
// [`k8s.hpa`] resource.
//
// [K8s HorizontalPodAutoscalerSpec]: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#horizontalpodautoscalerspec-v2-autoscaling
// [`k8s.hpa`]: ../resource/k8s.md#horizontalpodautoscaler
func (m HpaMaxPods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Add(ctx, incr)
	} else {
		m.inst.Add(ctx, incr, metric.WithAttributes(attrs...))
	}
}

// HpaMinPods is an instrument used to record metric values conforming to the
// "k8s.hpa.min_pods" semantic conventions. It represents the lower limit for the
// number of replica pods to which the autoscaler can scale down.
type HpaMinPods struct {
	inst metric.Int64UpDownCounter
}

// NewHpaMinPods returns a new HpaMinPods instrument.
func NewHpaMinPods(m metric.Meter) (HpaMinPods, error) {
	i, err := m.Int64UpDownCounter(
	    "k8s.hpa.min_pods",
	    metric.WithDescription("The lower limit for the number of replica pods to which the autoscaler can scale down"),
	    metric.WithUnit("{pod}"),
	)
	if err != nil {
	    return HpaMinPods{inst: noop.Int64UpDownCounter{}}, err
	}
	return HpaMinPods{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (HpaMinPods) Name() string {
	return "k8s.hpa.min_pods"
}

// Unit returns the semantic convention unit of the instrument
func (HpaMinPods) Unit() string {
	return "{pod}"
}

// Description returns the semantic convention description of the instrument
func (HpaMinPods) Description() string {
	return "The lower limit for the number of replica pods to which the autoscaler can scale down"
}

// Add adds incr to the existing count.
//
// This metric aligns with the `minReplicas` field of the
// [K8s HorizontalPodAutoscalerSpec]
//
// This metric SHOULD, at a minimum, be reported against a
// [`k8s.hpa`] resource.
//
// [K8s HorizontalPodAutoscalerSpec]: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#horizontalpodautoscalerspec-v2-autoscaling
// [`k8s.hpa`]: ../resource/k8s.md#horizontalpodautoscaler
func (m HpaMinPods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Add(ctx, incr)
	} else {
		m.inst.Add(ctx, incr, metric.WithAttributes(attrs...))
	}
}

// JobActivePods is an instrument used to record metric values conforming to the
// "k8s.job.active_pods" semantic conventions. It represents the number of
// pending and actively running pods for a job.
type JobActivePods struct {
	inst metric.Int64UpDownCounter
}

// NewJobActivePods returns a new JobActivePods instrument.
func NewJobActivePods(m metric.Meter) (JobActivePods, error) {
	i, err := m.Int64UpDownCounter(
	    "k8s.job.active_pods",
	    metric.WithDescription("The number of pending and actively running pods for a job"),
	    metric.WithUnit("{pod}"),
	)
	if err != nil {
	    return JobActivePods{inst: noop.Int64UpDownCounter{}}, err
	}
	return JobActivePods{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (JobActivePods) Name() string {
	return "k8s.job.active_pods"
}

// Unit returns the semantic convention unit of the instrument
func (JobActivePods) Unit() string {
	return "{pod}"
}

// Description returns the semantic convention description of the instrument
func (JobActivePods) Description() string {
	return "The number of pending and actively running pods for a job"
}

// Add adds incr to the existing count.
//
// This metric aligns with the `active` field of the
// [K8s JobStatus].
//
// This metric SHOULD, at a minimum, be reported against a
// [`k8s.job`] resource.
//
// [K8s JobStatus]: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#jobstatus-v1-batch
// [`k8s.job`]: ../resource/k8s.md#job
func (m JobActivePods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Add(ctx, incr)
	} else {
		m.inst.Add(ctx, incr, metric.WithAttributes(attrs...))
	}
}

// JobDesiredSuccessfulPods is an instrument used to record metric values
// conforming to the "k8s.job.desired_successful_pods" semantic conventions. It
// represents the desired number of successfully finished pods the job should be
// run with.
type JobDesiredSuccessfulPods struct {
	inst metric.Int64UpDownCounter
}

// NewJobDesiredSuccessfulPods returns a new JobDesiredSuccessfulPods instrument.
func NewJobDesiredSuccessfulPods(m metric.Meter) (JobDesiredSuccessfulPods, error) {
	i, err := m.Int64UpDownCounter(
	    "k8s.job.desired_successful_pods",
	    metric.WithDescription("The desired number of successfully finished pods the job should be run with"),
	    metric.WithUnit("{pod}"),
	)
	if err != nil {
	    return JobDesiredSuccessfulPods{inst: noop.Int64UpDownCounter{}}, err
	}
	return JobDesiredSuccessfulPods{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (JobDesiredSuccessfulPods) Name() string {
	return "k8s.job.desired_successful_pods"
}

// Unit returns the semantic convention unit of the instrument
func (JobDesiredSuccessfulPods) Unit() string {
	return "{pod}"
}

// Description returns the semantic convention description of the instrument
func (JobDesiredSuccessfulPods) Description() string {
	return "The desired number of successfully finished pods the job should be run with"
}

// Add adds incr to the existing count.
//
// This metric aligns with the `completions` field of the
// [K8s JobSpec].
//
// This metric SHOULD, at a minimum, be reported against a
// [`k8s.job`] resource.
//
// [K8s JobSpec]: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#jobspec-v1-batch
// [`k8s.job`]: ../resource/k8s.md#job
func (m JobDesiredSuccessfulPods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Add(ctx, incr)
	} else {
		m.inst.Add(ctx, incr, metric.WithAttributes(attrs...))
	}
}

// JobFailedPods is an instrument used to record metric values conforming to the
// "k8s.job.failed_pods" semantic conventions. It represents the number of pods
// which reached phase Failed for a job.
type JobFailedPods struct {
	inst metric.Int64UpDownCounter
}

// NewJobFailedPods returns a new JobFailedPods instrument.
func NewJobFailedPods(m metric.Meter) (JobFailedPods, error) {
	i, err := m.Int64UpDownCounter(
	    "k8s.job.failed_pods",
	    metric.WithDescription("The number of pods which reached phase Failed for a job"),
	    metric.WithUnit("{pod}"),
	)
	if err != nil {
	    return JobFailedPods{inst: noop.Int64UpDownCounter{}}, err
	}
	return JobFailedPods{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (JobFailedPods) Name() string {
	return "k8s.job.failed_pods"
}

// Unit returns the semantic convention unit of the instrument
func (JobFailedPods) Unit() string {
	return "{pod}"
}

// Description returns the semantic convention description of the instrument
func (JobFailedPods) Description() string {
	return "The number of pods which reached phase Failed for a job"
}

// Add adds incr to the existing count.
//
// This metric aligns with the `failed` field of the
// [K8s JobStatus].
//
// This metric SHOULD, at a minimum, be reported against a
// [`k8s.job`] resource.
//
// [K8s JobStatus]: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#jobstatus-v1-batch
// [`k8s.job`]: ../resource/k8s.md#job
func (m JobFailedPods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Add(ctx, incr)
	} else {
		m.inst.Add(ctx, incr, metric.WithAttributes(attrs...))
	}
}

// JobMaxParallelPods is an instrument used to record metric values conforming to
// the "k8s.job.max_parallel_pods" semantic conventions. It represents the max
// desired number of pods the job should run at any given time.
type JobMaxParallelPods struct {
	inst metric.Int64UpDownCounter
}

// NewJobMaxParallelPods returns a new JobMaxParallelPods instrument.
func NewJobMaxParallelPods(m metric.Meter) (JobMaxParallelPods, error) {
	i, err := m.Int64UpDownCounter(
	    "k8s.job.max_parallel_pods",
	    metric.WithDescription("The max desired number of pods the job should run at any given time"),
	    metric.WithUnit("{pod}"),
	)
	if err != nil {
	    return JobMaxParallelPods{inst: noop.Int64UpDownCounter{}}, err
	}
	return JobMaxParallelPods{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (JobMaxParallelPods) Name() string {
	return "k8s.job.max_parallel_pods"
}

// Unit returns the semantic convention unit of the instrument
func (JobMaxParallelPods) Unit() string {
	return "{pod}"
}

// Description returns the semantic convention description of the instrument
func (JobMaxParallelPods) Description() string {
	return "The max desired number of pods the job should run at any given time"
}

// Add adds incr to the existing count.
//
// This metric aligns with the `parallelism` field of the
// [K8s JobSpec].
//
// This metric SHOULD, at a minimum, be reported against a
// [`k8s.job`] resource.
//
// [K8s JobSpec]: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#jobspec-v1-batch
// [`k8s.job`]: ../resource/k8s.md#job
func (m JobMaxParallelPods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Add(ctx, incr)
	} else {
		m.inst.Add(ctx, incr, metric.WithAttributes(attrs...))
	}
}

// JobSuccessfulPods is an instrument used to record metric values conforming to
// the "k8s.job.successful_pods" semantic conventions. It represents the number
// of pods which reached phase Succeeded for a job.
type JobSuccessfulPods struct {
	inst metric.Int64UpDownCounter
}

// NewJobSuccessfulPods returns a new JobSuccessfulPods instrument.
func NewJobSuccessfulPods(m metric.Meter) (JobSuccessfulPods, error) {
	i, err := m.Int64UpDownCounter(
	    "k8s.job.successful_pods",
	    metric.WithDescription("The number of pods which reached phase Succeeded for a job"),
	    metric.WithUnit("{pod}"),
	)
	if err != nil {
	    return JobSuccessfulPods{inst: noop.Int64UpDownCounter{}}, err
	}
	return JobSuccessfulPods{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (JobSuccessfulPods) Name() string {
	return "k8s.job.successful_pods"
}

// Unit returns the semantic convention unit of the instrument
func (JobSuccessfulPods) Unit() string {
	return "{pod}"
}

// Description returns the semantic convention description of the instrument
func (JobSuccessfulPods) Description() string {
	return "The number of pods which reached phase Succeeded for a job"
}

// Add adds incr to the existing count.
//
// This metric aligns with the `succeeded` field of the
// [K8s JobStatus].
//
// This metric SHOULD, at a minimum, be reported against a
// [`k8s.job`] resource.
//
// [K8s JobStatus]: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#jobstatus-v1-batch
// [`k8s.job`]: ../resource/k8s.md#job
func (m JobSuccessfulPods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Add(ctx, incr)
	} else {
		m.inst.Add(ctx, incr, metric.WithAttributes(attrs...))
	}
}

// NamespacePhase is an instrument used to record metric values conforming to the
// "k8s.namespace.phase" semantic conventions. It represents the describes number
// of K8s namespaces that are currently in a given phase.
type NamespacePhase struct {
	inst metric.Int64UpDownCounter
}

// NewNamespacePhase returns a new NamespacePhase instrument.
func NewNamespacePhase(m metric.Meter) (NamespacePhase, error) {
	i, err := m.Int64UpDownCounter(
	    "k8s.namespace.phase",
	    metric.WithDescription("Describes number of K8s namespaces that are currently in a given phase."),
	    metric.WithUnit("{namespace}"),
	)
	if err != nil {
	    return NamespacePhase{inst: noop.Int64UpDownCounter{}}, err
	}
	return NamespacePhase{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (NamespacePhase) Name() string {
	return "k8s.namespace.phase"
}

// Unit returns the semantic convention unit of the instrument
func (NamespacePhase) Unit() string {
	return "{namespace}"
}

// Description returns the semantic convention description of the instrument
func (NamespacePhase) Description() string {
	return "Describes number of K8s namespaces that are currently in a given phase."
}

// Add adds incr to the existing count.
//
// The namespacePhase is the the phase of the K8s namespace.
//
// This metric SHOULD, at a minimum, be reported against a
// [`k8s.namespace`] resource.
//
// [`k8s.namespace`]: ../resource/k8s.md#namespace
func (m NamespacePhase) Add(
	ctx context.Context,
	incr int64,
	namespacePhase NamespacePhaseAttr,
	attrs ...attribute.KeyValue,
) {
	m.inst.Add(
		ctx,
		incr,
		metric.WithAttributes(
			append(
				attrs,
				attribute.String("k8s.namespace.phase", string(namespacePhase)),
			)...,
		),
	)
}

// NodeCPUTime is an instrument used to record metric values conforming to the
// "k8s.node.cpu.time" semantic conventions. It represents the total CPU time
// consumed.
type NodeCPUTime struct {
	inst metric.Float64Counter
}

// NewNodeCPUTime returns a new NodeCPUTime instrument.
func NewNodeCPUTime(m metric.Meter) (NodeCPUTime, error) {
	i, err := m.Float64Counter(
	    "k8s.node.cpu.time",
	    metric.WithDescription("Total CPU time consumed"),
	    metric.WithUnit("s"),
	)
	if err != nil {
	    return NodeCPUTime{inst: noop.Float64Counter{}}, err
	}
	return NodeCPUTime{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (NodeCPUTime) Name() string {
	return "k8s.node.cpu.time"
}

// Unit returns the semantic convention unit of the instrument
func (NodeCPUTime) Unit() string {
	return "s"
}

// Description returns the semantic convention description of the instrument
func (NodeCPUTime) Description() string {
	return "Total CPU time consumed"
}

// Add adds incr to the existing count.
//
// Total CPU time consumed by the specific Node on all available CPU cores
func (m NodeCPUTime) Add(ctx context.Context, incr float64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Add(ctx, incr)
	} else {
		m.inst.Add(ctx, incr, metric.WithAttributes(attrs...))
	}
}

// NodeCPUUsage is an instrument used to record metric values conforming to the
// "k8s.node.cpu.usage" semantic conventions. It represents the node's CPU usage,
// measured in cpus. Range from 0 to the number of allocatable CPUs.
type NodeCPUUsage struct {
	inst metric.Int64Gauge
}

// NewNodeCPUUsage returns a new NodeCPUUsage instrument.
func NewNodeCPUUsage(m metric.Meter) (NodeCPUUsage, error) {
	i, err := m.Int64Gauge(
	    "k8s.node.cpu.usage",
	    metric.WithDescription("Node's CPU usage, measured in cpus. Range from 0 to the number of allocatable CPUs"),
	    metric.WithUnit("{cpu}"),
	)
	if err != nil {
	    return NodeCPUUsage{inst: noop.Int64Gauge{}}, err
	}
	return NodeCPUUsage{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (NodeCPUUsage) Name() string {
	return "k8s.node.cpu.usage"
}

// Unit returns the semantic convention unit of the instrument
func (NodeCPUUsage) Unit() string {
	return "{cpu}"
}

// Description returns the semantic convention description of the instrument
func (NodeCPUUsage) Description() string {
	return "Node's CPU usage, measured in cpus. Range from 0 to the number of allocatable CPUs"
}

// Record records val to the current distribution.
//
// CPU usage of the specific Node on all available CPU cores, averaged over the
// sample window
func (m NodeCPUUsage) Record(ctx context.Context, val int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Record(ctx, val)
	} else {
		m.inst.Record(ctx, val, metric.WithAttributes(attrs...))
	}
}

// NodeMemoryUsage is an instrument used to record metric values conforming to
// the "k8s.node.memory.usage" semantic conventions. It represents the memory
// usage of the Node.
type NodeMemoryUsage struct {
	inst metric.Int64Gauge
}

// NewNodeMemoryUsage returns a new NodeMemoryUsage instrument.
func NewNodeMemoryUsage(m metric.Meter) (NodeMemoryUsage, error) {
	i, err := m.Int64Gauge(
	    "k8s.node.memory.usage",
	    metric.WithDescription("Memory usage of the Node"),
	    metric.WithUnit("By"),
	)
	if err != nil {
	    return NodeMemoryUsage{inst: noop.Int64Gauge{}}, err
	}
	return NodeMemoryUsage{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (NodeMemoryUsage) Name() string {
	return "k8s.node.memory.usage"
}

// Unit returns the semantic convention unit of the instrument
func (NodeMemoryUsage) Unit() string {
	return "By"
}

// Description returns the semantic convention description of the instrument
func (NodeMemoryUsage) Description() string {
	return "Memory usage of the Node"
}

// Record records val to the current distribution.
//
// Total memory usage of the Node
func (m NodeMemoryUsage) Record(ctx context.Context, val int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Record(ctx, val)
	} else {
		m.inst.Record(ctx, val, metric.WithAttributes(attrs...))
	}
}

// NodeNetworkErrors is an instrument used to record metric values conforming to
// the "k8s.node.network.errors" semantic conventions. It represents the node
// network errors.
type NodeNetworkErrors struct {
	inst metric.Int64Counter
}

// NewNodeNetworkErrors returns a new NodeNetworkErrors instrument.
func NewNodeNetworkErrors(m metric.Meter) (NodeNetworkErrors, error) {
	i, err := m.Int64Counter(
	    "k8s.node.network.errors",
	    metric.WithDescription("Node network errors"),
	    metric.WithUnit("{error}"),
	)
	if err != nil {
	    return NodeNetworkErrors{inst: noop.Int64Counter{}}, err
	}
	return NodeNetworkErrors{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (NodeNetworkErrors) Name() string {
	return "k8s.node.network.errors"
}

// Unit returns the semantic convention unit of the instrument
func (NodeNetworkErrors) Unit() string {
	return "{error}"
}

// Description returns the semantic convention description of the instrument
func (NodeNetworkErrors) Description() string {
	return "Node network errors"
}

// Add adds incr to the existing count.
//
// All additional attrs passed are included in the recorded value.
func (m NodeNetworkErrors) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	m.inst.Add(
		ctx,
		incr,
		metric.WithAttributes(
			attrs...,
		),
	)
}

// AttrNetworkInterfaceName returns an optional attribute for the
// "network.interface.name" semantic convention. It represents the network
// interface name.
func (NodeNetworkErrors) AttrNetworkInterfaceName(val string) attribute.KeyValue {
	return attribute.String("network.interface.name", val)
}

// AttrNetworkIODirection returns an optional attribute for the
// "network.io.direction" semantic convention. It represents the network IO
// operation direction.
func (NodeNetworkErrors) AttrNetworkIODirection(val NetworkIODirectionAttr) attribute.KeyValue {
	return attribute.String("network.io.direction", string(val))
}

// NodeNetworkIO is an instrument used to record metric values conforming to the
// "k8s.node.network.io" semantic conventions. It represents the network bytes
// for the Node.
type NodeNetworkIO struct {
	inst metric.Int64Counter
}

// NewNodeNetworkIO returns a new NodeNetworkIO instrument.
func NewNodeNetworkIO(m metric.Meter) (NodeNetworkIO, error) {
	i, err := m.Int64Counter(
	    "k8s.node.network.io",
	    metric.WithDescription("Network bytes for the Node"),
	    metric.WithUnit("By"),
	)
	if err != nil {
	    return NodeNetworkIO{inst: noop.Int64Counter{}}, err
	}
	return NodeNetworkIO{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (NodeNetworkIO) Name() string {
	return "k8s.node.network.io"
}

// Unit returns the semantic convention unit of the instrument
func (NodeNetworkIO) Unit() string {
	return "By"
}

// Description returns the semantic convention description of the instrument
func (NodeNetworkIO) Description() string {
	return "Network bytes for the Node"
}

// Add adds incr to the existing count.
//
// All additional attrs passed are included in the recorded value.
func (m NodeNetworkIO) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	m.inst.Add(
		ctx,
		incr,
		metric.WithAttributes(
			attrs...,
		),
	)
}

// AttrNetworkInterfaceName returns an optional attribute for the
// "network.interface.name" semantic convention. It represents the network
// interface name.
func (NodeNetworkIO) AttrNetworkInterfaceName(val string) attribute.KeyValue {
	return attribute.String("network.interface.name", val)
}

// AttrNetworkIODirection returns an optional attribute for the
// "network.io.direction" semantic convention. It represents the network IO
// operation direction.
func (NodeNetworkIO) AttrNetworkIODirection(val NetworkIODirectionAttr) attribute.KeyValue {
	return attribute.String("network.io.direction", string(val))
}

// NodeUptime is an instrument used to record metric values conforming to the
// "k8s.node.uptime" semantic conventions. It represents the time the Node has
// been running.
type NodeUptime struct {
	inst metric.Float64Gauge
}

// NewNodeUptime returns a new NodeUptime instrument.
func NewNodeUptime(m metric.Meter) (NodeUptime, error) {
	i, err := m.Float64Gauge(
	    "k8s.node.uptime",
	    metric.WithDescription("The time the Node has been running"),
	    metric.WithUnit("s"),
	)
	if err != nil {
	    return NodeUptime{inst: noop.Float64Gauge{}}, err
	}
	return NodeUptime{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (NodeUptime) Name() string {
	return "k8s.node.uptime"
}

// Unit returns the semantic convention unit of the instrument
func (NodeUptime) Unit() string {
	return "s"
}

// Description returns the semantic convention description of the instrument
func (NodeUptime) Description() string {
	return "The time the Node has been running"
}

// Record records val to the current distribution.
//
// Instrumentations SHOULD use a gauge with type `double` and measure uptime in
// seconds as a floating point number with the highest precision available.
// The actual accuracy would depend on the instrumentation and operating system.
func (m NodeUptime) Record(ctx context.Context, val float64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Record(ctx, val)
	} else {
		m.inst.Record(ctx, val, metric.WithAttributes(attrs...))
	}
}

// PodCPUTime is an instrument used to record metric values conforming to the
// "k8s.pod.cpu.time" semantic conventions. It represents the total CPU time
// consumed.
type PodCPUTime struct {
	inst metric.Float64Counter
}

// NewPodCPUTime returns a new PodCPUTime instrument.
func NewPodCPUTime(m metric.Meter) (PodCPUTime, error) {
	i, err := m.Float64Counter(
	    "k8s.pod.cpu.time",
	    metric.WithDescription("Total CPU time consumed"),
	    metric.WithUnit("s"),
	)
	if err != nil {
	    return PodCPUTime{inst: noop.Float64Counter{}}, err
	}
	return PodCPUTime{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (PodCPUTime) Name() string {
	return "k8s.pod.cpu.time"
}

// Unit returns the semantic convention unit of the instrument
func (PodCPUTime) Unit() string {
	return "s"
}

// Description returns the semantic convention description of the instrument
func (PodCPUTime) Description() string {
	return "Total CPU time consumed"
}

// Add adds incr to the existing count.
//
// Total CPU time consumed by the specific Pod on all available CPU cores
func (m PodCPUTime) Add(ctx context.Context, incr float64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Add(ctx, incr)
	} else {
		m.inst.Add(ctx, incr, metric.WithAttributes(attrs...))
	}
}

// PodCPUUsage is an instrument used to record metric values conforming to the
// "k8s.pod.cpu.usage" semantic conventions. It represents the pod's CPU usage,
// measured in cpus. Range from 0 to the number of allocatable CPUs.
type PodCPUUsage struct {
	inst metric.Int64Gauge
}

// NewPodCPUUsage returns a new PodCPUUsage instrument.
func NewPodCPUUsage(m metric.Meter) (PodCPUUsage, error) {
	i, err := m.Int64Gauge(
	    "k8s.pod.cpu.usage",
	    metric.WithDescription("Pod's CPU usage, measured in cpus. Range from 0 to the number of allocatable CPUs"),
	    metric.WithUnit("{cpu}"),
	)
	if err != nil {
	    return PodCPUUsage{inst: noop.Int64Gauge{}}, err
	}
	return PodCPUUsage{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (PodCPUUsage) Name() string {
	return "k8s.pod.cpu.usage"
}

// Unit returns the semantic convention unit of the instrument
func (PodCPUUsage) Unit() string {
	return "{cpu}"
}

// Description returns the semantic convention description of the instrument
func (PodCPUUsage) Description() string {
	return "Pod's CPU usage, measured in cpus. Range from 0 to the number of allocatable CPUs"
}

// Record records val to the current distribution.
//
// CPU usage of the specific Pod on all available CPU cores, averaged over the
// sample window
func (m PodCPUUsage) Record(ctx context.Context, val int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Record(ctx, val)
	} else {
		m.inst.Record(ctx, val, metric.WithAttributes(attrs...))
	}
}

// PodMemoryUsage is an instrument used to record metric values conforming to the
// "k8s.pod.memory.usage" semantic conventions. It represents the memory usage of
// the Pod.
type PodMemoryUsage struct {
	inst metric.Int64Gauge
}

// NewPodMemoryUsage returns a new PodMemoryUsage instrument.
func NewPodMemoryUsage(m metric.Meter) (PodMemoryUsage, error) {
	i, err := m.Int64Gauge(
	    "k8s.pod.memory.usage",
	    metric.WithDescription("Memory usage of the Pod"),
	    metric.WithUnit("By"),
	)
	if err != nil {
	    return PodMemoryUsage{inst: noop.Int64Gauge{}}, err
	}
	return PodMemoryUsage{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (PodMemoryUsage) Name() string {
	return "k8s.pod.memory.usage"
}

// Unit returns the semantic convention unit of the instrument
func (PodMemoryUsage) Unit() string {
	return "By"
}

// Description returns the semantic convention description of the instrument
func (PodMemoryUsage) Description() string {
	return "Memory usage of the Pod"
}

// Record records val to the current distribution.
//
// Total memory usage of the Pod
func (m PodMemoryUsage) Record(ctx context.Context, val int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Record(ctx, val)
	} else {
		m.inst.Record(ctx, val, metric.WithAttributes(attrs...))
	}
}

// PodNetworkErrors is an instrument used to record metric values conforming to
// the "k8s.pod.network.errors" semantic conventions. It represents the pod
// network errors.
type PodNetworkErrors struct {
	inst metric.Int64Counter
}

// NewPodNetworkErrors returns a new PodNetworkErrors instrument.
func NewPodNetworkErrors(m metric.Meter) (PodNetworkErrors, error) {
	i, err := m.Int64Counter(
	    "k8s.pod.network.errors",
	    metric.WithDescription("Pod network errors"),
	    metric.WithUnit("{error}"),
	)
	if err != nil {
	    return PodNetworkErrors{inst: noop.Int64Counter{}}, err
	}
	return PodNetworkErrors{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (PodNetworkErrors) Name() string {
	return "k8s.pod.network.errors"
}

// Unit returns the semantic convention unit of the instrument
func (PodNetworkErrors) Unit() string {
	return "{error}"
}

// Description returns the semantic convention description of the instrument
func (PodNetworkErrors) Description() string {
	return "Pod network errors"
}

// Add adds incr to the existing count.
//
// All additional attrs passed are included in the recorded value.
func (m PodNetworkErrors) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	m.inst.Add(
		ctx,
		incr,
		metric.WithAttributes(
			attrs...,
		),
	)
}

// AttrNetworkInterfaceName returns an optional attribute for the
// "network.interface.name" semantic convention. It represents the network
// interface name.
func (PodNetworkErrors) AttrNetworkInterfaceName(val string) attribute.KeyValue {
	return attribute.String("network.interface.name", val)
}

// AttrNetworkIODirection returns an optional attribute for the
// "network.io.direction" semantic convention. It represents the network IO
// operation direction.
func (PodNetworkErrors) AttrNetworkIODirection(val NetworkIODirectionAttr) attribute.KeyValue {
	return attribute.String("network.io.direction", string(val))
}

// PodNetworkIO is an instrument used to record metric values conforming to the
// "k8s.pod.network.io" semantic conventions. It represents the network bytes for
// the Pod.
type PodNetworkIO struct {
	inst metric.Int64Counter
}

// NewPodNetworkIO returns a new PodNetworkIO instrument.
func NewPodNetworkIO(m metric.Meter) (PodNetworkIO, error) {
	i, err := m.Int64Counter(
	    "k8s.pod.network.io",
	    metric.WithDescription("Network bytes for the Pod"),
	    metric.WithUnit("By"),
	)
	if err != nil {
	    return PodNetworkIO{inst: noop.Int64Counter{}}, err
	}
	return PodNetworkIO{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (PodNetworkIO) Name() string {
	return "k8s.pod.network.io"
}

// Unit returns the semantic convention unit of the instrument
func (PodNetworkIO) Unit() string {
	return "By"
}

// Description returns the semantic convention description of the instrument
func (PodNetworkIO) Description() string {
	return "Network bytes for the Pod"
}

// Add adds incr to the existing count.
//
// All additional attrs passed are included in the recorded value.
func (m PodNetworkIO) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	m.inst.Add(
		ctx,
		incr,
		metric.WithAttributes(
			attrs...,
		),
	)
}

// AttrNetworkInterfaceName returns an optional attribute for the
// "network.interface.name" semantic convention. It represents the network
// interface name.
func (PodNetworkIO) AttrNetworkInterfaceName(val string) attribute.KeyValue {
	return attribute.String("network.interface.name", val)
}

// AttrNetworkIODirection returns an optional attribute for the
// "network.io.direction" semantic convention. It represents the network IO
// operation direction.
func (PodNetworkIO) AttrNetworkIODirection(val NetworkIODirectionAttr) attribute.KeyValue {
	return attribute.String("network.io.direction", string(val))
}

// PodUptime is an instrument used to record metric values conforming to the
// "k8s.pod.uptime" semantic conventions. It represents the time the Pod has been
// running.
type PodUptime struct {
	inst metric.Float64Gauge
}

// NewPodUptime returns a new PodUptime instrument.
func NewPodUptime(m metric.Meter) (PodUptime, error) {
	i, err := m.Float64Gauge(
	    "k8s.pod.uptime",
	    metric.WithDescription("The time the Pod has been running"),
	    metric.WithUnit("s"),
	)
	if err != nil {
	    return PodUptime{inst: noop.Float64Gauge{}}, err
	}
	return PodUptime{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (PodUptime) Name() string {
	return "k8s.pod.uptime"
}

// Unit returns the semantic convention unit of the instrument
func (PodUptime) Unit() string {
	return "s"
}

// Description returns the semantic convention description of the instrument
func (PodUptime) Description() string {
	return "The time the Pod has been running"
}

// Record records val to the current distribution.
//
// Instrumentations SHOULD use a gauge with type `double` and measure uptime in
// seconds as a floating point number with the highest precision available.
// The actual accuracy would depend on the instrumentation and operating system.
func (m PodUptime) Record(ctx context.Context, val float64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Record(ctx, val)
	} else {
		m.inst.Record(ctx, val, metric.WithAttributes(attrs...))
	}
}

// ReplicaSetAvailablePods is an instrument used to record metric values
// conforming to the "k8s.replicaset.available_pods" semantic conventions. It
// represents the total number of available replica pods (ready for at least
// minReadySeconds) targeted by this replicaset.
type ReplicaSetAvailablePods struct {
	inst metric.Int64UpDownCounter
}

// NewReplicaSetAvailablePods returns a new ReplicaSetAvailablePods instrument.
func NewReplicaSetAvailablePods(m metric.Meter) (ReplicaSetAvailablePods, error) {
	i, err := m.Int64UpDownCounter(
	    "k8s.replicaset.available_pods",
	    metric.WithDescription("Total number of available replica pods (ready for at least minReadySeconds) targeted by this replicaset"),
	    metric.WithUnit("{pod}"),
	)
	if err != nil {
	    return ReplicaSetAvailablePods{inst: noop.Int64UpDownCounter{}}, err
	}
	return ReplicaSetAvailablePods{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (ReplicaSetAvailablePods) Name() string {
	return "k8s.replicaset.available_pods"
}

// Unit returns the semantic convention unit of the instrument
func (ReplicaSetAvailablePods) Unit() string {
	return "{pod}"
}

// Description returns the semantic convention description of the instrument
func (ReplicaSetAvailablePods) Description() string {
	return "Total number of available replica pods (ready for at least minReadySeconds) targeted by this replicaset"
}

// Add adds incr to the existing count.
//
// This metric aligns with the `availableReplicas` field of the
// [K8s ReplicaSetStatus].
//
// This metric SHOULD, at a minimum, be reported against a
// [`k8s.replicaset`] resource.
//
// [K8s ReplicaSetStatus]: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#replicasetstatus-v1-apps
// [`k8s.replicaset`]: ../resource/k8s.md#replicaset
func (m ReplicaSetAvailablePods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Add(ctx, incr)
	} else {
		m.inst.Add(ctx, incr, metric.WithAttributes(attrs...))
	}
}

// ReplicaSetDesiredPods is an instrument used to record metric values conforming
// to the "k8s.replicaset.desired_pods" semantic conventions. It represents the
// number of desired replica pods in this replicaset.
type ReplicaSetDesiredPods struct {
	inst metric.Int64UpDownCounter
}

// NewReplicaSetDesiredPods returns a new ReplicaSetDesiredPods instrument.
func NewReplicaSetDesiredPods(m metric.Meter) (ReplicaSetDesiredPods, error) {
	i, err := m.Int64UpDownCounter(
	    "k8s.replicaset.desired_pods",
	    metric.WithDescription("Number of desired replica pods in this replicaset"),
	    metric.WithUnit("{pod}"),
	)
	if err != nil {
	    return ReplicaSetDesiredPods{inst: noop.Int64UpDownCounter{}}, err
	}
	return ReplicaSetDesiredPods{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (ReplicaSetDesiredPods) Name() string {
	return "k8s.replicaset.desired_pods"
}

// Unit returns the semantic convention unit of the instrument
func (ReplicaSetDesiredPods) Unit() string {
	return "{pod}"
}

// Description returns the semantic convention description of the instrument
func (ReplicaSetDesiredPods) Description() string {
	return "Number of desired replica pods in this replicaset"
}

// Add adds incr to the existing count.
//
// This metric aligns with the `replicas` field of the
// [K8s ReplicaSetSpec].
//
// This metric SHOULD, at a minimum, be reported against a
// [`k8s.replicaset`] resource.
//
// [K8s ReplicaSetSpec]: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#replicasetspec-v1-apps
// [`k8s.replicaset`]: ../resource/k8s.md#replicaset
func (m ReplicaSetDesiredPods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Add(ctx, incr)
	} else {
		m.inst.Add(ctx, incr, metric.WithAttributes(attrs...))
	}
}

// ReplicationControllerAvailablePods is an instrument used to record metric
// values conforming to the "k8s.replicationcontroller.available_pods" semantic
// conventions. It represents the total number of available replica pods (ready
// for at least minReadySeconds) targeted by this replication controller.
type ReplicationControllerAvailablePods struct {
	inst metric.Int64UpDownCounter
}

// NewReplicationControllerAvailablePods returns a new
// ReplicationControllerAvailablePods instrument.
func NewReplicationControllerAvailablePods(m metric.Meter) (ReplicationControllerAvailablePods, error) {
	i, err := m.Int64UpDownCounter(
	    "k8s.replicationcontroller.available_pods",
	    metric.WithDescription("Total number of available replica pods (ready for at least minReadySeconds) targeted by this replication controller"),
	    metric.WithUnit("{pod}"),
	)
	if err != nil {
	    return ReplicationControllerAvailablePods{inst: noop.Int64UpDownCounter{}}, err
	}
	return ReplicationControllerAvailablePods{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (ReplicationControllerAvailablePods) Name() string {
	return "k8s.replicationcontroller.available_pods"
}

// Unit returns the semantic convention unit of the instrument
func (ReplicationControllerAvailablePods) Unit() string {
	return "{pod}"
}

// Description returns the semantic convention description of the instrument
func (ReplicationControllerAvailablePods) Description() string {
	return "Total number of available replica pods (ready for at least minReadySeconds) targeted by this replication controller"
}

// Add adds incr to the existing count.
//
// This metric aligns with the `availableReplicas` field of the
// [K8s ReplicationControllerStatus]
//
// This metric SHOULD, at a minimum, be reported against a
// [`k8s.replicationcontroller`] resource.
//
// [K8s ReplicationControllerStatus]: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#replicationcontrollerstatus-v1-core
// [`k8s.replicationcontroller`]: ../resource/k8s.md#replicationcontroller
func (m ReplicationControllerAvailablePods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Add(ctx, incr)
	} else {
		m.inst.Add(ctx, incr, metric.WithAttributes(attrs...))
	}
}

// ReplicationControllerDesiredPods is an instrument used to record metric values
// conforming to the "k8s.replicationcontroller.desired_pods" semantic
// conventions. It represents the number of desired replica pods in this
// replication controller.
type ReplicationControllerDesiredPods struct {
	inst metric.Int64UpDownCounter
}

// NewReplicationControllerDesiredPods returns a new
// ReplicationControllerDesiredPods instrument.
func NewReplicationControllerDesiredPods(m metric.Meter) (ReplicationControllerDesiredPods, error) {
	i, err := m.Int64UpDownCounter(
	    "k8s.replicationcontroller.desired_pods",
	    metric.WithDescription("Number of desired replica pods in this replication controller"),
	    metric.WithUnit("{pod}"),
	)
	if err != nil {
	    return ReplicationControllerDesiredPods{inst: noop.Int64UpDownCounter{}}, err
	}
	return ReplicationControllerDesiredPods{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (ReplicationControllerDesiredPods) Name() string {
	return "k8s.replicationcontroller.desired_pods"
}

// Unit returns the semantic convention unit of the instrument
func (ReplicationControllerDesiredPods) Unit() string {
	return "{pod}"
}

// Description returns the semantic convention description of the instrument
func (ReplicationControllerDesiredPods) Description() string {
	return "Number of desired replica pods in this replication controller"
}

// Add adds incr to the existing count.
//
// This metric aligns with the `replicas` field of the
// [K8s ReplicationControllerSpec]
//
// This metric SHOULD, at a minimum, be reported against a
// [`k8s.replicationcontroller`] resource.
//
// [K8s ReplicationControllerSpec]: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#replicationcontrollerspec-v1-core
// [`k8s.replicationcontroller`]: ../resource/k8s.md#replicationcontroller
func (m ReplicationControllerDesiredPods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Add(ctx, incr)
	} else {
		m.inst.Add(ctx, incr, metric.WithAttributes(attrs...))
	}
}

// StatefulSetCurrentPods is an instrument used to record metric values
// conforming to the "k8s.statefulset.current_pods" semantic conventions. It
// represents the number of replica pods created by the statefulset controller
// from the statefulset version indicated by currentRevision.
type StatefulSetCurrentPods struct {
	inst metric.Int64UpDownCounter
}

// NewStatefulSetCurrentPods returns a new StatefulSetCurrentPods instrument.
func NewStatefulSetCurrentPods(m metric.Meter) (StatefulSetCurrentPods, error) {
	i, err := m.Int64UpDownCounter(
	    "k8s.statefulset.current_pods",
	    metric.WithDescription("The number of replica pods created by the statefulset controller from the statefulset version indicated by currentRevision"),
	    metric.WithUnit("{pod}"),
	)
	if err != nil {
	    return StatefulSetCurrentPods{inst: noop.Int64UpDownCounter{}}, err
	}
	return StatefulSetCurrentPods{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (StatefulSetCurrentPods) Name() string {
	return "k8s.statefulset.current_pods"
}

// Unit returns the semantic convention unit of the instrument
func (StatefulSetCurrentPods) Unit() string {
	return "{pod}"
}

// Description returns the semantic convention description of the instrument
func (StatefulSetCurrentPods) Description() string {
	return "The number of replica pods created by the statefulset controller from the statefulset version indicated by currentRevision"
}

// Add adds incr to the existing count.
//
// This metric aligns with the `currentReplicas` field of the
// [K8s StatefulSetStatus].
//
// This metric SHOULD, at a minimum, be reported against a
// [`k8s.statefulset`] resource.
//
// [K8s StatefulSetStatus]: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#statefulsetstatus-v1-apps
// [`k8s.statefulset`]: ../resource/k8s.md#statefulset
func (m StatefulSetCurrentPods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Add(ctx, incr)
	} else {
		m.inst.Add(ctx, incr, metric.WithAttributes(attrs...))
	}
}

// StatefulSetDesiredPods is an instrument used to record metric values
// conforming to the "k8s.statefulset.desired_pods" semantic conventions. It
// represents the number of desired replica pods in this statefulset.
type StatefulSetDesiredPods struct {
	inst metric.Int64UpDownCounter
}

// NewStatefulSetDesiredPods returns a new StatefulSetDesiredPods instrument.
func NewStatefulSetDesiredPods(m metric.Meter) (StatefulSetDesiredPods, error) {
	i, err := m.Int64UpDownCounter(
	    "k8s.statefulset.desired_pods",
	    metric.WithDescription("Number of desired replica pods in this statefulset"),
	    metric.WithUnit("{pod}"),
	)
	if err != nil {
	    return StatefulSetDesiredPods{inst: noop.Int64UpDownCounter{}}, err
	}
	return StatefulSetDesiredPods{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (StatefulSetDesiredPods) Name() string {
	return "k8s.statefulset.desired_pods"
}

// Unit returns the semantic convention unit of the instrument
func (StatefulSetDesiredPods) Unit() string {
	return "{pod}"
}

// Description returns the semantic convention description of the instrument
func (StatefulSetDesiredPods) Description() string {
	return "Number of desired replica pods in this statefulset"
}

// Add adds incr to the existing count.
//
// This metric aligns with the `replicas` field of the
// [K8s StatefulSetSpec].
//
// This metric SHOULD, at a minimum, be reported against a
// [`k8s.statefulset`] resource.
//
// [K8s StatefulSetSpec]: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#statefulsetspec-v1-apps
// [`k8s.statefulset`]: ../resource/k8s.md#statefulset
func (m StatefulSetDesiredPods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Add(ctx, incr)
	} else {
		m.inst.Add(ctx, incr, metric.WithAttributes(attrs...))
	}
}

// StatefulSetReadyPods is an instrument used to record metric values conforming
// to the "k8s.statefulset.ready_pods" semantic conventions. It represents the
// number of replica pods created for this statefulset with a Ready Condition.
type StatefulSetReadyPods struct {
	inst metric.Int64UpDownCounter
}

// NewStatefulSetReadyPods returns a new StatefulSetReadyPods instrument.
func NewStatefulSetReadyPods(m metric.Meter) (StatefulSetReadyPods, error) {
	i, err := m.Int64UpDownCounter(
	    "k8s.statefulset.ready_pods",
	    metric.WithDescription("The number of replica pods created for this statefulset with a Ready Condition"),
	    metric.WithUnit("{pod}"),
	)
	if err != nil {
	    return StatefulSetReadyPods{inst: noop.Int64UpDownCounter{}}, err
	}
	return StatefulSetReadyPods{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (StatefulSetReadyPods) Name() string {
	return "k8s.statefulset.ready_pods"
}

// Unit returns the semantic convention unit of the instrument
func (StatefulSetReadyPods) Unit() string {
	return "{pod}"
}

// Description returns the semantic convention description of the instrument
func (StatefulSetReadyPods) Description() string {
	return "The number of replica pods created for this statefulset with a Ready Condition"
}

// Add adds incr to the existing count.
//
// This metric aligns with the `readyReplicas` field of the
// [K8s StatefulSetStatus].
//
// This metric SHOULD, at a minimum, be reported against a
// [`k8s.statefulset`] resource.
//
// [K8s StatefulSetStatus]: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#statefulsetstatus-v1-apps
// [`k8s.statefulset`]: ../resource/k8s.md#statefulset
func (m StatefulSetReadyPods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Add(ctx, incr)
	} else {
		m.inst.Add(ctx, incr, metric.WithAttributes(attrs...))
	}
}

// StatefulSetUpdatedPods is an instrument used to record metric values
// conforming to the "k8s.statefulset.updated_pods" semantic conventions. It
// represents the number of replica pods created by the statefulset controller
// from the statefulset version indicated by updateRevision.
type StatefulSetUpdatedPods struct {
	inst metric.Int64UpDownCounter
}

// NewStatefulSetUpdatedPods returns a new StatefulSetUpdatedPods instrument.
func NewStatefulSetUpdatedPods(m metric.Meter) (StatefulSetUpdatedPods, error) {
	i, err := m.Int64UpDownCounter(
	    "k8s.statefulset.updated_pods",
	    metric.WithDescription("Number of replica pods created by the statefulset controller from the statefulset version indicated by updateRevision"),
	    metric.WithUnit("{pod}"),
	)
	if err != nil {
	    return StatefulSetUpdatedPods{inst: noop.Int64UpDownCounter{}}, err
	}
	return StatefulSetUpdatedPods{i}, nil
}

// Name returns the semantic convention name of the instrument.
func (StatefulSetUpdatedPods) Name() string {
	return "k8s.statefulset.updated_pods"
}

// Unit returns the semantic convention unit of the instrument
func (StatefulSetUpdatedPods) Unit() string {
	return "{pod}"
}

// Description returns the semantic convention description of the instrument
func (StatefulSetUpdatedPods) Description() string {
	return "Number of replica pods created by the statefulset controller from the statefulset version indicated by updateRevision"
}

// Add adds incr to the existing count.
//
// This metric aligns with the `updatedReplicas` field of the
// [K8s StatefulSetStatus].
//
// This metric SHOULD, at a minimum, be reported against a
// [`k8s.statefulset`] resource.
//
// [K8s StatefulSetStatus]: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#statefulsetstatus-v1-apps
// [`k8s.statefulset`]: ../resource/k8s.md#statefulset
func (m StatefulSetUpdatedPods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	if len(attrs) == 0 {
		m.inst.Add(ctx, incr)
	} else {
		m.inst.Add(ctx, incr, metric.WithAttributes(attrs...))
	}
}