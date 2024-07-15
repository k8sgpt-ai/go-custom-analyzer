## go-custom-analyzer

This is an example of how to extend [k8sgpt](https://github.com/k8sgpt-ai/k8sgpt.git) analyzer with a custom implementation. This enables you to run additional behaviours outside of Kubernetes or with different languages.

[See the related tutorial](https://docs.k8sgpt.ai/tutorials/custom-analyzers/)
## How to run 

You will need to run this example locally.
This example is hard coded to port 8085.
e.g. 
```
go run main.go
```
Trigger it through `grpcurl`.
e.g.

```
❯ grpcurl --plaintext localhost:8085 schema.v1.AnalyzerService/Run
{
  "result": {
    "name": "example",
    "error": [
      {
        "text": "This is an example error message!"
      }
    ],
    "details": "example"
  }
}

```

## Calling from K8sGPT

When you have this service ready to go and running somewhere.
Add it to the K8sGPT `custom-analysis` config

```
❯ cat ~/Library/Application\ Support/k8sgpt/k8sgpt.yaml
custom_analyzers:
  - name: Example
    connection:
      url: localhost
      port: 8085

```