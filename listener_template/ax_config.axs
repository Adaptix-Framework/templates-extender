/// _LISTENER_

function ListenerUI(mode_create)
{

/// FORM HERE

    let container = form.create_container();

    let panel = form.create_panel();

    return {
        ui_panel: panel,
        ui_container: container,
        ui_height: 650,
        ui_width: 650
    }
}