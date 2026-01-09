# Instance Metadata Service

> The Instance Metadata Service (IMDS) is a specialized, internal service available on cloud virtual machines (e.g., AWS EC2, Azure VM) that allows instances to access configuration data about themselves, such as IP addresses, IAM role credentials, and user data 

```bash
go get github.com/si3nloong/imds
```

## Usage

```go
import "github.com/si3nloong/imds"

func main() {
    instanceID, _ := imds.InstanceID()
    println(instanceID) // A8AE896D-1C03-50A2-83CE-5FB4D52A6442
}
```