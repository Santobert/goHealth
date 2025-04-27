# GoHealth

**goHealth** is a lightweight system monitoring tool written in **Go**, designed to provide essential **metrics** and **health status** information for various aspects of the host machine.
Inspired by [Prometheus Node Exporter](https://github.com/prometheus/node_exporter), it integrates seamlessly with [Uptime Kuma](https://github.com/louislam/uptime-kuma), offering real-time status updates for system components.
The application is distributed as a **single binary** with **no dependencies**, ensuring easy deployment and minimal resource consumption.
Ideal for users who require a straightforward, efficient monitoring solution without the complexity of additional frameworks.

## Usage

To use **goHealth**, follow these steps:

1. **Download the Binary**
   Obtain the latest release of **goHealth** from the [releases page](https://github.com/Santobert/goHealth/releases).

2. **Prepare the Configuration File** (Optional)
   Create a configuration file (e.g., `config.yaml`) with the necessary settings for your system monitoring.

3. **Run the Application**
   Execute the binary with the following command-line parameters:

   ```bash
   ./goHealth -config /path/to/config.yaml -port 9100
   ```

   - `-config`: Specifies the path to the configuration file (default is `config.yaml` in your current directory)
   - `-port`: Sets the HTTP port for the monitoring service (default is `9100` if not specified)

## Configuration

An example configuration can be found [here](examples/config.yaml).

1. `load`
   - `max_load`: Specifies the maximum load threshold per CPU.
2. `memory`
   - `max_memory`: Sets the maximum percentage of physical memory usage allowed.
   - `max_swap`: Sets the maximum percentage of swap memory usage allowed.
   - `swap_enabled`: Enables the swap healthcheck. Default is `true`.
3. `disk`
   - `max_disk`: Defines the maximum percentage of disk usage allowed.
   - `paths`: Lists the directories or mount points to monitor for disk usage.
   - `ignore`: Specifies paths to be excluded from auto-discovery.
   - `auto`: Enable auto-discovery of physical partitions for disk checking. Default is `true`.

## Deployment

The released binaries usually run without any dependencies.
To configure **goHealth** to start automatically on system boot using **systemd**, follow these steps:

1. **Copy the Service File**  
   Locate the [example service file](examples/gohealth.service) and copy it to `/etc/systemd/system/`.

2. **Create the Configuration File**  
   Ensure that a configuration file exists at `/etc/gohealth/config.yaml`. You can create or copy the [example configuration](examples/config.yaml).

3. **Enable and Start Service**  
   Reload systemd to apply changes, enable the service to start on boot, and start it immediately.
