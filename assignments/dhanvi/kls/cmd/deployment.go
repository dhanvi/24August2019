package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// deploymentCmd represents the deployment command
var deploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "Get deployment list",
	Long:  `get deployment list`,
	Run:   getDeployment,
}

func init() {
	rootCmd.AddCommand(deploymentCmd)
}

func getDeployment(cmd *cobra.Command, args []string) {
	var kubeConfigPath, namespace string
	kubeConfigPath = viper.GetString("kube-config")
	namespace = viper.GetString("namespace")

	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		fmt.Printf("can't build config from kubeconfig at %s: %v\n", kubeConfigPath, err)
		os.Exit(-1)
	}

	cs, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("can't get kubernetes client: %v\n", err)
		os.Exit(-1)
	}

	dep, err := cs.AppsV1().Deployments(namespace).List(metaV1.ListOptions{})

	if err != nil {
		fmt.Printf("could not get deployments in %q namespace: %v\n", namespace, err)
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintln(w, "Name\tAvaliable Replicas")
	for _, dep := range dep.Items {
		replicas := fmt.Sprint(dep.Status.AvailableReplicas)
		fmt.Fprintf(w, "%s\t%s\n", dep.Name, replicas)
	}
	w.Flush()

}
