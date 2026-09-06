# Extender templates

Generic stubs for AdaptixC2 **v2** plugins. This directory is the source that is published as [templates-extender](https://github.com/Adaptix-Framework/templates-extender). `axtool template` clones that GitHub repo (or `--from` this path).

These are **not** copies of a specific agent or listener. They implement the same **plugin contract** every current v2 agent, listener, and service uses. Protocol, crypto, and implant code go between `START CODE HERE` / `END CODE HERE`.

Docs: [Extenders](https://adaptix-framework.gitbook.io/adaptix-framework/development/extenders).

---

## Scaffold

```bash
# default: clone templates-extender@main
axtool template agent my_agent
axtool template listener my_http --protocol http
axtool template service my_svc
axtool template axscript my_kit

# preview this tree before a GitHub push
axtool template listener my_http --from ./AdaptixServer/extenders_test --protocol http
```

`axtool` replaces placeholders, writes `axtool.spec`, then:

```bash
axtool <adaptix.spec> ext install ./my_http -f
```

Does not require `adaptix.spec` for `template` itself.

### Stubs:

| File                | Role                                                                      |
|---------------------|---------------------------------------------------------------------------|
| `config.yaml`       | Loader metadata (`extender_type`, names, protocol)                        |
| `pl_main.go`        | `InitPlugin` + interface methods                                          |
| `ax_config.axs`     | GUI: commands / listener form / service hooks                             |
| `Makefile`          | `go build -buildmode=plugin` → `dist/`                                    |
| `go.mod` / `go.sum` | axc2/v2                                                                   |
| `axtool.spec`       | Recreated by `axtool template`; kept here so a manual copy still installs |

### Placeholders (`axtool template`)

| Token                                  | Agent       | Listener                        | Service     |
|----------------------------------------|-------------|---------------------------------|-------------|
| `_AGENT_` / `_LISTENER_` / `_SERVICE_` | plugin name | plugin name                     | plugin name |
| `_SO_FILE_HERE_`                       | `<name>.so` | `<name>.so`                     | `<name>.so` |
| `_RANDOM_HEX_8_`                       | watermark   | —                               | —           |
| `_PROTOCOL_`                           | —           | `--protocol` (default `custom`) | —           |
| `adaptix_agent_NAME`                   | Go module   | —                               | —           |
| `adaptix_listener_NAME_PROTOCOL`       | —           | Go module                       | —           |
| `adaptix_service_NAME_PROTOCOL`        | —           | —                               | Go module   |
| `listener_LISTENER_`                   | —           | Makefile helper                 | —           |

---

## Agent (`agent_template`)

**Go — `InitPlugin(ts, moduleDir, watermark) adaptix.PluginAgent`**

| Method             |                                                                                                      |
|--------------------|------------------------------------------------------------------------------------------------------|
| `GenerateProfiles` | listener blobs → implant config                                                                      |
| `BuildPayload`     | `profile.AgentConfig` is JSON from `GenerateUI`                                                      |
| `CreateAgent`      | parse check-in beat                                                                                  |
| `AgentRestore`     | `CreateCommand`, `ProcessData`, `Encrypt`/`Decrypt`, `PackTasks`, `PivotPackData`, tunnels, terminal |
| `Call`             | GUI RPC (`TsPluginAgentSendDataClient`)                                                              |

Fill `START CODE HERE` for the wire format.

**AxScript**

| Function                         |                                                           |
|----------------------------------|-----------------------------------------------------------|
| `RegisterCommands(listenerType)` | command groups (`commands_windows` / `_linux` / `_macos`) |
| `GenerateUI(listeners_type)`     | payload dialog; values land in `BuildProfile.AgentConfig` |

**config.yaml:** `agent_name`, `agent_watermark`, `listeners` (uncomment names the agent supports), `multi_listeners`.

---

## Listener (`listener_template`)

**Go — `InitPlugin(ts, moduleDir, listenerDir) adaptix.PluginListener`**

| Method                    |                                                                                   |
|---------------------------|-----------------------------------------------------------------------------------|
| `Create`                  | parse UI JSON or restore `customData`; return `ExtenderListener` + `ListenerData` |
| `Call`                    | GUI RPC (`TsPluginListenerSendDataClient`)                                        |
| `Start` / `Stop` / `Edit` | bind / unbind / live config                                                       |
| `GetProfile`              | blob for the agent `GenerateProfiles`                                             |
| `InternalHandler`         | pivot path; returns **agent id `int64`**                                          |

`TransportConfig` is the JSON of the `ListenerUI` container (`host_bind`, `port_bind`, `protocol`). Split transport into `pl_transport.go` if needed (see Makefile comment).

**AxScript:** `ListenerUI(mode_create)` must return `{ ui_panel, ui_container, ui_height, ui_width }`. Put form fields into the container (`host_bind`, `port_bind`, …).

**config.yaml:** `listener_name`, `protocol`, `listener_type`: `external` (HTTP/DNS/…) or `internal` (pivot SMB/TCP).

---

## Service (`service_template`)

**Go — `InitPlugin(ts, moduleDir, serviceConfig) adaptix.PluginService`**

| Method    |                                                                            |
|-----------|----------------------------------------------------------------------------|
| `CallRPC` | sync JSON result                                                           |
| `Call`    | async; default stub forwards `CallRPC` via `TsPluginServiceSendDataClient` |

Persist settings with `TsExtenderDataSave` / `TsExtenderDataLoad` (helpers commented in `pl_main.go`).

**AxScript**

| Function                    |                                      |
|-----------------------------|--------------------------------------|
| `metadata`                  | `{ name, description }`              |
| `InitService()`             | plugin load                          |
| `data_handler(data)`        | push from `TsPluginServiceSendData*` |
| `RegisterServiceCommands()` | optional console commands            |

---

## Build

From a filled stub:

```bash
make          # → dist/*.so + config.yaml + ax_config.axs
```

Requires Go **1.26+** and a teamserver `go.work` that uses the same `axc2/v2` as the server (axtool `ext install` adds the module).
