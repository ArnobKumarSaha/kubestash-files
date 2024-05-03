func getAzureEnv(backupStorage *storageapi.BackupStorage, dbMeta metav1.ObjectMeta, subDir string) ([]core.EnvVar, error) {
	storageSecretName, err := getSecretName(backupStorage)
	if err != nil {
		return []core.EnvVar{}, err
	}
	azureEnv := []core.EnvVar{
		{
			Name:  "WALG_AZ_PREFIX",
			Value: getAzurePrefix(backupStorage, dbMeta, subDir),
		},
		{
			Name:  "AZURE_STORAGE_ACCOUNT",
			Value: backupStorage.Spec.Storage.Azure.StorageAccount,
		},
		{
			Name: "AZURE_STORAGE_ACCESS_KEY",
			ValueFrom: &core.EnvVarSource{
				SecretKeyRef: &core.SecretKeySelector{
					LocalObjectReference: core.LocalObjectReference{
						Name: storageSecretName,
					},
					Key: "AZURE_ACCOUNT_KEY",
				},
			},
		},
	}

	return azureEnv, nil
}

func getAzurePrefix(bs *storageapi.BackupStorage, dbMeta metav1.ObjectMeta, subDir string) string {
	return fmt.Sprintf("azure://%s", filepath.Join(bs.Spec.Storage.Azure.Container, getWalDir(bs, dbMeta, subDir)))
}