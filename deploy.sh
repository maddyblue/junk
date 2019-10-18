#!/bin/sh

set -ex

gcloud functions deploy Avg --env-vars-file env.yaml --runtime go111
