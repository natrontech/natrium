package helm

import (
	"context"
	"strings"
	"time"

	helmclient "github.com/mittwald/go-helm-client"
	"helm.sh/helm/v3/pkg/release"
	"helm.sh/helm/v3/pkg/repo"
)

var (
	Prefix = "natrium"
)

func GetNamespaceName(customer string) string {
	return Prefix + "-" + customer
}

func CreateHelmClient(namespace string) (helmclient.Client, error) {
	opt := &helmclient.Options{
		Namespace: namespace,
		Debug:     true,
		Linting:   true,
		DebugLog:  func(format string, v ...interface{}) {},
	}

	helmClient, err := helmclient.New(opt)
	if err != nil {
		return nil, err
	}
	return helmClient, nil
}

func AddHelmRepositoryToClient(helmClient helmclient.Client, repositoryName string, repositoryURL string) error {
	chartRepo := repo.Entry{
		Name: strings.ToLower(repositoryName),
		URL:  repositoryURL,
	}

	if err := helmClient.AddOrUpdateChartRepo(chartRepo); err != nil {
		return err
	}

	return nil
}

func CreateOrUpdateHelmRelease(helmClient helmclient.Client, chartName string, releaseName string, namespace string, version string, valuesYaml string) (*release.Release, error) {

	chartSpec := helmclient.ChartSpec{
		ChartName:       strings.ToLower(chartName),
		ReleaseName:     strings.ToLower(releaseName),
		Namespace:       namespace,
		CreateNamespace: true,
		Timeout:         32 * time.Second,
		Version:         version,
		ValuesYaml:      valuesYaml,
	}

	if release, err := helmClient.InstallOrUpgradeChart(context.Background(), &chartSpec, nil); err != nil {
		return nil, err
	} else {
		return release, nil
	}
}

func GetHelmRelease(helmClient helmclient.Client, releaseName string) (*release.Release, error) {
	return helmClient.GetRelease(releaseName)
}
