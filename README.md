Script made in golang to complete the "data" machine from VulnLab.

The script decodes the "password" present in the grafana.db file to a readable hash format, which can be decrypted using hashcat.

After inserting the "users" string, just run it using `go run decode.go` and it will create a file named "output.txt" with the desired content.

