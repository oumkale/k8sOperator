# /*
# Copyright 2022.

# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at

#     http://www.apache.org/licenses/LICENSE-2.0

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
# */

set -e

if [[ -z $GOPATH ]]; then
	echo "Setting GOPATH to ~/go"
	GOPATH=~/go
fi

if [[ ! -d "${GOPATH}/src/k8s.io/code-generator" ]]; then
  echo ">>>>>> k8s.io/code-generator of v0.20.0 missing from GOPATH"
  echo ">>>>>> Cloning https://github.com/kubernetes/code-generator with tag v0.20.0 under '${GOPATH}/src/k8s.io'"
  git clone -b v0.20.0 https://github.com/kubernetes/code-generator ${GOPATH}/src/k8s.io/code-generator
fi
# Switching to v0.15.12 if already cloned
git --git-dir=${GOPATH}/src/k8s.io/code-generator/.git  --work-tree=${GOPATH}/src/k8s.io/code-generator checkout v0.20.0

${GOPATH}/src/k8s.io/code-generator/generate-groups.sh all \
  git.com.info/users/oumk/repos/podtatohead/pkg/client git.com.info/users/oumk/repos/podtatohead/api \
  oumk:v1alpha1 \
  --go-header-file hack/boilerplate.go.txt -v 10
