## go-terraform-tester

Clone this repository and run:

```bash
go mod tidy
```

If you want to run this software properly, because of a bug, you will to run it this way:

```bash
go run . 1>/dev/null
```

Be mindful of variables with sensitive values. You will have to manually add the `sensitive` option to the right tests because there is no way for Terratest to get the `sensitive = true` attribute.
