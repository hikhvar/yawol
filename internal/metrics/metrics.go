package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

type LoadBalancerMetricList struct {
	OpenstackMetrics       *prometheus.CounterVec
	InfoMetrics            *prometheus.GaugeVec
	OpenstackInfoMetrics   *prometheus.GaugeVec
	ReplicasMetrics        *prometheus.GaugeVec
	ReplicasCurrentMetrics *prometheus.GaugeVec
	ReplicasReadyMetrics   *prometheus.GaugeVec
}

var LoadBalancerMetrics = LoadBalancerMetricList{
	OpenstackMetrics:       OpenstackMetrics,
	InfoMetrics:            LoadBalancerInfoMetrics,
	OpenstackInfoMetrics:   LoadBalancerOpenstackMetrics,
	ReplicasMetrics:        LoadBalancerReplicasMetrics,
	ReplicasCurrentMetrics: LoadBalancerReplicasCurrentMetrics,
	ReplicasReadyMetrics:   LoadBalancerReplicasReadyMetrics,
}

type LoadBalancerSetMetricList struct {
	ReplicasMetrics        *prometheus.GaugeVec
	ReplicasCurrentMetrics *prometheus.GaugeVec
	ReplicasReadyMetrics   *prometheus.GaugeVec
}

var LoadBalancerSetMetrics = LoadBalancerSetMetricList{
	ReplicasMetrics:        LoadBalancerSetReplicasMetrics,
	ReplicasCurrentMetrics: LoadBalancerSetReplicasCurrentMetrics,
	ReplicasReadyMetrics:   LoadBalancerSetReplicasReadyMetrics,
}

type LoadBalancerMachineMetricList struct {
	VM               *prometheus.GaugeVec
	Conditions       *prometheus.GaugeVec
	OpenstackMetrics *prometheus.CounterVec
}

var LoadBalancerMachineMetrics = LoadBalancerMachineMetricList{
	VM:               LoadBalancerMachineVMMetrics,
	Conditions:       LoadBalancerMachineConditionMetrics,
	OpenstackMetrics: OpenstackMetrics,
}

var (
	// OpenstackMetrics Openstack usage counter by API
	OpenstackMetrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "yawol_openstack",
		Help: "Openstack usage counter by api, object, operation",
	}, []string{"api", "object", "operation"})

	// LoadBalancerInfoMetrics Loadbalancer Info for LoadBalancer contains labels like isInternal, externalIP
	LoadBalancerInfoMetrics = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "loadbalancer_info",
		Help: "Loadbalancer Info for LoadBalancer contains labels like isInternal, externalIP",
	}, []string{"lb", "namespace", "isInternal", "tcpProxyProtocol", "externalIP"})
	// LoadBalancerOpenstackMetrics Openstack Info contains labels with the OpenStackIDs for LoadBalancer
	LoadBalancerOpenstackMetrics = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "loadbalancer_openstack_info",
		Help: "Openstack Info contains labels with the OpenStackIDs for LoadBalancer",
	}, []string{"lb", "namespace", "portID", "floatingID", "securityGroupID"})
	// LoadBalancerReplicasMetrics Replicas for LoadBalancer (from lb.spec.replicas)
	LoadBalancerReplicasMetrics = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "loadbalancer_replicas",
		Help: "Replicas for LoadBalancer (from lb.spec.replicas)",
	}, []string{"lb", "namespace"})
	// LoadBalancerReplicasCurrentMetrics Current replicas for LoadBalancer (from lb.status.replicas)
	LoadBalancerReplicasCurrentMetrics = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "loadbalancer_replicas_current",
		Help: "Current replicas for LoadBalancer (from lb.status.replicas)",
	}, []string{"lb", "namespace"})
	// LoadBalancerReplicasReadyMetrics Ready replicas for LoadBalancer (from lb.status.readyReplicas)
	LoadBalancerReplicasReadyMetrics = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "loadbalancer_replicas_ready",
		Help: "Ready replicas for LoadBalancer (from lb.status.readyReplicas)",
	}, []string{"lb", "namespace"})

	// LoadBalancerSetReplicasMetrics Replicas for LoadBalancerSet (from lbs.spec.replicas)
	LoadBalancerSetReplicasMetrics = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "loadbalancerset_replicas",
		Help: "Replicas for LoadBalancerSet (from lbs.spec.replicas)",
	}, []string{"lb", "lbs", "namespace"})
	// LoadBalancerSetReplicasCurrentMetrics Current replicas for LoadBalancerSet (from lbs.status.replicas)
	LoadBalancerSetReplicasCurrentMetrics = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "loadbalancerset_replicas_current",
		Help: "Current replicas for LoadBalancerSet (from lbs.status.replicas)",
	}, []string{"lb", "lbs", "namespace"})
	// LoadBalancerSetReplicasReadyMetrics Ready replicas for LoadBalancerSet (from lbs.status.readyReplicas)
	LoadBalancerSetReplicasReadyMetrics = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "loadbalancerset_replicas_ready",
		Help: "Ready replicas for LoadBalancerSet (from lbs.status.readyReplicas)",
	}, []string{"lb", "lbs", "namespace"})

	// LoadBalancerMachineVMMetrics Metrics of loadbalancermachine (all metrics from lbm.status.metrics)
	LoadBalancerMachineVMMetrics = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "loadbalancermachine",
		Help: "Metrics of loadbalancermachine (all metrics from lbm.status.metrics)",
	}, []string{"type", "lb", "lbm", "namespace"})
	// LoadBalancerMachineConditionMetrics Conditions of loadbalancermachine (lbm.status.conditions
	LoadBalancerMachineConditionMetrics = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "loadbalancermachine_condition",
		Help: "Conditions of loadbalancermachine (lbm.status.conditions)",
	}, []string{"lb", "lbm", "namespace", "condition", "reason", "status"})
)

func init() {
	// Register custom metrics with the global prometheus registry
	metrics.Registry.MustRegister(OpenstackMetrics)
	metrics.Registry.MustRegister(LoadBalancerInfoMetrics)
	metrics.Registry.MustRegister(LoadBalancerOpenstackMetrics)
	metrics.Registry.MustRegister(LoadBalancerReplicasMetrics)
	metrics.Registry.MustRegister(LoadBalancerReplicasCurrentMetrics)
	metrics.Registry.MustRegister(LoadBalancerReplicasReadyMetrics)
	metrics.Registry.MustRegister(LoadBalancerSetReplicasMetrics)
	metrics.Registry.MustRegister(LoadBalancerSetReplicasCurrentMetrics)
	metrics.Registry.MustRegister(LoadBalancerSetReplicasReadyMetrics)
	metrics.Registry.MustRegister(LoadBalancerMachineVMMetrics)
	metrics.Registry.MustRegister(LoadBalancerMachineConditionMetrics)
}
