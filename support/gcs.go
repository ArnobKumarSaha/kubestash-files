func getGcsEnv(backupStorage *storageapi.BackupStorage, dbMeta metav1.ObjectMeta, subDir string) ([]core.EnvVar, error) {
	gcsEnv := []core.EnvVar{
		{
			Name:  "WALG_GS_PREFIX",
			Value: getGCSPrefix(backupStorage, dbMeta, subDir),
		},
		{
			Name:  "GOOGLE_APPLICATION_CREDENTIALS",
			Value: fmt.Sprintf("%s/%s", googleCredMountPath, googleCredFileName),
		},
	}

	return gcsEnv, nil
}

func getGCSPrefix(bs *storageapi.BackupStorage, dbMeta metav1.ObjectMeta, subDir string) string {
	return fmt.Sprintf("gs://%s", filepath.Join(bs.Spec.Storage.GCS.Bucket, getWalDir(bs, dbMeta, subDir)))
}



if backupStorage.Spec.Storage.Provider == storageapi.ProviderGCS {
	volumes = core_util.UpsertVolume(volumes, core.Volume{
		Name: googleCredVolumeName,
		VolumeSource: core.VolumeSource{
			Secret: &core.SecretVolumeSource{
				SecretName: backupStorage.Spec.Storage.GCS.SecretName,
			},
		},
	})
}

if backupStorage.Spec.Storage.Provider == storageapi.ProviderGCS {
	volumeMounts = append(volumeMounts, core.VolumeMount{
		Name:      googleCredVolumeName,
		MountPath: googleCredMountPath,
	})
}