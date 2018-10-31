## Overview 
* This is a very simple executable to walk through Service Fabric's Application Mode. 
* This is based on the work of github.com/containous/traefik-extra-service-fabric and github.com/jjcollinge/servicefabric
* This was created this to validate a bug that I was seeing with Traefik and Service when using stateless servics with multiple partitions. 
    * https://github.com/containous/traefik-extra-service-fabric/issues/40

## Build 
* go get "github.com/jjcollinge/servicefabric"
* go build 

## Run 
* You need a client certificate that Service Fabric is configured to use for authentication. Can be either Admin or Read-Only certificate 
* The certificate has to be split into two files name "servicefabric.crt" and "servicefabric.key"
* If you have a pfx file then you can use openssl to generate them:
    * openssl pkcs12 -in SfAdminCert.pfx -nocerts -nodes -out servicefabric.key
    * openssl pkcs12 -in SfAdminCert.pfx -clcerts -nokeys -out servicefabric.crt
* Command Line - ./main localhost:19080 (or whatever your Service Fabric portal hostname is)