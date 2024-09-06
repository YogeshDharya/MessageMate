# Helm Charts For Message Mate 🌇:

These charts were originally created within the ssh session of an `AWS EC2 ubuntu instance` which was first installed with `kubectl` , `kind` , `kubeEdm`

Following values were modified from the default nginx-chart provided by helm from the command `helm create mysql-chart`


- **values.yaml**  :
NOTE : Some of the following fields might not be present in the yaml fileso just create them below the image or above the port 
o image :
o tag : "latest"
o port : 3306

- **deployment.yaml**  :
o

- **service.yaml**  :
o

### Deployment Instructions  🚀 :

1. Ensure that docker is up n running and docker daemon does not have any connection refused error 

2. Ensure that your k8s cluster is well configured and

3. Environment variable:

4. Deploy the chart using:

