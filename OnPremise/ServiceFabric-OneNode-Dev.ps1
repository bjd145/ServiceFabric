New-Item -Path C:\Installs -ItemType Directory
Set-Location C:\Installs
Invoke-WebRequest -UseBasicParsing -Uri http://go.microsoft.com/fwlink/?LinkId=730690 -OutFile SF.zip
Expand-Archive -Path .\SF.zip -DestinationPath C:\Installs\SF
Set-Location -Path c:\Installs\SF
.\CreateServiceFabricCluster.ps1 -ClusterConfigFilePath .\ClusterConfig.Unsecure.OneNode.json -AcceptEula