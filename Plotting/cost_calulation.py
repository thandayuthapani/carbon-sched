# Unit of cost is in a billionth of a dollar
# Check https://cloud.google.com/run/pricing
GCR_COST_VCPU_PER_SECOND = 24000
GCR_COST_MEM_GIB_PER_SECOND = 2500
GCR_COST_PER_INVOCATION = 400

# Cost source: https://cloud.google.com/compute/vm-instance-pricing
GKE_COST_CLUSTER_PER_HOUR = 100000000
GKE_COST_NODE_E2_HIGHCPU_4_PER_HOUR = 108836000

GKE_COST_AUTOPILOT_MILI_VCPU_PER_HOUR = 49000
GKE_COST_AUTOPILOT_MEM_MIB_PER_HOUR = 5293.1640625
GKE_COST_AUTOPILOT_STORAGE_MIB_PER_HOUR = 58.88671875

# Cost to internet in europe
GCP_COST_EGRESS_TRAFFIC_PER_BYTE = 0.11 * (10**9) / (2**30)


class Cost:
    # The cost is split up into traffic, compute (memory + cpu + persistent storage)
    # and management (e.g. invocations)
    def __init__(self, traffic, compute, management):
        self.traffic = traffic
        self.compute = compute
        self.management = management

    def total(self):
        return self.t() + self.c() + self.m()

    def t(self):
        return round(self.traffic / 10**9, 2)

    def c(self):
        return round(self.compute / 10**9, 2)

    def m(self):
        return round(self.management / 10**9, 2)

    # Normalizes the cost per 1,000,000 requests
    # Should only be done once all fields are completely filled
    def normalize(self, nr_front_end_requests):
        self.traffic = self.traffic * 1000000 // nr_front_end_requests
        self.compute = self.compute * 1000000 // nr_front_end_requests
        self.management = self.management * 1000000 // nr_front_end_requests


def get_cost_gcr(gcp_metrics, k6_metrics, static_dynamic=False):
    cost = Cost(0, 0, 0)
    invocations = 0
    for service in gcp_metrics.values():
        cost.compute += service["cpu_alloc_time"] * GCR_COST_VCPU_PER_SECOND
        cost.compute += service["memory_allocation_time"] * GCR_COST_MEM_GIB_PER_SECOND
        cost.management += service["request_count"] * GCR_COST_PER_INVOCATION
        invocations += service["request_count"]
    cost.traffic = (
        k6_metrics["metrics"]["data_received"]["values"]["count"]
        * GCP_COST_EGRESS_TRAFFIC_PER_BYTE
    )
    if static_dynamic:
        # In the experiment with static/dynamic we have not only requested follow/unfollow
        # but also the frontend. So we have double the many requests
        cost.normalize(k6_metrics["metrics"]["http_req_failed"]["values"]["fails"] / 2)
    else:
        cost.normalize(k6_metrics["metrics"]["http_req_failed"]["values"]["fails"])
    return cost


def get_cost_gke_autopilot(metrics, k6_metrics):
    cost = Cost(0, 0, 0)
    poll_interval = metrics["poll_interval_in_s"]
    cost.compute = (
        metrics["requests"]["cpu"]
        * poll_interval
        * GKE_COST_AUTOPILOT_MILI_VCPU_PER_HOUR
        / 3600
    )
    cost.compute += (
        metrics["requests"]["memory"]
        * poll_interval
        * GKE_COST_AUTOPILOT_MEM_MIB_PER_HOUR
        / 3600
    )
    cost.compute += (
        metrics["requests"]["ephemeral-storage"]
        * poll_interval
        * GKE_COST_AUTOPILOT_STORAGE_MIB_PER_HOUR
        / 3600
    )
    cost.traffic = (
        k6_metrics["metrics"]["data_received"]["values"]["count"]
        * GCP_COST_EGRESS_TRAFFIC_PER_BYTE
    )
    cost.management = metrics["time_running_in_s"] * GKE_COST_CLUSTER_PER_HOUR / 3600
    cost.normalize(k6_metrics["metrics"]["http_req_failed"]["values"]["fails"] / 2)
    return cost


def get_cost_gke_standard(metrics, k6_metrics):
    cost = Cost(0, 0, 0)
    cost.compute = (
        sum(metrics["nr_nodes"])
        * metrics["poll_interval_in_s"]
        * GKE_COST_NODE_E2_HIGHCPU_4_PER_HOUR
        / 3600
    )
    cost.traffic = (
        k6_metrics["metrics"]["data_received"]["values"]["count"]
        * GCP_COST_EGRESS_TRAFFIC_PER_BYTE
    )
    cost.management = metrics["time_running_in_s"] * GKE_COST_CLUSTER_PER_HOUR / 3600
    cost.normalize(k6_metrics["metrics"]["http_req_failed"]["values"]["fails"] / 2)
    return cost


# Returns the cost per 100000 made requests
def get_cost(gcp_metrics, k6_metrics, static_dynamic=False):
    infra = gcp_metrics["infra"]
    if infra == "gcr":
        return get_cost_gcr(gcp_metrics["metrics"], k6_metrics, static_dynamic)
    elif infra == "gke_standard":
        return get_cost_gke_standard(gcp_metrics["metrics"], k6_metrics)
    elif infra == "gke_autopilot":
        return get_cost_gke_autopilot(gcp_metrics["metrics"], k6_metrics)
    else:
        raise NotImplementedError(f"Cannot handle infra {infra}")
