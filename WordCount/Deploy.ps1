param(
    [string] $SFCluster,  #xyzcontainers.westus.cloudapp.azure.com:19000
    [string] $Package,
    [string] $ApplicationName,
    [string] $ServerCert,
    [string] $AdminCert,
    [switch] $Insecure
)

Import-Module "$ENV:ProgramFiles\Microsoft SDKs\Service Fabric\Tools\PSModule\ServiceFabricSDK\ServiceFabricSDK.psm1"

if($Insecure) {
    Connect-ServiceFabricCluster -ConnectionEndpoint $SFCluster
}
else {
    Connect-ServiceFabricCluster -ConnectionEndpoint $SFCluster `
        -X509Credential -ServerCertThumbprint $ServerCert `
        -FindType FindByThumbprint -FindValue $AdminCert `
        -StoreLocation CurrentUser -StoreName 'My'
}

$opts = @{
    ApplicationPackagePath =  (Get-Item $Package | Select-Object -ExpandProperty FullName)
    ApplicationName = ("fabric:/{0}" -f $ApplicationName)
}
Publish-NewServiceFabricApplication @opts
