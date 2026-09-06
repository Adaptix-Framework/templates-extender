/// _LISTENER_

function ListenerUI(mode_create)
{
    /// FORM HERE
    // mode_create: true = create, false = edit.

    let labelHost = form.create_label("Host:");
    let comboHost = form.create_combo();
    comboHost.setEnabled(mode_create);
    let addrs = ax.interfaces();
    for (let item of addrs) { comboHost.addItem(item); }

    let labelPort = form.create_label("Port:");
    let spinPort = form.create_spin();
    spinPort.setRange(1, 65535);
    spinPort.setValue(443);
    spinPort.setEnabled(mode_create);

    let layout = form.create_gridlayout();
    layout.addWidget(labelHost, 0, 0, 1, 1);
    layout.addWidget(comboHost, 0, 1, 1, 1);
    layout.addWidget(labelPort, 1, 0, 1, 1);
    layout.addWidget(spinPort, 1, 1, 1, 1);

    let container = form.create_container();
    container.put("host_bind", comboHost);
    container.put("port_bind", spinPort);

    let panel = form.create_panel();
    panel.setLayout(layout);

    return {
        ui_panel: panel,
        ui_container: container,
        ui_height: 360,
        ui_width: 520
    }
}
