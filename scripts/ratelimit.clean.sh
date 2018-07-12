#!/bin/bash

namespace=$1

if [ "$namespace" == "" ]; then
	echo "namespace needed!"
	exit 1
fi

ratelimit_resources=(quota memquota rule quotaspec quotaspecbinding)

echo "Removing '${ratelimit_resources[@]}' in namespace '${namespace}'"


for res in ${ratelimit_resources[@]}
do
	res_found=$(kubectl get $res -n $namespace 2>&1 |  grep -v 'No resources found' | grep -v NAME | awk '{print $1}')
	arystr=$(echo $res_found)
	if [ "$arystr" != "" ]; then
		kubectl delete $res -n $namespace $res_found
	fi
done
