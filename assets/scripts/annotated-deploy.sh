#!/bin/bash
# f:name=app f:verb=deploy
# f:description=Deploy the application to a target environment.
# f:tags=deployment|fromfile
# f:params=prompt:Target environment (staging/prod)?:ENV
# f:params=secretRef:demo/message:API_KEY
# f:args=flag:dry-run:DRY_RUN
echo "Deploying to ${ENV}..."
if [ "$DRY_RUN" = "true" ]; then
  echo "DRY RUN: skipping actual deploy"
else
  echo "Using API key: ${API_KEY:0:4}****"
  echo "Deploy to ${ENV} complete"
fi
