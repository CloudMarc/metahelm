package metahelm

import "time"

// HealthIndication describes how to decide if a deployment is successful
type HealthIndication int

const (
	// IgnorePodHealth indicates that we don't care about pod health
	IgnorePodHealth HealthIndication = iota
	// AllPodsHealthy indeicates that all pods are OK
	AllPodsHealthy
	// AtLeastOnePodHealthy indicates >= 1 pods are OK
	AtLeastOnePodHealthy
)

// DefaultDeploymentTimeout indicates the default time to wait for a deployment to be healthy
const DefaultDeploymentTimeout = 10 * time.Minute

// Chart models a single installable Helm chart
type Chart struct {
	Title                      string           `yaml:"title"`                        // unique name for this chart (must not collide with any dependencies)
	Location                   string           `yaml:"location"`                     // local filesystem location
	ValueOverrides             []byte           `yaml:"value_overrides"`              // value overrides as raw YAML stream
	WaitUntilDeployment        string           `yaml:"wait_until_deployment"`        // Deployment name that, when healthy, indicates chart install has succeeded
	WaitTimeout                time.Duration    `yaml:"wait_timeout"`                 // how long to wait for the deployment to become healthy. If unset, DefaultDeploymentTimeout is used
	DeploymentHealthIndication HealthIndication `yaml:"deployment_health_indication"` // How to determine if a deployment is healthy
	DependencyList             []string         `yaml:"dependencies"`
}

func (c *Chart) Name() string {
	return c.Title
}

func (c *Chart) String() string {
	return c.Title
}

func (c *Chart) Dependencies() []string {
	return c.DependencyList
}