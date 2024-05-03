func getS3Env(backupStorage *storageapi.BackupStorage, dbMeta metav1.ObjectMeta, subDir string) ([]core.EnvVar, error) {
	storageSecretName, err := getSecretName(backupStorage)
	if err != nil {
		return []core.EnvVar{}, err
	}
	return []core.EnvVar{
		{
			Name:  "AWS_S3_FORCE_PATH_STYLE",
			Value: "true",
		},
		{
			Name:  "WALG_S3_PREFIX",
			Value: getS3Prefix(backupStorage, dbMeta, subDir),
		},
		{
			Name:  "AWS_REGION",
			Value: backupStorage.Spec.Storage.S3.Region,
		},
		{
			Name:  "AWS_ENDPOINT",
			Value: backupStorage.Spec.Storage.S3.Endpoint,
		},
		{
			Name: "AWS_ACCESS_KEY_ID",
			ValueFrom: &core.EnvVarSource{
				SecretKeyRef: &core.SecretKeySelector{
					LocalObjectReference: core.LocalObjectReference{
						Name: storageSecretName,
					},
					Key: "AWS_ACCESS_KEY_ID",
				},
			},
		},
		{
			Name: "AWS_SECRET_ACCESS_KEY",
			ValueFrom: &core.EnvVarSource{
				SecretKeyRef: &core.SecretKeySelector{
					LocalObjectReference: core.LocalObjectReference{
						Name: storageSecretName,
					},
					Key: "AWS_SECRET_ACCESS_KEY",
				},
			},
		},
	}, nil
}

func getS3Prefix(bs *storageapi.BackupStorage, dbMeta metav1.ObjectMeta, subDir string) string {
	return fmt.Sprintf("s3://%s", filepath.Join(bs.Spec.Storage.S3.Bucket, getWalDir(bs, dbMeta, subDir)))
}
