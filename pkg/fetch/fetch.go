// Copyright 2015-2016 Apcera Inc. All rights reserved.

package fetch

import (
	"fmt"
	"net/url"

	"github.com/apcera/kurma/pkg/local"
	"github.com/apcera/kurma/pkg/remote/aci"
	"github.com/apcera/kurma/pkg/remote/docker"
	"github.com/apcera/kurma/pkg/remote/http"

	"github.com/apcera/util/tempfile"

	"github.com/appc/spec/schema/types"
)

// Fetch loads a container image. Images may be sourced from the local machine,
// or may be retrieved from a remote server.
func Fetch(imageURI string, labels map[types.ACIdentifier]string, insecure bool) (tempfile.ReadSeekCloser, error) {
	u, err := url.Parse(imageURI)
	if err != nil {
		return nil, err
	}

	// TODO: re-introduce local retrieval.
	switch u.Scheme {
	case "file":

	case "http", "https":
		puller := http.New()

		r, err := puller.Pull(imageURI)
		if err != nil {
			return nil, err
		}
		return tempfile.New(r)
	case "docker":
		puller := docker.New(insecure)

		r, err := puller.Pull(imageURI)
		if err != nil {
			return nil, err
		}
		return tempfile.New(r)
	case "aci", "":
		puller := aci.New(insecure, labels)

		r, err := puller.Pull(imageURI)
		if err != nil {
			return nil, err
		}
		return tempfile.New(r)
	default:
		return nil, fmt.Errorf("%q scheme not supported", u.Scheme)
	}
}
