#!/bin/bash
#strip -s  postpusher
scp -r /Users/mattik/work/src/gitlab.com/pasinworld/pms/frontend/index.html root@178.62.14.228:/opt/pms_front/index.html
scp -r /Users/mattik/work/src/gitlab.com/pasinworld/pms/frontend/src/* root@178.62.14.228:/opt/pms_front/src
scp -r /Users/mattik/work/src/gitlab.com/pasinworld/pms/frontend/static/* root@178.62.14.228:/opt/pms_front/static
scp -r /Users/mattik/work/src/gitlab.com/pasinworld/pms/frontend/package.json root@178.62.14.228:/opt/pms_front/package.json
