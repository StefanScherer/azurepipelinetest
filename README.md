# azurepipelinetest
[![Build Status](https://stefanscherer.visualstudio.com/azurepipelinetest/_apis/build/status/azurepipelinetest-CI?branchName=master)](https://stefanscherer.visualstudio.com/azurepipelinetest/_build/latest?definitionId=2?branchName=master)

Build a multi-arch Docker image with Azure Pipelines. All build steps are defined in the `azure-pipelines.yml` file.

The following jobs are configured:
- Linux AMD 64bit
- Linux ARM 32bit
- Linux ARM 64bit
- Windows Server 2016
- Windows Server 1803
- Windows Server 2019

After each Docker image is pushed to Docker Hub with a platform specific tag the pipeline
activates the `docker manifest` command to create a manifest list of all the six images.

The manifest list is pushed to Docker Hub [stefanscherer/hello:latest](https://cloud.docker.com/repository/docker/stefanscherer/hello)

For an introduction to Azure Pipelines watch this video https://www.youtube.com/watch?v=vlBuNM6Wzic 
