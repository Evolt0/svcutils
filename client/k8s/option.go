package k8s

type Option func(*K8sClient)

func WithConfigPath(configPath string) Option {
	return func(c *K8sClient) {
		c.ConfigPath = configPath
	}
}

func WithConfigData(configData []byte) Option {
	return func(c *K8sClient) {
		c.ConfigData = configData
	}
}

func WithInCluster(inCluster bool) Option {
	return func(c *K8sClient) {
		c.InCluster = inCluster
	}
}
