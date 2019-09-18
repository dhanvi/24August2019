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

// podsCmd represents the pods command
var podsCmd = &cobra.Command{
	Use:   "pods",
	Short: "Get the pods list",
	Long:  `get pods list`,
	Run:   getPods,
}

func init() {
	rootCmd.AddCommand(podsCmd)
}

func getPods(cmd *cobra.Command, args []string) {
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

	pods, err := cs.CoreV1().Pods(namespace).List(metaV1.ListOptions{})
	if err != nil {
		fmt.Printf("could not get pods in %q namespace: %v\n", namespace, err)
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintln(w, "Name\tStatus")
	for _, pod := range pods.Items {
		fmt.Fprintf(w, "%s\t%s\n", pod.Name, pod.Status.Phase)
	}
	w.Flush()

}
