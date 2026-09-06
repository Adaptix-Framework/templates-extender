/// _AGENT_

function RegisterCommands(listener_type)
{
    /// Commands Here
    // listenerType is the listener registration name (BeaconHTTP, GopherTCP, …).

    let cmd_pwd = ax.create_command("pwd", "Print working directory", "pwd", "Task: pwd");

    let commands = ax.create_commands_group("_AGENT_", [cmd_pwd]);

    return {
        commands_windows: commands,
        commands_linux: commands,
        commands_macos: commands
    }
}

function GenerateUI(listeners_type)
{
    /// Form Here
    // listeners_type: array of selected listener types.

    let labelArch = form.create_label("Arch:");
    let comboArch = form.create_combo();
    comboArch.addItems(["x64", "x86"]);

    let layout = form.create_gridlayout();
    layout.addWidget(labelArch, 0, 0);
    layout.addWidget(comboArch, 0, 1);

    let container = form.create_container();
    container.put("arch", comboArch);

    let panel = form.create_panel();
    panel.setLayout(layout);

    return {
        ui_panel: panel,
        ui_container: container,
        ui_height: 300,
        ui_width: 400
    }
}
