// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

// List of APIs:
//   - [accountApi.createAccount]
//   - [accountApi.deleteAccount]
//   - [accountApi.listAccounts]
//   - [accountApi.updateAccount]
//   - [accountApi.updateAccountPrivileges]
//   - [adminUserApi.createAdminUser]
//   - [adminUserApi.deleteAdminUser]
//   - [adminUserApi.listAdminUsers]
//   - [adminUserApi.patchAdminUser]
//   - [adminUserApi.readAdminUser]
//   - [alertConfigApi.getAlertConfig]
//   - [alertConfigApi.getAlertSMSConfig]
//   - [alertConfigApi.setAlertConfig]
//   - [alertConfigApi.updateAlertSMSConfig]
//   - [alertInhibitApi.createAlertInhibit]
//   - [alertInhibitApi.deleteAlertInhibit]
//   - [alertInhibitApi.getAlertInhibit]
//   - [alertInhibitApi.listAlertInhibits]
//   - [alertInhibitApi.patchAlertInhibit]
//   - [alertMetricsApi.listAlertMetrics]
//   - [alertObjectApi.listAlertObjects]
//   - [alertObjectApi.setAlertObjectStatus]
//   - [alertObjectApi.setAlertObjectsStatus]
//   - [alertReceiverApi.createAlertReceiver]
//   - [alertReceiverApi.deleteAlertReceiver]
//   - [alertReceiverApi.getAlertReceiver]
//   - [alertReceiverApi.listAlertReceivers]
//   - [alertReceiverApi.patchAlertReceiver]
//   - [alertRuleApi.createAlertRule]
//   - [alertRuleApi.deleteAlertRule]
//   - [alertRuleApi.getAlertRule]
//   - [alertRuleApi.listAlertRules]
//   - [alertRuleApi.updateAlertRule]
//   - [alertSMTPConfigApi.getAlertSMTPConfig]
//   - [alertSMTPConfigApi.updateAlertSMTPConfig]
//   - [alertStrategyApi.createAlertStrategy]
//   - [alertStrategyApi.deleteAlertStrategy]
//   - [alertStrategyApi.listAlertStrategies]
//   - [alertStrategyApi.patchAlertStrategy]
//   - [alertStrategyApi.updateAlertStrategy]
//   - [alertTemplateApi.createAlertTemplate]
//   - [alertTemplateApi.deleteAlertTemplate]
//   - [alertTemplateApi.getAlertTemplate]
//   - [alertTemplateApi.listAlertTemplates]
//   - [alertTemplateApi.patchAlertTemplate]
//   - [analyzeApi.analyzeBackup]
//   - [analyzeApi.analyzeClusterParam]
//   - [analyzeApi.analyzeClusterRestore]
//   - [analyzeApi.analyzeLogs]
//   - [analyzeApi.analyzeOps]
//   - [analyzeApi.analyzeParam]
//   - [analyzeApi.analyzeService]
//   - [analyzeApi.analyzeSlowLogs]
//   - [analyzeApi.analyzeView]
//   - [autohealingApi.getAutohealing]
//   - [backupApi.createClusterBackup]
//   - [backupApi.deleteBackup]
//   - [backupApi.downloadBackup]
//   - [backupApi.downloadMutipleBackups]
//   - [backupApi.getBackup]
//   - [backupApi.getBackupLog]
//   - [backupApi.getBackupStats]
//   - [backupApi.getClusterBackupPolicy]
//   - [backupApi.listBackups]
//   - [backupApi.patchBackupPolicy]
//   - [backupApi.viewBackup]
//   - [backupRepoApi.checkBackupRepo]
//   - [backupRepoApi.createBackupRepo]
//   - [backupRepoApi.deleteBackupRepo]
//   - [backupRepoApi.getBackupRepo]
//   - [backupRepoApi.listBackupRepos]
//   - [backupRepoApi.listStorageProvidersArchive]
//   - [backupRepoApi.updateBackupRepo]
//   - [backupRepoApi.viewBackupRepo]
//   - [classApi.batchClass]
//   - [classApi.createClass]
//   - [classApi.deleteClass]
//   - [classApi.listClasses]
//   - [classApi.patchClass]
//   - [clusterApi.createCluster]
//   - [clusterApi.deleteCluster]
//   - [clusterApi.describeClusterHaHistory]
//   - [clusterApi.getCluster]
//   - [clusterApi.getClusterByID]
//   - [clusterApi.getClusterInstanceLog]
//   - [clusterApi.getClusterManifest]
//   - [clusterApi.getInstacesMetrics]
//   - [clusterApi.listCluster]
//   - [clusterApi.listClusters]
//   - [clusterApi.listEndpoints]
//   - [clusterApi.listInstance]
//   - [clusterApi.patchCluster]
//   - [clusterAlertSwitchApi.getClusterAlertDisabled]
//   - [clusterAlertSwitchApi.setClusterAlertDisabled]
//   - [clusterLogApi.queryAuditLogs]
//   - [clusterLogApi.queryErrorLogs]
//   - [clusterLogApi.queryRunningLogs]
//   - [clusterLogApi.querySlowLogs]
//   - [databaseApi.createDatabase]
//   - [databaseApi.deleteDatabase]
//   - [databaseApi.listDatabases]
//   - [dmsApi.DataExport]
//   - [dmsApi.DataImport]
//   - [dmsApi.GetObjectInfo]
//   - [dmsApi.GetTaskList]
//   - [dmsApi.GetTaskProgress]
//   - [dmsApi.ListObjectNamesByType]
//   - [dmsApi.ListObjectTypesInSchema]
//   - [dmsApi.alterParameter]
//   - [dmsApi.closeSessions]
//   - [dmsApi.createDataSourceV2]
//   - [dmsApi.deleteDataSourceV2]
//   - [dmsApi.generateDDL]
//   - [dmsApi.getDataSourceV2]
//   - [dmsApi.getSchemaList]
//   - [dmsApi.listDataSourceV2]
//   - [dmsApi.listParameters]
//   - [dmsApi.listQueryHistory]
//   - [dmsApi.listSessions]
//   - [dmsApi.query]
//   - [dmsApi.showData]
//   - [dmsApi.sqlExplain]
//   - [dmsApi.tenantParameterHistory]
//   - [dmsApi.testDataSourceV2]
//   - [dmsApi.updateDataSourceV2]
//   - [engineApi.engineAction]
//   - [engineApi.engineActionInOrg]
//   - [engineApi.getEngineByNameInEnv]
//   - [engineApi.listEnginesInEnv]
//   - [engineApi.listEnginesInOrg]
//   - [engineLicenseApi.createEngineLicense]
//   - [engineLicenseApi.deleteEngineLicense]
//   - [engineLicenseApi.engineLicense]
//   - [engineLicenseApi.listEngineLicenses]
//   - [engineOptionApi.ListUpgradeableServiceVersion]
//   - [engineOptionApi.createEngineOption]
//   - [engineOptionApi.getEngineOption]
//   - [engineOptionApi.listEngineOptionHistory]
//   - [engineOptionApi.listEngineOptions]
//   - [engineOptionApi.updateEngineOption]
//   - [engineVersionApi.createEngineVersion]
//   - [engineVersionApi.deleteEngineVersion]
//   - [engineVersionApi.listEngineVersions]
//   - [engineVersionApi.patchEngineVersion]
//   - [environmentApi.addNodes]
//   - [environmentApi.checkEnvironmentAddRecord]
//   - [environmentApi.cordonEnvironmentNode]
//   - [environmentApi.createEnvironment]
//   - [environmentApi.createNodeGroup]
//   - [environmentApi.createWorkflow]
//   - [environmentApi.deleteEnvironment]
//   - [environmentApi.deleteNodeGroup]
//   - [environmentApi.getEnvironment]
//   - [environmentApi.getEnvironmentBackupRepo]
//   - [environmentApi.getEnvironmentBootstrapManifests]
//   - [environmentApi.getEnvironmentKubeconfig]
//   - [environmentApi.getEnvironmentMetricsMonitorStats]
//   - [environmentApi.getEnvironmentNamespaceStats]
//   - [environmentApi.getEnvironmentProvisioningProgress]
//   - [environmentApi.getEnvironmentStatus]
//   - [environmentApi.getEnvironmentStatusHistory]
//   - [environmentApi.getLatestComponentVersion]
//   - [environmentApi.getNode]
//   - [environmentApi.getWorkflow]
//   - [environmentApi.getWorkflowList]
//   - [environmentApi.getWorkflowLog]
//   - [environmentApi.listEnvNodeZone]
//   - [environmentApi.listEnvironment]
//   - [environmentApi.listEnvironmentObjectStorage]
//   - [environmentApi.listKubernetesNode]
//   - [environmentApi.listKubernetesStorageClass]
//   - [environmentApi.listNodeGroup]
//   - [environmentApi.listNodePod]
//   - [environmentApi.listNodes]
//   - [environmentApi.patchEnvironment]
//   - [environmentApi.patchNodeGroup]
//   - [environmentApi.preflightEnvironment]
//   - [environmentApi.preflightRestoreEnvironment]
//   - [environmentApi.restoreEnvironment]
//   - [environmentApi.uncordonEnvironmentNode]
//   - [eventApi.queryClusterEvents]
//   - [faultApi.createClusterNetworkChaos]
//   - [faultApi.createClusterPodChaos]
//   - [faultApi.deleteFault]
//   - [faultApi.getChaos]
//   - [faultApi.listFault]
//   - [featureApi.listFeature]
//   - [featureApi.readFeature]
//   - [imageRegistryApi.createImageRegistry]
//   - [imageRegistryApi.deleteImageRegistry]
//   - [imageRegistryApi.getImageRegistry]
//   - [imageRegistryApi.listImageRegistries]
//   - [imageRegistryApi.patchImageRegistry]
//   - [inspectionApi.createAutoInspection]
//   - [inspectionApi.createInspectionScript]
//   - [inspectionApi.deleteInspectionScript]
//   - [inspectionApi.listAutoInspection]
//   - [inspectionApi.listInspectionScripts]
//   - [inspectionApi.listInspections]
//   - [inspectionApi.updateAutoInspection]
//   - [inspectionApi.updateInspection]
//   - [inspectionApi.updateInspectionScript]
//   - [invitationApi.listInvitation]
//   - [ipWhitelistApi.createIPWhitelist]
//   - [ipWhitelistApi.listIPWhitelist]
//   - [ipWhitelistApi.updateIPWhitelist]
//   - [licenseApi.getLicense]
//   - [licenseApi.updateLicense]
//   - [llmApi.createLLM]
//   - [llmApi.deleteLLM]
//   - [llmApi.listLLM]
//   - [llmApi.updateLLM]
//   - [loadBalancerApi.checkLoadBalancer]
//   - [loadBalancerApi.getLoadBalancer]
//   - [loadBalancerApi.installLoadBalancer]
//   - [loadBalancerApi.uninstallLoadBalancer]
//   - [markClusterApi.markClusterRestoreCompleted]
//   - [memberApi.addOrgMember]
//   - [memberApi.deleteOrgMember]
//   - [memberApi.listOrgMember]
//   - [memberApi.patchOrgMember]
//   - [memberApi.readOrgMember]
//   - [metadbApi.getMetadbInstancesMetrics]
//   - [metadbApi.getPostgresClusterInfo]
//   - [metadbApi.listMetadbBackups]
//   - [metadbApi.metadbBackup]
//   - [metadbApi.metadbBackupConfig]
//   - [metadbApi.metadbBackupGetConfig]
//   - [metadbApi.metadbListInstances]
//   - [metadbApi.metadbRestore]
//   - [metadbApi.metadbVerticalScale]
//   - [metadbApi.metadbVolumeExpand]
//   - [metricsApi.getAggregateMetaData]
//   - [metricsApi.getEnvironmentStats]
//   - [metricsApi.queryClusterMetrics]
//   - [monitorDataSinkApi.createMonitorDataSink]
//   - [monitorDataSinkApi.deleteMonitorDataSink]
//   - [monitorDataSinkApi.listMonitorDataSinks]
//   - [monitorDataSinkApi.patchMonitorDataSink]
//   - [oceanbaseApi.getTenant]
//   - [oceanbaseApi.listTenants]
//   - [opsrequestApi.cancelOps]
//   - [opsrequestApi.clusterVolumeExpand]
//   - [opsrequestApi.customOps]
//   - [opsrequestApi.deleteOps]
//   - [opsrequestApi.exposeCluster]
//   - [opsrequestApi.horizontalScaleCluster]
//   - [opsrequestApi.promoteCluster]
//   - [opsrequestApi.rebuildInstance]
//   - [opsrequestApi.reconfigureCluster]
//   - [opsrequestApi.restartCluster]
//   - [opsrequestApi.startCluster]
//   - [opsrequestApi.stopCluster]
//   - [opsrequestApi.updateClusterLicense]
//   - [opsrequestApi.upgradeCluster]
//   - [opsrequestApi.verticalScaleCluster]
//   - [organizationApi.disableOrg]
//   - [organizationApi.enableOrg]
//   - [organizationApi.listOrganizations]
//   - [organizationApi.patchOrg]
//   - [organizationApi.readOrg]
//   - [paramTplApi.createParamTpl]
//   - [paramTplApi.createParamTplFromCluster]
//   - [paramTplApi.deleteParamTpl]
//   - [paramTplApi.getClusterParamTpls]
//   - [paramTplApi.listParamTpls]
//   - [paramTplApi.patchParamTpl]
//   - [paramTplApi.readParamTpl]
//   - [parameterApi.listConfigurations]
//   - [parameterApi.listParameterSpecs]
//   - [parameterApi.listParametersHistory]
//   - [platformParameterApi.getPlatformParameter]
//   - [platformParameterApi.listPlatformParameters]
//   - [platformParameterApi.updatePlatformParameters]
//   - [providerApi.createCloudProvider]
//   - [providerApi.deleteCloudProvider]
//   - [providerApi.getCloudProvider]
//   - [providerApi.listCloudProviders]
//   - [providerApi.updateCloudProvider]
//   - [recycleBinClusterApi.deleteRecycleBinCluster]
//   - [recycleBinClusterApi.getRecycleBinCluster]
//   - [recycleBinClusterApi.listRecycleBinCluster]
//   - [recycleBinClusterApi.listRecycleBinClusters]
//   - [recycleBinClusterApi.restoreRecycleBinCluster]
//   - [regionApi.createRegion]
//   - [regionApi.deleteRegion]
//   - [regionApi.getRegion]
//   - [regionApi.listRegions]
//   - [regionApi.updateRegion]
//   - [regionGroupApi.createRegionGroup]
//   - [regionGroupApi.deleteRegionGroup]
//   - [regionGroupApi.listRegionGroups]
//   - [regionGroupApi.updateRegionGroup]
//   - [resourceStatsApi.getResourceStats]
//   - [resourceStatsApi.listInstancesResourceStats]
//   - [resourceStatsApi.listNodesResourceStats]
//   - [restoreApi.GetRestoreLog]
//   - [restoreApi.deleteRestoreObject]
//   - [restoreApi.doRestore]
//   - [restoreApi.getRestoreTimeRange]
//   - [restoreApi.listClusterRestore]
//   - [restoreApi.listRestores]
//   - [restoreApi.restoreCluster]
//   - [roleApi.createRole]
//   - [roleApi.deleteRoleByName]
//   - [roleApi.getRoleByName]
//   - [roleApi.listPermissions]
//   - [roleApi.listRolePermissions]
//   - [roleApi.listRoles]
//   - [roleApi.updateRoleByName]
//   - [serviceVersionApi.ListServiceVersion]
//   - [storageApi.checkStorage]
//   - [storageApi.createStorage]
//   - [storageApi.deleteStorage]
//   - [storageApi.getStorage]
//   - [storageApi.listStorageProviders]
//   - [storageApi.listStorages]
//   - [storageApi.updateStorage]
//   - [storageClassApi.createStorageClass]
//   - [storageClassApi.deleteStorageClass]
//   - [storageClassApi.getStorageClass]
//   - [storageClassApi.listStorageClassPvcs]
//   - [storageClassApi.listStorageClasses]
//   - [storageClassApi.listStorageProvisioners]
//   - [storageClassApi.updateStorageClass]
//   - [tagApi.createTag]
//   - [tagApi.deleteTags]
//   - [tagApi.getTags]
//   - [tagApi.listOrgTags]
//   - [tagApi.updateTag]
//   - [tlsApi.getTLSCertificate]
//   - [tlsApi.tlsSwitcher]
//   - [userApi.createUserApikey]
//   - [userApi.deleteApikey]
//   - [userApi.patchAPIkey]
//   - [userApi.readUserApikeys]
//   - [viewApi.getTreeView]
//   - [viewApi.getViewByCluster]
//   - [vipPoolApi.createVIPPool]
//   - [vipPoolApi.deleteVIPPool]
//   - [vipPoolApi.listVIPPool]
//   - [whitelistApi.deleteIPWhiteList]
//   - [zoneApi.createZone]
//   - [zoneApi.deleteZone]
//   - [zoneApi.getZones]
//   - [zoneApi.listZones]
//   - [zoneApi.updateZone]
package admin