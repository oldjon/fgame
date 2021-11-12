APPS = gatewayserver

GCFLAGS = #-gcflags "all=-N -l"
BUILDVER = `git rev-parse HEAD`
BUILDDATE = `date +%F`
TAG = -ldflags "-X 'oldjon.com/glog.VERSION=$(BUILDVER)' -X 'oldjon.com/glog.DATE=$(BUILDDATE)' -X 'oldjon.com/glog.DEBUG=$(DEBUG)'"

install:
	export GOPATH=$(PWD)\
	&& for ser in $(APPS);\
	do \
		go install $(GCFLAGS) $(TAG) -x $$ser...\
		&& echo -e '#' build $(PACKAGE)$$ser success! || exit 1;\
	done