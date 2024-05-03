func getLocalEnv(backupStorage *storageapi.BackupStorage, dbMeta metav1.ObjectMeta, subDir string) ([]core.EnvVar, error) {
	localEnvs := []core.EnvVar{
		{
			Name:  "WALG_FILE_PREFIX",
			Value: getLocalPrefix(backupStorage, dbMeta, subDir),
		},
	}
	return localEnvs, nil
}

func getLocalPrefix(bs *storageapi.BackupStorage, dbMeta metav1.ObjectMeta, subDir string) string {
	return getWalDir(bs, dbMeta, subDir)
}


if backupStorage.Spec.Storage.Provider == storageapi.ProviderLocal {
	volumeSource := backupStorage.Spec.Storage.Local.VolumeSource
	volumes = core_util.UpsertVolume(volumes, core.Volume{
		Name:         string(storageapi.ProviderLocal),
		VolumeSource: *volumeSource.ToAPIObject(),
	})
}
if backupStorage.Spec.Storage.Provider == storageapi.ProviderLocal {
	volumeMounts = append(volumeMounts, core.VolumeMount{
		Name:      string(storageapi.ProviderLocal),
		MountPath: backupStorage.Spec.Storage.Local.MountPath,
		SubPath:   backupStorage.Spec.Storage.Local.SubPath,
	})
}