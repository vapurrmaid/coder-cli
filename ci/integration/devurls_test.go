package integration

import (
	"context"
	"testing"
	"time"

	"cdr.dev/coder-cli/ci/tcli"
	"cdr.dev/slog/sloggers/slogtest/assert"
)

func TestDevURLCLI(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()

	c, err := tcli.NewContainerRunner(ctx, &tcli.ContainerConfig{
		Image: "codercom/enterprise-dev",
		Name:  "coder-cli-devurl-tests",
		BindMounts: map[string]string{
			binpath: "/bin/coder",
		},
	})
	assert.Success(t, "new run container", err)
	defer c.Close()

	c.Run(ctx, "which coder").Assert(t,
		tcli.Success(),
		tcli.StdoutMatches("/usr/sbin/coder"),
		tcli.StderrEmpty(),
	)

	c.Run(ctx, "coder urls ls").Assert(t,
		tcli.Error(),
	)

	// The following cannot be enabled nor verified until either the
	// integration testing dogfood target has environments created, or
	// we implement the 'env create' command for coder-cli to create our
	// own here.

	// If we were to create an env ourselves ... we could test devurls something like

	// // == Login
	// headlessLogin(ctx, t, c)

	// // == urls ls should fail w/o supplying an envname
	// c.Run(ctx, "coder urls ls").Assert(t,
	// 	tcli.Error(),
	// )

	// // == env creation should succeed
	// c.Run(ctx, "coder envs create env1 --from image1 --cores 1 --ram 2gb --disk 10gb --nogpu").Assert(t,
	// 	tcli.Success())

	// // == urls ls should succeed for a newly-created environment
	// var durl entclient.DevURL
	// c.Run(ctx, `coder urls ls -o json`).Assert(t,
	// 	tcli.Success(),
	// 	jsonUnmarshals(&durl), // though if a new env, durl should be empty
	// )

	// // == devurl creation w/default PRIVATE access
	// c.Run(ctx, `coder urls create env1 3000`).Assert(t,
	// 	tcli.Success())

	// // == devurl create w/access == AUTHED
	// c.Run(ctx, `coder urls create env1 3001 --access=AUTHED`).Assert(t,
	// 	tcli.Success())

	// // == devurl create with name
	// c.Run(ctx, `coder urls create env1 3002 --access=PUBLIC --name=foobar`).Assert(t,
	// 	tcli.Success())

	// // == devurl ls should return well-formed entries incl. one with AUTHED access
	// c.Run(ctx, `coder urls ls env1 -o json | jq -c '.[] | select( .access == "AUTHED")'`).Assert(t,
	// 	tcli.Success(),
	// 	jsonUnmarshals(&durl))

	// // == devurl ls should return well-formed entries incl. one with name 'foobar'
	// c.Run(ctx, `coder urls ls env1 -o json | jq -c '.[] | select( .name == "foobar")'`).Assert(t,
	// 	tcli.Success(),
	// 	jsonUnmarshals(&durl))

	// // == devurl del should function
	// c.Run(ctx, `coder urls del env1 3002`).Assert(t,
	// 	tcli.Success())

	// // == removed devurl should no longer be there
	// c.Run(ctx, `coder urls ls env1 -o json | jq -c '.[] | select( .name == "foobar")'`).Assert(t,
	// 	tcli.Error(),
	// 	jsonUnmarshals(&durl))

}
