
# How connect to server

See [docs](https://adaptix-framework.gitbook.io/adaptix-framework/development/new-extenders)



# Agent Extender Template

**config.json**
* Change the path to .so goplaings (`_SO_FILE_HERE_`) in the `extender_file` parameter.
* Set the agent's registration name (`_AGENT_`) in the `agent_name` parameter.
* Set the 8-character hex value of the agent watermark (`_RANDOM_HEX_8_`) in the `agent_watermark`parameter.
* Set the listeners supported by the agent (`_LISTENER_1_`, `_LISTENER_2_`) in the `listeners` parameter.

**Makefile**
* Replace `_AGENT_` with the agent's registration name.

**ax_config.axs**
* Register commands in the `RegisterCommands` function. 
* Create an agent generation form in the `GenerateUI` function.

**pl_agent.go**
* Specify your code inside the functions in the `START CODE HERE` and `END CODE HERE` tags.

By default, no modification to the **pl_main.go** file is required.



# Listener (External) Extender Template

**config.json**
* Change the path to .so goplaings (`_SO_FILE_HERE_`) in the `extender_file` parameter.
* Set the listener's registration name (`_LISTENER_`) in the `listener_name` parameter.
* Set protocol designation (`_PROTOCOL_`) in the `protocol` parameter.

**Makefile**
* Replace `_LISTENER_` with the listener's registration name.

**ax_config.axs**
* Create a listener creation form in the `ListenerUI` function.

**pl_listener.go**
* Specify your code inside the functions in the `START CODE HERE` and `END CODE HERE` tags.

By default, no modification to the **pl_main.go** file is required.



# Listener (Internal) Extender Template

**config.json**
* Change the path to .so goplaings (`_SO_FILE_HERE_`) in the `extender_file` parameter.
* Set the listener's registration name (`_LISTENER_`) in the `listener_name` parameter.
* Set protocol designation (`_PROTOCOL_`) in the `protocol` parameter.

**Makefile**
* Replace `_LISTENER_` with the listener's registration name.

**ax_config.axs**
* Create a listener creation form in the `ListenerUI` function.

**pl_listener.go**
* Specify your code inside the functions in the `START CODE HERE` and `END CODE HERE` tags.

By default, no modification to the **pl_main.go** file is required.
