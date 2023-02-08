# Create cobra app
cobra-cli init --viper

# Add commands
cobra-cli add create
cobra-cli add container -p createCmd
cobra-cli add blob -p createCmd
cobra-cli add list
cobra-cli add containers -p listCmd
cobra-cli add blobs -p listCmd
cobra-cli add download
cobra-cli add delete

# Test commands

# Create container
go run . create container --name code
go run . create container --name pics --public

# List containers
go run . list containers

# Upload file
go run . create blob -n main.go -c code -p main.go
go run . create blob -n licencia -c code -p LICENSE


# List blobs
go run . list blobs -c code

# Download blob
go run . download blob -c code -n main.go -p prueba_main.go
go run . download blob -c code -n licencia -p prueba_licencia

# Delete blob
go run . delete blob -c code -n main.go