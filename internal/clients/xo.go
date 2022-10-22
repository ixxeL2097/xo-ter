package clients

import (
        "context"
        "encoding/json"
	"fmt"

        "github.com/crossplane/crossplane-runtime/pkg/resource"
        "github.com/pkg/errors"
        "k8s.io/apimachinery/pkg/types"
        "sigs.k8s.io/controller-runtime/pkg/client"

        "github.com/crossplane/terrajet/pkg/terraform"

        "github.com/crossplane-contrib/provider-jet-xo/apis/v1alpha1"
)


const (
	keyUsername = "username"
	keyPassword = "password"
	keyInsecure = "insecure"
	keyUrl      = "url"

	envUsername = "USERNAME"
        envPassword = "PASSWORD"
        envInsecure = "INSECURE"
        envUrl      = "URL"
)


const (
        fmtEnvVar = "%s=%s"
	// error messages
        errNoProviderConfig     = "no providerConfigRef provided"
        errGetProviderConfig    = "cannot get referenced ProviderConfig"
        errTrackUsage           = "cannot track ProviderConfig usage"
        errExtractCredentials   = "cannot extract credentials"
        errUnmarshalCredentials = "cannot unmarshal xo credentials as JSON"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
        return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
                ps := terraform.Setup{
                        Version: version,
                        Requirement: terraform.ProviderRequirement{
                                Source:  providerSource,
                                Version: providerVersion,
                        },
                }

                configRef := mg.GetProviderConfigReference()
                if configRef == nil {
                        return ps, errors.New(errNoProviderConfig)
                }
                pc := &v1alpha1.ProviderConfig{}
                if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
                        return ps, errors.Wrap(err, errGetProviderConfig)
                }

                t := resource.NewProviderConfigUsageTracker(client, &v1alpha1.ProviderConfigUsage{})
                if err := t.Track(ctx, mg); err != nil {
                        return ps, errors.Wrap(err, errTrackUsage)
                }

                data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
                if err != nil {
                        return ps, errors.Wrap(err, errExtractCredentials)
                }
                xoCreds := map[string]string{}
                if err := json.Unmarshal(data, &xoCreds); err != nil {
                        return ps, errors.Wrap(err, errUnmarshalCredentials)
                }

		ps.Configuration = map[string]interface{}{
			"url": xoCreds[keyUrl],
		}
		ps.Env = []string{
			fmt.Sprintf(fmtEnvVar, envUsername, xoCreds[keyUsername]),
			fmt.Sprintf(fmtEnvVar, envPassword, xoCreds[keyPassword]),
			fmt.Sprintf(fmtEnvVar, envInsecure, xoCreds[keyInsecure]),
			fmt.Sprintf(fmtEnvVar, envUrl, xoCreds[keyUrl]),
		}
                // set environment variables for sensitive provider configuration
                // Deprecated: In shared gRPC mode we do not support injecting
                // credentials via the environment variables. You should specify
                // credentials via the Terraform main.tf.json instead.
                //ps.Env = []string{
                //        fmt.Sprintf("%s=%s", "HASHICUPS_USERNAME", xoCreds["username"]),
                //        fmt.Sprintf("%s=%s", "HASHICUPS_PASSWORD", xoCreds["password"]),
                //}
                // set credentials in Terraform provider configuration
                // ps.Configuration = map[string]interface{}{
                //         "username": xoCreds["username"],
                //         "password": xoCreds["password"],
                // }

                return ps, nil
        }
}
