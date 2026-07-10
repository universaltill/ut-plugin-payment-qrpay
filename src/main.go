// QR Pay event handler — compiled to a WASI command (GOOS=wasip1 GOARCH=wasm)
// and executed in-process by the till's wazero runtime. The till passes the
// event as JSON on stdin; stdout (JSON) is logged to the audit trail and
// stderr goes to the POS log.
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type event struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Timestamp string `json:"timestamp"`
	Payload   struct {
		SaleID    string `json:"sale_id"`
		Method    string `json:"method"`
		Amount    int64  `json:"amount"`
		Reference string `json:"reference"`
	} `json:"payload"`
}

func main() {
	var ev event
	if err := json.NewDecoder(os.Stdin).Decode(&ev); err != nil {
		fmt.Fprintf(os.Stderr, "qrpay: bad event: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stderr, "qrpay: settling sale %s — %d minor units (%s)\n",
		ev.Payload.SaleID, ev.Payload.Amount, ev.Type)

	// v1: emit the payment QR content for the customer display / audit.
	// A real gateway call needs the (permission-gated) HTTP host function.
	_ = json.NewEncoder(os.Stdout).Encode(map[string]any{
		"handled": true,
		"sale_id": ev.Payload.SaleID,
		"qr":      fmt.Sprintf("unitill://pay?sale=%s&amount=%d", ev.Payload.SaleID, ev.Payload.Amount),
	})
}
