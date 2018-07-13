#!/bin/bash

namespace=$1

if [ "$namespace" == "" ]; then
	echo "namespace needed!"
	exit 1
fi

ratelimit_resources=(quotas memquotas rules quotaspecs quotaspecbindings)

echo "Getting '${ratelimit_resources[@]}' in namespace '${namespace}'"


for res in ${ratelimit_resources[@]}
do
	echo '-----------------' $res:
	kubectl get $res -n $namespace 2>&1
done
