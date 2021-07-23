package kubernetes_pod

// Ready implements the ready.Readiness interface.
func (k *Kubernetes) Ready() bool { return k.APIConn.HasSynced() }
