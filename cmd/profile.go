package cmd

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/textileio/textile-go/core"
	"gopkg.in/abiosoft/ishell.v2"
)

func publishProfile(c *ishell.Context) {
	entry, err := core.Node.PublishAccountProfile(nil)
	if err != nil {
		c.Err(err)
		return
	}
	if entry == nil {
		c.Println(color.New(color.FgHiRed).SprintFunc()("profile does not exist"))
		return
	}

	green := color.New(color.FgHiGreen).SprintFunc()
	c.Println(green(fmt.Sprintf("ok, published %s -> %s", entry.Name, entry.Value)))
}

func resolveProfile(c *ishell.Context) {
	var name string
	if len(c.Args) == 0 {
		id, err := core.Node.Id()
		if err != nil {
			c.Err(err)
			return
		}
		name = id.Pretty()
	} else {
		name = c.Args[0]
	}

	entry, err := core.Node.ResolveName(name)
	if err != nil {
		c.Err(err)
		return
	}

	green := color.New(color.FgHiGreen).SprintFunc()
	c.Println(green(entry.String()))
}

func getProfile(c *ishell.Context) {
	var id string
	if len(c.Args) == 0 {
		pid, err := core.Node.Id()
		if err != nil {
			c.Err(err)
			return
		}
		id = pid.Pretty()
	} else {
		id = c.Args[0]
	}

	prof, err := core.Node.GetAccountProfile(id)
	if err != nil {
		c.Err(err)
		return
	}

	green := color.New(color.FgHiGreen).SprintFunc()
	if prof.Address != "" {
		c.Println(green(fmt.Sprintf("address:         %s", prof.Address)))
	}
	if prof.Username != "" {
		c.Println(green(fmt.Sprintf("username:   %s", prof.Username)))
	}
	if prof.AvatarUri != "" {
		c.Println(green(fmt.Sprintf("avatar_uri: %s", prof.AvatarUri)))
	}
}

func setAvatarId(c *ishell.Context) {
	if len(c.Args) == 0 {
		c.Err(errors.New("missing photo id"))
		return
	}
	id := c.Args[0]

	if err := core.Node.SetAvatar(id); err != nil {
		c.Err(err)
		return
	}

	green := color.New(color.FgHiGreen).SprintFunc()
	c.Println(green("ok, updated"))
}
