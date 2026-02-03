
# How connect to server

See [docs](https://adaptix-framework.gitbook.io/adaptix-framework/development/new-extenders)



# Agent Extender Template

**config.yaml**
* Change the path to .so goplagins (`_SO_FILE_HERE_`) in the `extender_file` parameter.
* Set the agent's registration name (`_AGENT_`) in the `agent_name` parameter.
* Set the 8-character hex value of the agent watermark (`_RANDOM_HEX_8_`) in the `agent_watermark`parameter.
* Set the listeners supported by the agent (`_LISTENER_1_`, `_LISTENER_2_`) in the `listeners` parameter.

**go.mod**
* Replace `_AGENT_` with the agent's registration name.

**Makefile**
* Replace `_AGENT_` with the agent's registration name.

**ax_config.axs**
* Register commands in the `RegisterCommands` function.
* Create an agent generation form in the `GenerateUI` function.

**pl_main.go**
* Specify your code inside the functions in the `START CODE HERE` and `END CODE HERE` tags.



# Listener Extender Template

**config.yaml**
* Change the path to .so goplagins (`_SO_FILE_HERE_`) in the `extender_file` parameter.
* Set the listener's registration name (`_LISTENER_`) in the `listener_name` parameter.
* Set protocol designation (`_PROTOCOL_`) in the `protocol` parameter.

**go.mod**
* Replace `_LISTENER_` with the listener's registration name.

**Makefile**
* Replace `_LISTENER_` with the listener's registration name.

**ax_config.axs**
* Create a listener creation form in the `ListenerUI` function.

**pl_main.go**
* Specify your code inside the functions in the `START CODE HERE` and `END CODE HERE` tags.



# Service Extender Template

**config.yaml**
* Change the path to .so goplagins (`_SO_FILE_HERE_`) in the `extender_file` parameter.
* Set the service's registration name (`_SERVICE_`) in the `service_name` parameter.

**go.mod**
* Replace `_SERVICE_` with the services's registration name.

**Makefile**
* Replace `_SERVICE_` with the services's registration name.

**ax_config.axs**
* Set `InitService` and `data_handler` functions.

**pl_main.go**
* Specify your code inside the functions in the `START CODE HERE` and `END CODE HERE` tags.
