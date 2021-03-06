// Copyright (c) 2019-present Mattermost, Inc. All Rights Reserved.
// See License for license information.

package msgraph

import (
	"net/http"

	"github.com/mattermost/mattermost-plugin-mscalendar/server/remote"
)

func (c *client) GetMailboxSettings(remoteUserID string) (*remote.MailboxSettings, error) {
	u := c.rbuilder.Users().ID(remoteUserID).URL() + "/mailboxSettings"
	out := &remote.MailboxSettings{}

	_, err := c.CallJSON(http.MethodGet, u, nil, out)
	return out, err
}
