# Real-World Example: Corporate Air-Gapped Network

## The Scenario

GlobalCorp has a highly secure department working on trade secrets.
- **The Network:** Completely disconnected from the corporate WAN and the internet.
- **The Policy:** No data can be transferred via USB. All transfers must go through a secure, audited gateway.
- **The Tool:** GitSynq is used over the secure gateway (which supports SSH).

## The Setup

- Developers work on "Low-Side" workstations (internet-connected).
- Production code lives on "High-Side" servers (air-gapped).

## The Workflow

1. **Approval:** Source code is reviewed on the Low-Side.
2. **Deployment:** Once approved, a developer runs `gitsync push`.
   - The bundle passes through the automated security scanner on the gateway.
   - If clean, it's transferred to the High-Side server.
3. **Audit:** GitSynq logs the transfer, providing a clear audit trail of exactly which commits were moved to the secure network.

## Why this is a "Good Fit"

- **Auditability:** Git history provides a much better audit trail than a simple file list.
- **Predictability:** The High-Side environment is guaranteed to match the approved Low-Side state.
- **Efficiency:** Secure gateways are often slow; incremental bundles minimize the bandwidth needed.

## Why this is a "Bad Fit"

- If the "security scanner" cannot inspect binary `.bundle` files, this workflow might be blocked by policy.
- If the gateway requires a multi-factor authentication (MFA) that GitSynq's automated SSH client cannot handle.

## Real-World Pro-Tip

GlobalCorp uses a custom `post-push` hook in GitSynq to automatically trigger a build and test pipeline on the High-Side as soon as a bundle is merged.
