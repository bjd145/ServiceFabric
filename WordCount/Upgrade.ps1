param(
    [string] $SFCluster,  #xyzcontainers.westus.cloudapp.azure.com:19000
    [string] $Package,
    [string] $ApplicationName,
    [string] $ServerCert,
    [string] $AdminCert,
    [string] $Version = "2.0.0",
    [switch] $Insecure
)

if($Insecure) {
    Connect-ServiceFabricCluster -ConnectionEndpoint $SFCluster
}
else {
    Connect-ServiceFabricCluster -ConnectionEndpoint $SFCluster `
        -X509Credential -ServerCertThumbprint $ServerCert `
        -FindType FindByThumbprint -FindValue $AdminCert `
        -StoreLocation CurrentUser -StoreName 'My'
}

$ImageStore = $ApplicationName
$FQDNApplicationName = "fabric:/{0}" -f $ApplicationName

Copy-ServiceFabricApplicationPackage -ApplicationPackagePath $ApplicationPackage -ApplicationPackagePathInImageStore $ImageStore
Register-ServiceFabricApplicationType -ApplicationPathInImageStore $ImageStore
Remove-ServiceFabricApplicationPackage -ApplicationPackagePathInImageStore $ImageStore -ImageStoreConnectionString fabric:ImageStore

Start-ServiceFabricApplicationUpgrade -ApplicationName $FQDNApplicationName -ApplicationTypeVersion $Version `
    -HealthCheckStableDurationSec 60 -UpgradeDomainTimeoutSec 1200 -UpgradeTimeout 3000  -FailureAction Rollback -Monitored