**configcomb** is a tool that process the configs between ServiceComb chassis and other micro service frameworks(like Istio)

**current status: demo**

The tool is still in demo state, we still need more discussions on the idea



**play with configcomb**

- Install dependencies

  The project dependencies are managed by [glide](https://github.com/Masterminds/glide):

  `glide intall`

  And then all the dependencies are just there.

- Run it!

  The project is shipped with a sample chassis config from a demo project [go-bmi](https://github.com/ServiceComb-samples/go-bmi). Just upload the `calculator.bmi.chassis.yaml` to the demo web API, pass `service_name` and `namespace` in headers, then the equivalent Istio **rate limit rules** will be created(this is the only config that the demo can handle for now).

  Make sure the web daemon is running under a environment that it can access a kubernetes cluster(`~/.kube/config` is ready or running inside a k8s cluster) 

  Test it with curl:

  ```bash
  curl -X POST --data-binary @./sample-configs/calculator.bmi.chassis.yaml \
  	 -H 'service_name: calculator' \
  	 -H 'namespace: bmi' \
  	 http://$YOUR_IP_ADDR:8882/ratelimit
  ```
