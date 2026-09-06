var metadata = {
    name: "_SERVICE_",
    description: "TODO: service extender"
};

function RegisterServiceCommands() {
    let cmd_get = ax.create_command("get_config", "Return service config");
    let group = ax.create_commands_group("_SERVICE_", [cmd_get]);
    ax.register_service_commands(group);
}

function InitService() {
}

function data_handler(data) {
    /// START CODE HERE
    // let response = JSON.parse(data);
    /// END CODE HERE
}
