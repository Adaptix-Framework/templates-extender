/// _AGENT_

function RegisterCommands(listener_type)
{

/// Commands Here

    if(listener_type == "_LISTENER_") {
        let commands_external = ax.create_commands_group("_AGENT_", [] );

        return { commands_windows: commands_external }
    }
    return ax.create_commands_group("none",[]);
}

function GenerateUI(listeners_type)
{

/// Form Here

    let container = form.create_container()

    let panel = form.create_panel()

    return {
        ui_panel: panel,
        ui_container: container,
        ui_height: 450,
        ui_width: 550
    }
}
