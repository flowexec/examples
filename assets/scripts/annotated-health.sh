#!/bin/bash
# f:name=service f:verb=check f:timeout=30s
# f:description=Check that a deployed service is healthy.
# f:args=pos:1:SERVICE_URL
curl -sf "${SERVICE_URL}/health" \
  && echo "Service is healthy" \
  || (echo "Health check failed" && exit 1)
