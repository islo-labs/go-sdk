# Reference
## tenants
<details><summary><code>client.Tenants.ListTenantComputeRegions() -> *gosdk.TenantRegionsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Return the compute regions the authenticated tenant may use, including the API and WebSocket base URLs for each region.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Tenants.ListTenantComputeRegions(
    context.TODO(),
)
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## knowledge
<details><summary><code>client.Knowledge.ListKnowledge() -> *gosdk.PaginatedKnowledgeResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ListKnowledgeRequest{}
client.Knowledge.ListKnowledge(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**level:** `*gosdk.KnowledgeLevel` 
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*gosdk.KnowledgeLevel` 
    
</dd>
</dl>

<dl>
<dd>

**tag:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**repository:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` — Search identifier or body text
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Knowledge.CreateKnowledge(request) -> *gosdk.KnowledgeItemResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.KnowledgeItemCreate{
    Slug: "slug",
}
client.Knowledge.CreateKnowledge(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**slug:** `string` — Unique lowercase identifier (letters, digits, hyphens). Set at creation and cannot be changed.
    
</dd>
</dl>

<dl>
<dd>

**level:** `*gosdk.KnowledgeLevel` 
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*gosdk.KnowledgeLevel` 
    
</dd>
</dl>

<dl>
<dd>

**format:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**body:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]any` 
    
</dd>
</dl>

<dl>
<dd>

**links:** `[]*gosdk.KnowledgeLinkInput` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Knowledge.ListKnowledgeTags() -> *gosdk.KnowledgeTagsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Knowledge.ListKnowledgeTags(
    context.TODO(),
)
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Knowledge.CreateKnowledgeMedia(request) -> *gosdk.KnowledgeItemResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.BodyCreateKnowledgeMedia{
    File: strings.NewReader(
        "",
    ),
    Item: "item",
}
client.Knowledge.CreateKnowledgeMedia(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Knowledge.GetKnowledge(Identifier) -> *gosdk.KnowledgeItemResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetKnowledgeRequest{
    Identifier: "identifier",
}
client.Knowledge.GetKnowledge(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**identifier:** `string` — Unique lowercase identifier (letters, digits, hyphens). Set at creation and cannot be changed.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Knowledge.DeleteKnowledge(Identifier) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.DeleteKnowledgeRequest{
    Identifier: "identifier",
}
client.Knowledge.DeleteKnowledge(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**identifier:** `string` — Unique lowercase identifier (letters, digits, hyphens). Set at creation and cannot be changed.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Knowledge.UpdateKnowledge(Identifier, request) -> *gosdk.KnowledgeItemResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.KnowledgeItemUpdate{
    Identifier: "identifier",
}
client.Knowledge.UpdateKnowledge(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**identifier:** `string` — Unique lowercase identifier (letters, digits, hyphens). Set at creation and cannot be changed.
    
</dd>
</dl>

<dl>
<dd>

**level:** `*gosdk.KnowledgeLevel` 
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*gosdk.KnowledgeLevel` 
    
</dd>
</dl>

<dl>
<dd>

**format:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**body:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]any` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*gosdk.KnowledgeStatus` 
    
</dd>
</dl>

<dl>
<dd>

**links:** `[]*gosdk.KnowledgeLinkInput` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Knowledge.GetKnowledgeContent(Identifier) -> any</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetKnowledgeContentRequest{
    Identifier: "identifier",
}
client.Knowledge.GetKnowledgeContent(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**identifier:** `string` — Unique lowercase identifier (letters, digits, hyphens). Set at creation and cannot be changed.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Knowledge.PutKnowledgeContent(Identifier, request) -> *gosdk.KnowledgeItemResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.BodyPutKnowledgeContent{
    Identifier: "identifier",
    File: strings.NewReader(
        "",
    ),
}
client.Knowledge.PutKnowledgeContent(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**identifier:** `string` — Unique lowercase identifier (letters, digits, hyphens). Set at creation and cannot be changed.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Knowledge.ListKnowledgeVersions(Identifier) -> *gosdk.PaginatedKnowledgeVersionResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ListKnowledgeVersionsRequest{
    Identifier: "identifier",
}
client.Knowledge.ListKnowledgeVersions(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**identifier:** `string` — Unique lowercase identifier (letters, digits, hyphens). Set at creation and cannot be changed.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Knowledge.GetKnowledgeVersion(Identifier, VersionNumber) -> *gosdk.KnowledgeVersionResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetKnowledgeVersionRequest{
    Identifier: "identifier",
    VersionNumber: 1,
}
client.Knowledge.GetKnowledgeVersion(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**identifier:** `string` — Unique lowercase identifier (letters, digits, hyphens). Set at creation and cannot be changed.
    
</dd>
</dl>

<dl>
<dd>

**versionNumber:** `int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Knowledge.GetKnowledgeVersionContent(Identifier, VersionNumber) -> any</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetKnowledgeVersionContentRequest{
    Identifier: "identifier",
    VersionNumber: 1,
}
client.Knowledge.GetKnowledgeVersionContent(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**identifier:** `string` — Unique lowercase identifier (letters, digits, hyphens). Set at creation and cannot be changed.
    
</dd>
</dl>

<dl>
<dd>

**versionNumber:** `int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Knowledge.RestoreKnowledgeVersion(Identifier, request) -> *gosdk.KnowledgeItemResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.KnowledgeRestoreRequest{
    Identifier: "identifier",
    VersionNumber: 1,
}
client.Knowledge.RestoreKnowledgeVersion(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**identifier:** `string` — Unique lowercase identifier (letters, digits, hyphens). Set at creation and cannot be changed.
    
</dd>
</dl>

<dl>
<dd>

**versionNumber:** `int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Credits
<details><summary><code>client.Credits.GetCreditBalance() -> *gosdk.CreditBalance</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Return the tenant's available prepaid credit balance in cents.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Credits.GetCreditBalance(
    context.TODO(),
)
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## integrations
<details><summary><code>client.Integrations.ListIntegrationProviders() -> *gosdk.IntegrationProvidersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Return the integration providers available to connect from Islo, including the supported authentication methods and connection scopes.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Integrations.ListIntegrationProviders(
    context.TODO(),
)
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Integrations.ListIntegrationTriggers() -> *gosdk.TriggerCatalogListResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Integrations.ListIntegrationTriggers(
    context.TODO(),
)
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Integrations.GetIntegrationTrigger(Provider, TriggerName) -> *gosdk.TriggerCatalogItem</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetIntegrationTriggerRequest{
    Provider: "provider",
    TriggerName: "trigger_name",
}
client.Integrations.GetIntegrationTrigger(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**provider:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**triggerName:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Integrations.ListIntegrations() -> *gosdk.IntegrationListResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List the integrations the user/tenant has connected.

Includes preset providers (from the PROVIDERS registry) and tenant-scoped
custom outbound apps (filtered out of Descope's load_all_applications).
Returns one entry per connected (provider, scope, auth_type) slot, so a
provider with both a personal api_key and a personal oauth token will
appear twice. Disconnected slots are not emitted; clients that need a
list of available-but-not-connected providers should call
``GET /integrations/providers`` instead.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Integrations.ListIntegrations(
    context.TODO(),
)
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Integrations.ListCustomServices() -> *gosdk.CustomServicesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List custom service definitions in the current tenant (catalog view).

Returns every custom Descope app belonging to the tenant regardless of
connection status, so the Add Integration picker can surface them for
any tenant member to connect to. Connection state (per-user/per-workspace
tokens) lives on ``GET /integrations``; this endpoint is purely the
service catalog.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Integrations.ListCustomServices(
    context.TODO(),
)
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Integrations.CreateCustomService(request) -> *gosdk.CustomServiceCreateResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a tenant-scoped custom Descope outbound app.

Returns the ``app_id`` so the frontend can immediately kick off the
connect flow (OAuth) or surface the API key form. Presets do not pass
through this endpoint -- their app ids come straight from
``GET /integrations/providers``.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.CustomServiceCreateRequest{
    Custom: &gosdk.CustomIntegration{
        Name: "name",
        Slug: "slug",
    },
}
client.Integrations.CreateCustomService(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**custom:** `*gosdk.CustomIntegration` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Integrations.DisconnectCustomIntegration(DescopeAppID) -> map[string]any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Disconnect a custom integration by its Descope app ID.

Authorization is by deterministic-ID prefix: only apps whose ID matches
``cust-{tenant-prefix}-`` are accepted, which scopes the operation to the
caller's workspace without a DB lookup. ``scope`` selects which side's
tokens to revoke (per-user vs tenant-wide); ``delete_app=true`` removes
the Descope app entirely (affects every user in the workspace).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.DisconnectCustomIntegrationRequest{
    DescopeAppID: "descope_app_id",
}
client.Integrations.DisconnectCustomIntegration(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**descopeAppID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**scope:** `*gosdk.IntegrationLevel` — Which token to revoke: 'user' (this user's personal) or 'tenant' (workspace)
    
</dd>
</dl>

<dl>
<dd>

**deleteApp:** `*bool` — Also remove the Descope outbound app entirely (affects every user in this workspace)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Integrations.GetIntegrationStatus(Provider) -> *gosdk.IntegrationDetailResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the detailed status of a specific integration.

Returns both user-level and tenant-level connection status independently.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetIntegrationStatusRequest{
    Provider: "provider",
}
client.Integrations.GetIntegrationStatus(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**provider:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Integrations.DisconnectIntegration(Provider) -> map[string]any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Disconnect/revoke an integration.

Args:
    provider: Provider name
    level: Which level to disconnect (USER or TENANT)
    auth_type: Optional. Defaults to provider's primary type.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.DisconnectIntegrationRequest{
    Provider: "provider",
}
client.Integrations.DisconnectIntegration(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**provider:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**level:** `*gosdk.IntegrationLevel` 
    
</dd>
</dl>

<dl>
<dd>

**authType:** `*gosdk.AuthMethod` — oauth or api_key
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Integrations.ListConnectedIntegrationTriggers() -> *gosdk.TriggerCatalogListResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Integrations.ListConnectedIntegrationTriggers(
    context.TODO(),
)
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## gateway-profiles
<details><summary><code>client.GatewayProfiles.ListGatewayProfiles() -> []*gosdk.GatewayProfileResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.GatewayProfiles.ListGatewayProfiles(
    context.TODO(),
)
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.GatewayProfiles.CreateGatewayProfile(request) -> *gosdk.GatewayProfileResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GatewayProfileCreate{
    Name: "name",
}
client.GatewayProfiles.CreateGatewayProfile(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**defaultAction:** `*gosdk.GatewayAction` 
    
</dd>
</dl>

<dl>
<dd>

**internetEnabled:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**isDefault:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**cloudRole:** `*string` — Cloud role public ID (UUID)
    
</dd>
</dl>

<dl>
<dd>

**integrationPolicy:** `*gosdk.GatewayProfileCreateIntegrationPolicy` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.GatewayProfiles.GetGatewayProfile(ProfileID) -> *gosdk.GatewayProfileDetailResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetGatewayProfileRequest{
    ProfileID: "profile_id",
}
client.GatewayProfiles.GetGatewayProfile(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**profileID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.GatewayProfiles.DeleteGatewayProfile(ProfileID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.DeleteGatewayProfileRequest{
    ProfileID: "profile_id",
}
client.GatewayProfiles.DeleteGatewayProfile(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**profileID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.GatewayProfiles.UpdateGatewayProfile(ProfileID, request) -> *gosdk.GatewayProfileResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GatewayProfileUpdate{
    ProfileID: "profile_id",
}
client.GatewayProfiles.UpdateGatewayProfile(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**profileID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**defaultAction:** `*gosdk.GatewayAction` 
    
</dd>
</dl>

<dl>
<dd>

**internetEnabled:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**isDefault:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**cloudRole:** `*string` — Cloud role public ID (UUID), empty string to unset
    
</dd>
</dl>

<dl>
<dd>

**integrationPolicy:** `*gosdk.GatewayProfileUpdateIntegrationPolicy` — Omit to leave unchanged; send {"mode": "all"} to allow all integrations
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.GatewayProfiles.CreateGatewayRule(ProfileID, request) -> *gosdk.GatewayRuleResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GatewayRuleCreate{
    ProfileID: "profile_id",
    HostPattern: "host_pattern",
}
client.GatewayProfiles.CreateGatewayRule(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**profileID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**hostPattern:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**pathPattern:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**methods:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**rateLimitRpm:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**authStrategy:** `*gosdk.AuthStrategySchema` 
    
</dd>
</dl>

<dl>
<dd>

**contentFilter:** `*gosdk.GatewayRuleCreateContentFilter` 
    
</dd>
</dl>

<dl>
<dd>

**priority:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**action:** `*gosdk.GatewayAction` 
    
</dd>
</dl>

<dl>
<dd>

**providerKey:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.GatewayProfiles.DeleteGatewayRule(ProfileID, RuleID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.DeleteGatewayRuleRequest{
    ProfileID: "profile_id",
    RuleID: "rule_id",
}
client.GatewayProfiles.DeleteGatewayRule(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**profileID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**ruleID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.GatewayProfiles.UpdateGatewayRule(ProfileID, RuleID, request) -> *gosdk.GatewayRuleResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GatewayRuleUpdate{
    ProfileID: "profile_id",
    RuleID: "rule_id",
}
client.GatewayProfiles.UpdateGatewayRule(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**profileID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**ruleID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**priority:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**hostPattern:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**pathPattern:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**methods:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**action:** `*gosdk.GatewayAction` 
    
</dd>
</dl>

<dl>
<dd>

**rateLimitRpm:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**providerKey:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**authStrategy:** `*gosdk.AuthStrategySchema` 
    
</dd>
</dl>

<dl>
<dd>

**contentFilter:** `*gosdk.GatewayRuleUpdateContentFilter` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.GatewayProfiles.ReorderGatewayRules(ProfileID, request) -> []*gosdk.GatewayRuleResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.RuleReorderRequest{
    ProfileID: "profile_id",
    Rules: []*gosdk.RuleReorderItem{
        &gosdk.RuleReorderItem{
            RuleID: "rule_id",
            Priority: 1,
        },
    },
}
client.GatewayProfiles.ReorderGatewayRules(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**profileID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**rules:** `[]*gosdk.RuleReorderItem` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Environments
<details><summary><code>client.Environments.ListEnvironments() -> []*gosdk.EnvironmentListItem</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ListEnvironmentsRequest{}
client.Environments.ListEnvironments(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Environments.CreateEnvironment(request) -> *gosdk.EnvironmentResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.EnvironmentCreate{
    Name: "name",
}
client.Environments.CreateEnvironment(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**isDefault:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**entries:** `[]*gosdk.EnvironmentCreateEntriesItem` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Environments.GetEnvironment(EnvironmentRef) -> *gosdk.EnvironmentResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetEnvironmentRequest{
    EnvironmentRef: "environment_ref",
}
client.Environments.GetEnvironment(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**environmentRef:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Environments.DeleteEnvironment(EnvironmentRef) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.DeleteEnvironmentRequest{
    EnvironmentRef: "environment_ref",
}
client.Environments.DeleteEnvironment(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**environmentRef:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Environments.UpdateEnvironment(EnvironmentRef, request) -> *gosdk.EnvironmentResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.EnvironmentUpdate{
    EnvironmentRef: "environment_ref",
}
client.Environments.UpdateEnvironment(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**environmentRef:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**isDefault:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**entries:** `[]*gosdk.EnvironmentUpdateEntriesItem` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Environments.SetDefaultEnvironment(EnvironmentRef) -> *gosdk.EnvironmentResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.SetDefaultEnvironmentRequest{
    EnvironmentRef: "environment_ref",
}
client.Environments.SetDefaultEnvironment(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**environmentRef:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Environments.UnsetDefaultEnvironment(EnvironmentRef) -> *gosdk.EnvironmentResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.UnsetDefaultEnvironmentRequest{
    EnvironmentRef: "environment_ref",
}
client.Environments.UnsetDefaultEnvironment(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**environmentRef:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## CloudRoles
<details><summary><code>client.CloudRoles.ListCloudRoles() -> []*gosdk.CloudRoleResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ListCloudRolesRequest{}
client.CloudRoles.ListCloudRoles(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**type_:** `*gosdk.CloudRoleType` — Filter by role type
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CloudRoles.CreateCloudRole(request) -> *gosdk.CloudRoleResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.CloudRoleCreate{
    Provider: gosdk.CloudProviderAws,
    RoleArn: "role_arn",
}
client.CloudRoles.CreateCloudRole(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**provider:** `*gosdk.CloudProvider` 
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*gosdk.CloudRoleType` 
    
</dd>
</dl>

<dl>
<dd>

**roleArn:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**sessionDurationSeconds:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CloudRoles.GetCloudRole(RoleID) -> *gosdk.CloudRoleResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetCloudRoleRequest{
    RoleID: "role_id",
}
client.CloudRoles.GetCloudRole(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**roleID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CloudRoles.DeleteCloudRole(RoleID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.DeleteCloudRoleRequest{
    RoleID: "role_id",
}
client.CloudRoles.DeleteCloudRole(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**roleID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CloudRoles.UpdateCloudRole(RoleID, request) -> *gosdk.CloudRoleResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.CloudRoleUpdate{
    RoleID: "role_id",
}
client.CloudRoles.UpdateCloudRole(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**roleID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**roleArn:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**sessionDurationSeconds:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**isEnabled:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## byo
<details><summary><code>client.Byo.GetByoInferenceStatus() -> *gosdk.ByoStatusResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Byo.GetByoInferenceStatus(
    context.TODO(),
)
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Byo.StartByoInferenceSetup(request) -> *gosdk.ByoSetupResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ByoSetupRequest{
    SourceKind: gosdk.ByoSourceKindDatabricks,
}
client.Byo.StartByoInferenceSetup(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sourceKind:** `*gosdk.ByoSourceKind` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## inference
<details><summary><code>client.Inference.ListInferenceModels() -> *gosdk.InferenceModelsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Inference.ListInferenceModels(
    context.TODO(),
)
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## ContainerRegistries
<details><summary><code>client.ContainerRegistries.ListContainerRegistries() -> []*gosdk.ContainerRegistryResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.ContainerRegistries.ListContainerRegistries(
    context.TODO(),
)
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ContainerRegistries.CreateContainerRegistry(request) -> *gosdk.ContainerRegistryResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ContainerRegistryCreate{
    Provider: gosdk.RegistryProviderEcr,
    RegistryHost: "registry_host",
    CloudRoleID: "cloud_role_id",
    Region: "region",
}
client.ContainerRegistries.CreateContainerRegistry(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**provider:** `*gosdk.RegistryProvider` 
    
</dd>
</dl>

<dl>
<dd>

**registryHost:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**repositoryPrefixes:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**cloudRoleID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**region:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ContainerRegistries.GetContainerRegistry(ID) -> *gosdk.ContainerRegistryResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetContainerRegistryRequest{
    ID: "id",
}
client.ContainerRegistries.GetContainerRegistry(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ContainerRegistries.DeleteContainerRegistry(ID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.DeleteContainerRegistryRequest{
    ID: "id",
}
client.ContainerRegistries.DeleteContainerRegistry(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ContainerRegistries.UpdateContainerRegistry(ID, request) -> *gosdk.ContainerRegistryResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ContainerRegistryUpdate{
    ID: "id",
}
client.ContainerRegistries.UpdateContainerRegistry(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**repositoryPrefixes:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**cloudRoleID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**isEnabled:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## jobs
<details><summary><code>client.Jobs.ValidateJobManifest(Name, request) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ValidateJobManifestRequest{
    Name: "name",
    Body: &gosdk.JobDeployRequest{
        Manifest: &gosdk.JobManifestInput{
            Job: &gosdk.JobSection{
                Name: "name",
            },
            Run: &gosdk.RunSectionInput{
                Tasks: []*gosdk.TaskInput{
                    &gosdk.TaskInput{
                        Name: "name",
                        Steps: []*gosdk.TaskStepInput{
                            &gosdk.TaskStepInput{},
                        },
                    },
                },
            },
        },
    },
}
client.Jobs.ValidateJobManifest(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*gosdk.JobDeployRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Jobs.DeployJob(Name, request) -> *gosdk.JobVersionResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.DeployJobRequest{
    Name: "name",
    Body: &gosdk.JobDeployRequest{
        Manifest: &gosdk.JobManifestInput{
            Job: &gosdk.JobSection{
                Name: "name",
            },
            Run: &gosdk.RunSectionInput{
                Tasks: []*gosdk.TaskInput{
                    &gosdk.TaskInput{
                        Name: "name",
                        Steps: []*gosdk.TaskStepInput{
                            &gosdk.TaskStepInput{},
                        },
                    },
                },
            },
        },
    },
}
client.Jobs.DeployJob(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*gosdk.JobDeployRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Jobs.GetJob(Name) -> *gosdk.JobResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetJobRequest{
    Name: "name",
}
client.Jobs.GetJob(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Jobs.DeleteJob(Name) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.DeleteJobRequest{
    Name: "name",
}
client.Jobs.DeleteJob(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Jobs.ListJobs() -> []*gosdk.JobListItem</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ListJobsRequest{}
client.Jobs.ListJobs(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Jobs.ListJobVersions(Name) -> []*gosdk.JobVersionResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ListJobVersionsRequest{
    Name: "name",
}
client.Jobs.ListJobVersions(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Jobs.GetJobVersion(Name, VersionID) -> *gosdk.JobVersionResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetJobVersionRequest{
    Name: "name",
    VersionID: "version_id",
}
client.Jobs.GetJobVersion(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**versionID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Jobs.ListJobRuns(Name) -> []*gosdk.JobRunListItem</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ListJobRunsRequest{
    Name: "name",
}
client.Jobs.ListJobRuns(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Jobs.TriggerJobRun(Name, request) -> *gosdk.JobRunResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.JobRunCreate{
    Name: "name",
}
client.Jobs.TriggerJobRun(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**versionID:** `*string` — Deployed version to run; defaults to latest
    
</dd>
</dl>

<dl>
<dd>

**region:** `*string` — Compute region override
    
</dd>
</dl>

<dl>
<dd>

**params:** `map[string]any` — Run-time parameter values (validated against [job.params])
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Jobs.GetJobRun(Name, RunID) -> *gosdk.JobRunResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetJobRunRequest{
    Name: "name",
    RunID: "run_id",
}
client.Jobs.GetJobRun(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**runID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Jobs.StopJobRun(Name, RunID, request) -> *gosdk.JobRunResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.JobRunStopRequest{
    Name: "name",
    RunID: "run_id",
}
client.Jobs.StopJobRun(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**runID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**reason:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Jobs.GetJobSchedule(Name) -> *gosdk.JobScheduleResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetJobScheduleRequest{
    Name: "name",
}
client.Jobs.GetJobSchedule(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Jobs.DeleteJobSchedule(Name) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.DeleteJobScheduleRequest{
    Name: "name",
}
client.Jobs.DeleteJobSchedule(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## JobRuns
<details><summary><code>client.JobRuns.ListAllJobRuns() -> *gosdk.ListPageJobRunListItem</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ListAllJobRunsRequest{}
client.JobRuns.ListAllJobRuns(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**sort:** `*string` — Sort order. Allowed: -created_at, created_at
    
</dd>
</dl>

<dl>
<dd>

**include:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**jobName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**createdAt:** `*gosdk.TimestampRange` — created_at range. Operators: gte, gt, lte, lt. Serialized as created_at[gte]=…&created_at[lt]=…
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.JobRuns.ListJobRunFacets() -> *gosdk.FacetsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ListJobRunFacetsRequest{
    Fields: []*string{
        gosdk.String(
            "fields",
        ),
    },
}
client.JobRuns.ListJobRunFacets(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fields:** `*string` — Facet fields to return (e.g. job_name, status)
    
</dd>
</dl>

<dl>
<dd>

**status:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**jobName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**createdAt:** `*gosdk.TimestampRange` — created_at range. Operators: gte, gt, lte, lt. Serialized as created_at[gte]=…&created_at[lt]=…
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.JobRuns.GetJobRunByID(RunID) -> *gosdk.JobRunResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetJobRunByIDRequest{
    RunID: "run_id",
}
client.JobRuns.GetJobRunByID(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**runID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## factory
<details><summary><code>client.Factory.ValidateFactoryLineManifest(Name, request) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ValidateFactoryLineManifestRequest{
    Name: "name",
    Body: &gosdk.LineDeployRequest{
        Manifest: &gosdk.LineManifestInput{
            Line: &gosdk.LineSection{
                Name: "name",
            },
            Trigger: &gosdk.LineManifestInputTrigger{
                IntegrationTrigger: &gosdk.IntegrationTriggerSectionInput{
                    Provider: "provider",
                    Name: "name",
                    Selector: &gosdk.IntegrationTriggerSectionInputSelector{
                        Github: &gosdk.GitHubRepositorySelector{},
                    },
                },
            },
            Stages: []*gosdk.LineStage{
                &gosdk.LineStage{
                    ID: "id",
                    Job: "job",
                },
            },
        },
    },
}
client.Factory.ValidateFactoryLineManifest(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*gosdk.LineDeployRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Factory.DeployFactoryLine(Name, request) -> *gosdk.LineVersionResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.DeployFactoryLineRequest{
    Name: "name",
    Body: &gosdk.LineDeployRequest{
        Manifest: &gosdk.LineManifestInput{
            Line: &gosdk.LineSection{
                Name: "name",
            },
            Trigger: &gosdk.LineManifestInputTrigger{
                IntegrationTrigger: &gosdk.IntegrationTriggerSectionInput{
                    Provider: "provider",
                    Name: "name",
                    Selector: &gosdk.IntegrationTriggerSectionInputSelector{
                        Github: &gosdk.GitHubRepositorySelector{},
                    },
                },
            },
            Stages: []*gosdk.LineStage{
                &gosdk.LineStage{
                    ID: "id",
                    Job: "job",
                },
            },
        },
    },
}
client.Factory.DeployFactoryLine(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*gosdk.LineDeployRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Factory.ListFactoryLines() -> []*gosdk.LineResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ListFactoryLinesRequest{}
client.Factory.ListFactoryLines(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Factory.GetFactoryLine(Name) -> *gosdk.LineResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetFactoryLineRequest{
    Name: "name",
}
client.Factory.GetFactoryLine(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Factory.UpdateFactoryLine(Name, request) -> *gosdk.LineResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.LineUpdate{
    Name: "name",
}
client.Factory.UpdateFactoryLine(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*gosdk.LineUpdateStatus` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Factory.ListFactoryLineVersions(Name) -> []*gosdk.LineVersionResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ListFactoryLineVersionsRequest{
    Name: "name",
}
client.Factory.ListFactoryLineVersions(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Factory.ListFactoryLineRunsForLine(Name) -> []*gosdk.LineRunSummary</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ListFactoryLineRunsForLineRequest{
    Name: "name",
}
client.Factory.ListFactoryLineRunsForLine(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Factory.TriggerFactoryLineRun(Name, request) -> *gosdk.LineRunDetail</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.LineRunCreate{
    Name: "name",
}
client.Factory.TriggerFactoryLineRun(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**versionID:** `*string` — Deployed line version to run; defaults to latest
    
</dd>
</dl>

<dl>
<dd>

**region:** `*string` — Compute region override
    
</dd>
</dl>

<dl>
<dd>

**params:** `map[string]any` 
    
</dd>
</dl>

<dl>
<dd>

**triggerPayload:** `map[string]any` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Factory.ListFactoryLineRuns() -> *gosdk.ListPageLineRunSummary</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ListFactoryLineRunsRequest{}
client.Factory.ListFactoryLineRuns(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**sort:** `*string` — Sort order. Allowed: -created_at, created_at
    
</dd>
</dl>

<dl>
<dd>

**include:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**lineName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**createdAt:** `*gosdk.TimestampRange` — created_at range. Operators: gte, gt, lte, lt. Serialized as created_at[gte]=…&created_at[lt]=…
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Factory.ListFactoryLineRunFacets() -> *gosdk.FacetsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ListFactoryLineRunFacetsRequest{
    Fields: []*string{
        gosdk.String(
            "fields",
        ),
    },
}
client.Factory.ListFactoryLineRunFacets(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fields:** `*string` — Facet fields to return (e.g. line_name, status)
    
</dd>
</dl>

<dl>
<dd>

**status:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**lineName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**createdAt:** `*gosdk.TimestampRange` — created_at range. Operators: gte, gt, lte, lt. Serialized as created_at[gte]=…&created_at[lt]=…
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Factory.GetFactoryLineRun(RunID) -> *gosdk.LineRunDetail</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetFactoryLineRunRequest{
    RunID: "run_id",
}
client.Factory.GetFactoryLineRun(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**runID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Factory.GetFactoryLineRunDebug(RunID) -> *gosdk.LineRunDebugResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Per-stage and per-step diagnostics for one line run, including the last failed stage attempt's first failing step, each step's exit code and output tails, and the sandbox environment each stage ran in.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetFactoryLineRunDebugRequest{
    RunID: "run_id",
}
client.Factory.GetFactoryLineRunDebug(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**runID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Factory.GetFactoryLineSchedule(Name) -> *gosdk.LineScheduleResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetFactoryLineScheduleRequest{
    Name: "name",
}
client.Factory.GetFactoryLineSchedule(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Factory.UpsertFactoryLineSchedule(Name, request) -> *gosdk.LineScheduleResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.LineScheduleUpdate{
    Name: "name",
    Cron: "cron",
}
client.Factory.UpsertFactoryLineSchedule(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cron:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**timezone:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**enabled:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**inputs:** `map[string]any` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Factory.DeleteFactoryLineSchedule(Name) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.DeleteFactoryLineScheduleRequest{
    Name: "name",
}
client.Factory.DeleteFactoryLineSchedule(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## ComputeEvents
<details><summary><code>client.ComputeEvents.GetComputeEvent(CommandID) -> *gosdk.ComputeEventDetailResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetComputeEventRequest{
    CommandID: "command_id",
}
client.ComputeEvents.GetComputeEvent(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**commandID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## sandboxes
<details><summary><code>client.Sandboxes.ListSandboxes() -> *gosdk.PaginatedSandboxResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List sandboxes for the authenticated tenant with optional filters and pagination.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ListSandboxesRequest{}
client.Sandboxes.ListSandboxes(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**q:** `*string` — Search term for sandbox name, image, creator, or public ID. Takes precedence over `search` when both are provided.
    
</dd>
</dl>

<dl>
<dd>

**search:** `*string` — Search term for sandbox name, image, creator, or public ID. Alias for `q`.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**namePrefix:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**createdBy:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sandboxes.CreateSandbox(request) -> *gosdk.SandboxResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new sandbox VM with the requested resources.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.CreateSandboxRequest{}
client.Sandboxes.CreateSandbox(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cacheKey:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**diskGb:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**env:** `map[string]*string` 
    
</dd>
</dl>

<dl>
<dd>

**environment:** `*string` — Environment for sandbox env and environment-owned gateway injection.
    
</dd>
</dl>

<dl>
<dd>

**gatewayProfile:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**image:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**init:** `*gosdk.SandboxInit` 
    
</dd>
</dl>

<dl>
<dd>

**internetEnabled:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**lifecycle:** `*gosdk.LifecyclePolicy` 
    
</dd>
</dl>

<dl>
<dd>

**memoryMb:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**requestID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**setupScripts:** `[]*gosdk.SetupScript` 
    
</dd>
</dl>

<dl>
<dd>

**snapshotName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**sources:** `[]*gosdk.GitSource` 
    
</dd>
</dl>

<dl>
<dd>

**vcpus:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**workdir:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sandboxes.GetSandboxByID(ID) -> *gosdk.SandboxResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Return details for a sandbox by public ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetSandboxByIDRequest{
    ID: "id",
}
client.Sandboxes.GetSandboxByID(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Sandbox public ID (UUID)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sandboxes.GetSandbox(SandboxName) -> *gosdk.SandboxResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Return details for a sandbox by name.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetSandboxRequest{
    SandboxName: "sandbox_name",
}
client.Sandboxes.GetSandbox(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sandboxName:** `string` — Sandbox name
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sandboxes.DeleteSandbox(SandboxName) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a sandbox and clean up its running VM, if any.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.DeleteSandboxRequest{
    SandboxName: "sandbox_name",
}
client.Sandboxes.DeleteSandbox(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sandboxName:** `string` — Sandbox name
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sandboxes.SandboxCreationEvents(SandboxName) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Server-sent events for live sandbox creation progress.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.SandboxCreationEventsRequest{
    SandboxName: "sandbox_name",
}
client.Sandboxes.SandboxCreationEvents(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sandboxName:** `string` — Sandbox name
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sandboxes.ExecInSandbox(SandboxName, request) -> *gosdk.ExecResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Start a command in a sandbox and return an exec ID for polling results.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ExecInSandboxRequest{
    SandboxName: "sandbox_name",
    Body: &gosdk.ExecRequest{
        Command: []string{
            "command",
        },
    },
}
client.Sandboxes.ExecInSandbox(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sandboxName:** `string` — Sandbox name
    
</dd>
</dl>

<dl>
<dd>

**request:** `*gosdk.ExecRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sandboxes.ExecInSandboxStream(SandboxName, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Stream command stdout, stderr, and exit events as Server-Sent Events.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ExecInSandboxStreamRequest{
    SandboxName: "sandbox_name",
    Body: &gosdk.ExecRequest{
        Command: []string{
            "command",
        },
    },
}
client.Sandboxes.ExecInSandboxStream(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sandboxName:** `string` — Sandbox name
    
</dd>
</dl>

<dl>
<dd>

**request:** `*gosdk.ExecRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sandboxes.GetExecResult(SandboxName, ExecID) -> *gosdk.ExecResultResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Return the captured result for a previously started sandbox command.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetExecResultRequest{
    SandboxName: "sandbox_name",
    ExecID: "exec_id",
}
client.Sandboxes.GetExecResult(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sandboxName:** `string` — Sandbox name
    
</dd>
</dl>

<dl>
<dd>

**execID:** `string` — Exec ID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sandboxes.DownloadFile(SandboxName) -> string</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Download a file from a sandbox.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.DownloadFileRequest{
    SandboxName: "sandbox_name",
    Path: "path",
}
client.Sandboxes.DownloadFile(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sandboxName:** `string` — Sandbox name
    
</dd>
</dl>

<dl>
<dd>

**path:** `string` — File path inside the sandbox
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sandboxes.UploadFile(SandboxName, request) -> *gosdk.FileUploadStatusResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Upload a file to a path inside a sandbox.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.UploadFileRequest{
    SandboxName: "sandbox_name",
    Path: "path",
    File: strings.NewReader(
        "",
    ),
}
client.Sandboxes.UploadFile(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sandboxName:** `string` — Sandbox name
    
</dd>
</dl>

<dl>
<dd>

**path:** `string` — Destination path inside the sandbox
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sandboxes.DownloadArchive(SandboxName) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Download a sandbox directory as an archive.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.DownloadArchiveRequest{
    SandboxName: "sandbox_name",
    Path: "path",
}
client.Sandboxes.DownloadArchive(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sandboxName:** `string` — Sandbox name
    
</dd>
</dl>

<dl>
<dd>

**path:** `string` — Directory path to archive inside the sandbox
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sandboxes.UploadArchive(SandboxName, request) -> *gosdk.FileUploadStatusResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Upload and extract an archive into a sandbox directory.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.UploadArchiveRequest{
    SandboxName: "sandbox_name",
    Path: "path",
    File: strings.NewReader(
        "",
    ),
}
client.Sandboxes.UploadArchive(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sandboxName:** `string` — Sandbox name
    
</dd>
</dl>

<dl>
<dd>

**path:** `string` — Destination directory inside the sandbox
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sandboxes.PauseSandbox(SandboxName) -> *gosdk.SandboxResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Pause a running sandbox VM.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.PauseSandboxRequest{
    SandboxName: "sandbox_name",
}
client.Sandboxes.PauseSandbox(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sandboxName:** `string` — Sandbox name
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sandboxes.ResumeSandbox(SandboxName) -> *gosdk.SandboxResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Resume a paused sandbox VM.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ResumeSandboxRequest{
    SandboxName: "sandbox_name",
}
client.Sandboxes.ResumeSandbox(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sandboxName:** `string` — Sandbox name
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sandboxes.ListSessions(SandboxName) -> *gosdk.ListSessionsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List persistent shell sessions in a sandbox.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ListSessionsRequest{
    SandboxName: "sandbox_name",
}
client.Sandboxes.ListSessions(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sandboxName:** `string` — Sandbox name
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sandboxes.CreateSession(SandboxName, request) -> *gosdk.CreateSessionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a persistent shell session in a sandbox.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.CreateSessionRequest{
    SandboxName: "sandbox_name",
    Name: "name",
}
client.Sandboxes.CreateSession(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sandboxName:** `string` — Sandbox name
    
</dd>
</dl>

<dl>
<dd>

**command:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**env:** `map[string]*string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**ttl:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**user:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**workdir:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sandboxes.KillSession(SandboxName, Session) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Terminate a persistent shell session in a sandbox.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.KillSessionRequest{
    SandboxName: "sandbox_name",
    Session: "session",
}
client.Sandboxes.KillSession(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sandboxName:** `string` — Sandbox name
    
</dd>
</dl>

<dl>
<dd>

**session:** `string` — Session name
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sandboxes.StopSandbox(SandboxName) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Stop the sandbox VM while keeping the sandbox record available.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.StopSandboxRequest{
    SandboxName: "sandbox_name",
}
client.Sandboxes.StopSandbox(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sandboxName:** `string` — Sandbox name
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## shares
<details><summary><code>client.Shares.ListShares(SandboxName) -> []*gosdk.ShareResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List active public shares for a sandbox.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ListSharesRequest{
    SandboxName: "sandbox_name",
}
client.Shares.ListShares(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sandboxName:** `string` — Sandbox name
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Shares.CreateShare(SandboxName, request) -> *gosdk.ShareResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a temporary public share for a sandbox port.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.CreateShareRequest{
    SandboxName: "sandbox_name",
    Port: 1,
}
client.Shares.CreateShare(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sandboxName:** `string` — Sandbox name
    
</dd>
</dl>

<dl>
<dd>

**port:** `int` 
    
</dd>
</dl>

<dl>
<dd>

**ttlSeconds:** `*int64` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Shares.RevokeShare(SandboxName, ShareID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Revoke a sandbox port share.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.RevokeShareRequest{
    SandboxName: "sandbox_name",
    ShareID: "share_id",
}
client.Shares.RevokeShare(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sandboxName:** `string` — Sandbox name
    
</dd>
</dl>

<dl>
<dd>

**shareID:** `string` — Share ID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## snapshots
<details><summary><code>client.Snapshots.ListSnapshots() -> *gosdk.PaginatedSnapshotResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List all snapshots for the current tenant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ListSnapshotsRequest{}
client.Snapshots.ListSnapshots(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Snapshots.CreateSnapshot(request) -> *gosdk.SnapshotResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a snapshot from a running sandbox. By default, waits for capture and returns a ready snapshot. Send `Prefer: respond-async` to return immediately with a saving snapshot.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.SnapshotCreate{
    SandboxName: "sandbox_name",
}
client.Snapshots.CreateSnapshot(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**sandboxName:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Snapshots.GetSnapshot(Name) -> *gosdk.SnapshotResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get snapshot details by name.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetSnapshotRequest{
    Name: "name",
}
client.Snapshots.GetSnapshot(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` — Snapshot name
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Snapshots.DeleteSnapshot(Name) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a snapshot by name.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.DeleteSnapshotRequest{
    Name: "name",
}
client.Snapshots.DeleteSnapshot(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` — Snapshot name
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## webhooks
<details><summary><code>client.Webhooks.ListIncomingWebhooks() -> []*gosdk.IncomingWebhook</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List active incoming webhooks for the tenant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Webhooks.ListIncomingWebhooks(
    context.TODO(),
)
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.CreateIncomingWebhook(request) -> *gosdk.IncomingWebhook</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a tenant-scoped incoming webhook receiver. The receiver URL accepts external webhook deliveries and routes them to a resolved sandbox.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.IncomingWebhookCreate{
    Auth: &gosdk.IncomingWebhookAuth{
        IncomingWebhookAuthZero: &gosdk.IncomingWebhookAuthZero{
            AuthType: gosdk.IncomingWebhookAuthZeroAuthTypeNone,
        },
    },
    Idempotency: &gosdk.IdempotencyConfig{
        Header: &gosdk.IdempotencyConfigHeader{
            Name: "name",
        },
    },
    Name: "name",
    Target: &gosdk.IncomingWebhookTarget{
        FixedSandboxName: &gosdk.IncomingWebhookTargetFixedSandboxName{
            SandboxName: "sandbox_name",
        },
    },
}
client.Webhooks.CreateIncomingWebhook(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**auth:** `*gosdk.IncomingWebhookAuth` 
    
</dd>
</dl>

<dl>
<dd>

**idempotency:** `*gosdk.IdempotencyConfig` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**rules:** `[]*gosdk.IncomingWebhookRule` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*gosdk.IncomingWebhookStatus` 
    
</dd>
</dl>

<dl>
<dd>

**target:** `*gosdk.IncomingWebhookTarget` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.GetIncomingWebhook(WebhookID) -> *gosdk.IncomingWebhook</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get one incoming webhook by ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetIncomingWebhookRequest{
    WebhookID: "webhook_id",
}
client.Webhooks.GetIncomingWebhook(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**webhookID:** `string` — Incoming webhook ID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.DeleteIncomingWebhook(WebhookID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Soft-delete an incoming webhook receiver.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.DeleteIncomingWebhookRequest{
    WebhookID: "webhook_id",
}
client.Webhooks.DeleteIncomingWebhook(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**webhookID:** `string` — Incoming webhook ID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.UpdateIncomingWebhook(WebhookID, request) -> *gosdk.IncomingWebhook</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Partially update an incoming webhook receiver. Provided top-level fields replace the existing values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.IncomingWebhookUpdate{
    WebhookID: "webhook_id",
}
client.Webhooks.UpdateIncomingWebhook(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**webhookID:** `string` — Incoming webhook ID
    
</dd>
</dl>

<dl>
<dd>

**auth:** `*gosdk.IncomingWebhookAuth` 
    
</dd>
</dl>

<dl>
<dd>

**idempotency:** `*gosdk.IdempotencyConfig` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**rules:** `[]*gosdk.IncomingWebhookRule` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*gosdk.IncomingWebhookStatus` 
    
</dd>
</dl>

<dl>
<dd>

**target:** `*gosdk.IncomingWebhookTarget` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.ListWebhookDeliveries(WebhookID) -> []*gosdk.WebhookDeliverySummary</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List delivery events for an incoming webhook, with optional status and date filters.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.ListWebhookDeliveriesRequest{
    WebhookID: "webhook_id",
}
client.Webhooks.ListWebhookDeliveries(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**webhookID:** `string` — Incoming webhook ID
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int64` 
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int64` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*gosdk.IngressEventStatus` 
    
</dd>
</dl>

<dl>
<dd>

**from:** `*time.Time` 
    
</dd>
</dl>

<dl>
<dd>

**to:** `*time.Time` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.GetWebhookDelivery(WebhookID, EventID) -> *gosdk.WebhookDeliveryDetail</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get full detail for a single webhook delivery event, including a truncated body preview and action attempts.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &gosdk.GetWebhookDeliveryRequest{
    WebhookID: "webhook_id",
    EventID: "event_id",
}
client.Webhooks.GetWebhookDelivery(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**webhookID:** `string` — Incoming webhook ID
    
</dd>
</dl>

<dl>
<dd>

**eventID:** `string` — Delivery event ID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

