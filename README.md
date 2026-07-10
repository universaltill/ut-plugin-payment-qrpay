# ut-plugin-payment-qrpay

A `payment`-type plugin (ADR-0009 naming: `ut-plugin-{type}-{name}`): adds a
**QR Pay** tender method to the Universal Till Pay tab.

When a sale completes with this method, the POS records the payment and
publishes `payment.qrpay.requested` (`sale_id`, `amount` in minor units,
`reference`) on the plugin event bus. A future `runtime:"wasm"` handler
(ADR-0001) will subscribe to render the QR and confirm settlement; today the
event is audited so back-office reconciliation can pick it up.

## Release
Tag `v<version>` → CI validates the manifest, packages a universal
`tar.gz`, publishes to the marketplace and (dev) auto-approves.
Secrets: `MARKETPLACE_BASE_URL`, `MARKETPLACE_UPLOAD_TOKEN`.
Vars: `AUTO_APPROVE`, `MARKETPLACE_LISTING_ID` (set after first publish).
