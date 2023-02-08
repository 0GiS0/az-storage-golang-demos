# Variables
RESOURCE_GROUP="az-storage-golang-demos"
LOCATION="northeurope"
STORAGE_ACCOUNT="golangdemos"

# Create resource group
az group create -n $RESOURCE_GROUP -l $LOCATION

# Create azure storage account
az storage account create -n $STORAGE_ACCOUNT -g $RESOURCE_GROUP 

# Assign Storage Blob Data Contributor role to my user
az role assignment create --assignee $(az ad signed-in-user show --query id -o tsv) --role "Storage Blob Data Contributor" --scope $(az storage account show -n $STORAGE_ACCOUNT -g $RESOURCE_GROUP --query id -o tsv)
