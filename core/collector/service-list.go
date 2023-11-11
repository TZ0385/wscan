/**
* @Author: shaochuyu
* @Date: 5/7/2022 11:30
 */
package collector

import (
	"context"
	"io"
	"wscan/core/resource"
)

type serviceListCollect struct {
	r io.ReadCloser
}

func (*serviceListCollect) FitOut(context.Context, []string) (chan resource.Resource, error) {
	resource.ServiceFromAddr()
	return nil, nil
}
