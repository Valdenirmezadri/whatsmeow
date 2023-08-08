package whatsmeow

import (
	"sync"

	"go.mau.fi/whatsmeow/types/events"
)

type receiptRetryType struct {
	l    sync.Mutex
	list map[*Client]map[string]uint
}

var receiptRetry = receiptRetryType{
	l:    sync.Mutex{},
	list: make(map[*Client]map[string]uint),
}

func (r *receiptRetryType) _total(cli *Client, messageID string) uint {
	if list, ok := r.list[cli]; ok {
		if total, ok := list[messageID]; ok {
			return total
		}

		list[messageID] = uint(1)
		return list[messageID]
	}

	counter := make(map[string]uint)
	counter[messageID] = uint(1)
	r.list[cli] = counter
	return 1
}

func (r *receiptRetryType) _add(cli *Client, messageID string) {
	total := r._total(cli, messageID)
	total = total + 1

	r.list[cli][messageID] = total
}

func (r *receiptRetryType) stopRetry(cli *Client, messageID string) bool {
	r.l.Lock()
	defer r.l.Unlock()

	if r._total(cli, messageID) > 3 {
		return true
	}

	r._add(cli, messageID)
	return false
}

func (r *receiptRetryType) canRetry(cli *Client, messageID string, receipt *events.Receipt) bool {
	if r.stopRetry(cli, messageID) {
		return false
	}

	users, err := cli.IsOnWhatsApp([]string{"+" + receipt.Chat.String()})

	if err != nil {
		return false
	}

	if len(users) == 0 {
		return false
	}

	user := users[0]

	if user.JID == cli.GetOwnID() {
		return false
	}

	if user.VerifiedName == nil {
		return true
	}

	if user.VerifiedName.Certificate.ServerSignature == nil {
		return true
	}

	//ServerSignature not null = verified account
	return false
}
