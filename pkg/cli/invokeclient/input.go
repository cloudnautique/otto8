package invokeclient

import (
	"context"

	"github.com/otto8-ai/otto8/apiclient"
	"github.com/otto8-ai/otto8/apiclient/types"
	"github.com/otto8-ai/otto8/pkg/cli/textio"
)

type QuietInputter struct {
}

func (d QuietInputter) Next(_ context.Context, previous string, resp *types.InvokeResponse) (string, bool, error) {
	if resp == nil {
		return previous, true, nil
	}
	return "", false, nil
}

type VerboseInputter struct {
	client *apiclient.Client
}

func nextInput() (string, bool, error) {
	x, err := textio.Ask("Input", "")
	if err != nil {
		return "", false, err
	}
	return x, true, nil
}

func (d VerboseInputter) Next(ctx context.Context, previous string, resp *types.InvokeResponse) (string, bool, error) {
	if resp == nil {
		if previous == "" {
			return nextInput()
		}
		return previous, true, nil
	}

	return nextInput()
}
